package lmap

import (
	"fmt"
)

// Map - структура для хранения данных
type Map struct {
	nextIndex int64
	mp        map[int64]int64
}

// NewMap создает новый экземпляр Map и возвращает его
func NewMap() *Map {
	return &Map{
		mp: make(map[int64]int64),
	}
}

// Len возвращает количество элементов в Map.
func (m *Map) Len() int64 {
	return int64(len(m.mp))
}

// Add добавляет значение в Map и возвращает его уникальный индекс
func (m *Map) Add(value int64) int64 {
	index := m.nextIndex
	m.mp[index] = value
	m.nextIndex++
	return index
}

// RemoveByIndex удаляет значение по указанному индексу из Map
func (m *Map) RemoveByIndex(id int64) {
	delete(m.mp, id)
}

// RemoveByValue удаляет первое найденное значение из Map
func (m *Map) RemoveByValue(value int64) {
	for id, v := range m.mp {
		if v == value {
			delete(m.mp, id)
			break
		}
	}
}

// RemoveAllByValue удаляет все значения из Map по указанному значению
func (m *Map) RemoveAllByValue(value int64) {
	for id, v := range m.mp {
		if v == value {
			delete(m.mp, id)
		}
	}
}

// GetByIndex возвращает значение по указанному индексу из Map
func (m *Map) GetByIndex(id int64) (int64, bool) {
	value, ok := m.mp[id]
	return value, ok
}

// GetByValue возвращает индекс первого найденного значения из Map
func (m *Map) GetByValue(value int64) (int64, bool) {
	for id, v := range m.mp {
		if v == value {
			return id, true
		}
	}
	return 0, false
}

// GetAllByValue возвращает индексы всех значений, равных указанному значению, из Map
func (m *Map) GetAllByValue(value int64) (ids []int64, ok bool) {
	for id, v := range m.mp {
		if v == value {
			ids = append(ids, id)
		}
	}
	if len(ids) > 0 {
		return ids, true
	}
	return nil, false
}

// GetAll возвращает все значения из Map
func (m *Map) GetAll() (values []int64, ok bool) {
	for _, v := range m.mp {
		values = append(values, v)
	}
	if len(values) > 0 {
		return values, true
	}
	return nil, false
}

// Clear очищает Map, удаляя все элементы
func (m *Map) Clear() {
	m.mp = make(map[int64]int64)
	m.nextIndex = 0
}

// Print выводит содержимое Map в консоль
func (m *Map) Print() {
	for index, value := range m.mp {
		fmt.Printf("Индекс: %d, Значение: %d\n", index, value)
	}
}
