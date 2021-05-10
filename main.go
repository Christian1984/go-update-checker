package main

import (
	"christian1984-updatechecker/updatechecker"
	"fmt"
)

func main() {
	res := updatechecker.IsUpdateAvailable("Christian1984", "vfrmap-for-vr", "1.1.0", 3, true)
	fmt.Println(res)
}
