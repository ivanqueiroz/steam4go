package steam4go

import (
	"encoding/json"
	"net/url"
	"strconv"
)

const ifSteam string = "http://api.steampowered.com"

//Appnews type represent return of appnews
type Appnews struct {
	AppNews struct {
		AppID     float64 `json:"appid"`
		NewsItems []struct {
			Author        string  `json:"author"`
			Contents      string  `json:"contents"`
			Date          float64 `json:"date"`
			FeedLabel     string  `json:"feedlabel"`
			FeedName      string  `json:"feedname"`
			Gid           string  `json:"gid"`
			IsExternalURL bool    `json:"is_external_url"`
			Title         string  `json:"title"`
			URL           string  `json:"url"`
		} `json:"newsitems"`
	} `json:"appnews"`
}

//GetNewsForApp returns the latest of a game specified by its appID.
func GetNewsForApp(appID float64, count int, maxlength int) (data Appnews) {
	u, _ := url.Parse(ifSteam)
	u.Path = "/ISteamNews/GetNewsForApp/v0002/"
	q := u.Query()
	q.Set("appid", strconv.FormatFloat(appID, 'f', 2, 64))
	q.Set("count", strconv.Itoa(count))
	q.Set("maxlength", strconv.Itoa(maxlength))
	q.Set("format", "json")
	u.RawQuery = q.Encode()
	var news Appnews
	jsonSrc := navigateToByte(u.String())
	json.Unmarshal(jsonSrc, &news)
	return news
}
