package main

import (
	"fmt"
	"math"
)

type hitung interface {
	luas() float64
	keliling() float64
}

type lingkaran struct {
	diameter float64
}

func (l lingkaran) jarijari() float64 {
	return l.diameter / 2
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jarijari(), 2)
}

func (l lingkaran) keliling() float64 {
	return math.Pi * math.Pow(l.jarijari(), 2)
}

type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (p persegi) keliling() float64 {
	return 4 * p.sisi
}

func main() {

	var bangunDatar hitung

	bangunDatar = persegi{sisi: 10.0}
	fmt.Println("====== Persegi")
	fmt.Println("luas : ", bangunDatar.luas())
	fmt.Println("keliling : ", bangunDatar.keliling())

	bangunDatar = lingkaran{diameter: 14.0}
	fmt.Println("===== lingkaran")
	fmt.Println("luas      :", bangunDatar.luas())
	fmt.Println("keliling  :", bangunDatar.keliling())
	fmt.Println("jari-jari :", bangunDatar.(lingkaran).jarijari())

}
