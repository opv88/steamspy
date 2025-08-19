package client_test

import (
	"testing"

	"github.com/opv88/steamspy/client"
	"github.com/stretchr/testify/assert"
)

func Test_GetByTag_GetValues_Success(t *testing.T) {
	// Arrange
	steamSpyClient := client.NewClient()
	tag := "Early+Access"

	// Act
	actual, err := steamSpyClient.GetByTag(tag)

	// Assert
	assert.NoError(t, err, "GetByTag shouldn't return an error.")
	assert.NotEmpty(t, actual, "GetByTag response should not be empty.")
	assert.Greater(t, len(actual), 1, "GetByTag response should contain more than 1 item.")
}
