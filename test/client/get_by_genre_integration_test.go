package client_test

import (
	"testing"

	"github.com/opv88/steamspy/client"
	"github.com/stretchr/testify/assert"
)

func Test_GetByGenre_GetValues_Success(t *testing.T) {
	// Arrange
	steamSpyClient := client.NewClient()
	genre := "Action"

	// Act
	actual, err := steamSpyClient.GetByGenre(genre)

	// Assert
	assert.NoError(t, err, "GetByGenre shouldn't return an error.")
	assert.NotEmpty(t, actual, "GetByGenre response should not be empty.")
	assert.Greater(t, len(actual), 1, "GetByGenre response should contain more than 1 item.")
}
