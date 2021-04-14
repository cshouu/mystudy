/*
	go mod init module_name 创建go.mod文件
	go run . 在main包中运行main()函数
	go mod tidy 在import导入三方包之后，自动下载该包
*/

package main

import (
	"fmt"

	"rsc.io/quote"
)

func main1() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
}
