package employee

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) CreateEmployee(ctx context.Context, employee *Employee) error {
	args := m.Called(ctx, employee)
	return args.Error(0)
}

func (m *MockRepository) GetEmployeeByID(ctx context.Context, id int) (*Employee, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*Employee), args.Error(1)
}

func (m *MockRepository) ListEmployees(ctx context.Context, filter *ListEmployeesFilter) ([]Employee, error) {
	args := m.Called(ctx, filter)
	return args.Get(0).([]Employee), args.Error(1)
}

func (m *MockRepository) UpdateEmployee(ctx context.Context, employee *Employee) error {
	args := m.Called(ctx, employee)
	return args.Error(0)
}

func (m *MockRepository) DeleteEmployee(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

type ServiceTestSuite struct {
	suite.Suite
	repo    *MockRepository
	redis   *miniredis.Miniredis
	cache   *redis.Client
	service Service
}

func (suite *ServiceTestSuite) SetupTest() {
	suite.repo = new(MockRepository)

	redisServer, err := miniredis.Run()
	suite.NoError(err)
	suite.redis = redisServer
	suite.cache = redis.NewClient(&redis.Options{
		Addr: redisServer.Addr(),
	})

	suite.service = NewService(suite.repo, suite.cache, 1*time.Hour)
}

func (suite *ServiceTestSuite) TearDownTest() {
	suite.cache.Close()
	suite.redis.Close()
}

func (suite *ServiceTestSuite) TestAddEmployee() {
	ctx := context.Background()
	employee := &Employee{Name: "John Doe", Position: "Engineer", Status: Active, ContactInfo: "john.doe@example.com", Salary: 50000}

	suite.repo.On("CreateEmployee", ctx, employee).Return(nil)

	result, err := suite.service.AddEmployee(ctx, employee.Name, employee.Position, employee.ContactInfo, employee.Salary)

	suite.NoError(err)
	suite.Equal(employee.Name, result.Name)
	suite.Equal(employee.Position, result.Position)
	suite.Equal(employee.ContactInfo, result.ContactInfo)
	suite.Equal(employee.Salary, result.Salary)

	suite.repo.AssertCalled(suite.T(), "CreateEmployee", ctx, employee)
}

func (suite *ServiceTestSuite) TestGetEmployeeDetails_FromCache() {
	ctx := context.Background()
	employee := &Employee{ID: 1, Name: "John Doe", Position: "Engineer", ContactInfo: "john.doe@example.com", Salary: 50000}
	data, _ := json.Marshal(employee)
	suite.cache.Set(ctx, "employee:1", data, 1*time.Hour)

	result, err := suite.service.GetEmployeeDetails(ctx, employee.ID)

	suite.NoError(err)
	suite.Equal(employee.Name, result.Name)
	suite.Equal(employee.Position, result.Position)
	suite.Equal(employee.ContactInfo, result.ContactInfo)
	suite.Equal(employee.Salary, result.Salary)
}

func (suite *ServiceTestSuite) TestGetEmployeeDetails_FromRepo() {
	ctx := context.Background()
	employee := &Employee{ID: 1, Name: "John Doe", Position: "Engineer", ContactInfo: "john.doe@example.com", Salary: 50000}

	suite.repo.On("GetEmployeeByID", ctx, employee.ID).Return(employee, nil)

	result, err := suite.service.GetEmployeeDetails(ctx, employee.ID)

	suite.NoError(err)
	suite.Equal(employee.Name, result.Name)
	suite.Equal(employee.Position, result.Position)
	suite.Equal(employee.ContactInfo, result.ContactInfo)
	suite.Equal(employee.Salary, result.Salary)

	// 確認數據被緩存
	cached, err := suite.cache.Get(ctx, "employee:1").Result()
	suite.NoError(err)
	var cachedEmployee Employee
	err = json.Unmarshal([]byte(cached), &cachedEmployee)
	suite.NoError(err)
	suite.Equal(employee.ID, cachedEmployee.ID)
}

func (suite *ServiceTestSuite) TestUpdateEmployeeDetails() {
	ctx := context.Background()
	employee := &Employee{ID: 1, Name: "John Doe", Position: "Engineer", ContactInfo: "john.doe@example.com", Salary: 50000}

	suite.repo.On("UpdateEmployee", ctx, employee).Return(nil)

	err := suite.service.UpdateEmployeeDetails(ctx, employee)
	suite.NoError(err)

	// 確認數據被緩存
	cached, err := suite.cache.Get(ctx, "employee:1").Result()
	suite.NoError(err)
	var cachedEmployee Employee
	err = json.Unmarshal([]byte(cached), &cachedEmployee)
	suite.NoError(err)
	suite.Equal(employee.ID, cachedEmployee.ID)
}

func (suite *ServiceTestSuite) TestRemoveEmployee() {
	ctx := context.Background()
	employee := &Employee{ID: 1, Name: "John Doe", Position: "Engineer", ContactInfo: "john.doe@example.com", Salary: 50000}

	suite.repo.On("DeleteEmployee", ctx, employee.ID).Return(nil)

	err := suite.service.RemoveEmployee(ctx, employee.ID)
	suite.NoError(err)

	// 確認數據被從緩存中刪除
	_, err = suite.cache.Get(ctx, "employee:1").Result()
	suite.Error(err)
	suite.Equal(redis.Nil, err)
}

func (suite *ServiceTestSuite) TestListEmployees_FromCache() {
	ctx := context.Background()
	filter := &ListEmployeesFilter{Name: "John"}
	employees := []Employee{
		{ID: 1, Name: "John Doe", Position: "Engineer", ContactInfo: "john.doe@example.com", Salary: 50000},
	}
	data, _ := json.Marshal(employees)
	suite.cache.Set(ctx, "employees:name=John", data, 1*time.Hour)

	result, err := suite.service.ListEmployees(ctx, filter)
	suite.NoError(err)
	suite.Len(result, 1)
	suite.Equal(employees[0].Name, result[0].Name)
}

func (suite *ServiceTestSuite) TestListEmployees_FromRepo() {
	ctx := context.Background()
	filter := &ListEmployeesFilter{Name: "John"}
	employees := []Employee{
		{ID: 1, Name: "John Doe", Position: "Engineer", ContactInfo: "john.doe@example.com", Salary: 50000},
	}

	suite.repo.On("ListEmployees", ctx, filter).Return(employees, nil)

	result, err := suite.service.ListEmployees(ctx, filter)
	suite.NoError(err)
	suite.Len(result, 1)
	suite.Equal(employees[0].Name, result[0].Name)

	// 確認數據被緩存
	cached, err := suite.cache.Get(ctx, "employees:name=John").Result()
	suite.NoError(err)
	var cachedEmployees []Employee
	err = json.Unmarshal([]byte(cached), &cachedEmployees)
	suite.NoError(err)
	suite.Len(cachedEmployees, 1)
	suite.Equal(employees[0].Name, cachedEmployees[0].Name)
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}
