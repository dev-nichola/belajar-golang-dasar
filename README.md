# Pengenalan Singkat

- penulisan go mirip c/c#
- go hanya memiliki struc bukan class
- di go titik koma ; tidak wajib sama seperti js
- go secara default mendukung cuncurrency

# Compile Golang

```
go build
```

Cari file executablenya sesuai modulenya biasanya

```
./belajar-golang dasar
```

## Menjalankan Tanpa Compile

```
go run helloworld.go
```

# Multiple Main Function

function main di go itu unique, jadi di setiap file tidak boleh ada function yang sama

Contoh ada di helloworld.go dan sample.go

# Tipe Data Number

- INT = Bilangan Bulat
- FLOAT = Bilangan Desimal

# Tipe Data Boolean

- Benar = True
- Salah = False

# Tipe Data String

- merupakan tipe kumpulan karakter atau teks
- diawali dengan petik dua (di go tidak bisa petik satu)

## Function Untuk String

- len("isinya") = hitung jumlah karakter
- "string"[number] = mengambil karakter pada posisi yang di tentukan

# Variable

- tempat menyimpan data
- di golang sifatnya global artinya bisa digunakan di mana saja

Contoh Deklarasi :

```
var name string

	name = "Nichola Sapitri"
	fmt.Println(name)

	name = "Nichola Saputra"
	fmt.Println(name)
```

- walaupun tidak kita sebutkan tipe datanya golang akan otomatis tau tipe data yang kita gunakan di variablenya
- best practicenya kita sebutkan saja tipe data variablenya

## Kata Kunci Var

- kata kunci var di golang tidak wajib
- asalkan kita langsung menginisialisasikan isinya
- dan menggunakan kata kunci := saat deklarasi variablenya artinya hanya di gunakan pertama kali saja saat deklarasi variablenya serta langsung di inisiasi datanya

## Deklarasi Multiple Variable

membuat variable banyak sekaligus

```
var (
		hobby = "Makan Tidur Boker"
		gf    = "tak punya"
	)

	fmt.Println(hobby)
	fmt.Println(gf)
```

Di golang kalau variable tidak digunakan otomatis error

# Constant

- variable yang nilainya tidak bisa di ubah
- wajib di inisasi datanya ketika di buat
- constant tidak di panggil tidak apa apa beda dengan variable

## Deklarasi Multiple Constant

```
const (
		price = 120000
		value = 10000
	)

	fmt.Println(price)
	fmt.Println(value)
```

# Konversi Tipe Data

mengkonversi tipe data satu ke tipe data lain

```
      var nilai32 int32 = 32178000
	var nilai64 int64 = int64(nilai32)

	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai64)
	fmt.Println(nilai16)
```

# Type Declaration

Kemampuan membuat ulang tipe data baru dari tipe data yang sudah ada, biasanya digunakan membuat alias, tujuanya supaya mudah di mengerti.
deklarasinya seperti ini

```
type NoKTP string
```

# Operasi Matematika

- Penjumlahan, Pengurangan, Perkalian, Pembagian, Sisa Bagi Atau Modulus
- Augmented assignment Atau operasi penugasan
- Unary operator

# Operasi Perbandingan

untuk membandingan dua buah data, hasil nya adalah boolean true atau false

- ">" lebih dari
- "<" kurang dari
- "==" sama dengan
- dan lain lain

# Operasi Boolean

- && AND
- || ATAU
- ! Kebalikan

# Tipe Data Array

- array adalah tipe data yang kumpulan dengan tipe data yang sama
- urutan data di array di sebut index dan di buat dari 0
- ada 2 cara membuatnya yaitu membuat dulu variablenya dan yang kedua membuat array secara langsung
- di golang itu tidak ada operasi menghapus index di array
- jika kita tidak ingin menentukan jumlahnya secara pasti kita bisa menggunakan `...` saat deklarasi array nya

## Function Array

- len(array) = mendapatkan panjang array
- array[index] = mendapatkan data di posisi index
- array[index] = mengubah data di posisi index

# Tipe Data Slice

- adalah potongan dari data array
- slice mirip array, yang membedakan ukuranya bisa berusah

## Detail Tipe Data Slice

- memiliki 3 data, pointer, length, dan capacity
- pointer adalah petunjuk data pertama
- lenght adalah panjang dari slice
- capacity adalah kapasitas dari slice, dimana length tidak boleh keluar dari capacity

## Function Slice

sama seperti array ada function yang bisa digunakan

- len(slice) = mendapatkan panjang slicenya!, ingat ya bukan panjang arraynya tapi panjang slicenya
- cap(slice) = untuk mendapatkan kapasitas
- append(slice, data) = menambahkan slice baru
- make([]TypeData, length, capacity) = membuat slice baru
- copy(destination, source) = menyalin slice

## Hati Hati Saat Membuat Array Dan Slice

```
	iniArray := [...]string{"Nichola", "Saputra"}
	iniSlice := []string{"Nichola Saputra"}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
```
