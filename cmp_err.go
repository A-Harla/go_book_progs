// Программа вызывается из терминала и получает на вход 2 числа (int или float) и сравнивает их
// вывыодит или сообщение о том что они равны, или min max, или сообщение об ошибке

package main

import (
	"errors"
	"io"
	"log"
	"os"
	"strconv"
)

var LOGFILE = "/tmp/mGo.log"

func catchIoError(err error) {
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// Здесь мы создаем новый журнальный файл, используя функцию os.OpenFile() с необходимыми правами доступа к UNIX-файлам (0644).

	iLog := log.New(f, "customLogLineNumber ", log.LstdFlags)
	iLog.SetFlags(log.LstdFlags)
	iLog.Println("Error while trying to use io.WriteString!")
	iLog.Println(err.Error())
	// функция SetFlags позволяет устанавливать выходные флаги (варианты) для текущего средства журналирования.
	//По умолчанию функция предлагает значения LstdFlags: Ldate и Ltime, то есть в каждой записи журнала,
	//которая записывается в журнальный файл, будут указаны текущая дата и время.
}

func main() {
	arguments := os.Args // считываем количество переданных функции аргументов
	// Возможная ошибка 1 - неверное количество переданных аргументов
	if len(arguments) < 2 { // 1 арг - имя программы, 2 и 3 - числа для сравнения
		err := errors.New("You should give at least 2 arguments!\n")
		_, errIo := io.WriteString(os.Stderr, err.Error()) // выводим сообщени об ошибке в консоль
		if errIo != nil {
			catchIoError(errIo)
			os.Exit(10)
		}
		_, errIo = io.WriteString(os.Stdout, "\n")
		if errIo != nil {
			catchIoError(errIo)
			os.Exit(10)
		}
	}

	// Возможная ошибка 2 - аргументы не числа
	arg1, err := strconv.ParseFloat(arguments[1], 32)
	if err == nil { // если первое число преобразовалось, то пробуем преобразовать второе
		arg2, err := strconv.ParseFloat(arguments[2], 32)
		if err == nil { // если второе тоже преобразовалось, то сравниваем их
			if arg1 == arg2 { // при равенстве аргументов выводим в консоль сообщение об этом
				_, errIo := io.WriteString(os.Stdout, "Arguments are equal!\n")
				if errIo != nil {
					catchIoError(errIo)
					os.Exit(10)
				}
				return
			} else {
				if arg1 < arg2 { // В случае если аргументы не равны, определяем какой из аргументов больше
					result := "Arguments are not equal!\n Min: " + arguments[2] + "\n Max: " + arguments[1] + "\n"
					_, errIo := io.WriteString(os.Stdout, result)
					if errIo != nil {
						catchIoError(errIo)
						os.Exit(10)
					}
					return
				} else {
					result := "Arguments are not equal!\n Min: " + arguments[1] + "\n Max: " + arguments[2] + "\n"
					_, errIo := io.WriteString(os.Stdout, result)
					if errIo != nil {
						catchIoError(errIo)
						os.Exit(10)
					}
					return
				}
			}
		} else {
			_, errIo := io.WriteString(os.Stdout, "Second argument is incorrect!\n")
			if errIo != nil {
				catchIoError(errIo)
				os.Exit(10)
			}
			_, errIo = io.WriteString(os.Stderr, err.Error())
			if errIo != nil {
				catchIoError(errIo)
				os.Exit(10)
			}
			_, errIo = io.WriteString(os.Stdout, "\n")
			if errIo != nil {
				catchIoError(errIo)
				os.Exit(10)
			}
			return
		}
	} else {
		_, errIo := io.WriteString(os.Stdout, "First argument is incorrect!\n")
		if errIo != nil {
			catchIoError(errIo)
			os.Exit(10)
		}
		_, errIo = io.WriteString(os.Stderr, err.Error())
		if errIo != nil {
			catchIoError(errIo)
			os.Exit(10)
		}
		_, errIo = io.WriteString(os.Stdout, "\n")
		if errIo != nil {
			catchIoError(errIo)
			os.Exit(10)
		}
		return
	}
}
