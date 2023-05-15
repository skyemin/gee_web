package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	url := "https://www.ziroom.com/z/nl/z2.html?qwd=&p=1"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36 Edge/16.16299")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 使用正则表达式匹配房源信息
	reg := regexp.MustCompile(`<div class="clearfix">.*?</div>`)
	matches := reg.FindAllStringSubmatch(string(body), -1)

	for _, match := range matches {
		// 提取房源信息
		regTitle := regexp.MustCompile(`<h3 class="i-tit">(.*?)</h3>`)
		title := regTitle.FindStringSubmatch(match[0])[1]
		regLocation := regexp.MustCompile(`<p class="i-address">(.*?)</p>`)
		location := regLocation.FindStringSubmatch(match[0])[1]
		regPrice := regexp.MustCompile(`<p class="i-price">(.*?)</p>`)
		price := regPrice.FindStringSubmatch(match[0])[1]
		fmt.Println(title, location, price)
	}
}
