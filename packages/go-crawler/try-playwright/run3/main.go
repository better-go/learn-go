package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gocolly/colly"
)

var (
	LoginAPI  = "https://ais.usvisa-info.com/en-ca/niv/users/sign_in"
	OKPageUrl = "https://ais.usvisa-info.com/en-ca/niv/account/settings/update_email"
)

func getCAScheduleCheckAPI(scheduleID string, locationID string) string {
	return fmt.Sprintf("https://ais.usvisa-info.com/en-ca/niv/schedule/%s/appointment/days/%s.json?appointments[expedite]=false", scheduleID, locationID)

}

func main() {

	email := os.Getenv("EMAIL")
	password := os.Getenv("PASSWORD")
	scheduleID := os.Getenv("SCHEDULE_ID")
	locationID := os.Getenv("LOCATION_ID") // [95 = Vancouver, ]

	log.Printf("email: %s, password: %s\n", email, password)
	log.Printf("scheduleID: %s, locationID: %s\n", scheduleID, locationID)

	// 创建一个新的 Collector
	c := colly.NewCollector()

	// 设置 user-agent
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36"

	// 设置超时时间
	c.SetRequestTimeout(30 * time.Second)

	// 登录表单提交的 URL
	// 模拟登录
	err := c.Post(LoginAPI, map[string]string{
		"user[email]":      email,
		"user[password]":   password,
		"policy_confirmed": "1",
		"commit":           "Sign in",
	})
	if err != nil {
		log.Fatal(err)
	}
	// 访问需要登录才能访问的页面
	checkUrl := getCAScheduleCheckAPI(scheduleID, locationID)
	err = c.Visit(checkUrl)
	log.Printf("Visit: %v, err: %v\n", checkUrl, err)

	// 获取 cookie
	cookies := c.Cookies(OKPageUrl)
	// 打印 cookie
	for _, cookie := range cookies {
		fmt.Println(cookie.Name, cookie.Value)
	}

}
