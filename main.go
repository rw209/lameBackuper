package main

import (
	"github.com/BurntSushi/toml"
)

type title struct {
	System         sys
	OldArchRemover oldarch
	ArchConf       ArchConf
}
type sys struct {
	TimeFormat string
	LogMode    int
}
type oldarch struct {
	HoursOld       float64
	StorageArchDir string
}
type ArchConf struct {
	SourceFilesDir  []string
	ExecutableName  []string
	ExitArchiveFile []string
	ExitArchivePath []string
}

var config title

func main() {

	if _, err := toml.DecodeFile("backupper.conf", &config); err != nil {
		loger("Configuration no found") //Сделать возможность прием дефолтных параметров, для
		//логирования, если не найден конфиг
	}
	loger("Program started")
	dirlistener(config.OldArchRemover.StorageArchDir) //Вызов OldCleaner

	for i, _ := range config.ArchConf.ExecutableName {

		packer(
			config.ArchConf.ExecutableName[i],
			config.ArchConf.SourceFilesDir[i],
			config.ArchConf.ExitArchivePath[i],
			config.ArchConf.ExitArchiveFile[i],
		) //пакуем нужную нам папку
	}
	loger("Finish")
}

///usr/lib/p7zip/7za a /home/user/arch.7z /usr/lib/p7zip/7zCon.sfx
