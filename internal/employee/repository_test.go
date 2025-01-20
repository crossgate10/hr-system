package employee

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type RepositoryTestSuite struct {
	suite.Suite
	db   *gorm.DB
	repo Repository
	ctx  context.Context
}

func (suite *RepositoryTestSuite) SetupTest() {
	var err error
	suite.db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	suite.NoError(err)

	err = suite.db.AutoMigrate(&Employee{})
	suite.NoError(err)

	suite.repo = NewRepository(suite.db)
	suite.ctx = context.Background()
}

func (suite *RepositoryTestSuite) TearDownTest() {
	db, err := suite.db.DB()
	suite.NoError(err)
	db.Close()
}

func (suite *RepositoryTestSuite) TestCreateEmployee() {
	employee := &Employee{
		Name:        "John Doe",
		Position:    "Engineer",
		ContactInfo: "john.doe@example.com",
		Salary:      50000,
		Status:      Active,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err := suite.repo.CreateEmployee(suite.ctx, employee)
	suite.NoError(err)

	var result Employee
	suite.db.First(&result, employee.ID)
	assert.Equal(suite.T(), employee.Name, result.Name)
	assert.Equal(suite.T(), employee.Position, result.Position)
	assert.Equal(suite.T(), employee.ContactInfo, result.ContactInfo)
	assert.Equal(suite.T(), employee.Salary, result.Salary)
}

func (suite *RepositoryTestSuite) TestGetEmployeeByID() {
	employee := &Employee{
		Name:        "Jane Doe",
		Position:    "Manager",
		ContactInfo: "jane.doe@example.com",
		Salary:      60000,
		Status:      Active,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	suite.db.Create(employee)

	result, err := suite.repo.GetEmployeeByID(suite.ctx, employee.ID)
	suite.NoError(err)
	assert.Equal(suite.T(), employee.Name, result.Name)
	assert.Equal(suite.T(), employee.Position, result.Position)
	assert.Equal(suite.T(), employee.ContactInfo, result.ContactInfo)
	assert.Equal(suite.T(), employee.Salary, result.Salary)
}

func (suite *RepositoryTestSuite) TestListEmployees() {
	employees := []Employee{
		{Name: "Alice", Position: "Developer", ContactInfo: "alice@example.com", Salary: 55000, Status: Active, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Bob", Position: "Designer", ContactInfo: "bob@example.com", Salary: 50000, Status: Active, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}
	suite.db.Create(&employees)

	filter := &ListEmployeesFilter{Name: "Alice"}
	results, err := suite.repo.ListEmployees(suite.ctx, filter)
	suite.NoError(err)
	assert.Len(suite.T(), results, 1)
	assert.Equal(suite.T(), "Alice", results[0].Name)
}

func (suite *RepositoryTestSuite) TestUpdateEmployee() {
	employee := &Employee{
		Name:        "Charlie",
		Position:    "Tester",
		ContactInfo: "charlie@example.com",
		Salary:      45000,
		Status:      Active,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	suite.db.Create(employee)

	employee.Salary = 47000
	err := suite.repo.UpdateEmployee(suite.ctx, employee)
	suite.NoError(err)

	var result Employee
	suite.db.First(&result, employee.ID)
	assert.Equal(suite.T(), 47000.0, result.Salary)
}

func (suite *RepositoryTestSuite) TestDeleteEmployee() {
	employee := &Employee{
		Name:        "Dave",
		Position:    "Admin",
		ContactInfo: "dave@example.com",
		Salary:      52000,
		Status:      Active,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	suite.db.Create(employee)

	err := suite.repo.DeleteEmployee(suite.ctx, employee.ID)
	suite.NoError(err)

	var result Employee
	err = suite.db.First(&result, employee.ID).Error
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), gorm.ErrRecordNotFound, err)
}

func TestRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}
