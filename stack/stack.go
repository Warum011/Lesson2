package stack

import (
	"errors"
	"fmt"
)

// StackOnSlice - стек на базе slice-а
type StackOnSlice struct {
	items []int // элементы стека
	size  int   // размер стека
}

// NewStackOnSlice - конструктор стека
func NewStackOnSlice() *StackOnSlice {
	return &StackOnSlice{
		items: make([]int, 0), // в момент создания стек пустой
		size:  0,              // в момент создания размер стека равен 0
	}
}

// Print - возвращает строковое представление стека
func (s *StackOnSlice) Print() string {
	result := ""
	for i := s.size - 1; i >= 0; i-- {
		result += fmt.Sprintf("%d ", s.items[i])
	}
	return result
}

// Empty - возвращает признак пустой ли стек
func (s *StackOnSlice) Empty() bool {
	return s.size == 0
}

func (s *StackOnSlice) increaseSlice() {
	newCount := s.size * 2 // в newCount кладем новый размер слайса
	if s.size == 0 {       // если текущий размер 0, то по умолчанию задаем размер 4
		newCount = 4
	}
	newSlice := make([]int, newCount) // создаем новый слайс
	copy(newSlice, s.items)           // копируем элементы из старого слайса в новый
	s.items = newSlice                // обновляем ссылку на элементы стека новым слайсом
}

func (s *StackOnSlice) Push(item int) {
	if s.size == len(s.items) {
		s.increaseSlice()
	}
	s.items[s.size] = item
	s.size++
}

func (s *StackOnSlice) Pop() (int, error) {
	if s.size == 0 {
		return 0, errors.New("stack is empty")
	}
	s.size--
	return s.items[s.size], nil
}
