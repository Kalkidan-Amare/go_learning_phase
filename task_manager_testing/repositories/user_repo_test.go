package repositories_test

import (
   "task_manager/domain"
   "task_manager/repositories"
   "task_manager/infrastructure"
   "task_manager/mocks"
   "testing"

   // "github.com/stretchr/testify/assert"
   "github.com/stretchr/testify/mock"
   "github.com/stretchr/testify/suite"
   "go.mongodb.org/mongo-driver/bson/primitive"
   "go.mongodb.org/mongo-driver/mongo"


)

// var mockUser = mock.AnythingOfType("*domain.User")

type UserRepositorySuite struct {
   suite.Suite
   repo *repositories.UserRepository
   collection *mocks.Collection
}

func (suite *UserRepositorySuite) SetupTest() {
	suite.collection = new(mocks.Collection)
   suite.repo = repositories.NewUserRepository(suite.collection)

   infrastructure.HashPassword = func(password string) (string, error) {
       return "password123", nil
   }
}

// func (suite *UserRepositorySuite) SetupSuite() {
//    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
//    client, err := mongo.Connect(context.TODO(), clientOptions)
//    assert.Nil(suite.T(), err)
   
//    suite.db = client.Database("test_db")
//    suite.repo = repositories.UserRepository(mocks)
// }



func (suite *UserRepositorySuite) TearDownSuite() {
   suite.collection.AssertExpectations(suite.T())
}

func (suite *UserRepositorySuite) TestAddUser() {
	// success case
	suite.Run("Create success", func() {
		user := &domain.User{
         ID:       primitive.NewObjectID(),
         Username: "testuser",
         Password: "password123",
     }

     insertResult := &mongo.InsertOneResult{InsertedID: user.ID}
     suite.collection.On("InsertOne", mock.Anything, user).Return(insertResult, nil).Once()

      resID,err := suite.repo.AddUser(user)


		suite.NoError(err)
      suite.Equal(user.ID, resID)
	})

	// insertion failure case
	suite.Run("AddUser_Failure", func() {
		user := &domain.User{
         ID:       primitive.NewObjectID(),
         Username: "testuser",
         Password: "password123",
     }
      suite.collection.On("InsertOne", mock.Anything, user).Return(&mongo.InsertOneResult{}, mongo.ErrClientDisconnected).Once()
		
		
		resID,err := suite.repo.AddUser(user)

		suite.Error(err)
      suite.Nil(resID)
	})

}

// func (suite *UserRepositorySuite) TestAddUser() {
//    user := &domain.User{
//        ID:       primitive.NewObjectID(),
//        Username: "testuser",
//        Password: "password123",
//    }

//    id, err := suite.repo.AddUser(user)
//    assert.Nil(suite.T(), err)
//    assert.NotNil(suite.T(), id)
// }

func TestUserRepositorySuite(t *testing.T) {
   suite.Run(t, new(UserRepositorySuite))
}
