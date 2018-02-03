## 基于映射配置的地域英文转中文服务(配置根据微信里的wx.getUserInfo返回的location信息)[golang]

## Install
	go get -u -v github.com/liufuqiang/location



## Examples
参考examples
[main.go][1]

	
	package main

	import (
		"fmt"

		"github.com/liufuqiang/location"
	)

	func main() {
		fmt.Println(location.GetLocation("China", "Beijing", "Chaoyang"))
	}
