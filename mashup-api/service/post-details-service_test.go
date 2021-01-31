package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var postDetailsService = NewPostDetailsService()

func TestGetDetails(t *testing.T) {
	postDetails := postDetailsService.GetDetails()

	assert.NotNil(t, postDetails)
	assert.Equal(t, 1, postDetails.ID)
}
