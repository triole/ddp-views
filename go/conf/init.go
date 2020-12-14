package conf

import (
	"encoding/json"
	"fmt"
	"olibs/rx"
	"olibs/syslib"
	"os"
	"path"
	"strings"
)

const (
	token   = "58114e78e6c3488247148a3e5c9e6fa462a9e4c9"
	baseURL = "http://localhost:8280/api/v1/views/views/"
)

// Conf is the exported configuration
type Conf struct {
	Dirs    tDirs
	Token   string
	BaseURL string
	Files   tFiles
}

type tFiles map[string]tFile

type tFile struct {
	HTML string
	Tpl  string
	Tmp  string
}

type tDirs struct {
	Base string
	HTML string
	Tpl  string
	Tmp  string
}

// Init is where it all starts
func Init() (conf Conf) {
	b, _ := os.Getwd()
	conf = Conf{
		Dirs: tDirs{
			Base: b,
			HTML: path.Join(b, "content/html"),
			Tpl:  path.Join(b, "content/tpl"),
			Tmp:  path.Join(b, "content/tmp"),
		},
		Token:   token,
		BaseURL: baseURL,
	}
	conf.Files = conf.detectFiles()
	return
}

func (c Conf) detectFiles() (ff tFiles) {
	ff = tFiles{}
	for _, tpl := range syslib.Find(c.Dirs.Tpl, ".*", "f", true) {
		no := strings.Split(rx.Find("[0-9]+.json$", tpl), ".")[0]
		html := strings.Replace(
			strings.Replace(tpl, ".json", ".html", -1),
			"/tpl/", "/html/", -1,
		)
		tmp := strings.Replace(tpl, "/content/tpl/", "/tmp/", -1)
		f := tFile{
			HTML: html,
			Tpl:  tpl,
			Tmp:  tmp,
		}
		ff[no] = f
	}
	return
}

// Print prints config to terminal
func (c Conf) Print() {
	pprint(c)
}

func pprint(i interface{}) {
	s, _ := json.MarshalIndent(i, "", "\t")
	fmt.Println(string(s))
}
