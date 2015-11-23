package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func main()  {

	fmt.Println("start");

	yamlData, err := ioutil.ReadFile("./yaml/sample.yml")

	if err != nil {
		panic(err)
		return
	}

	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal(yamlData, &m)
	if err != nil {
		panic(err)
	}

	//fmt.Println(m)

	fmt.Printf("%s\n", m["machine"].(map[interface {}]interface {})["timezone"])
	fmt.Printf("%s\n", m["machine"].
	(map[interface {}]interface {})["hosts"].
	(map[interface {}]interface {})["circlehost"])
	//fmt.Printf("%d\n", m["b"].(map[interface {}]interface {})["c"])

}
