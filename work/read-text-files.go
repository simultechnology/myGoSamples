package main
import (
	"os"
	"log"
	"fmt"
	"io/ioutil"
)

func main() {

	targetFile := "./tmp/test.txt"

	file, err := os.Open(targetFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buffs := make([]byte, 35)

	_, err = file.Read(buffs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buffs))
	fmt.Println("---------------------------------")

	contents, err :=ioutil.ReadFile(targetFile)
	fmt.Println(string(contents));
}
