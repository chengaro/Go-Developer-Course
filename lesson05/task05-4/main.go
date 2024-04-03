package main

import "fmt"

func main() {
	// Все переменные имеют свой уникальный адрес в памяти
	// Переменные примитивных и собственных типов при присваивании копируют значения
	// Последующее изменение значения в переменной-источнике не меняет значение в переменной-приемнике
	b1, b2 := false, false
	b3 := b1
	fmt.Printf("&b1: %v, &b2: %v, &b3: %v\n", &b1, &b2, &b3)
	b1 = true
	fmt.Printf("%v %v %v\n\n", b1, b2, b3)

	type t struct{ name string }
	t1, t2 := t{name: "myname"}, t{name: "myname"}
	t3 := t1
	fmt.Printf("&t1: %v, &t2: %v, &t2: %v\n", &t1, &t2, &t3)
	fmt.Println(&t1 == &t3)
	t1.name = "another name"
	fmt.Printf("%v %v %v\n\n", t1, t2, t3)

	// Переменные типа map и slice также имеют уникальный адрес
	// Но при присваивании они под капотом ссылаются на один и тот же "объект"
	// И последующее изменение значения в переменной-источнике изменяет значение в переменной-приемнике
	s1, s2 := []int{0, 1}, []int{0, 1}
	s3 := s1
	fmt.Printf("&s1: %v, &s2: %v, &s2: %v\n", &s1, &s2, &s3)
	fmt.Println(&s1 == &s3)
	s1[0] = -100
	fmt.Printf("%v %v %v\n\n", s1, s2, s3)

	m1, m2 := map[int]int{0: 0, 1: 1}, map[int]int{0: 0, 1: 1}
	m3 := m1
	fmt.Printf("&m1: %v, &m2: %v, &m2: %v\n", &m1, &m2, &m3)
	fmt.Println(&m1 == &m3)
	m1[0] = -100
	fmt.Printf("%v %v %v\n\n", m1, m2, m3)

}
