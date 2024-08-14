package usecase_test

import (
	"errors"
	"testing"
	"task_manager/domain"
	"task_manager/infrastructure"
	"task_manager/usecase"
	"task_manager/mocks"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserUsecaseTestSuite is a test suite for UserUsecase
type UserUsecaseTestSuite struct {
	suite.Suite
	mockRepo    *mocks.UserRepositoryInterface
	userUsecase *usecase.UserUsecase
}

// SetupTest sets up the test environment
func (suite *UserUsecaseTestSuite) SetupTest() {
	suite.mockRepo = new(mocks.UserRepositoryInterface)
	suite.userUsecase = usecase.NewUserUsecase(suite.mockRepo)
}

// TestCreateUser tests the CreateUser method
func (suite *UserUsecaseTestSuite) TestCreateUser() {
	// Success case
	suite.Run("TestCreateUser_Success", func() {
		user := &domain.User{
			ID:       primitive.NewObjectID(),
			Username: "testuser",
			Password: "password123",
		}

		suite.mockRepo.On("GetUserByUsername", user.Username).Return(nil, nil).Once()
		suite.mockRepo.On("AddUser", user).Return(primitive.NewObjectID(), nil).Once()
		infrastructure.HashPassword = func(password string) (string, error) {
			return "hashedPassword", nil
		}

		createdUser, err := suite.userUsecase.CreateUser(user)
		suite.Nil(err)
		suite.NotNil(createdUser)
	})

	// User already exists case
	suite.Run("TestCreateUser_UserExists", func() {
		user := &domain.User{
			Username: "testuser",
			Password: "password123",
		}

		suite.mockRepo.On("GetUserByUsername", user.Username).Return(user, nil).Once()

		createdUser, err := suite.userUsecase.CreateUser(user)
		suite.NotNil(err)
		suite.Nil(createdUser)
		suite.Equal("username already exists", err.Error())
	})

	// Failed to hash password case
	suite.Run("TestCreateUser_HashPasswordFailure", func() {
		user := &domain.User{
			Username: "testuser",
			Password: "password123",
		}

		suite.mockRepo.On("GetUserByUsername", user.Username).Return(nil, nil).Once()
		infrastructure.HashPassword = func(password string) (string, error) {
			return "", errors.New("failed to hash password")
		}

		createdUser, err := suite.userUsecase.CreateUser(user)
		suite.NotNil(err)
		suite.Nil(createdUser)
		suite.Equal("failed to hash password", err.Error())
	})

	// Internal server error case
	suite.Run("TestCreateUser_InternalServerError", func() {
		user := &domain.User{
			Username: "testuser",
			Password: "password123",
		}

		suite.mockRepo.On("GetUserByUsername", user.Username).Return(nil, nil).Once()
		infrastructure.HashPassword = func(password string) (string, error) {
			return "hashedPassword", nil
		}
		suite.mockRepo.On("AddUser", user).Return(nil, errors.New("Internal server error")).Once()

		createdUser, err := suite.userUsecase.CreateUser(user)
		suite.NotNil(err)
		suite.Nil(createdUser)
		suite.Equal("Internal server error", err.Error())
	})
}

// TestLoginUser tests the LoginUser method
func (suite *UserUsecaseTestSuite) TestLoginUser() {
	// Success case
	suite.Run("TestLoginUser_Success", func() {
		authUser := &domain.AuthUser{
			Username: "testuser",
			Password: "password123",
		}

		user := &domain.User{
			Username: "testuser",
			Password: "hashedPassword",
		}

		suite.mockRepo.On("GetUserByUsername", authUser.Username).Return(user, nil).Once()
		infrastructure.ComparePassword = func(hashedPassword, plainPassword string) error {
			return nil
		}


		infrastructure.GenerateJWT = func(user *domain.User) (string, error) {
			return "sampletoken", nil
		}

		token, err := suite.userUsecase.LoginUser(authUser)
		suite.Nil(err)
		suite.Equal("sampletoken", token)
	})

	// User not found case
	suite.Run("TestLoginUser_UserNotFound", func() {
		authUser := &domain.AuthUser{
			Username: "testuser",
			Password: "password123",
		}

		suite.mockRepo.On("GetUserByUsername", authUser.Username).Return(nil, errors.New("not found")).Once()

		token, err := suite.userUsecase.LoginUser(authUser)
		suite.NotNil(err)
		suite.Empty(token)
		suite.Equal("not found", err.Error())
	})

	// Invalid credentials case
	suite.Run("TestLoginUser_InvalidCredentials", func() {
		authUser := &domain.AuthUser{
			Username: "testuser",
			Password: "password123",
		}

		user := &domain.User{
			Username: "testuser",
			Password: "hashedPassword",
		}

		suite.mockRepo.On("GetUserByUsername", authUser.Username).Return(user, nil).Once()
		infrastructure.ComparePassword = func(hashedPassword, plainPassword string) error {
			return errors.New("invalid credentials")
		}

		token, err := suite.userUsecase.LoginUser(authUser)
		suite.NotNil(err)
		suite.Empty(token)
		suite.Equal("invalid credentials", err.Error())
	})

	// JWT generation failure case
	suite.Run("TestLoginUser_JWTGenerationFailure", func() {
		authUser := &domain.AuthUser{
			Username: "testuser",
			Password: "password123",
		}

		user := &domain.User{
			Username: "testuser",
			Password: "hashedPassword",
		}

		suite.mockRepo.On("GetUserByUsername", authUser.Username).Return(user, nil).Once()
		infrastructure.ComparePassword = func(hashedPassword, plainPassword string) error {
			return nil
		}
		infrastructure.GenerateJWT = func(user *domain.User) (string, error) {
			return "", errors.New("Internal server error")
		}

		token, err := suite.userUsecase.LoginUser(authUser)
		suite.NotNil(err)
		suite.Empty(token)
		suite.Equal("Internal server error", err.Error())
	})
}

// TestUserUsecaseTestSuite runs the test suite
func TestUserUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseTestSuite))
}
