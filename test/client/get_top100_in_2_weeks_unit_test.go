package client_test

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/opv88/steamspy/client"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetTop100In2Weeks_ShouldBeAsExpected(t *testing.T) {
	// Arrange
	httpmock.Activate(t)
	defer httpmock.DeactivateAndReset()

	uri := "https://steamspy.com/api.php?request=top100in2weeks"
	httpmock.RegisterResponder("GET", uri,
		httpmock.NewStringResponder(200, `{"9450": {
		"appid": 9450,
		"name": "Warhammer 40,000: Dawn of War - Soulstorm",
		"developer": "Relic Entertainment",
		"publisher": "SEGA",
		"score_rank": 123,
		"positive": 16764,
		"negative": 847,
		"userscore": 5,
		"owners": "1,000,000 .. 2,000,000",
		"average_forever": 2571,
		"average_2weeks": 880,
		"median_forever": 810,
		"median_2weeks": 1262,
		"price": "100",
		"initialprice": "200",
		"discount": "50",
		"ccu": 673
	}}`))

	steamSpyClient := client.NewClient()

	expected := []client.SteamItem{
		client.SteamItem{
			AppId:           9450,
			Name:            "Warhammer 40,000: Dawn of War - Soulstorm",
			Developer:       "Relic Entertainment",
			Publisher:       "SEGA",
			ScoreRank:       123,
			Positive:        16764,
			Negative:        847,
			UserScore:       5,
			Owners:          "1,000,000 .. 2,000,000",
			AverageForever:  2571,
			AverageTwoWeeks: 880,
			MedianForever:   810,
			MedianTwoWeeks:  1262,
			Price:           "100",
			InitialPrice:    "200",
			Discount:        "50",
			CCU:             673,
		},
	}

	// Act
	actual, err := steamSpyClient.GetTop100In2Weeks()
	httpmock.GetTotalCallCount()
	info := httpmock.GetCallCountInfo()
	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
	assert.Equal(t, 1, info[fmt.Sprintf("GET %s", uri)])
}

func Test_GetTop100In2WeeksWithEmpty_ShouldReturnEmpty(t *testing.T) {
	// Arrange
	httpmock.Activate(t)
	defer httpmock.DeactivateAndReset()

	uri := "https://steamspy.com/api.php?request=top100in2weeks"
	httpmock.RegisterResponder("GET", uri,
		httpmock.NewStringResponder(200, `{}`))

	steamSpyClient := client.NewClient()

	// Act
	actual, err := steamSpyClient.GetTop100In2Weeks()
	httpmock.GetTotalCallCount()
	info := httpmock.GetCallCountInfo()

	// Assert
	assert.NoError(t, err)
	assert.Empty(t, actual)
	assert.Equal(t, 1, info[fmt.Sprintf("GET %s", uri)])
}

func Test_GetTop100In2WeeksWithError_ShouldReturnError(t *testing.T) {
	// Arrange
	httpmock.Activate(t)
	defer httpmock.DeactivateAndReset()

	someError := fmt.Errorf("some error")
	uri := "https://steamspy.com/api.php?request=top100in2weeks"
	httpmock.RegisterResponder("GET", uri,
		httpmock.NewErrorResponder(someError))

	steamSpyClient := client.NewClient()
	expectedErrorMessage := fmt.Sprintf("Error getting top 100 in 2 weeks: Get \"%s\": %v.", uri, someError.Error())

	// Act
	actual, err := steamSpyClient.GetTop100In2Weeks()
	httpmock.GetTotalCallCount()
	info := httpmock.GetCallCountInfo()

	// Assert
	assert.EqualError(t, err, expectedErrorMessage)
	assert.Equal(t, 1, info[fmt.Sprintf("GET %s", uri)])
	assert.Nil(t, actual)
}

func Test_GetTop100In2WeeksWithInvalidStatusCode_ShouldReturnError(t *testing.T) {
	// Arrange
	httpmock.Activate(t)
	defer httpmock.DeactivateAndReset()

	uri := "https://steamspy.com/api.php?request=top100in2weeks"
	httpmock.RegisterResponder("GET", uri,
		httpmock.NewStringResponder(400, ``))

	steamSpyClient := client.NewClient()
	expectedErrorMessage := "Error getting top 100 in 2 weeks, incorrect response status code: 400 Bad Request."

	// Act
	actual, err := steamSpyClient.GetTop100In2Weeks()
	httpmock.GetTotalCallCount()
	info := httpmock.GetCallCountInfo()

	// Assert
	assert.EqualError(t, err, expectedErrorMessage)
	assert.Equal(t, 1, info[fmt.Sprintf("GET %s", uri)])
	assert.Nil(t, actual)
}

func Test_GetTop100In2WeeksWithNonMarshalledBody_ShouldReturnError(t *testing.T) {
	// Arrange
	httpmock.Activate(t)
	defer httpmock.DeactivateAndReset()

	uri := "https://steamspy.com/api.php?request=top100in2weeks"
	httpmock.RegisterResponder("GET", uri,
		httpmock.NewStringResponder(200, `invalid`))

	steamSpyClient := client.NewClient()
	expectedErrorMessage := "Error reading body of top 100 in 2 weeks: invalid character 'i' looking for beginning of value."

	// Act
	actual, err := steamSpyClient.GetTop100In2Weeks()
	httpmock.GetTotalCallCount()
	info := httpmock.GetCallCountInfo()

	// Assert
	assert.EqualError(t, err, expectedErrorMessage)
	assert.Equal(t, 1, info[fmt.Sprintf("GET %s", uri)])
	assert.Nil(t, actual)
}
