package usecase_test

import (
   "task_manager/domain"
   "task_manager/usecase"
   "task_manager/repository"
   "testing"

   "github.com/stretchr/testify/assert"
   "github.com/stretchr/testify/mock"
   "github.com/stretchr/testify/suite"
)

type MockUserRepository struct {
   mock.Mock
}

func (m *MockUserRepository) AddUser(user *domain.User) (interface{}, error) {
   args := m.Called(user)
   return args.Get(0), args.Error(1)
}

func (m *MockUserRepository) GetUserByUsername(username string) (*domain.User, error) {
   args := m.Called(username)
   return args.Get(0).(*domain.User), args.Error(1)
}

type UserUsecaseSuite struct {
   suite.Suite
   usecase *usecase.UserUsecase
   repo    *MockUserRepository
}

func (suite *UserUsecaseSuite) SetupTest() {
   suite.repo = new(MockUserRepository)
   suite.usecase = usecase.NewUserUsecase(suite.repo)
}

func (suite *UserUsecaseSuite) TestCreateUser_Success() {
   user := &domain.User{Username: "newUser", Password: "password"}

   suite.repo.On("GetUserByUsername", "newUser").Return(nil, nil)
   suite.repo.On("AddUser", user).Return(user.ID, nil)

   id, err := suite.usecase.CreateUser(user)
   assert.Nil(suite.T(), err)
   assert.NotNil(suite.T(), id)
}

func (suite *UserUsecaseSuite) TestCreateUser_ExistingUser() {
   existingUser := &domain.User{Username: "existingUser"}

   suite.repo.On("GetUserByUsername", "existingUser").Return(existingUser, nil)

   _, err := suite.usecase.CreateUser(existingUser)
   assert.NotNil(suite.T(), err)
   assert.Equal(suite.T(), "username already exists", err.Error())
}

func TestUserUsecaseSuite(t *testing.T) {
   suite.Run(t, new(UserUsecaseSuite))
}
