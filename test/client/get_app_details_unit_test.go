package client_test

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/opv88/steamspy/client"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetAppDetails_ShouldBeAsExpected(t *testing.T) {
	// Arrange
	httpmock.Activate(t)
	defer httpmock.DeactivateAndReset()

	appId := int64(9450)
	uri := fmt.Sprintf("https://steamspy.com/api.php?request=appdetails&appId=%d", appId)
	httpmock.RegisterResponder("GET", uri,
		httpmock.NewStringResponder(200, `{
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
			"ccu": 673,
			"languages": "English, French, German",
			"genre": "Strategy",
			"tags": {
				"Warhammer 40K": 321,
				"Strategy": 296,
				"RTS": 290,
				"Base-Building": 188,
				"Sci-fi": 146,
				"Multiplayer": 131,
				"Real-Time": 118,
				"Games Workshop": 96,
				"Singleplayer": 85,
				"Futuristic": 68,
				"Classic": 65,
				"Action": 54,
				"Atmospheric": 46,
				"Great Soundtrack": 38,
				"Gore": 36,
				"Moddable": 35,
				"Adventure": 21,
				"Fantasy": 19,
				"Co-op": 17,
				"Memes": 17
			}
		}`))

	steamSpyClient := client.NewClient()

	expected := client.SteamItemDetailed{
		SteamItem: client.SteamItem{
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
		Languages: "English, French, German",
		Genre:     "Strategy",
		Tags: map[string]int64{
			"Warhammer 40K":    321,
			"Strategy":         296,
			"RTS":              290,
			"Base-Building":    188,
			"Sci-fi":           146,
			"Multiplayer":      131,
			"Real-Time":        118,
			"Games Workshop":   96,
			"Singleplayer":     85,
			"Futuristic":       68,
			"Classic":          65,
			"Action":           54,
			"Atmospheric":      46,
			"Great Soundtrack": 38,
			"Gore":             36,
			"Moddable":         35,
			"Adventure":        21,
			"Fantasy":          19,
			"Co-op":            17,
			"Memes":            17,
		},
	}

	// Act
	actual, err := steamSpyClient.GetAppDetails(appId)
	httpmock.GetTotalCallCount()
	info := httpmock.GetCallCountInfo()

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
	assert.Equal(t, 1, info[fmt.Sprintf("GET %s", uri)])
}

func Test_GetAppDetailsTagsIsNull_ShouldBeAsExpected(t *testing.T) {
	// Arrange
	httpmock.Activate(t)
	defer httpmock.DeactivateAndReset()

	appId := int64(9450)
	uri := fmt.Sprintf("https://steamspy.com/api.php?request=appdetails&appId=%d", appId)
	httpmock.RegisterResponder("GET", uri,
		httpmock.NewStringResponder(200, `{
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
			"ccu": 673,
			"languages": "English, French, German",
			"genre": "Strategy",
			"tags": null
		}`))

	steamSpyClient := client.NewClient()

	expected := client.SteamItemDetailed{
		SteamItem: client.SteamItem{
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
		Languages: "English, French, German",
		Genre:     "Strategy",
		Tags:      nil,
	}

	// Act
	actual, err := steamSpyClient.GetAppDetails(appId)
	httpmock.GetTotalCallCount()
	info := httpmock.GetCallCountInfo()

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
	assert.Equal(t, 1, info[fmt.Sprintf("GET %s", uri)])
}

func Test_GetAppDetailsTagsIsInvalidFormat_ShouldReturnError(t *testing.T) {
	// Arrange
	httpmock.Activate(t)
	defer httpmock.DeactivateAndReset()

	appId := int64(9450)
	uri := fmt.Sprintf("https://steamspy.com/api.php?request=appdetails&appId=%d", appId)
	httpmock.RegisterResponder("GET", uri,
		httpmock.NewStringResponder(200, `{
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
			"ccu": 673,
			"languages": "English, French, German",
			"genre": "Strategy",
			"tags": "invalid data"
		}`))

	steamSpyClient := client.NewClient()
	expectedErrorMessage := "Error reading body of app details: json: cannot unmarshal string into Go struct field SteamItemDetailed.tags of type map[string]int64."

	// Act
	actual, err := steamSpyClient.GetAppDetails(appId)
	httpmock.GetTotalCallCount()
	info := httpmock.GetCallCountInfo()

	// Assert
	assert.Error(t, err)
	assert.EqualError(t, err, expectedErrorMessage)
	assert.NotNil(t, actual)
	assert.Equal(t, 1, info[fmt.Sprintf("GET %s", uri)])
}

func Test_GetAppDetailsWithEmpty_ShouldReturnEmpty(t *testing.T) {
	// Arrange
	httpmock.Activate(t)
	defer httpmock.DeactivateAndReset()

	appId := int64(9450)
	uri := fmt.Sprintf("https://steamspy.com/api.php?request=appdetails&appId=%d", appId)
	httpmock.RegisterResponder("GET", uri,
		httpmock.NewStringResponder(200, `{}`))

	steamSpyClient := client.NewClient()

	// Act
	actual, err := steamSpyClient.GetAppDetails(appId)
	httpmock.GetTotalCallCount()
	info := httpmock.GetCallCountInfo()

	// Assert
	assert.NoError(t, err)
	assert.Empty(t, actual)
	assert.Equal(t, 1, info[fmt.Sprintf("GET %s", uri)])
}

func Test_GetAppDetailsWithZeroAppId_ShouldReturnError(t *testing.T) {
	// Arrange
	appId := int64(0)
	steamSpyClient := client.NewClient()
	expectedErrorMessage := client.ValidationError_AppIdIsRequired

	// Act
	actual, err := steamSpyClient.GetAppDetails(appId)

	// Assert
	assert.EqualError(t, err, expectedErrorMessage)
	assert.Empty(t, actual)
}

func Test_GetAppDetailsWithError_ShouldReturnError(t *testing.T) {
	// Arrange
	httpmock.Activate(t)
	defer httpmock.DeactivateAndReset()

	appId := int64(9450)
	someError := fmt.Errorf("some error")
	uri := fmt.Sprintf("https://steamspy.com/api.php?request=appdetails&appId=%d", appId)
	httpmock.RegisterResponder("GET", uri,
		httpmock.NewErrorResponder(someError))

	steamSpyClient := client.NewClient()
	expectedErrorMessage := fmt.Sprintf("Error getting app details: Get \"%s\": %v.", uri, someError.Error())

	// Act
	actual, err := steamSpyClient.GetAppDetails(appId)
	httpmock.GetTotalCallCount()
	info := httpmock.GetCallCountInfo()

	// Assert
	assert.EqualError(t, err, expectedErrorMessage)
	assert.Equal(t, 1, info[fmt.Sprintf("GET %s", uri)])
	assert.Empty(t, actual)
}

func Test_GetAppDetailsWithInvalidStatusCode_ShouldReturnError(t *testing.T) {
	// Arrange
	httpmock.Activate(t)
	defer httpmock.DeactivateAndReset()

	appId := int64(9450)
	uri := fmt.Sprintf("https://steamspy.com/api.php?request=appdetails&appId=%d", appId)
	httpmock.RegisterResponder("GET", uri,
		httpmock.NewStringResponder(400, ``))

	steamSpyClient := client.NewClient()
	expectedErrorMessage := "Error getting app details, incorrect response status code: 400 Bad Request."

	// Act
	actual, err := steamSpyClient.GetAppDetails(appId)
	httpmock.GetTotalCallCount()
	info := httpmock.GetCallCountInfo()

	// Assert
	assert.EqualError(t, err, expectedErrorMessage)
	assert.Equal(t, 1, info[fmt.Sprintf("GET %s", uri)])
	assert.Empty(t, actual)
}

func Test_GetAppDetailsWithNonMarshalledBody_ShouldReturnError(t *testing.T) {
	// Arrange
	httpmock.Activate(t)
	defer httpmock.DeactivateAndReset()

	appId := int64(9450)
	uri := fmt.Sprintf("https://steamspy.com/api.php?request=appdetails&appId=%d", appId)
	httpmock.RegisterResponder("GET", uri,
		httpmock.NewStringResponder(200, `invalid`))

	steamSpyClient := client.NewClient()
	expectedErrorMessage := "Error reading body of app details: invalid character 'i' looking for beginning of value."

	// Act
	actual, err := steamSpyClient.GetAppDetails(appId)
	httpmock.GetTotalCallCount()
	info := httpmock.GetCallCountInfo()

	// Assert
	assert.EqualError(t, err, expectedErrorMessage)
	assert.Equal(t, 1, info[fmt.Sprintf("GET %s", uri)])
	assert.Empty(t, actual)
}
