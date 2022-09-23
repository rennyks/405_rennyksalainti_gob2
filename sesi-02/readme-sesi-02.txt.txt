Nomor peserta : 149368582100-405 
Nama : Renny Kristina Salainti

Tgl : 21/09/22
SESI 02 : FUNDAMENTAL GO PROGRAMMING

=> Conditions : dalam pemrogramman digunakan untuk mengatur alur dari suatu program, analogi yang biasa du=igunakan adalah lampu merah.
Misalkan, merah : berhenti, kuning : bersiap, hijau : jalan. Sama halnya, Conditions, kita bisa mengatur program untuk mengetahui kapan
harus mengeksekusi suatu kode.
- temporary variabel ( Jika kita perhatikan pada baris 8, age merupakan sebuah variable yang terbuat pada scope blockkondisional if
            else nya.Kita membuat variable age yang dimana kita membuat  perkurangan antara variable currentYeardengan angka 1998. Lalu
            setelah tanda semicolon; kita langsung membuat kondisinya yangmengatakan jika variable age lebih kecil daripada angka 17. Kondisi
            dibawah akan menghasilkankalimat yang ada pada blok else karena variable age menghasilkan nilai dengan angka 2 )
- switch ( switch pda Go tdk memerlukan keyword break, jadi jika suatu case telah terpenuhi maka dia tdk akan
            berlanjutan pda case berikutnya )
            * switch with relational operators, seperti halnya sebuah kondisi dengan keyword if, else if dan else.
            * switch falthrough keyword, ini digunakan ketika kita ingin melanjutkan pengecekan pada case selanjutnya
            walaupun case sebelumnya telah terpenuhi kondisinya
- nested conditions ( kondisi bersarangn yaitu dimana sebuah proses kondisional yang didalamnya terdapat proses kondisional
                        kembali. Misalkan, kita bisa menggunakkan if, else if, else, switch atau menggabungkan semuanya )

=> Loopings : perulangan merupakan proses berulang, dimana prosesnya akan berhenti jika memnuhi suatu kondisi.Bahasa Gi hanya memilik satu
looping yaitu keyword for atau dikenal for loop.

=> Array : merupakan ti pe data untuk menyimpan sebuah kumpulan dari data - data. Data yang disimpan pada array harus memiliki tipe data yang sama,
kecuali kita menyimpan sebagai suatu interface kosong. Array pd Go memiliki sifat fixed-lenght[ panjang yang tetap, yang harus kita tentukan dari awal
saat membuatnya ]. Dua macam cara membuat array :
            1) deklarasih terlebih dahulu tanpa membero nilai apanpun, dan
            2) mendeklarasi dan langsung menginialisasikan dengan memberikan sebuah nilai
            3) memodifikasi elemen yang terdapat dalam sebuah array dengan mengakses indexnya
            4) 2 cara penulisan agar kt bisa melakukan looping untuk mengakses element - element yg terdapat pada sebuah array.
            cara 1 mengunakan range loop, 2 menggunakkan looping biasa dan perlu menggunakkan fungsi len untuk mendapat panjang dari array
            5) array multidimensional yg brtarti terdapat sebuah array dalam sebuah array

=> Slice : merupakan tipe data yang mirip array, yg menyipan satu atau lebih data, perbedaannya slice tidak memiliki sifat fixed-lenght
artinya  panjang dari slice tidak tetap sehingga kita bisa dengan leluasa menentukan panjang dr slicenya.
            - membuat sebuah slice dengan menggunakan fungsi make
            - memanfaatkan fungsi append ( ini digunakan untuk mengembalikan nilai dari slice yg ditambahkan )
            - fungsi ellipsis(...) untuk memasukkan seluruh element pada satu array kedalam array lain
            - fungsi copy untuk mengcopy sebuah slice kedalam slice lain, maka seluruh elemen slice lain akan ter replace oleh element copykannya
            - slicing untuk mendapatkan element dari sebuah slice dan menentukan elemen dari index ke beberapa yang ingin kt dapatkan 
            - combining slicing and append
            - nacking array untuk menyimpan element pda slice, bukan slicenya sendiri

                        Bahasa Go mengimplementasikan slice sebagai sebuah struktur data yang disebut dengan slice header. Slice headerterdiri dari:
                        -Alamat memori/address dari backing array.
                        -Panjang dari slice yang bisa didapatkan dari fungsi len.
                        -Kapasitas dari slice yang bisa didapatkan dari fungsi cap.
            - cap function digunakan untuk mengetahui kapasitas dari sebuah array maupun slice
            - creating new backing array jika kt ingin mendapatkan element dari slice yg sudah ada, namun ingin juga
             membuat backing array yg baru kt dapat menggunakan fungsi append

=> Map : sama seperti array dan slice, map juga untuk menyimpab satu atau lebih data namun, map disimpan sebagai key:value pairs (pasang key dan value)
jika pernah belajar bahasa js dan php, maka ini mirip seperti tipe data object pd js dan tipe data associative array pda php.
            - looping with map menggunakan range loop
            - deleting value menggunakan fungsi delete
            - detecting a value dengan cara mengakses key dari mapnya lalu memberikan 2 variabel sebagai penampungnya
            - combining slice with map 

=> Aliase : merupakan fitur Go yang digunakan sebagai alternatif dari tipe data yang sudah ada
=> String in depth :