package main

import (
	"encoding/json"

	"./conf"
	"./req"
)

var ()

func main() {
	conf := conf.Init()
	req := req.Init(conf)

	for no, file := range conf.Files {
		tpl := readJSONFile(file.Tpl)
		html := readFile(file.HTML)
		tpl["template"] = html

		d, err := json.Marshal(tpl)
		check(err)
		req.Put(no, d)

	}
}
