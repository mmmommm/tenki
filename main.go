package main

import (
	"flag"
	"github.com/mmmommm/tenki/tenki"
)

func main() {
	const defaultPrefecture = ""
	var	prefecture string
	flag.StringVar(&prefecture, "current", defaultPrefecture, "today's weather where you are")
	flag.StringVar(&prefecture, "c", defaultPrefecture, "today's weather where you are")
	flag.StringVar(&prefecture, "next", defaultPrefecture, "tommorow's weather where you are")
	flag.StringVar(&prefecture, "n", defaultPrefecture, "tomorrow's weather where you are")
	flag.Parse()
	tenki.Current(prefecture)
}
