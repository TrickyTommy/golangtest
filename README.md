# Soal Esai 1
## Go-Lang
Bahasa pemrograman yang telah dibuat oleh google dengn menggunakan bahasa pemrograman C
Go-Lang dirilis ke public pada tahun 2009 jadi sudah 12 tahun lamanya
Go-Lang sangat populer sejak digunakan untuk membuat Docker pada tahun 2011, jadi golang didonkrak populeritasnya pada saat ini
Saat ini mulai banyak teknologi baru yang dibuat menggunakan bahasa Go-Lang dibanding dengan bahasa pemrograman C, seperti Kubernetes, Promotheus,CockroachDB, dll.Jadi lebih cenderung membuat pada sistem.
Go-Lang sangan ringan jadi lebih populer digunakan pada Start Up Start-Up dibandingkan dengan C, dan sering digunakan pada Microservice
### Kenapa Belajar Go-Lang
Bahasa Go-Lang sangat sederhana tidak membutuhkan oop dll,tidak membutuhkan waktu yang lama dalam mempelajarinya
Go-Lang mendukung concurrency programming
Go-Lang mendukung garbage colector jadi tidak membutuhkan management memory secara manual
Bahasa pemrograman yang sedang naik daun
### SDK
Download Go-Lang
Install Compiler Go-Lang
cek apakah sudah terinstall menggunakan go version
### Prosses Development Go-Lang
Go-Lang bahasa pemrograman yang di compile tetapi beda dengan C/Java kalau Java compile nya harus menggunakan java virtual mesin sedangkan Go tidak memakai itu. 
Cara menggunakannya download terlebih dahulu Go Compiler lalu bikin file dengan extension .go contoh: main.go, compilernya digunakan untuk menjalankan file Go hasilnya menjadi file Binary yang native sesuai komputer yang digunakan.
### Text Editor
Visual Studio Code (install plugin Go)
JetBrain GoLand
# Soal Esai 2
### Aturan Penulisan Go-Lang

```go
package main



import (
    "fmt"
)
 
func main(){
    fmt.Println("Belajar Golang di Website Kodingin.com")
}
```
Struktur Dasar Program Golang

Strukur Golang pada Dasarnya di bagi menjadi 4 yaitu

1. Dekrakasi Package
2. Import Library
3. Bagian Fungsi utama / Method Main.
4. Ekspresi.

Baris 1 , merupakan deklarasi package
Baris 2 – 5, merupakan import library fmt..
Beris 7, merupakan fungsi utama yang pertama kali di jalankan.

1. Deklarasi Package

Package adalah suatu cara untuk mengorganisasikan file yang ada atau bisa di bilang nama folder yang tersedia. Di Golang satu file dapat di jalankan ketika mempunya package main.

Nah bagaimana jika kasusnya file tersebut berada di root projek kita ?

Deklarasi package ketika berada di main projek yaitu dengan mendeklarasikan package dengan nama main
package main

```go	
package main
```
Ketika projek kita mempunyai modul yang banyak tentunya kita membuat folder guna agar struktur program lebih rapi dan mudah di pahami.

Lalu bagaimana contoh implementasi package di dalam folder ?

asumsikan di dalam projek kita punya folder dengan nama “model” nah di dalam folder terdapat file dengan nama “posts.go”, maka anda dapat menuliskan package seperti kode di bawah ini.
package model
```go
package model
```
Bagaimana jika kita di mendeklarasikan Packege ?

Program akan menimbulkan error / tidak dapat di jalankan
2. Bagian Import

Import digunakan untuk memanggil library yang ingin kita gunakan, contoh kode dia atas kita memanggil library dengan nama fmt.

Apa itu Library ?

Library merupakan Sekumpulan fungsi yang di proses sedemikian rupa untuk kita gunakan fungsi nya. Dengan menggunakan library kita lebih di mudahkan untuk membuat suatu program.

Di Golang berbeda dengan bahasa lain dalam hal import library. Pada Golang setiap perintah banyak menggunakan import walaupun hanya penulisan, misalnya hanya mencetak string saja.

Sintaks penulisannya ada dua yaitu
```go
import "nama librarynya satu"
import "nama librarynya dua"

import "nama librarynya satu"
import "nama librarynya dua"

atau
import(
 "nama librarynya satu"
 "nama librarynya dua"
)	
import(
 "nama librarynya satu"
 "nama librarynya dua"
)

Contohnya.
import "fmt"
	
import "fmt"
```
Artinya file main.go mengimport library fmt.
3. Method Main

Di dalam bahasa Golang harus ada satu nama method utama, dimana method tersebut di eskteksusi pertama kali saat program tersebut di jalankan, yaitu method main()

Tanpa method main() maka program tidak dapat di jalankan.

Contoh method main()
```go
func main(){
fmt.Println("Belajar Golang di Website Kodingin.com")
}
func main(){
fmt.Println("Belajar Golang di Website Kodingin.com")
}

fmt.Println artinya mencetak tulisan di sebuah command di layar.
```
4. Ekspresi

Bahassa lain eksrepsi yaitu statement, eksrepsi merupakan sebuah kode bagian terkecil. Kode tersebut bisa di artikan suatu kegiatan yang dilakukan.

Pada contoh di atas saya menggunakan ekspresi penulisan string yaitu

fmt.println("Belajar Golang di Website Kodingin.com")
```go
fmt.println("Belajar Golang di Website Kodingin.com")
```
Kode diatas akan menghasilkan tulisan “Belajar Golang di Website Kodingin.com” di layar monitor anda.

Ekspresi bisa di bilang sebuah baris kode. Di golang berbeda dengan bahasa lain, dimana akhir baris tidak perlu menambahkan kode titik kode(;).
Blok Program Golang

Blok kode adalah sebuah tanda yang digunakan untuk membungkus kode lainnya. Sifatnya tetap yaitu pembuat kode di awali dengan tanda { dan di akhiri tanda } pada sebuah fungsi.

Contohnya…
```go
//Blok Program If pada Golang
if i < i {

}
//Blok Program If pada Golang
if i < i {
 
}
```
Banyak bentuk blok kode pada golang. Pada intinya setiap ada pembuka { harus ada penutup }.
Penulisan Bagian Komentar pada Golang

Komentar merupakan satu teks / tulisan dimana ketika program itu di jalankan maka tidak akan ikut di eksekusi.

Manfaat komentar pada sebuah kode prorgram yaitu membuat dokumentasi, Mematikan sebuah ekspresi mapun fungsi tertentu.

Untuk gaya penulisan di Golang di bagi menjadi 2 yaitu

    inline, menggunakan 2 garus miring(//… dalam satu baris.
    multiline, menggunakan garis miring di ikuti bintang (/*…*/) untuk komentar lebih dari satuu baris kode program.

Contohnya.
```go
package main

import (
	"fmt"
)

func main() {
	// fmt.Println("Belajar Golang di Website Kodingin.com") | inline

	fmt.Println("Belajar Golang di Website Kodingin.com")
	/*
	Belajar Pemrograman Golang
	di Kodingin / multiline
	*/
}
package main
 
import (
    "fmt"
)
 
func main() {
    // fmt.Println("Belajar Golang di Website Kodingin.com") | inline
 
    fmt.Println("Belajar Golang di Website Kodingin.com")
    /*
    Belajar Pemrograman Golang
    di Kodingin / multiline
    */
}
```
