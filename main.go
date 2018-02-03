package main

import (
	"fmt"

	"github.com/liufuqiang/location"
)

func main() {
	fmt.Println(location.GetLocation("China", "Beijing", "Chaoyang"))
}
