// dirlistener.go
package main

import (
	"os"
	"path/filepath"
	"strings"
)

func dirlist(adress string) []string {
	adress, ext, wCardBool := wldChk(adress) //wildcard cheker (have || not)
	rawlist := []string{}                    //for all files
	outlist := []string{}                    //for files after wildcard filter
	fileinf := []os.FileInfo{}               //for FileIsDir method
	err := filepath.Walk(adress, func(oneFile string, info os.FileInfo, err error) error {
		rawlist = append(rawlist, oneFile)
		fileinf = append(fileinf, info)
		return nil
	})
	if err != nil {
		loger("dirlist() adress filepath.Walk", err)
	}
	if wCardBool == true { //wildCard filter

		for _, file := range rawlist {
			if filepath.Ext(file) == ext {
				outlist = append(outlist, file)
			}
		}
		return outlist

	} else {
		return rawlist
	}

}

func wldChk(adress string) (string, string, bool) {
	if strings.Contains(adress, "*") {
		return strings.Split(adress, "*")[0], strings.Split(adress, "*")[1], true
	} else {
		return adress, adress, false
	}
}
