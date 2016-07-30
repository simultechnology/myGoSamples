package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"regexp"
)

func main() {

	fp, err := os.Open("addresses.csv");
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	input := bufio.NewScanner(fp)

	var lineCount int

	for input.Scan() {
		line := input.Text()
		elems := strings.Split(line, ",")
		lineCount += 1
		//contains := strings.Contains(elems[0], elems[2])
		matched, _ := regexp.MatchString(fmt.Sprintf("^%s", elems[2]), elems[0])
		if matched {
			prefix := strings.TrimPrefix(elems[0], elems[2])
			prefix = strings.TrimPrefix(prefix, "大字")
			prefix = strings.TrimPrefix(prefix, "字")
			fmt.Println(prefix)
		} else {
			fmt.Println(elems[0])
		}
	}



}
