package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mmmommm/tenki/tenki"
)

func main() {
	const defaultPrefecture = ""
	var	cPrefecture, nPrefecture string
	flag.StringVar(&cPrefecture, "current", defaultPrefecture, "--curennt <県名> returns current weather information where you are.")
	flag.StringVar(&cPrefecture, "c", defaultPrefecture, "-c <県名> returns weather information where you are.")
	flag.StringVar(&nPrefecture, "next", defaultPrefecture, "--next <県名> returns tomorrow's weather forecast where you are.")
	flag.StringVar(&nPrefecture, "n", defaultPrefecture, "-n <県名> returns tomorrow's weather forecast where you are.")
	flag.Parse()
	// -cオプションの時の処理を書く
	if cPrefecture != "" && nPrefecture == "" {
		tenki.Current(cPrefecture)
	// -nオプションの時の処理を書く
	} else if cPrefecture == "" && nPrefecture != "" {
		fmt.Print("明日の予定を表示します")
	} else {
		fmt.Printf("%s: requestType is not correct!\n", os.Args[0])
		fmt.Printf("%s: try 'tenki --help' or 'tenki -h' for more information\n", os.Args[0])
		os.Exit(1)
	}
}
