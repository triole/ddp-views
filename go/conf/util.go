package conf

import (
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func find(basedir string, rxFilter string, dirOrFile string) []string {
	filelist := []string{}
	// index files
	filelist = listFolderContentSimple(basedir)

	// gather file infos and filter by type
	newlist := []string{}
	for _, item := range filelist {
		fileinfo, err := os.Stat(item)
		if err != nil {
			log.Fatal(err)
		}
		if fileinfo.IsDir() == true && dirOrFile == "d" {
			newlist = append(newlist, item)
		} else if fileinfo.IsDir() == false && dirOrFile == "f" {
			newlist = append(newlist, item)
		}
	}
	sort.Strings(newlist)
	return newlist
}

func listFolderContentSimple(basedir string) []string {
	if strings.HasSuffix(basedir, "/") == false {
		basedir += "/"
	}
	filelist, err := filepath.Glob(basedir + "*")
	if err != nil {
		log.Fatal(err)
	}
	return filelist
}
