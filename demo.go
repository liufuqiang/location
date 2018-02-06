package mypackage
import (
	"fmt"

	"github.com/liufuqiang/location"
)

func myFunc() {
	fmt.Println(location.GetLocation("China", "Beijing", "Chaoyang"))
}
