package steam4go

import (
	"net/url"
	"strconv"
)

const ifSteam string = "http://api.steampowered.com/ISteamNews"

//GetNewsForApp returns the latest of a game specified by its appID.
func GetNewsForApp(appID int, count int, maxlength int, format string) (data string) {
	u, _ := url.Parse(ifSteam)
	u.Path = "/ISteamNews/GetNewsForApp/v0002/"
	q := u.Query()
	q.Set("appid", strconv.Itoa(appID))
	q.Set("count", strconv.Itoa(count))
	q.Set("maxlength", strconv.Itoa(maxlength))
	q.Set("format", format)
	u.RawQuery = q.Encode()
	return Navigate(u.String())

}
