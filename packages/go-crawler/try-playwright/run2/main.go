package main

import (
	"fmt"
	"github.com/playwright-community/playwright-go"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"net/http"
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

func getCaScheduleCheckAPI(scheduleID string, locationID string) string {
	return fmt.Sprintf("https://ais.usvisa-info.com/en-ca/niv/schedule/%s/appointment/days/%s.json?appointments[expedite]=false", scheduleID, locationID)
}

func main() {
	email := os.Getenv("EMAIL")
	password := os.Getenv("PASSWORD")

	scheduleID := os.Getenv("SCHEDULE_ID")
	locationID := os.Getenv("LOCATION_ID") // [95 = Vancouver, ]
	fmt.Printf("email: %s, password: %s\n", email, password)
	fmt.Printf("scheduleID: %s, locationID: %s\n", scheduleID, locationID)

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

	// get cookies
	cookies, err := page.Context().Cookies()
	if err != nil {
		log.Fatal(err)
	}

	// for range cookies
	for _, cookie := range cookies {
		fmt.Printf("cookie: %v\n", cookie)
	}

	//
	// http request
	//
	// call check API
	checkUrl := getCaScheduleCheckAPI(scheduleID, locationID)
	//page2, _ := context.NewPage()

	// 创建一个新的请求
	// 创建一个新的 HTTP 客户端
	client := &http.Client{}
	req, err := http.NewRequest("GET", checkUrl, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, cookie := range cookies {
		req.AddCookie(&http.Cookie{
			Name:  cookie.Name,
			Value: cookie.Value,
		})
	}

	//resp
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return
	}

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 打印响应体
	fmt.Println(string(body))

	// 读取响应体
	// 解析 JSON 响应

	time.Sleep(150 * time.Second)

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
