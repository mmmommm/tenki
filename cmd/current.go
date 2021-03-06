package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	sub "github.com/mmmommm/tenki/sub"
)

type OpenWeatherMapAPIResponse struct {
	Main    Main      `json:"main"`
	Weather []Weather `json:"weather"`
	Wind    Wind      `json:"wind"`
	Dt      int64     `json:"dt"`
}
type Main struct {
	Temp     float64 `json:"temp"`
	TempMin  float64 `json:"temp_min"`
	TempMax  float64 `json:"temp_max"`
	Pressuer int     `json:"pressure"`
	Humidity int     `json:"humidity"`
}
type Weather struct {
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}
type Wind struct {
	Speed float64 `json:"speed"`
}
type Env struct {
	Token string `required:"true"`
}

func CurrentWeather(cPrefecture, env string) {
	var city string
	city = sub.Prefecture(cPrefecture)

	token := env                                                  // APIトークン
	endPoint := "https://api.openweathermap.org/data/2.5/weather" // APIのエンドポイント

	// パラメータを設定
	values := url.Values{}
	values.Set("q", city)
	values.Set("APPID", token)
	// リクエストを投げる
	res, err := http.Get(endPoint + "?" + values.Encode())
	if res != nil {
		defer res.Body.Close()
	}
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	// レスポンスを読み取り
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	// JSONパース
	var apiRes OpenWeatherMapAPIResponse
	if err := json.Unmarshal(bytes, &apiRes); err != nil {
		fmt.Println(err)
	}

	// aa := sub.ASCIIart(apiRes.Weather[0].Icon)
	emoji := sub.GetEmoji(apiRes.Weather[0].Icon)

	fmt.Printf("----------------------------------------\n")
	fmt.Printf("時刻:     %s\n", time.Unix(apiRes.Dt, 0))
	fmt.Printf("天気:     %s   %s\n", emoji, apiRes.Weather[0].Main)
	fmt.Printf("詳細:     %s\n", apiRes.Weather[0].Description)
	fmt.Printf("平均気温: %s °C\n", fmt.Sprintf("%.1f", sub.Change(apiRes.Main.Temp))) // ケルビンで取得される
	fmt.Printf("最高気温: %s °C\n", fmt.Sprintf("%.1f", sub.Change(apiRes.Main.TempMax)))
	fmt.Printf("最低気温: %s °C\n", fmt.Sprintf("%.1f", sub.Change(apiRes.Main.TempMin)))
	fmt.Printf("気圧:     %d hPa\n", apiRes.Main.Pressuer)
	fmt.Printf("湿度:     %d ％\n", apiRes.Main.Humidity)
	fmt.Printf("風速:     %s m/s\n", fmt.Sprintf("%.1f", apiRes.Wind.Speed))
	// fmt.Printf("%s\n", aa)
	fmt.Printf("----------------------------------------\n")
}
