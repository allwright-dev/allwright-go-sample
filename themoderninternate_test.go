package main

import (
	"context"
	"log"
	"testing"
	"time"

	allwright "allwright.dev"
)

func TestAdd(t *testing.T) {

	defer func() {
		if err := allwright.Shutdown(); err != nil {
			log.Printf("shutdown allwright Go client: %v", err)
		}
	}()

	ctx, cancel := context.WithTimeout(context.TODO(), 60*time.Second)
	defer cancel()
	if browser, err := allwright.LaunchChrome(ctx, allwright.LaunchOptions{}); err != nil {
		t.Fatalf("Error: %s", err.Error())
	} else {
		defer browser.Close(ctx)
		tab := browser.InitialTab()
		_, _ = tab.Navigate(ctx, "https://themoderninternet.vercel.app")
		_, _ = tab.Click(ctx, "//*[@data-slot='card' and .//*[text()='Form Inputs']]//button")
		toBeVisible := true
		if _, err := tab.WaitForSelector(ctx, "//h1[text()='Form Inputs']", allwright.WaitForSelectorOptions{
			Visible: &toBeVisible,
		}); err != nil {
			t.Fatalf("Error: %s", err.Error())
		}
	}

}
