package main

import (
	"bytes"
	"crypto/tls"
	"github/h0ld1rs/xianzhj_crawl/htmlcontent"
	"github/h0ld1rs/xianzhj_crawl/page"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/JohannesKaufmann/html-to-markdown/plugin"
)

func main() {

	headers := map[string]string{
		"Host":                      "xz.aliyun.com",
		"Cookie":                    "_uab_collina=168967217002875816551932; t=55bb3741ce4883b96f0a4763aab03a3a; csrftoken=D2dbe0WRjnicSet0CBlfemTo3lWNddPQsiAlXh1xNU0hKK4v9xoaqAyV146qopnn; login_aliyunid_pk=1791792327537544; login_current_pk=1791792327537544; aliyun_site=CN; l=fBP4evyeNgMgv5zJBOfCFurza77TSIRvmuPzaNbMi9fPOr1RxSwOW1BZRr8vCnGVFsT6R3yI_B67BeYBcQd-nxvTZjFFs_DmnmOk-Wf..; tfstk=eBfDjDT-_3iQOWhCyidbwebzKXeRlmO6fGh9XCKaU3-7lhBA7C4GDZDvBNrfj1xP8FKvkRaisNS2DnKOsTGGRN1YlCUXGZO6_kEL9ObflCNH_wHU9q0xxzrLvWFRlZO6_kBva_bhcmH4M1m5GrAkx2vVm2a341tFun7Ak_kH3x_vq3-lNEbkfZyE_f5koKjyBYkU-9h6Ko6nCA92PUxLEjdlKLawn6aurvQ63UTBvz4oCqp2PUxLrzDKkK8WRHC..; isg=BJCQT14AwbzJDpwbXA0MhQovYdjiWXSjtmpiJYpgpeuXxTNvNm3jMSS2mY0lFSx7; acw_tc=781bad2517060761273695438e6d47e59fb6ffdaae94124dc143fa97d7dc09; ssxmod_itna=Yq+OGKBKAKY55BchDooDQDkUEQxQqxDtO=ao5aDlrjDxA5D8D6DQeGTiu5+HiqHKjB+AefK/gy+=jtCB+3hQoQSGfNdWgMT3D=xYQDwxYoDUxGtDpxG63/Den=D5xGoDPxDeDAlKDCZHd=D7xXBfdXUbMjXFqNDm36RDGjnDitdxiv4VdilS+DlDDf3xB=4xAfRD7U/bdQDbqDuz==o2qDLjWcHo6QDbrPgYpWDtLTu4qDH9uNDTuysj4xYlODrGYdiKEhv/YwWAbmKb0SdBKMqK0qNtv4sCNW4U0Z3eD===; ssxmod_itna2=Yq+OGKBKAKY55BchDooDQDkUEQxQqxDtO=ao5D6p3UD0Hox035qR=qH5UoR9wU8BTKhQUYChYrw4vxjHb4M/Ydf2OD82O4HoTIMAQ6dMFpBtk7V1alXKUbVwk96paTItrUoXHQl7GBuEN/TtMfdLsext2A4AjCm6/IL6u3nvd/uE1beuqoSplYBP/CTTZ/nNxr1KLb1NujYabwZ6a56jAB8Q1TcA2De9zZ+Y4c8DkjgGN/BpIt057/Oe+FmhfHGAyRDS+SIudIxgUvkDhiie6q86IQfDO/A5VDFfnTAgw+8xA5Y/PRCDOqd+0a2TQ=Y1B8xWrd3GQ+D8fxez8h+hDb2b=iiQuTOGhOizXieE2wA4m+EKmOxEoeG4KiRoXODZB7j3q+rzCrGk+rD+=rOr14YrnrCgI7ok/o+opdQQbH4H14Qjj75o1=nTg0+t+isAWZ846rzd8DHpD=oW5P6UidOirrgjXDdRkKM3htS=yoEVbCA+7R+757rlgErQ1Go+NqKFjtrWzOYPD7jQxGcDG7HiDD==",
		"Pragma":                    "no-cache",
		"Cache-Control":             "no-cache",
		"Sec-Ch-Ua":                 "\"Not_A Brand\";v=\"8\", \"Chromium\";v=\"120\", \"Google Chrome\";v=\"120\"",
		"Sec-Ch-Ua-Mobile":          "?0",
		"Sec-Ch-Ua-Platform":        "\"macOS\"",
		"Upgrade-Insecure-Requests": "1",
		"User-Agent":                "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
		"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
		"Sec-Fetch-Site":            "same-origin",
		"Sec-Fetch-Mode":            "navigate",
		"Sec-Fetch-User":            "?1",
		"Sec-Fetch-Dest":            "document",
		"Referer":                   "https://xz.aliyun.com/t/13099",
		"Accept-Language":           "zh-CN,zh;q=0.9",
		"Connection":                "close",
	}
	var data = []byte(nil)

	response := httpRequest(page.GetPage("13099"), "GET", data, headers)

	converter := md.NewConverter("", true, nil)

	converter.Use(plugin.GitHubFlavored())
	markdown, err := converter.ConvertString(htmlcontent.GetHTMLBetweenMarkers(response))
	mdContent := strings.TrimSpace(markdown)
	if err != nil {
		log.Fatal(err)
	}

	//测试输出文件的标题
	// fmt.Println(htmlcontent.GetTitleFromHTML(response))
	// 将云图片替换为本地图片
	LocalContent := strings.Replace(mdContent, "https://xzfile.aliyuncs.com/media/upload/picture/", "./assert/", -1)
	//指定父目录
	docParentDir := "doc"

	//在doc文件夹下创建文件夹
	docFolder := htmlcontent.GetTitleFromHTML(response)

	fullPath := filepath.Join(docParentDir, docFolder)
	err = os.MkdirAll(fullPath, os.ModePerm)
	if err != nil {
		panic(err)
	}
	// 创建 Markdown 文件
	mdFileName := filepath.Join(fullPath, docFolder+".md")
	mdFile, err := os.Create(mdFileName)
	if err != nil {
		panic(err)
	}
	defer mdFile.Close()

	// 将内容写入 Markdown 文件
	_, err = io.WriteString(mdFile, LocalContent)
	if err != nil {
		panic(err)
	}
	// fmt.Println(mdContent)

	// 获取图片链接
	links := htmlcontent.GetSpecificLinksFromHTML(mdContent)

	// fmt.Println(links)
	htmlcontent.DownloadLinks(links, fullPath+"/assert")

}

func httpRequest(targetUrl string, method string, data []byte, headers map[string]string) string {

	request, error := http.NewRequest(method, targetUrl, bytes.NewBuffer(data))
	for k, v := range headers {
		request.Header.Set(k, v)

	}

	customTransport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: customTransport}
	response, error := client.Do(request)
	defer response.Body.Close()

	if error != nil {
		panic(error)
	}

	body, _ := ioutil.ReadAll(response.Body)
	// fmt.Println("response Status:", response.Status)
	// fmt.Println("response Body:", string(body))
	return string(body)
}
