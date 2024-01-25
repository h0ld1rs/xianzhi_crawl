package page

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func GetPage(id string) string {
	// 启动 Chrome 浏览器
	opts := []selenium.ServiceOption{}
	service, err := selenium.NewChromeDriverService("./chromedriver", 9515, opts...)
	if err != nil {
		fmt.Println("Failed to start the ChromeDriver server:", err)
		panic(err)
	}
	defer service.Stop()

	// 创建 Chrome 浏览器实例
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}
	chromeOpts := chrome.Capabilities{
		Args: []string{
			"--headless", // 无头模式
		},
	}
	caps.AddChrome(chromeOpts)
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 9515))
	if err != nil {
		fmt.Println("Failed to connect to the WebDriver:", err)
		panic(err)
	}
	defer wd.Quit()
	request, err := http.NewRequest("GET", "https://xz.aliyun.com/t/"+id, nil)
	if err != nil {
		panic(err)
	}
	request.Header.Add("Cookie", "_uab_collina=168967217002875816551932; t=55bb3741ce4883b96f0a4763aab03a3a; csrftoken=D2dbe0WRjnicSet0CBlfemTo3lWNddPQsiAlXh1xNU0hKK4v9xoaqAyV146qopnn; login_aliyunid_pk=1791792327537544; login_current_pk=1791792327537544; aliyun_site=CN; l=fBP4evyeNgMgv5zJBOfCFurza77TSIRvmuPzaNbMi9fPOr1RxSwOW1BZRr8vCnGVFsT6R3yI_B67BeYBcQd-nxvTZjFFs_DmnmOk-Wf..; tfstk=eBfDjDT-_3iQOWhCyidbwebzKXeRlmO6fGh9XCKaU3-7lhBA7C4GDZDvBNrfj1xP8FKvkRaisNS2DnKOsTGGRN1YlCUXGZO6_kEL9ObflCNH_wHU9q0xxzrLvWFRlZO6_kBva_bhcmH4M1m5GrAkx2vVm2a341tFun7Ak_kH3x_vq3-lNEbkfZyE_f5koKjyBYkU-9h6Ko6nCA92PUxLEjdlKLawn6aurvQ63UTBvz4oCqp2PUxLrzDKkK8WRHC..; isg=BJCQT14AwbzJDpwbXA0MhQovYdjiWXSjtmpiJYpgpeuXxTNvNm3jMSS2mY0lFSx7; acw_tc=781bad2317060945338983443e3cdbc7a4ae97e95140a3293c7c3a4928f95d; acw_sc__v3=65b0f0686df015ef12f869d98310646beeee4dd3; ssxmod_itna=YuitqAOD7KAKDtDXiG7ma0=dG=L44mPb9aaeKDsvbDSxGKidDqxBmWCB+euu4WIWtFD7Kztlj4x+PEzmjbPrt1m8Hp1D84i7DKqibDCqD1D3qDk79xYA8Dt4DTD34DYDiO=DBRsUQDFAL6BfdLRKkLuAqDE3dQDYLQDmTbDno0ZSwUjfGtDG3bDzqiDf3QDIdXRS4GnD0jadurSD03o=1aiqqGWb2=ixPGuRWjUD0pMnwXRnp5eC4xo9OGoff4+m4m+BDSAGAGQI/qAu0x+3GwbFYqQ+IZPSeLd9hYxD; ssxmod_itna2=YuitqAOD7KAKDtDXiG7ma0=dG=L44mPb9aaeikAaqCiqDlZtDj4WK1P4vtCPZ48tDQqQCquhD9GsiRaDK4egizRBn3CR22vK8qCf2HY8Y43/etugRBu2rjCOEp2T7PfFYOn0BeMVj9DApstOd7eeEhnbQ4P+29CB2CG0vGBA35mhKH7qKog0Ngjoier4+bojI4ePrFAEG45notnddev85NPjIetW0k4b0qnqx5mf6hezYm73fEemmq5480lm88B6fm2hy8U1YoojN77tk1U/vhvKoxpWQu1ezoNwLhCfWpNXn3nqBqPyy8w0mP+7nUQnyB4XWuUKRN0RGlcOneh4YpimVfbp+F9RmEtvvYAYxRh4bi7DQe9yWRR/bPWO9/FQCml8uN44u3vQZWP8yMjvSloKZ+GnbPrijRrtf04fBDiRWf7ARrfK4tzvepm4rW1e7HcuHTmz9IxqbCimqmWOPWoAraFBKF8wSWflG5VpqeKgbmSlqWBPn4wbYen+ALeKiuYzLD5MKHNgkH4z3ozfmqjPbGbTqMHZETzBhnxANjNHkxO2WQGAryDlTkRDmteOfDfkYlIMDdOFixDKkx0DDLxD2D/4loq9CTuGDD==")
	request.Header.Add("User-Agent", getRandomUserAgent())
	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	request.Header.Add("Referer", "https://xz.aliyun.com")

	_ = wd.Get("https://xz.aliyun.com/t/" + id)

	// 获取页面加载后的当前URL
	currentURL, _ := wd.CurrentURL()
	return currentURL
}

func getRandomUserAgent() string {
	userAgents := []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.0) AppleWebKit/5341 (KHTML, like Gecko) Chrome/38.0.815.0 Mobile Safari/5341",
		"Mozilla/5.0 (compatible; MSIE 6.0; Windows NT 5.1; Trident/3.1)",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/5360 (KHTML, like Gecko) Chrome/36.0.836.0 Mobile Safari/5360",
		"Mozilla/5.0 (Windows; U; Windows 95) AppleWebKit/532.40.5 (KHTML, like Gecko) Version/4.0 Safari/532.40.5",
		// Add more user agents as needed
	}

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(userAgents))

	return userAgents[randomIndex]
}
