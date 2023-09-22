package list

import (
	"fmt"
)

type List struct {
	len       int64
	firstNode *node
}

// NewList создает новый список
func NewList() (l *List) {
	return &List{len: 0, firstNode: nil}
}

// Len возвращает длину списка
func (l *List) Len() (len int64) {
	return l.len
}

// Add добавляет элемент в список и возвращает его индекс
func (l *List) Add(value int64) (id int64) {
	newNode := &node{value: value}
	
	if l.firstNode == nil {
		l.firstNode = newNode
	} else {
		current := l.firstNode
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
	
	l.len++
	
	return l.len - 1
}

// RemoveByIndex удаляет элемент из списка по индексу
func (l *List) RemoveByIndex(id int64) {
	if id < 0 || id >= l.len {
		return
	}

	if id == 0 {
		l.firstNode = l.firstNode.next
	} else {
		current := l.firstNode
		for i := int64(0); i < id-1; i++ {
			current = current.next
		}
		current.next = current.next.next
	}
	l.len--
}

// RemoveByValue удаляет элемент из списка по значению
func (l *List) RemoveByValue(value int64) {
	current := l.firstNode
	if current != nil && current.value == value {
		l.firstNode = current.next
		l.len--
		return
	}

	for current != nil && current.next != nil {
		if current.next.value == value {
			current.next = current.next.next
			l.len--
			return
		}
		current = current.next
	}
}

// RemoveAllByValue удаляет все элементы из списка по значению
func (l *List) RemoveAllByValue(value int64) {
	for l.firstNode != nil && l.firstNode.value == value {
		l.firstNode = l.firstNode.next
		l.len--
	}

	current := l.firstNode
	for current != nil && current.next != nil {
		if current.next.value == value {
			current.next = current.next.next
			l.len--
		} else {
			current = current.next
		}
	}
}

// GetByIndex возвращает значение элемента по индексу.
//
// Если элемента с таким индексом нет, то возвращается 0 и false.
func (l *List) GetByIndex(id int64) (value int64, ok bool) {
	if id < 0 || id >= l.len {
		return 0, false
	}

	current := l.firstNode
	for i := int64(0); i < id; i++ {
		current = current.next
	}
	return current.value, true
}

// GetByValue возвращает индекс первого найденного элемента по значению.
//
// Если элемента с таким значением нет, то возвращается 0 и false.
func (l *List) GetByValue(value int64) (id int64, ok bool) {
	current := l.firstNode
	index := int64(0)

	for current != nil {
		if current.value == value {
			return index, true
		}
		current = current.next
		index++
	}

	return 0, false
}

// GetAllByValue возвращает индексы всех найденных элементов по значению
//
// Если элементов с таким значением нет, то возвращается nil и false.
func (l *List) GetAllByValue(value int64) (ids []int64, ok bool) {
	var indices []int64
	current := l.firstNode
	index := int64(0)

	for current != nil {
		if current.value == value {
			indices = append(indices, index)
		}
		current = current.next
		index++
	}

	if len(indices) > 0 {
		return indices, true
	}

	return nil, false
}

// GetAll возвращает все элементы списка
//
// Если список пуст, то возвращается nil и false.
func (l *List) GetAll() (values []int64, ok bool) {
	if l.len == 0 {
		return nil, false
	}

	values = make([]int64, 0, l.len)
	current := l.firstNode

	for current != nil {
		values = append(values, current.value)
		current = current.next
	}

	return values, true
}

// Clear очищает список
func (l *List) Clear() {
	l.len = 0
	l.firstNode = nil
}

// Print выводит список в консоль
func (l *List) Print() {
	current := l.firstNode
	for current != nil {
		fmt.Printf("%d ", current.value)
		current = current.next
	}
	fmt.Println()
}
