package main

import (
	"fmt"
	"github.com/smgolangdev/godev/goutils/utils"
)

func main() {
	fmt.Println("Running Utils Main.")
	_, err := utils.CopyFile("/home/smdeveloper/.Xdefaults", "/home/smdeveloper/smXdefaults.txt")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Generating ID: ", utils.GenerateId())
}
