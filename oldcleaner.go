package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func oldcleaner(input string) { //Получаем путь к нужной папке, и понеслась
	files, err := ioutil.ReadDir(input) //читаем папку, получаем массив из списка обЪектов в ней

	if err != nil {
		loger("oldcleaner() files ioutil.ReadDir", err)
	}

	for _, file := range files { // проходимся по этому массиву
		if filepath.Ext(file.Name()) == ".zip" { //и если расширение .zip

			CompareTime(file.ModTime().Format(config.System.TimeFormat), file.Name())
			//получили нужную нам дату в нужном формате но в типе string и передаем
			//ее в функцию сравнения с текущей, с именем файла
		}
	}
}
func CompareTime(creationTime, filename string) {
	archname := config.OldArchRemover.StorageArchDir + filename //Add file adress

	dateStamp, err := time.Parse(config.System.TimeFormat, creationTime) /* перевели дату из типа string
	в тип Time. Первый аргумент - формат даты. второе - переменная с string даты*/
	if err != nil {
		loger("CompareTime() dateStamp", err)
	}
	timeNow, _ := time.Parse(config.System.TimeFormat, time.Now().Format(config.System.TimeFormat))
	/*конструкция выше - приведение к нужному типу для сравнивалки .Sub*/

	difference := timeNow.Sub(dateStamp)        // сравниваем (Sub.) текущее время с dateStamp файла
	hours := float64(difference.Hours())        // Разницу в часах конвертируем в тип float
	if hours > config.OldArchRemover.HoursOld { //сравниваем с нужным числом часов (месяц - 720ч)
		FileRemover(archname)
	}

}
func FileRemover(name string) {
	err := os.Remove(name)
	if err != nil {
		loger("FileRemover() os.Remove()", err)
	} else {
		loger("Deleted:", name)
	}
}
