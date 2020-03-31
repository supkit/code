package lib

import (
	"code/go-lang/controllers/home"
	"code/go-lang/controllers/user"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

// Go并发请求
func Fetch() {
	urls := []string{
		"http://www.baidu.com",
		"http://www.qq.com",
		"http://www.weibo.com",
		"http://www.hao123.com",
		"http://www.163.com",
		"http://www.126.com",
		"http://www.mi.com",
		"http://www.youku.com",
		"http://v.qq.com",
		"http://m.qq.com",
	}

	start := time.Now()

	// 初始化一个管道
	message := make(chan string)

	for _, value := range urls {
		go get(value, message)
	}

	for range urls {
		text := <-message
		fmt.Println(text)
	}

	elapsed := time.Since(start).Seconds()
	fmt.Printf("%.2fs全部执行完 \n", elapsed)
}

// 发起HTTP GET请求
func get(url string, message chan string) bool {
	start := time.Now()
	res, _ := http.Get(url)
	ioutil.ReadAll(res.Body)
	elapsed := time.Since(start).Seconds()
	message <- fmt.Sprintf("%.2fs %s", elapsed, url)
	return true
}

// Go语言中的正则
func Regx() {
	// 正则表达式
	content := "https://www.jianshu.com/p/7d507ccb5eec"
	reg, err := regexp.Compile("https?://[a-z]{1,9}.[a-z]*.[a-z]{0,9}[/a-z0-9]*")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(reg.MatchString(content))
	fmt.Println(reg.FindString(content))
	fmt.Println(reg.FindStringSubmatchIndex(content))
	result := reg.FindStringSubmatch(content)

	for index, value := range result {
		fmt.Printf("[%d]:%v\n", index, value)
	}

	for i := 0; i < len(result); i++ {
		fmt.Printf("[%d]:%v\n", i, result[i])
	}
}

// HTTP服务
func HttpServer() {
	http.HandleFunc("/", home.Index)
	http.HandleFunc("/user/index", user.Index)

	http.ListenAndServe("127.0.0.1:8001", nil)
}
