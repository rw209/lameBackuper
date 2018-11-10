package main

import (
	"log"
	"os"
)

func loger(action ...interface{}) {
	writercall := false //если config.Logmode = 0, так и останется false, логов не будет.
	for i := range action {

		switch action[i].(type) { //проверка на тип того что пришло в интерфейс
		case error: /*если хоть одна из ячеек []interface имеет тип error, а LogMode = 1,2 то: */
			if config.System.LogMode == 1 || config.System.LogMode == 2 {
				writercall = true // это.
			}
		default: //если ничего с типом error нет:
			if config.System.LogMode == 2 { //при этом ЛогМод = 2, это запишется в лог.
				writercall = true
			}

		}
	}
	if writercall == true {
		WriteLog(action) //запись всех ячеек массива action
	}
}

func WriteLog(action interface{}) { //
	logfile, _ := os.OpenFile("backupper.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	text := log.New(logfile, "", log.Ldate|log.Ltime)
	text.Printf("%v \r", action)
}
