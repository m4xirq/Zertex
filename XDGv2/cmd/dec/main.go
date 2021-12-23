package main

import (
	"fmt"
	"github.com/m4xirq/Zertex/XDGv2/utils"
	"io/ioutil"
)

func main() {
	bytes, _ := ioutil.ReadFile(".xdg")
	fmt.Println(utils.Decrypt(string(bytes), "xNz#'%/2n4SZsB>m"))
}

/*func main() {
	ioutil.WriteFile(".xdg", utils.Encrypt(""))
}*/
