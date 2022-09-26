Nomor peserta : 149368582100-405 
Nama : Renny Kristina Salainti

Tgl : 26/09/22

Interface =>
Interface adalah sebuah tipe data pada bahasa Go yang merupakan kumpulan dari definisi satu atau lebih method. Kita dapat menggunakan tipe
data interface jika kita telah mengimplementasikan method-method yang telah didefinisikan oleh interface tersebut melalui tipe data lainnya.

Method area merupakan sebuah method yang me-return nilai dengan tipe data float64 dan method perimeter merupakan sebuah method yang me-return
nilai dengan tipe data float64. 

Empty interface =>
Empty interface merupakan suatu tipe data yang dapat menerima segala tipe data pada bahasa Go. Maka dari itu, sebuah variable dengan tipe data empty
interface dapat diberikan nilai dengan tipe data yang berbeda-beda. Contohnya seperti pada gambar di sebelah kanan.

Reflect =>
Reflect digunakan untuk melakukan  inspeksi variabel, mengambil informasi dari variabel tersebut atau bahkanmemanipulasinya.
Cakupan informasi yang bisa didapatkan lewat reflection sangat luas, seperti melihat struktur variabel, tipe, nilai pointer,dan banyak lagi.
Go menyediakan package reflect, berisikan banyak sekali fungsi untuk keperluan reflection.

Dari banyak fungsi yang tersedia di dalam package tersebut, ada 2 fungsi yang paling penting untuk diketahui, yaitureflect.ValueOf() danreflect.TypeOf().

- Fungsi reflect.ValueOf() akan mengembalikan objek dalam tipe reflect.Value, yang berisikan informasi yangberhubungan dengan nilai pada variabel yang dicari
- Sedangkan reflect.TypeOf() mengembalikan objek dalam tipe reflect.Type. Objek tersebut berisikan informasi yangberhubungan dengan tipe data variabel yang dicari,

            dengan reflection tipe data dari nilai variabel dapat diketahui dengan mudah.
            fungsu reflect.ValueOf() memiliki parameter yang bisa menam[ung segala jenis
            tipe data, fungsi ini mengembalikanobjek dalam tipe reflect.value yg berisikan informasi
            mengenai variabel yg bersangkutan. Salah satu methodnya, Type().

Concurrency & go routines =>
            Arti dari concurrency adalah mengeksekusi sebuah proses secara independen atau berhadapan dengan lebih dari satu tugas dalam waktu yang sama.
Perlu diingat disini bahwa concurrency  berbeda dengan parallelism, karena parallelism memiliki arti mengerjakan tugas yang banyak secara bersamaan
dari awal hingga akhir. Sedangkan pada concurrency, kita tidak akan tahu tentang siapa yang akan menyelesaikan tugasnya terlebih dahulu.
Sedangkan, Goroutine merupakan sebuah thread ringan pada bahasa Go untuk melakukan concurrency. Jika  dibandingkan dengan thread biasa, 
            Goroutine memiliki ukuran thread yang jauh lebih ringan. Pada saat kita mengeksekusi sebuah Goroutine, maka satu Goroutine hanya membutuhkan 
2kb memori saja, sedangkan satu thread biasa dapat menghabiskan 1-2mb memori.Goroutine bersifat asynchronous sehingga proses nya tidak saling 
tunggu dengan Goroutine lainnya. Untuk membuat sebuah Goroutine, maka kita harus terlebih dahulu membuat sebuah function. Lalu ketika kita ingin 
memanggil function tersebut, maka kita perlu menggunakan keyword go sebelum kita memanggil function tersebut

Waitgroup => WaitgroupIntroductionWaitgroup merupakan sebuah struct dari package sync, yang digunakan untuk melakukan sinkronisasi terhadap goroutine.