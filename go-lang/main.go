package main

import (
	"fmt"
	"regexp"
)

func main() {
	// https://www.jianshu.com/p/7d507ccb5eec
	// 正则表达式
	reg, err := regexp.Compile("^([a-z]+)([0-9]+)$")

	if err != nil {
		fmt.Println(err)
	}

	content := "abc9009"

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

	// http.HandleFunc("/[a-z]*$", home.Index)
	// http.HandleFunc("/home/index", home.Index)
	// http.HandleFunc("/home/hello", home.Hello)
	// http.HandleFunc("/user/index", user.Index)

	// http.ListenAndServe("127.0.0.1:8001", nil)
}
