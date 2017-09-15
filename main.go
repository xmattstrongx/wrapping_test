package wrapping_test

import (
	"fmt"

	"github.com/nesv/go-dynect/dynect"
	"github.com/xmattstrongx/wrapping_test/dyn"
)

func main() {
	fmt.Println("Hello World")

	myDyn := dyn.NewMyDyn(&dynect.ConvenientClient{})
	myDyn.CreateDNS()

}
