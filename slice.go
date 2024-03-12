package main

import "fmt"

func main() {
	names := [...]string{"Nichola", "Saputra", "Joko", "Morro", "Budi", "Nugraha"}

	sliceA := names[0:2]

	fmt.Println(sliceA)

	sliceB := names[2:4]
	fmt.Println(sliceB)

	sliceC := names[:4]
	fmt.Println(sliceC)

	sliceD := names[4:]
	fmt.Println(sliceD)

	// var sliceE []string = names[3:]
	// println(sliceE)

	// Append
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	daySlice1 := days[5:] //

	fmt.Println(daySlice1)

	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Sabtu Lama"

	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	// Membuat Slice Dari Awal Sekali
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Nichola"
	newSlice[1] = "Saputra"
	newSlice2 := append(newSlice, "Joko")

	fmt.Println(newSlice)
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// Copy Slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)
	fmt.Println(len(fromSlice))
	fmt.Println(len(toSlice))

	// Hati Hati
	iniArray := [...]string{"Nichola", "Saputra"}
	iniSlice := []string{"Nichola Saputra"}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
