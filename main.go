package main

import (
	"fmt"
	"os"
)

func main() {

	// 3. Задание

	for i := 0; i < 1000000; i++ {
		newFile, err := os.Create(fmt.Sprintf("%v.txt", i))
		if err != nil {
			panic(err)
		}
		defer newFile.Close()
	}

}
