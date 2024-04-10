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
			Args:     []string{"--start-maximized"}, // fullscreen
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

	// 定位到目标预约, 点击 Continue 按钮, 进行日期改签
	// 定位到 Continue button, click it
	//<a class="button primary small" href="/en-ca/niv/schedule/530xxxxx/continue_actions">Continue</a>
	btnContinueUrl := fmt.Sprintf("/en-ca/niv/schedule/%s/continue_actions", scheduleID)
	btnContinue := fmt.Sprintf("//a[@class='button primary small'] [@href='%s']", btnContinueUrl)
	log.Printf("btnContinue: %s\n", btnContinue)
	_ = page.Locator(btnContinue).Click()

	// 下一页, 改签, 跳转到城市+日期选择页面
	// <h5>
	//            <span class="fas fa-calendar-minus fa-lg fa-fw"></span>
	//            Reschedule Appointment
	//          </h5>
	// <a class="accordion-title" aria-controls="2z6g8s-accordion" role="tab" id="2z6g8s-accordion-label" aria-expanded="true" aria-selected="true">
	//          <h5>
	//            <span class="fas fa-calendar-minus fa-lg fa-fw"></span>
	//            Reschedule Appointment
	//          </h5>
	//        </a>

	// <h5>
	//            <span class="fas fa-calendar-minus fa-lg fa-fw"></span>
	//            Reschedule Appointment
	//          </h5>

	//_ = page.Locator(":has-text('Reschedule Appointment')").Click()
	//_ = page.GetByText("Reschedule Appointment").Click()
	//_ = page.Locator("//a[contains(., 'Reschedule Appointment')]").Click()
	//_ = page.Locator("//section[@id='forms").GetByRole("listitem").Filter(
	//	playwright.LocatorFilterOptions{HasText: "Reschedule Appointment"}).Click()

	// <a class="accordion-title" aria-controls="7hbgqy-accordion" role="tab" id="7hbgqy-accordion-label" aria-expanded="true" aria-selected="true">
	//          <h5>
	//            <span class="fas fa-calendar-minus fa-lg fa-fw"></span>
	//            Reschedule Appointment
	//          </h5>
	//        </a>

	// 展开待点击的按钮
	// <span class="fas fa-calendar-minus fa-lg fa-fw"></span>
	secRes := page.Locator("//span[@class='fas fa-calendar-minus fa-lg fa-fw']")
	_ = secRes.Click()

	p, _ := secRes.Page()
	log.Printf("p: %v\n", p)

	////////////////////////////////////////////////////////////////////////////////////

	// 点击 Reschedule Appointment
	// <a class="button small primary small-only-expanded" href="/en-ca/niv/schedule/53081715/appointment">Reschedule Appointment</a>
	btnRescheduleUrl := fmt.Sprintf("/en-ca/niv/schedule/%s/appointment", scheduleID)
	btnReschedule := fmt.Sprintf("//a[@class='button small primary small-only-expanded'] [@href='%s']", btnRescheduleUrl)
	log.Printf("btnReschedule: %s\n", btnReschedule)
	_ = page.Locator(btnReschedule, playwright.PageLocatorOptions{HasText: "Reschedule Appointment"}).Click()

	////////////////////////////////////////////////////////////////////////////////////

	// 地区+时间选项查看
	// <select name="appointments[consulate_appointment][facility_id]" id="appointments_consulate_appointment_facility_id" class="required"><option value="" label=" "></option>
	//<option data-collects-biometrics="false" value="89">Calgary</option>
	//<option data-collects-biometrics="false" value="90">Halifax</option>
	//<option data-collects-biometrics="false" value="91">Montreal</option>
	//<option data-collects-biometrics="false" value="92">Ottawa</option>
	//<option data-collects-biometrics="false" value="93">Quebec City</option>
	//<option data-collects-biometrics="false" value="94">Toronto</option>
	//<option data-collects-biometrics="false" selected="selected" value="95">Vancouver</option></select>

	locationIDs := []string{"89", "90", "91", "92", "93", "94", "95"}

	for _, id := range locationIDs {

		elSelect := page.Locator("//select[@id='appointments_consulate_appointment_facility_id']")
		_, _ = elSelect.SelectOption(playwright.SelectOptionValues{Values: &[]string{id}})
		_ = elSelect.Click()
		time.Sleep(3 * time.Second)

		// fill select option
		log.Printf("location id: %s, waiting 3 seconds, then click next option\n", id)
	}

	////////////////////////////////////////////////////////////////////////////////////

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
