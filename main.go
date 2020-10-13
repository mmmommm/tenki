package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mmmommm/tenki/tenki"
)

var (
	Version  = "unset"
	Revision = "unset"
)


func main() {
	const defaultPrefecture = ""
	var	cPrefecture, nPrefecture, fPrefecture string
	flag.StringVar(&cPrefecture, "current", defaultPrefecture, "--curennt <県名> returns current weather information where you are.")
	flag.StringVar(&cPrefecture, "c", defaultPrefecture, "-c <県名> returns current weather information where you are.")
	flag.StringVar(&fPrefecture, "forecast", defaultPrefecture, "--forecast <県名> returns current weather information where you are.")
	flag.StringVar(&fPrefecture, "f", defaultPrefecture, "-f <県名> returns current weather information where you are.")
	flag.StringVar(&nPrefecture, "next", defaultPrefecture, "--next <県名> returns tomorrow's weather forecast where you are.")
	flag.StringVar(&nPrefecture, "n", defaultPrefecture, "-n <県名> returns tomorrow's weather forecast where you are.")
	flag.Parse()
	// -cオプションの時の処理を書く
	if cPrefecture != "" && nPrefecture == "" && nPrefecture == "" {
		tenki.CurrentWeather(cPrefecture)
	// -fオプションの時の処理を書く
	} else if cPrefecture == "" && fPrefecture != "" && nPrefecture == "" {
		tenki.ForecastWeather(fPrefecture)
	// -nオプションの時の処理を書く
	} else if cPrefecture == "" && fPrefecture == "" && nPrefecture != "" {
		tenki.NextWeather(nPrefecture)
	} else {
		fmt.Printf("%s: option arguments is empty !! please type prefecture name. \n", os.Args[0])
		fmt.Printf("%s: try 'tenki --help' or 'tenki -h' for more information\n", os.Args[0])
		os.Exit(1)
	}
}
