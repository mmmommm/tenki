package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main()  {
	//カレントディレクトリ取得
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	if 2 <= len(os.Args) {
		if os.Args[1] == "-a" {
			for _, fileInfo := range fileInfos {
				fmt.Println(fileInfo.Name())
			}
			return
		}
	}

	for _, fileInfo := range fileInfos {
		if strings.HasPrefix(fileInfo.Name(), ".") {
			continue
		}
		fmt.Println(fileInfo.Name())
	}
}
