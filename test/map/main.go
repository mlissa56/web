package main

import (
	"fmt"
	lmap "map/storage/map"
)

type Storage interface {
	NewStorage() interface{}
	Add(value interface{}) int64
	Len() int64
	RemoveByIndex(id int64)
	RemoveByValue(value interface{})
	RemoveAllByValue(value interface{})
	GetByIndex(id int64) (interface{}, bool)
	GetByValue(value interface{}) (int64, bool)
	GetAllByValue(value interface{}) (ids []int64, ok bool)
	GetAll() (values []interface{}, ok bool)
	Clear()
	Print()
}

func main() {
	// Создаем новый экземпляр Map
	m := lmap.NewMap()

	// Добавляем значения в Map
	m.Add(1)
	m.Add(2)
	m.Add(3)
	m.Add(2) // Добавляем дублирующее значение

	// Выводим все значения в Map
	values, ok := m.GetAll()
	if ok {
		fmt.Println("Все переменные:", values)
	} else {
		fmt.Println("Map пуст")
	}

	// Удаляем значение по индексу
	m.RemoveByIndex(1)

	// Выводим содержимое Map после удаления
	m.Print()

	// Поиск значения по существующему индексу
	indexToFind := int64(3)
	value, found := m.GetByIndex(indexToFind)
	if found {
		fmt.Printf("Значение под индексом %d: %d\n", indexToFind, value)
	} else {
		fmt.Printf("Значение под индексом %d не найдено\n", indexToFind)
	}

	indexToFind = int64(4) // Проверяем на поиск значения под несуществующим индексом
	value, found = m.GetByIndex(indexToFind)
	if found {
		fmt.Printf("Значение под индексом %d: %d\n", indexToFind, value)
	} else {
		fmt.Printf("Значение под индексом %d не найдено\n", indexToFind)
	}
}
