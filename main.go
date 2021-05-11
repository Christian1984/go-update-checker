package main

import (
	"christian1984-updatechecker/updatechecker"
)

func main() {
	//uc := updatechecker.New("Christian1984", "vfrmap-for-vr", "FSKneeboard", "", 3, false, false)
	uc := updatechecker.New("Christian1984", "vfrmap-for-vr", "FSKneeboard", "https://fskneeboard.com/download-latest", 3, false, false)
	uc.CheckForUpdate("v1.0.3")

	/*
		//a := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		a := "|/-\\"
		bt := "\b"

		for j := 0; j < 4; j++ {
			for i := 0; i < len(a); i++ {
				fmt.Print(bt + string(a[i]))
				time.Sleep(500 * time.Millisecond)
			}
		}
	*/
}
