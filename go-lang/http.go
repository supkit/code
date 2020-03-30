package main

import (
	"code/go-lang/controllers"
	"code/go-lang/demo"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"time"
)

func main() {
	urls := []string{
		"http://www.qq.com",
		"http://www.baidu.com",
		"http://www.supbin.com",
		"http://www.sohu.com",
		// "http://www.yahoo.com",
		"http://www.tencent.com",
		"http://www.hao123.com",
		"http://www.weibo.com",
	}
	// 初始化一个管道
	message := make(chan string)
	start := time.Now()

	for _, url := range urls {
		go get(url, message)
	}

	for range urls {
		text := <-message
		fmt.Println(text)
	}

	// 总消耗的时间
	elapsed := time.Since(start).Seconds()
	fmt.Printf("%.2fs all over\n", elapsed)

	demo.Study()
	fmt.Println(controllers.Index())
}

func get(url string, message chan string) {
	start := time.Now()

	res, err := http.Get(url)

	if err != nil {
		message <- fmt.Sprint(err)
		return
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		os.Exit(1)
	}

	writeError := ioutil.WriteFile(getFileName(url), body, 0644)

	if writeError != nil {
		os.Exit(2)
	}

	// 消耗的时间
	elapsed := time.Since(start).Seconds()

	// 输出单个URL消耗的时间
	message <- fmt.Sprintf("%.2fs %s", elapsed, url)
	return
}

// 获取文件名
func getFileName(url string) string {
	var Re = regexp.MustCompile("\\w+\\.\\w+$")
	// 从URL中匹配域名部分
	return "storage/" + Re.FindString(url) + ".html"
}
