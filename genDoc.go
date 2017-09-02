package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	//open file
	if fileList, er := os.Open("files.txt"); er == nil {
		defer fileList.Close()
		f, _ := os.Create(os.Args[1])
		scannerF := bufio.NewScanner(fileList)
		for scannerF.Scan() {
			addFile := scannerF.Text()
			fmt.Println(os.Args)
			if file, err := os.Open(addFile); err == nil {
				defer file.Close()
				scanner := bufio.NewScanner(file)
				for scanner.Scan() {
					l := scanner.Text()
					if len(l) > 2 {
						if string(l[0]) == "/" && string(l[1]) == "/" {
							var str string
							for i := 2; i < len(l); i++ {
								str = str + string(l[i])
							}
							str = str + "\r"
							io.WriteString(f, str)
						}
					}
				}

				// check for errors
				if err = scanner.Err(); err != nil {
					log.Fatal(err)
				}

			} else {
				log.Fatal(err)
			}
		}
	}

}
