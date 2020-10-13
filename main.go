package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type M map[string]string
type T struct {
	Packages []M
}

func main() {
	t := T{}
	dat, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	err = yaml.Unmarshal([]byte(dat), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", t)
	for i, val := range t.Packages {
		println(i)
		println(val["url"])
		println(val["version"])
	}
}
