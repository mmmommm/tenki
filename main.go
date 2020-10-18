package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mmmommm/tenki/cmd"
)

var version = "1.0.0"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	env := os.Getenv("WEATHER_TOKEN")

	const defaultPrefecture = ""
	var cPrefecture, nPrefecture, fPrefecture string
	var showVersion bool
	flag.StringVar(&cPrefecture, "current", defaultPrefecture, "--curennt <県名> returns current weather information where you are.")
	flag.StringVar(&cPrefecture, "c", defaultPrefecture, "-c <県名> returns current weather information where you are.")
	flag.StringVar(&fPrefecture, "forecast", defaultPrefecture, "--forecast <県名> returns current weather information where you are.")
	flag.StringVar(&fPrefecture, "f", defaultPrefecture, "-f <県名> returns current weather information where you are.")
	flag.StringVar(&nPrefecture, "next", defaultPrefecture, "--next <県名> returns tomorrow's weather forecast where you are.")
	flag.StringVar(&nPrefecture, "n", defaultPrefecture, "-n <県名> returns tomorrow's weather forecast where you are.")
	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.BoolVar(&showVersion, "version", false, "show version")
	flag.Parse()

	// -cオプションの時の処理を書く
	if cPrefecture != "" && fPrefecture == "" && nPrefecture == "" {
		cmd.CurrentWeather(cPrefecture, env)
		// -fオプションの時の処理を書く
	} else if cPrefecture == "" && fPrefecture != "" && nPrefecture == "" {
		cmd.ForecastWeather(fPrefecture, env)
		// -nオプションの時の処理を書く
	} else if cPrefecture == "" && fPrefecture == "" && nPrefecture != "" {
		cmd.NextWeather(nPrefecture, env)
	} else if showVersion {
		fmt.Println("version:", version)
		return
	} else {
		fmt.Printf("%s: option arguments is empty !! please type prefecture name. \n", os.Args[0])
		fmt.Printf("%s: try 'tenki --help' or 'tenki -h' for more information\n", os.Args[0])
		os.Exit(1)
	}
}
