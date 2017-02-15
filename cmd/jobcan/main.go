package main

import (
	"flag"
	"log"
	"os"

	"github.com/sclevine/agouti"

	"github.com/gotokatsuya/punch/src/jobcan"
)

func main() {
	flag.Parse()

	driver := agouti.ChromeDriver()
	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start driver:%v", err)
	}
	page, err := driver.NewPage(agouti.Browser("chrome"))
	if err != nil {
		log.Fatalf("Failed to open page:%v", err)
	}
	status, err := jobcan.Do(page)
	if err != nil {
		log.Fatalf("Failed to do jobcan:%v", err)
	}
	if status == 0 {
		driver.Stop()
	}

	os.Exit(0)
}
