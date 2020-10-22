package round

import (
	"fmt"
)

func GetEmoji(icon string) string {
	var text string
	if icon == "01d" || icon == "01n" {
		text = "☀️"
	} else if icon == "02d" || icon == "02n" {
		text = "🌤"
	} else if icon == "03d" || icon == "03n" {
		text = "⛅️"
	} else if icon == "04d" || icon == "04n" {
		text = "☁️"
	} else if icon == "09d" || icon == "09n" {
		text = "⛈"
	} else if icon == "10d" || icon == "10n" {
		text = "🌧"
	} else if icon == "11d" || icon == "11n" {
		text = "🌩"
	} else if icon == "13d" || icon == "13n" {
		text = "⛄"
	} else if icon == "50d" || icon == "50n" {
		text = "🌫"
	} else {
		fmt.Print("-----------------------")
	}
	return text
}
