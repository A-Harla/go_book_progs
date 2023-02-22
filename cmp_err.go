// Программа вызывается из терминала и получает на вход 2 числа (int или float) и сравнивает их
// вывыодит или сообщение о том что они равны, или min max, или сообщение об ошибке

package main

import (
	"errors"
	"io"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args // считываем количество переданных функции аргументов
	// Возможная ошибка 1 - неверное количество переданных аргументов
	if len(arguments) < 2 { // 1 арг - имя программы, 2 и 3 - числа для сравнения
		err := errors.New("You should give at least 2 arguments!\n")
		io.WriteString(os.Stderr, err.Error()) // выводим сообщени об ошибке в консоль
		io.WriteString(os.Stdout, "\n")
	}
	// io.WriteString также возвращает одним из аргументов ошибку, но мы не будем её отлавливать, в целях экономии времени

	// Возможная ошибка 2 - аргументы не числа
	arg1, err := strconv.ParseFloat(arguments[1], 32)
	if err == nil { // если первое число преобразовалось, то пробуем преобразовать второе
		arg2, err := strconv.ParseFloat(arguments[2], 32)
		if err == nil { // если второе тоже преобразовалось, то сравниваем их
			if arg1 == arg2 { // при равенстве аргументов выводим в консоль сообщение об этом
				io.WriteString(os.Stdout, "Arguments are equal!\n")
				return
			} else {
				if arg1 < arg2 { // В случае если аргументы не равны, определяем какой из аргументов больше
					result := "Arguments are not equal!\n Min: " + arguments[2] + "\n Max: " + arguments[1] + "\n"
					io.WriteString(os.Stdout, result)
					return
				} else {
					result := "Arguments are not equal!\n Min: " + arguments[1] + "\n Max: " + arguments[2] + "\n"
					io.WriteString(os.Stdout, result)
					return
				}
			}
		} else {
			io.WriteString(os.Stdout, "Second argument is incorrect!\n")
			io.WriteString(os.Stderr, err.Error())
			io.WriteString(os.Stdout, "\n")
			return
		}
	} else {
		io.WriteString(os.Stdout, "First argument is incorrect!\n")
		io.WriteString(os.Stderr, err.Error())
		io.WriteString(os.Stdout, "\n")
		return
	}

}
