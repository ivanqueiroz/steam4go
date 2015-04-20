package steam4go

import (
	"encoding/json"
	"net/url"
	"strconv"
)

//Stats is a Struct that represents the return of GetGlobalAchievementPercentagesForApp
type Stats struct {
	AchievementPercentages struct {
		Achievements []struct {
			Name    string  `json:"name"`
			Percent float64 `json:"percent"`
		} `json:"achievements"`
	} `json:"achievementpercentages"`
}

//GetGlobalAchievementPercentagesForApp returns all stats for app.
func GetGlobalAchievementPercentagesForApp(gameID float64) (data Stats) {
	u, _ := url.Parse(ifSteam)
	u.Path = "/ISteamUserStats/GetGlobalAchievementPercentagesForApp/v0002/"
	q := u.Query()
	q.Set("gameid", strconv.FormatFloat(gameID, 'f', 2, 64))
	u.RawQuery = q.Encode()
	var stats Stats
	jsonSrc := navigateToByte(u.String())
	json.Unmarshal(jsonSrc, &stats)
	return stats
}
