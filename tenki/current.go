package tenki

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"

	round "github.com/mmmommm/tenki/round"
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

func Current(prefecture string) {
	var city string
	switch prefecture {
		case "北海道":
			city = "Hokkaido,jp"
		case "青森":
			city = "Aomori,jp"
		case "岩手":
			city = "Iwate,jp"
		case "宮城":
			city = "Miyagi,jp"
		case "秋田":
			city = "Akita,jp"
		case "山形":
			city = "Yamagata,jp"
		case "福島":
			city = "Hukushima,jp"
		case "茨城":
			city = "Ibaragi,jp"
		case "栃木":
			city = "Tochigi,jp"
		case "群馬":
			city = "Gunma,jp"
		case "埼玉":
			city = "Saitama,jp"
		case "千葉":
			city = "Chiba,jp"
		case "東京":
			city = "Tokyo,jp"
		case "神奈川":
			city = "Kanagawa,jp"
		case "新潟":
			city = "Nigata,jp"
		case "富山":
			city = "Toyama,jp"
		case "石川":
			city = "Ishikawa,jp"
		case "福井":
			city = "Hukui,jp"
		case "山梨":
			city = "Yamanashi,jp"
		case "岐阜":
			city = "Gifu,jp"
		case "静岡":
			city = "Shizuoka,jp"
		case "愛知":
			city = "Aichi,jp"
		case "三重":
			city = "Mie,jp"
		case "滋賀":
			city = "Shiga,jp"
		case "京都":
			city = "Kyoto,jp"
		case "大阪":
			city = "Osaka,jp"
		case "兵庫":
			city = "Hyogo,jp"
		case "奈良":
			city = "Nara,jp"
		case "和歌山":
			city = "Wakayama,jp"
		case "鳥取":
			city = "Tottori,jp"
		case "島根":
			city = "Shimane,jp"
		case "岡山":
			city = "Okayama,jp"
		case "広島":
			city = "Hiroshima,jp"
		case "山口":
			city = "Yamaguchi,jp"
		case "徳島":
			city = "Tokushima,jp"
		case "香川":
			city = "Kagawa,jp"
		case "愛媛":
			city = "Ehime,jp"
		case "高知":
			city = "Kochi,jp"
		case "福岡":
			city = "Fukuoka,jp"
		case "佐賀":
			city = "Saga,jp"
		case "長崎":
			city = "Nagasaki,jp"
		case "熊本":
			city = "Kumamoto,jp"
		case "大分":
			city = "Oita,jp"
		case "宮崎":
			city = "Miyazaki,jp"
		case "鹿児島":
			city = "Kagoshima,jp"
		case "沖縄":
			city = "Okinawa,jp"
	}

	token := os.Getenv("WEATHER_TOKEN")                   // APIトークン
	endPoint := "https://api.openweathermap.org/data/2.5/weather" // APIのエンドポイント

	// パラメータを設定
	values := url.Values{}
	values.Set("q", city)
	values.Set("APPID", token)
	// リクエストを投げる
	res, err := http.Get(endPoint + "?" + values.Encode())
	if err != nil {
			panic(err)
	}
	defer res.Body.Close()

	// レスポンスを読み取り
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
			panic(err)
	}

	// JSONパース
	var apiRes OpenWeatherMapAPIResponse
	if err := json.Unmarshal(bytes, &apiRes); err != nil {
			panic(err)
	}

	fmt.Printf("時刻:     %s\n", time.Unix(apiRes.Dt, 0))
	fmt.Printf("天気:     %s\n", apiRes.Weather[0].Main)
	fmt.Printf("アイコン: https://openweathermap.org/img/wn/%s@2x.png\n", apiRes.Weather[0].Icon)
	fmt.Printf("説明:     %s\n", apiRes.Weather[0].Description)
	fmt.Printf("気温:     %s °C\n", fmt.Sprintf("%.1f", round.Change(apiRes.Main.Temp))) // ケルビンで取得される
	fmt.Printf("最高気温: %s °C\n", fmt.Sprintf("%.1f", round.Change(apiRes.Main.TempMax)))
	fmt.Printf("最低気温: %s °C\n", fmt.Sprintf("%.1f", round.Change(apiRes.Main.TempMin)))
	fmt.Printf("気圧:     %d hPa\n", apiRes.Main.Pressuer)
	fmt.Printf("湿度:     %d ％\n", apiRes.Main.Humidity)
	fmt.Printf("風速:     %s m/s\n", fmt.Sprintf("%.1f", apiRes.Wind.Speed))
}