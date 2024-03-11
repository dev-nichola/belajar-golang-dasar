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