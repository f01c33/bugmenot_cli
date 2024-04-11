package main

import (
	"fmt"
	"log"
	"os"

	"github.com/playwright-community/playwright-go"
)

func main() {
	err := playwright.Install(&playwright.RunOptions{Browsers: []string{"firefox"}})
	if err != nil {
		log.Fatal("can't install")
	}
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	browser, err := pw.Firefox.Launch()
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	if _, err = page.Goto("https://bugmenot.com/view/" + os.Args[1]); err != nil {
		log.Fatalf("could not goto: %v", err)
	}
	entries, err := page.Locator("article.account").All()
	if err != nil {
		log.Fatalf("could not get entries: %v", err)
	}
	fmt.Println("Index\tUsername\tPassword\tSuccess rate")
	for i, entry := range entries {
		accs, err := entry.Locator("kbd").All()
		if err != nil {
			log.Fatalf("could not get entries: %v", err)
		}
		success, err := entry.Locator(".success_rate").All()
		if err != nil {
			log.Fatalf("could not get text content: %v", err)
		}
		txt0, err := accs[0].TextContent()
		if err != nil {
			log.Fatalf("could not get entries: %v", err)
		}
		txt1, err := accs[1].TextContent()
		if err != nil {
			log.Fatalf("could not get entries: %v", err)
		}
		succs, err := success[0].TextContent()
		if err != nil {
			log.Fatalf("could not get entries: %v", err)
		}
		fmt.Printf("%v\t%v\t%v\t%v\n", i+1, txt0, txt1, succs)
	}
	if err = browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}
	if err = pw.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}
}
