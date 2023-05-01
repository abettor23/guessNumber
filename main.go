package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Игра \"Угадай число!\"")
	fmt.Println("У меня 4 попытки, я попробую угадать, какое число вы загадали.")

	fmt.Println("Загадайте число от 1 до 10 включительно.")
	time.Sleep(5 * time.Second)

	minNum := 1
	maxNum := 10
	tries := 0

	for tries < 4 {
		rand.Seed(time.Now().UnixNano())
		responseNum := rand.Intn(maxNum-minNum+1) + minNum

		fmt.Println("Загаданное число -", responseNum)
		fmt.Println("Верно? Введите ответ в формате да, больше, меньше")
		var userAnswer string
		fmt.Scan(&userAnswer)

		if userAnswer != "да" && userAnswer != "больше" && userAnswer != "меньше" {
			fmt.Println("Ошибка! Ваш ответ некорректен! Будьте внимательны!")
			return
		}

		if userAnswer == "да" {
			fmt.Println("Отлично! Я отгадал ваше число!")
			return
		} else if userAnswer == "больше" {
			minNum = responseNum + 1
		} else if userAnswer == "меньше" {
			maxNum = responseNum - 1
		}

		tries++
	}

	fmt.Println("Мои попытки кончились, я не смог отгадать ваше число :(")
}
