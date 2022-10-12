package file_operation

import (
	"fmt"
	"os"
)

func MyOpenFile(filename string) {
	f, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Something wrong!")
	} else {
		fmt.Println(f)
	}
}
