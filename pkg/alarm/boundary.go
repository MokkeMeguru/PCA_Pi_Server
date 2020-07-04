package alarm

import (
	"fmt"
	"os"
)

func ConfigWrite(confStrList []string, path string) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("file open exception: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()
	for _, confStr := range confStrList {
		_, _ = fmt.Fprintln(file, confStr)
		fmt.Println(confStr)
	}
}
