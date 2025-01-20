package employee

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

type Service interface {
	AddEmployee(ctx context.Context, name, position, contactInfo string, salary float64) (*Employee, error)
	GetEmployeeDetails(ctx context.Context, id int) (*Employee, error)
	ListEmployees(ctx context.Context, filter *ListEmployeesFilter) ([]Employee, error)
	UpdateEmployeeDetails(ctx context.Context, employee *Employee) error
	RemoveEmployee(ctx context.Context, id int) error
}

type service struct {
	repo     Repository
	cache    *redis.Client
	cacheTTL time.Duration
}

func NewService(repo Repository, cache *redis.Client, cacheTTL time.Duration) Service {
	return &service{repo: repo, cache: cache, cacheTTL: cacheTTL}
}

func (s *service) AddEmployee(ctx context.Context, name, position, contactInfo string, salary float64) (*Employee, error) {
	employee := &Employee{
		Name:        name,
		Position:    position,
		ContactInfo: contactInfo,
		Salary:      salary,
		Status:      Active,
	}
	if err := s.repo.CreateEmployee(ctx, employee); err != nil {
		return nil, err
	}
	return employee, nil
}

func (s *service) GetEmployeeDetails(ctx context.Context, id int) (*Employee, error) {
	// 從快取中查找
	cached, err := s.cache.Get(ctx, s.getEmployeeCacheKey(id)).Result()
	if err == nil {
		var employee Employee
		if err := json.Unmarshal([]byte(cached), &employee); err == nil {
			return &employee, nil
		}
	}

	// 從資料庫中查找
	employee, err := s.repo.GetEmployeeByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// 將結果存入快取
	data, _ := json.Marshal(employee)
	s.cache.Set(ctx, s.getEmployeeCacheKey(id), data, s.cacheTTL)

	return employee, nil
}

func (s *service) ListEmployees(ctx context.Context, filter *ListEmployeesFilter) ([]Employee, error) {
	// 快取鍵
	cacheKey := s.getEmployeesCacheKey(filter)

	// 從快取中查找
	cached, err := s.cache.Get(ctx, cacheKey).Result()
	if err == nil {
		var employees []Employee
		if err := json.Unmarshal([]byte(cached), &employees); err == nil {
			return employees, nil
		}
	}

	// 從資料庫中查找
	employees, err := s.repo.ListEmployees(ctx, filter)
	if err != nil {
		return nil, err
	}

	// 將結果存入快取
	data, _ := json.Marshal(employees)
	s.cache.Set(ctx, cacheKey, data, s.cacheTTL)

	return employees, nil
}

func (s *service) UpdateEmployeeDetails(ctx context.Context, employee *Employee) error {
	if err := s.repo.UpdateEmployee(ctx, employee); err != nil {
		return err
	}

	// 更新快取
	data, _ := json.Marshal(employee)
	s.cache.Set(ctx, s.getEmployeeCacheKey(employee.ID), data, s.cacheTTL)

	return nil
}

func (s *service) RemoveEmployee(ctx context.Context, id int) error {
	if err := s.repo.DeleteEmployee(ctx, id); err != nil {
		return err
	}

	// 刪除快取
	s.cache.Del(ctx, s.getEmployeeCacheKey(id))

	return nil
}

func (s *service) getEmployeeCacheKey(id int) string {
	return "employee:" + strconv.Itoa(id)
}

func (s *service) getEmployeesCacheKey(filter *ListEmployeesFilter) string {
	if filter == nil {
		return "employees:all"
	}

	var parts []string
	if filter.Name != "" {
		parts = append(parts, fmt.Sprintf("name=%s", filter.Name))
	}
	if filter.Position != "" {
		parts = append(parts, fmt.Sprintf("position=%s", filter.Position))
	}
	if filter.MinSalary > 0 {
		parts = append(parts, fmt.Sprintf("minSalary=%.2f", filter.MinSalary))
	}
	if filter.MaxSalary > 0 {
		parts = append(parts, fmt.Sprintf("maxSalary=%.2f", filter.MaxSalary))
	}

	if len(parts) == 0 {
		return "employees:all"
	}

	return "employees:" + strings.Join(parts, ":")
}
