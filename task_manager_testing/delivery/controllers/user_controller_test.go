package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"task_manager/delivery/controllers"
	"task_manager/domain"
	"task_manager/mocks"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockUserUsecase struct {
   mock.Mock
}

func (m *MockUserUsecase) CreateUser(user *domain.User) (interface{}, error) {
   args := m.Called(user)
   return args.Get(0), args.Error(1)
}

type UserControllerSuite struct {
   suite.Suite
   controller *controllers.UserController
   mockUsecase    *mocks.UserUsecaseInterface
   // router     *mux.Router
}

func (suite *UserControllerSuite) SetupTest() {
   suite.mockUsecase = new(mocks.UserUsecaseInterface)
   suite.controller = controllers.NewUserController(suite.mockUsecase)
}

func (suite *UserControllerSuite) TestRegisterController() {
	// Success case
	suite.Run("TestRegister_success", func() {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		user := &domain.User{
         ID:       primitive.NewObjectID(),
         Username: "testuser",
         Password: "password123",
     }

		// registerduser := &domain.User{
		// 	ID:       user.ID,
		// 	Username: user.Username,
		// 	Role:     user.Role,
		// }

		suite.mockUsecase.On("CreateUser", user).Return(user, nil).Once()

		body, err := json.Marshal(user)
		suite.Nil(err)
		ctx.Request = httptest.NewRequest("POST", "/register", bytes.NewBuffer(body))

		suite.controller.Register(ctx)
      expectedResponse, err := json.Marshal(user)
      suite.Nil(err)

		suite.Equal(201, w.Code)
		suite.Equal(string(expectedResponse), w.Body.String())
	})

	// Invalid request case
	suite.Run("TestRegister_invalid_request", func() {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		ctx.Request = httptest.NewRequest("POST", "/register", nil)

		suite.controller.Register(ctx)
      expected, err := json.Marshal(gin.H{"error": "Invalid user"})
		suite.Nil(err)

		suite.Equal(400, w.Code)
		suite.Equal(string(expected), w.Body.String())
	})

	// Failure case
	suite.Run("TestRegister_failure", func() {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		user := &domain.User{
         ID:       primitive.NewObjectID(),
         Username: "testuser",
         Password: "password123",
     }

		suite.mockUsecase.On("CreateUser", user).Return(nil, errors.New("Internal server error")).Once()

		body, err := json.Marshal(user)
		suite.Nil(err)
		ctx.Request = httptest.NewRequest("POST", "/register", bytes.NewReader(body))

		suite.controller.Register(ctx)
		expected, err := json.Marshal(gin.H{"error": "Internal server error"})
		suite.Nil(err)

		suite.Equal(500, w.Code)
		suite.Equal(string(expected), w.Body.String())
	})
}

func (suite *UserControllerSuite) TestLoginController() {
	// Success case
	suite.Run("TestLogin_success", func() {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		authUser := &domain.AuthUser{
			Username: "testuser",
			Password: "password123",
		}
      token := "sampletoken"

		suite.mockUsecase.On("LoginUser", authUser).Return("token", nil).Once()

		body, err := json.Marshal(authUser)
		suite.Nil(err)
		ctx.Request = httptest.NewRequest("POST", "/login", bytes.NewReader(body))

		suite.controller.Login(ctx)
		expectedResponse, err := json.Marshal(gin.H{"token":token})
      suite.Nil(err)

		suite.Equal(200, w.Code)
		suite.Equal(string(expectedResponse), w.Body.String())
	})

	// Invalid request case
	suite.Run("TestLogin_invalid_request", func() {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		ctx.Request = httptest.NewRequest("POST", "/login", nil)

		suite.controller.Login(ctx)
		expected, err := json.Marshal(gin.H{"error": "Invalid auth user"})
		suite.Nil(err)

		suite.Equal(400, w.Code)
		suite.Equal(string(expected), w.Body.String())
	})

	// Failure case
	suite.Run("TestLogin_failure", func() {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		authUser := &domain.AuthUser{
			Username: "testuser",
			Password: "password123",
		}

		suite.mockUsecase.On("LoginUser", authUser).Return("", errors.New("Internal server error")).Once()

		body, err := json.Marshal(authUser)
		suite.Nil(err)
		ctx.Request = httptest.NewRequest("POST", "/login", bytes.NewReader(body))

		suite.controller.Login(ctx)
		expected, err := json.Marshal(gin.H{"error": "Internal server error"})
		suite.Nil(err)

		suite.Equal(500, w.Code)
		suite.Equal(string(expected), w.Body.String())
	})
}



func (suite *UserControllerSuite) TearDownSuite() {
   suite.mockUsecase.AssertExpectations(suite.T())
}

func TestUserControllerSuite(t *testing.T) {
   suite.Run(t, new(UserControllerSuite))
}
