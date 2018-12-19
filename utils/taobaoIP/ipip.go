package taobaoIP

import (
	"self-wechat/utils/logging"
	"self-wechat/utils/snapHttp"
	"strings"
)

var (
	logger = logging.GetLogger()
)

func GetCountryAndCity(ip string) (country, city string, err error) {
	taoHost := "http://ip.taobao.com/service/getIpInfo.php?ip="
	tl := taoHost + ip
	sh := new(snapHttp.SnapHttp)

	var resp autoGenerated

	err = sh.GetJson(tl, &resp)
	if err != nil {
		logger.Error(err)
		return
	}
	country = strings.TrimSpace(resp.Data.Country)
	city = strings.TrimSpace(resp.Data.City)
	return
}

type autoGenerated struct {
	Code int `json:"code"`
	Data struct {
		IP        string `json:"ip"`
		Country   string `json:"country"`
		Area      string `json:"area"`
		Region    string `json:"region"`
		City      string `json:"city"`
		County    string `json:"county"`
		Isp       string `json:"isp"`
		CountryID string `json:"country_id"`
		AreaID    string `json:"area_id"`
		RegionID  string `json:"region_id"`
		CityID    string `json:"city_id"`
		CountyID  string `json:"county_id"`
		IspID     string `json:"isp_id"`
	} `json:"data"`
}
