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

# Tipe Data Map

- map adalah tipe data kumpulan key-value, yang keynya harus unique
- map tidak ada batasnya, tidak seperti array dan slice
- deklarasinya `map[tipeDataKeynya]tipeDatanya`

## Function Map

- len(map) = mendapatkan jumlah data
- map[key] = mengambil data
- map[key] = value - mengubah data di map
- make(map[typeKey]TypeValue) = membuat map baru
- delete(map, key) = menghapus map

# Percabangan If Else

- blok if akan di eksekusi ketika nilainya true
- else digunakan jika kondisinya false
- else if untuk menambhakan kondisi lainya

## If dengan short statement

if mendukung pembuatan statement sederhanan sebelum melakukan pengecekan

# Switch Statement

switch statemen merupakan bentuk sederahana dari if, biasanya digunakan untuk pengecekan ke kondisi dalam satu variable

## Switch Tanpa Kondisi

switch tidak wajib menggunakan kondisi, kita bisa menambahkan kondisi tersebut di setiap casenya

# For Loops / Perulangan

```
func main() {

	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan Ke = ", counter)
		counter++
	}

	fmt.Println("Selesai")
}
```

Atau

```
func main() {

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan Ke = ", counter)
	}

	fmt.Println("Selesai")

}
```

## For Range

- For bisa digunakan untuk mengiterasi semua data yang ada di collection
- data collection itu contohnya, array, slice, dan map

```
	for index, name := range names {
		fmt.Println("INDEX = ", index, "NAME = ", name)
	}
```
Atau Kalau Tidak Butuh Indexnya
```	for _, name := range names {
		fmt.Println("NAME = ", name)
	}
```

# Break & Continue
- break = menghentikan seluruh perulangan
- continue = menghentikan perulangan dan melanjutkan ke perulangan selanjutnya

# Function 
- adalah blok atau baris kode yang dapat digunakan secara berulang ulang
- pembuatanya menggunakan kata kunci func di ikuti nama functionya dan tanda buka plus tanda kurung

# Function Parameter
- kadang kita membutuhkan data dari luar, atau di sebut parameter
- bisa lebih dari satu menambahkan parameter
- parameter tidak wajib
- parameter wajib mengirim data kalau parameternya ada
- nama parameter itu unique

# Function Return 
- function itu bisa mengembalikan data
- dengan kata kunci return

# Returning Multiple Values
- function tidak hanya bisa mengembalikan satu nilai saja artinya bisa lebih banyak
- untuk bisa melakukanya kita harus memberitahukan semua tipe data returnya

## Menghiraukan Return Value
- kadang kita tidak ingin menangkap semua value maka dari itu kita bisa gunakan ```_``` garis bawah

# Named Return Value
- membuat variable secara langsung di tipe data return functionya
- artinya tidak usah membuat variablenya secara langsung di dalam functionya

# Variadic Function 
- parameter di function itu yang terakhir itu memiliki kemampuan variable argumen
- atau datanya bisa memiliki satu input, atau anggap saja seperti array
- mirip spread operator di javascript
- kita juga bisa menjadikan slice sebagai argumen untuk parameternya

# Function Value
- di golang function itu first class citizen
- function juga dianggapnya tipe data, bisa disimpan di variable

# Function as Parameter
- function juga bisa digunakan sebagai parameter
- kita juga bisa membuat type sendiri untuk function kita contoh ada di ```function_parameter.go```

# Anonymous Function
- function yang tidak ada identifiernya di sebut anonymous function 
- contoh penggunaan yang sering adalah penggunaan callback atau parameter di function yang berbentuk function

# Recursive Function
- adalah function yang bisa memanggil dirinya sendiri
- kadang kita lebih mudah menyelesaikan masalah menggunakan factorial di banding dengan perulangan
- contoh kasusnya adalah factorial

# Closure
- closure adalah kemampuan sebuah function berinteraksi dengan data-data disekitar dalam scope yang sama
- <b>Gunakan closure dengan bijak</b>


