package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line:=scanner.Text()
		//fmt.Println(scanner.Text())

		if strings.Contains(line, "cannot find package") {

			begin :=strings.Index(line, "\"")+1
			end   :=strings.LastIndex(line, "\"")
			fmt.Println("go get " + line[begin:end])
		}
	}

	if scanner.Err() != nil {
		// handle error.
	}

}