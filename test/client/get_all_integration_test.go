package client_test

import (
	"testing"

	"github.com/opv88/steamspy/client"
	"github.com/stretchr/testify/assert"
)

func Test_GetAll_GetValues_Success(t *testing.T) {
	// Arrange
	steamSpyClient := client.NewClient()
	page := int64(0)
	expectedLen := 1000

	// Act
	actual, err := steamSpyClient.GetAll(page)

	// Assert
	assert.NoError(t, err, "GetAll shouldn't return an error.")
	assert.NotEmpty(t, actual, "GetAll response should not be empty.")
	assert.Len(t, actual, expectedLen, "GetAll response should contain 1000 items.")
}
