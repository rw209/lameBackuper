package main

import (
	"os"

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
	ExitArchiveFile []string
	ExitArchivePath []string
}

var config title

func main() {

	if _, err := toml.DecodeFile("backupper.conf", &config); err != nil {
		defaultConfig := `[system]
TimeFormat = "2006-01-02 15:04:05"##2006-01-02 15:04:05
LogMode = 2 ## 0 - all off, 1 - errors only, 2 - actions ##
[OldArchRemover]
HoursOld = 1.1
StorageArchDir = "/home/user/Backupper/Arch/"
[ArchConf]
SourceFilesDir = ["/home/user/1","/home/user/2","/home/user/3"]
ExitArchiveFile = ["BDarch1","BDarch2","BDarch3"]
ExitArchivePath = ["/home/user/Backup1","/home/user/Backup2","/home/user/Backup3"]
## For Win systems use \\ in config: "C:\\Backup\\Arch" ##`
		file, _ := os.Create("backupper.conf")
		defer file.Close()
		file.WriteString(defaultConfig)
		_, _ = toml.DecodeFile("backupper.conf", &config)
		loger("Config created. First run?")
	}
	loger("Program started")
	oldcleaner(config.OldArchRemover.StorageArchDir) //Вызов OldCleaner

	for i, _ := range config.ArchConf.SourceFilesDir {

		packer(
			config.ArchConf.SourceFilesDir[i],
			config.ArchConf.ExitArchivePath[i],
			config.ArchConf.ExitArchiveFile[i],
		) //пакуем нужную нам папку
	}
	loger("Finish")
}

///usr/lib/p7zip/7za a /home/user/arch.7z /usr/lib/p7zip/7zCon.sfx
