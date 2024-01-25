package htmlcontent

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

func httpRequest(targetUrl string, method string, data []byte, headers map[string]string) *http.Response {

	request, error := http.NewRequest(method, targetUrl, bytes.NewBuffer(data))
	for k, v := range headers {
		request.Header.Set(k, v)
	}

	customTransport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: customTransport}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}

	return response
}
func GetSpecificLinksFromHTML(htmlContent string) []string {
	// 正则表达式匹配链接
	linkPattern := "https://xzfile.aliyuncs.com/media/upload" + `[^"'\s\)]*`
	re := regexp.MustCompile(linkPattern)
	matches := re.FindAllString(htmlContent, -1)
	return matches
}

func DownloadLinks(links []string, path string) error {
	for _, link := range links {
		err := DownloadLink(link, path)
		if err != nil {
			return err
		}
	}
	return nil
}

func DownloadLink(link, path string) error {

	headers := map[string]string{
		"Host":               "xzfile.aliyuncs.com",
		"Sec-Ch-Ua":          "\"Not_A Brand\";v=\"8\", \"Chromium\";v=\"120\", \"Google Chrome\";v=\"120\"",
		"Sec-Ch-Ua-Mobile":   "?0",
		"User-Agent":         "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
		"Sec-Ch-Ua-Platform": "\"macOS\"",
		"Accept":             "image/avif,image/webp,image/apng,image/svg+xml,image/*,*/*;q=0.8",
		"Sec-Fetch-Site":     "same-origin",
		"Sec-Fetch-Mode":     "no-cors",
		"Sec-Fetch-Dest":     "image",
		"Accept-Language":    "zh-CN,zh;q=0.9",
		"Referer":            "https://xz.aliyun.com/",
	}
	var data = []byte(nil)

	// response, err := http.Get(link)
	response := httpRequest(link, "GET", data, headers)
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download: %s", response.Status)
	}

	// 创建 doc 文件夹
	docFolder := path
	err := os.MkdirAll(docFolder, os.ModePerm)
	if err != nil {
		return err
	}

	// 从链接中提取文件名
	_, fileName := filepath.Split(link)

	// 创建文件在 doc 文件夹中
	filePath := filepath.Join(docFolder, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 将响应体写入文件
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	fmt.Printf("Downloaded: %s\n", filePath)
	return nil
}
