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
- Merupakan tipe kumpulan karakter atau teks
- Diawali dengan petik dua (di go tidak bisa petik satu)

## Function Untuk String
- len("isinya") = hitung jumlah karakter
- "string"[number] = mengambil karakter pada posisi yang di tentukan

