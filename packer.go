package main

import (
	"os/exec"
	"time"
)

func packer(exeName, sourceFiles, outputPath, ArchName string) {
	timestamp := time.Now().Format(config.System.TimeFormat)                       //Даташтамп.Лишние детали удалены форматированием Format(вид)
	destinationArch := outputPath + ArchName + "-" + timestamp + ".7z"             //имя и путь файла-бекапа
	flags := []string{"a", destinationArch, sourceFiles, "-r", "-ssw"}             //А, имя и путь получаемого файла, путь того что надо заархивировать
	cmd := exec.Command(exeName, flags[0], flags[1], flags[2], flags[3], flags[4]) //запускаем архиватор с нашими парам.
	if err := cmd.Run(); err != nil {
		loger("packer(), cmd.Run error", err)
	} else {
		loger("Archive created", flags[1])
	}
}
