149368582100-405 Renny Kristina Salainti

A. Introduction

>> Go Programming Introduction

Go adalah sebuah bahasa pemrograman open-source yang dibuat di Google pada tahun 2007 oleh RobertGriesemer, Rob Pike, 
dan Ken Thompson. Berikut merupakan kelebihan dari bahasa Go:
- Clean And Simple: Go membuat segala sesuatu menjadi kompak dan indah.
- Concurrent: Go dapat menjalankan fungsi-fungsi dengan ringan seperti thread dan dapat berkomunikasi antara
fungsi yang satu dan fungsi yang lain.
- Fast: Proses compile di Go dapat dilakukan dengan cepat seperti proses compile pada bahasa pemrogramanC.

>> Installation Go di Windows
 
 >Dapatkan installer dari Go pada link berikut https://golang.org/dl/, pilih installer sesuai sistem operasi.
 >Selanjutnya, menjalankan installernya, klik next hingga prosesinstalasi selesai.
 >Pada dasarnya, jika kalian tidak merubah path pada saat instalasi, Go akan ter-install di C:\go. Path
 tersebut secara otomatis akan didaftarkan dalam PATH environment variable.
 >Bukalah Command Prompt / CMD, kemudian eksekusi 'go version' untuk mengecek versi Go.
 
 >>Create First Project  with Go

 >Keyword package (setiap program harus memiliki package minimal satu file dengan nama package main, file ini 
 akan dieksekusi pertama kali ketika program dijalankan), 
 caranya : [ package <nama-package> enter package main ]
 >Keyword import ( ini digunakan intuk import/memasukan package lain ke dalam file program, package "ftm" merupakanpa
 package bawaan yang disediakan oleh Go, fungsinya untuk keperluan I/O ), berikut skema penulisan [ import <nama_package> enter import "ftm"]
 >Main function ( merupakan yang dipanggil pertama kali pada saat eksekusi program )
 >fmt.Println() ( untuk memunculkan teks di layar terminal, cmd fungsi ini berisi didalam package "fmt")
 >fmt.Print() ( sama seperti sebelumnya, bedanya ini tidak menghasilkan baris baru di akhir output dan nilai pd parameter
 yang dimasukkan difungsi ini digabungkan tanpa pemisah )
 { fmt.Println tidak perlu menambahkan spasi di tiap kata karena akan dibuat secara otomatis ditambahkan disela  nilai. fmt.Print perlu ditambahkan spasi
 karena fungsi ini tidak menambahkan secara otomatis
 > cara menambahkan komentar, caranya : // dan /* */.

 B. Variable
 
 >> Variable with data type
 >> Variable without data type
 >> Variable without data type ( Short Declaration )
 >> Multiple variable declarations
 >> Underscore variable
 >> fmt.Print Usage

 C. Data Types
 >> Number(Integer) (non desimal/non floating-point)
 Secara umum integer dibagi menjadisub-kategori yaitu uint dan int.
 -int: Bilangan cacah (Bilangan positif)
 -uint: Bilangan bulat (bilangan positif maupunnegatif)
 >> Number(Float), secara umum float dibagi menjadi 2 jenis yaitufloat32 dan float34.
 verb %f : untuk format angka desimal/tipe data float menjadi 6 digit angka desimal.
 verb %.nf (n : bisa diganti dengan jumlah digit yang ingin dihasilkan)
 >>String ( nilainya di apit oleh tanda petik dua "")

 D. Constants & Operator
 >> Constant
 >> Operators
 >> Arithmetic Operators
 >> Rasional Operators
 >> Logical Operators