package main

import (
	"fmt"
)

func main() {
	var first string = "first string" // объявляем переменную var имя переменной тип = значение
	fmt.Println(first)

	var second = "Вторая строка" //можно и без явнгого объявления типа
	fmt.Println(second)

	var i, k, j int //инициализируем несколько переменных типа int

	i, k, j = 45, 32, 17 // присваеваем переменным значение массово
	fmt.Println(j, i, k)

	var third, tempcel, varbool = "третья строка", 36.6, true //объявляем несколько переменных разных типов, не указывая тип
	fmt.Println(third, tempcel, varbool)

	fourth := "Четвертая строка" //локальная переменная, доступен краткий формат инициализации имя := значение
	fmt.Println(fourth)

	const pi = 3.14 // объявление константы
	fmt.Println(pi)

	fmt.Println(float32(j / 2))

	fmt.Printf("Первое число - %f32 \n", float32(i)*2.2) //приведение типа и форматируемая строка
}
