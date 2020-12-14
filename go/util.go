package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func pprint(i interface{}) {
	s, _ := json.MarshalIndent(i, "", "\t")
	fmt.Println(string(s))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(filename string) (t string) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	t = string(content)
	return
}

func readJSONFile(filename string) map[string]interface{} {
	content := readFile(filename)
	mmap := make(map[string]interface{})
	err := json.Unmarshal([]byte(content), &mmap)
	if err != nil {
		panic(err)
	}
	return mmap
}

func writeJSONToFile(filename string, data map[string]interface{}) {
	rankingsJSON, _ := json.Marshal(data)
	err := ioutil.WriteFile(filename, rankingsJSON, 0644)
	check(err)
}
