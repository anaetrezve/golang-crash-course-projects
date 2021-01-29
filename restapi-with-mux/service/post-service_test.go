package service

import (
	"testing"

	"../entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) Save(post *entity.Post) (*entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)

	return result.(*entity.Post), args.Error(1)
}

func (mock *MockRepository) FindAll() ([]entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)

	return result.([]entity.Post), args.Error(1)
}

func TestValidateEmptyPost(t *testing.T) {
	testService := NewPostService(nil)

	err := testService.Validate(nil)

	assert.NotNil(t, err)

	assert.Equal(t, "The post is empty", err.Error())
}

func TestValidatePostTitle(t *testing.T) {
	testService := NewPostService(nil)
	post := entity.Post{1, "", "Text"}

	err := testService.Validate(&post)

	assert.NotNil(t, err)
	assert.Equal(t, "The post title is empty", err.Error())

}

func TestFindAll(t *testing.T) {
	mockRepo := new(MockRepository)
	post := entity.Post{ID: 1, Title: "A", Text: "B"}

	// Setup Expections
	mockRepo.On("FindAll").Return([]entity.Post{post}, nil)

	testService := NewPostService(mockRepo)
	result, _ := testService.FindAll()

	// Mock Assertion: Behavioral
	mockRepo.AssertExpectations(t)

	// Data Assertion
	assert.Equal(t, int64(1), result[0].ID)
	assert.Equal(t, "A", result[0].Title)
	assert.Equal(t, "B", result[0].Text)
}

func TestCreatePost(t *testing.T) {
	mockRepo := new(MockRepository)
	post := entity.Post{Title: "A", Text: "B"}

	// Setup Expections
	mockRepo.On("Save").Return(&post, nil)

	testService := NewPostService(mockRepo)
	result, err := testService.Create(&post)

	// Mock Assertion: Behavioral
	mockRepo.AssertExpectations(t)

	// Data Assertion
	assert.NotNil(t, result.ID)
	assert.Equal(t, "A", result.Title)
	assert.Equal(t, "B", result.Text)
	assert.Nil(t, err)
}
