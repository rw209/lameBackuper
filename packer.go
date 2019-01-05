package main

import (
	"time"
)

func packer(sourceFiles, outputPath, ArchName string) {
	timestamp := time.Now().Format(config.System.TimeFormat)
	name := outputPath + ArchName + "-" + timestamp + ".zip"
	zipper(sourceFiles, name)
}

func zipper(filesPath, ArchFileName string) {

	err := zipThis(filesPath, ArchFileName)

	if err != nil {
		loger("zipper(), zipThis() error", err)
	} else {
		loger("Archive created", ArchFileName)
	}

}
