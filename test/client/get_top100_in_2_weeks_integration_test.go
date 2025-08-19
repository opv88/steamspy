package client_test

import (
	"testing"

	"github.com/opv88/steamspy/client"
	"github.com/stretchr/testify/assert"
)

func Test_GetTop100In2Weeks_GetValues_Success(t *testing.T) {
	// Arrange
	steamSpyClient := client.NewClient()
	expectedCnt := 100

	// Act
	actual, err := steamSpyClient.GetTop100In2Weeks()

	// Assert
	assert.NoError(t, err, "GetTop100In2Weeks shouldn't return an error.")
	assert.NotEmpty(t, actual, "GetTop100In2Weeks response should not be empty.")
	assert.Len(t, actual, expectedCnt, "GetTop100In2Weeks response should contain of 100 items.")
}
