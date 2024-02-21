package main

import (
	"fmt"
	"github.com/playwright-community/playwright-go"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"time"
)

var (
	HostUrl = "https://ais.usvisa-info.com/en-ca/niv/users/sign_in"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(os.Stdout)
}

func main() {
	email := os.Getenv("EMAIL")
	password := os.Getenv("PASSWORD")
	fmt.Printf("email: %s, password: %s\n", email, password)

	if email == "" || password == "" {
		log.Fatal("EMAIL or PASSWORD is empty")
		return
	}

	pw, _ := playwright.Run()
	browser, _ := pw.Chromium.Launch(
		playwright.BrowserTypeLaunchOptions{
			Headless: playwright.Bool(false),
		})
	context, _ := browser.NewContext()
	page, _ := context.NewPage()

	logrus.Debug("ready to login")

	// 打开登录页面
	_, err := page.Goto(HostUrl,
		playwright.PageGotoOptions{
			//WaitUntil: playwright.WaitUntilStateNetworkidle,
		})
	if err != nil {
		log.Fatal(err)
	}

	//_ = page.WaitForLoadState(
	//	playwright.PageWaitForLoadStateOptions{
	//		State: playwright.LoadStateLoad,
	//	})

	_ = page.Locator("#header").Click()

	//entries, err := page.Locator("form[id='sign_in_form']").All()
	//if err != nil {
	//	log.Fatalf("could not get entries: %v", err)
	//}
	//for _, entry := range entries {
	//	fmt.Printf("entry: %v", entry)
	//	//
	//	//// submit
	//	//_ = page.Locator("input[type=submit]").Click()
	//}

	//_ = page.Locator("//a[@class=\"down-arrow bounce\"]").Click()
	_ = page.Locator("input[id='user_email']").Fill(email)
	_ = page.Locator("input[id='user_password']").Fill(password)
	err = page.Locator(".icheckbox").Click() // todo x: select class name

	if err != nil {
		log.Fatal(err)
	}

	//submit
	err = page.Locator("input[type=submit]").Click()
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(10 * time.Second)

	// 等待登录完成
	//page.WaitForNavigation()

	if err = browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}
	if err = pw.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}
	// 在这里可以继续进行其他操作，如验证登录是否成功等
}
