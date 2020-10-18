package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/joho/godotenv"
	sub "github.com/mmmommm/tenki/sub"
)

type OpenWeatherMapAPILists struct {
	List []Forecast `json:"list"`
}

type Forecast struct {
	Main    fMain      `json:"main"`
	Weather []fWeather `json:"weather"`
	Wind    fWind      `json:"wind"`
	Dt      int64      `json:"dt"`
	DtText  string     `json:"dt_txt"`
}

type fMain struct {
	Temp     float64 `json:"temp"`
	TempMin  float64 `json:"temp_min"`
	TempMax  float64 `json:"temp_max"`
	Pressuer int     `json:"pressure"`
	Humidity int     `json:"humidity"`
}
type fWeather struct {
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type fWind struct {
	Speed float64 `json:"speed"`
}

func ForecastWeather(fPrefecture string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var city string
	city = sub.Prefecture(fPrefecture)

	token := os.Getenv("WEATHER_TOKEN")                            // APIトークン
	endPoint := "https://api.openweathermap.org/data/2.5/forecast" // APIのエンドポイント

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
	var apiRes OpenWeatherMapAPILists
	if err := json.Unmarshal(bytes, &apiRes); err != nil {
		fmt.Println(err)
	}
	times := [...]int{2, 4, 6}
	for i := range times {
		aa := sub.ASCIIart(apiRes.List[i].Weather[0].Icon)
		fmt.Printf("----------------------------------------\n")
		fmt.Printf("時刻:     %s\n", time.Unix(apiRes.List[i].Dt, 0))
		fmt.Printf("天気:     %s\n", apiRes.List[i].Weather[0].Description)
		fmt.Printf("最高気温: %s °C\n", fmt.Sprintf("%.1f", sub.Change(apiRes.List[i].Main.TempMax)))
		fmt.Printf("最低気温: %s °C\n", fmt.Sprintf("%.1f", sub.Change(apiRes.List[i].Main.TempMin)))
		fmt.Printf("湿度:     %d ％\n", apiRes.List[i].Main.Humidity)
		fmt.Printf("風速:     %s m/s\n", fmt.Sprintf("%.1f", apiRes.List[i].Wind.Speed))
		fmt.Printf("%s\n", aa)
	}
}
