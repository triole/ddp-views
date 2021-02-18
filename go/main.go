package main

import (
	"encoding/json"

	"ddpviews/conf"
	"ddpviews/req"
)

var ()

func main() {
	conf := conf.Init()
	req := req.Init(conf)

	for no, file := range conf.Files {
		tpl := readJSONFile(file.Tpl)
		html := readFile(file.HTML)
		tpl["template"] = html

		// save to file to be able to use them from shell
		writeJSONToFile(file.Tmp, tpl)

		d, err := json.Marshal(tpl)
		check(err)
		req.Put(no, d)

	}
}
