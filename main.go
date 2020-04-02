package main

import (
	"bufio"
	"fmt"
	"github.com/satori/go.uuid"
	"os"
	"time"
)

func main() {
	for arg := range os.Args {
		fmt.Println(arg)
	}

	go func() {
		f, _ := os.OpenFile("ok.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
		defer f.Close()
		for true {
			f.WriteString(fmt.Sprintf("%s\n", uuid.NewV4()))
		}
	}()

	ticker := time.NewTicker(time.Second * 5)
	go func() {
		for _ = range ticker.C {
			fileInfo, _ := os.Stat("ok.txt")
			size := fileInfo.Size() // B *1024 =KB *1024 = MB *1024 = GB
			unit := []string{"B", "KB", "MB", "GB"}
			unit_p := 0
			for size > 1024 {
				size /= 1024
				unit_p++
				if unit_p == 3 {
					break
				}
			}

			fmt.Printf("file sieze %d %s\n", size, unit[unit_p])
		}
	}()

	reader := bufio.NewReader(os.Stdin)
	reader.ReadLine()
}
