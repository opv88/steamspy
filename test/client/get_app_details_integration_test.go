package client_test

import (
	"testing"

	"github.com/opv88/steamspy/client"

	"github.com/stretchr/testify/assert"
)

func Test_GetAppDetails_GetValues_Success(t *testing.T) {
	// Arrange
	steamSpyClient := client.NewClient()
	csAppId := int64(730)

	expected := client.SteamItemDetailed{
		SteamItem: client.SteamItem{
			AppId:        csAppId,
			Name:         "Counter-Strike: Global Offensive",
			Developer:    "Valve",
			Publisher:    "Valve",
			ScoreRank:    client.ScoreRank(int64(0)),
			Owners:       "100,000,000 .. 200,000,000",
			Price:        "0",
			InitialPrice: "0",
			Discount:     "0",
		},
		Underscore: 0,
		Languages:  "English, Czech, Danish, Dutch, Finnish, French, German, Hungarian, Italian, Japanese, Korean, Norwegian, Polish, Portuguese - Portugal, Portuguese - Brazil, Romanian, Russian, Simplified Chinese, Spanish - Spain, Swedish, Thai, Traditional Chinese, Turkish, Bulgarian, Ukrainian, Greek, Spanish - Latin America, Vietnamese, Indonesian",
		Genre:      "Action, Free To Play",
		Tags: map[string]int64{
			"FPS":          91172,
			"Shooter":      65634,
			"Multiplayer":  62536,
			"Competitive":  53536,
			"Action":       47634,
			"Team-Based":   46549,
			"e-sports":     43682,
			"Tactical":     41468,
			"First-Person": 39540,
			"PvP":          34587,
			"Online Co-Op": 34056,
			"Co-op":        30342,
			"Strategy":     30189,
			"Military":     28762,
			"War":          28060,
			"Difficult":    26037,
			"Trading":      25813,
			"Realistic":    25497,
			"Fast-Paced":   25369,
			"Moddable":     6667,
		},
	}

	// Act
	actual, err := steamSpyClient.GetAppDetails(csAppId)

	// Assert
	assert.NoError(t, err, "GetAppDetails shouldn't return an error.")
	assert.NotEmpty(t, actual, "GetAppDetails response should not be empty.")
	assert.Equal(t, expected.AppId, actual.AppId, "AppId should be as expected.")
	assert.Equal(t, expected.Name, actual.Name, "Name should be as expected.")
	assert.Equal(t, expected.Developer, actual.Developer, "Developer should be as expected.")
	assert.Equal(t, expected.Publisher, actual.Publisher, "Publisher should be as expected.")
	assert.Equal(t, expected.ScoreRank, actual.ScoreRank, "ScoreRank should be as expected.")
	assert.Equal(t, expected.Owners, actual.Owners, "Owners should be as expected.")
	assert.Equal(t, expected.Price, actual.Price, "Price should be as expected.")
	assert.Equal(t, expected.InitialPrice, actual.InitialPrice, "InitialPrice should be as expected.")
	assert.Equal(t, expected.Discount, actual.Discount, "Discount should be as expected.")
	assert.Equal(t, expected.Underscore, actual.Underscore, "Underscore should be as expected.")
	assert.Equal(t, expected.Languages, actual.Languages, "Languages should be as expected.")
	assert.Equal(t, expected.Genre, actual.Genre, "Genre should be as expected.")
}
