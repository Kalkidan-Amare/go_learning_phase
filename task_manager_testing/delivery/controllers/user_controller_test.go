package controller_test

import (
   "task_manager/controller"
   "task_manager/domain"
   "task_manager/usecase"
   "bytes"
   "encoding/json"
   "net/http"
   "net/http/httptest"
   "testing"

   "github.com/gorilla/mux"
   "github.com/stretchr/testify/assert"
   "github.com/stretchr/testify/mock"
   "github.com/stretchr/testify/suite"
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
   controller *controller.UserController
   usecase    *MockUserUsecase
   router     *mux.Router
}

func (suite *UserControllerSuite) SetupTest() {
   suite.usecase = new(MockUserUsecase)
   suite.controller = &controller.UserController{usecase: suite.usecase}
   suite.router = mux.NewRouter()
   suite.router.HandleFunc("/users", suite.controller.CreateUser).Methods("POST")
}

func (suite *UserControllerSuite) TestCreateUser_Success() {
   user := &domain.User{Username: "newUser", Password: "password"}
   jsonUser, _ := json.Marshal(user)

   suite.usecase.On("CreateUser", user).Return(user.ID, nil)

   req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonUser))
   res := httptest.NewRecorder()
   suite.router.ServeHTTP(res, req)

   assert.Equal(suite.T(), http.StatusOK, res.Code)
   assert.Contains(suite.T(), res.Body.String(), user.ID.Hex())
}

func (suite *UserControllerSuite) TestCreateUser_Conflict() {
   user := &domain.User{Username: "existingUser"}
   jsonUser, _ := json.Marshal(user)

   suite.usecase.On("CreateUser", user).Return(nil, errors.New("username already exists"))

   req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonUser))
   res := httptest.NewRecorder()
   suite.router.ServeHTTP(res, req)

   assert.Equal(suite.T(), http.StatusConflict, res.Code)
   assert.Contains(suite.T(), res.Body.String(), "username already exists")
}

func TestUserControllerSuite(t *testing.T) {
   suite.Run(t, new(UserControllerSuite))
}
