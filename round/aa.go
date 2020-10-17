package round
import (
	"io/ioutil"
	"os"
	"fmt"
)

func readfile (url string) string {
	file, err := os.Open(url)
	if err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
	lines := string(b)
	return lines
}

func ASCIIart(icon string) string {
	var text string
	if icon == "01d" || icon == "01n" {
		text = readfile("./aa/01.txt")
	} else if icon == "02d" || icon == "02n" {
		text = readfile("./aa/02.txt")
	} else if icon == "03d" || icon == "03n" {
		text = readfile("./aa/03.txt")
	} else if icon == "04d" || icon == "04n" {
		text = readfile("./aa/04.txt")
	} else if icon == "09d" || icon == "09n" {
		text = readfile("./aa/09.txt")
	} else if icon == "10d" || icon == "10n" {
		text = readfile("./aa/10.txt")
	} else if icon == "11d" || icon == "11n" {
		text = readfile("./aa/11.txt")
	} else if icon == "13d" || icon == "13n" {
		text = readfile("./aa/13.txt")
	} else if icon == "50d" || icon == "50n" {
		text = readfile("./aa/50.txt")
	} else {
		fmt.Print("-----------------------")
	}
	return text
}
