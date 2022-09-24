Nomor peserta : 149368582100-405 
Nama : Renny Kristina Salainti

Tgl : 23/09/22

1) Function : Bahasa Go juga memiliki function atau fungsi nya sendiri. Cara menulis sebuah function pada Go cukup
mudah. Caranyaadalah dengan menggunakan keyword func lalu diikuti dengan nama function dan parameter yang digunakan. 
Contohnyaseperti pada gambar pertama dibawah. Kita telah membuat function bernama greet yang menerima  2 parameter dengan 
tipedata string dan int8. Function greet merupakan sebuah function yang tidak mengembalikan/me-return nilai apapun,
melainkanfunction ini hanya bertugas untuk menampilkan text pada terminal kita menggunakan fmt.Printf.'

2). Clouser : Closure merupakan merupakan sebuah anonymous function atau function tanpa nama yang dapat disimpan sebagai 
sebuahvariable maupun dapat dijadikan sebagai nilai return dari sebuah functionUntuk pertama-tama, 
kita akan membahas bagaimana cara menggunakan closure yang disimpan sebagai variable.

3). Pointer : Pointer merupakan sebuah variable spesial yang digunakan untuk menyimpan alamat memori variable lainnya. 
Sebagai contoh sebuah variabel bertipe integer memiliki nilai 4, maka yang dimaksud pointer adalah alamat memori dimana 
nilai 4disimpan, bukan nilai 4 itu sendiri.Variabel-variabel yang memiliki reference atau alamat memori yang sama, saling 
berhubungan satu sama lain dan nilainyapasti sama. Ketika ada perubahan nilai, maka akan memberikan efek kepada variabel lain 
(yang alamat memorinya sama) yaitunilainya juga ikut berubah.Cara membuat suatu variable menjadi sebuah pointer cukup mudah.
Caranya adalah dengan menggunakan tanda asterisk *sebelum penulisan tipe datanya.

4). Struct : Struct adalah sebuah tipe data berupa kumpulan/koleksi dari berbagai macam property/field dan juga method.
Cara membuat struct pada Go cukup mudah. Contohnya seperti pada gambar di sebelah kanan.

5). Exported & Unexported : Pada bahasa Go, tiap folder yang berbeda akan dianggap sebagai suatu package.
Kita dapat menggunakan berbagai variable ataupun tipe data dari package lainnya asalkan variable atau tipe
data lainnya tersebut telah ter-eksport dari package nya. Kemudian cara kita meng-eksport suatu variable atau
suatu tipe data adalah dengan mengawali penulisan variable maupun tipe data lainnya dengan huruf kapital atau upper case.