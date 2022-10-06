Nomor peserta : 149368582100-405 
Nama : Renny Kristina Salainti

Tgl : 05/10/22

1. Decoding & Parsing JSON Data
- membuat data JSON sederhana menggunakan tanda bcktick
- kita menggunakan function json.Ummarshal untuk mendecode data JSON kepada struct Employee
argumen pertama dari function json.Ummarshal menerima sebuah nilai dengan tipe data slice of byte, di argumen pertama
kt meletakan data JSON tetapi harus kt ubah dahulu menjadi tipe data slice of byte,
pda argumen kedua, kt meletakan pointer dari variabel result agar setelah data JSON berhasil di decode, datanya disimpan
kepada variabel result.
- kita juga bisa men decode data JSON kepada sebuah tipe data map dan emty interface, slice of struct.

2. URL Parsing
- data string url bisa dikonversi kedalam bentuk url.URL, dengan tipe tersebut akan ada banyak informasi yang bisa kita manfaatkan,
diantaranya adlh jenis protokol yg digunakan path yang diakses, query dll.
- function url.Parse digunakan untuk Parsing ke bentuk url, mengembalikan 2 data, variabel objek bertipe url.URL dan error(jika ada)
- melalui variabel objek tersebut pengaksesan informasi url akan menjadi lebih mudah, contoh : nama host bisa didapatkan lewat u.Host, protokol lewat u.Scheme, dll.
- query yang ada pada url akan otomatis di parsing juga, menjadi bentuk map(string)[]string, dengan key adlh nama elemen query dan value array string yg berisikan value elemen query

3. Swaggo
- Swagger UI memberikan cara yang nyaman bagi consumer untuk menjelajahi API
- Selain itu, API berkembang dari waktu ke waktu dan dokumentasi harus mencerminkan perubahan yang sesuai
(Jumlah bug yang muncul karena komunikasi perubahan API yang tidak tepat (atau tidak ada) terlalu tinggi!
- Dengan Swagger, memperbarui / memelihara dokumentasi API sangat mudah - pengembang 
hanya perlumenambahkan / mengubah anotasi dalam kode, dan perubahannya digabungkan saat dokumen API dibuat berikutnya.
- Swaggo dan go-swagger adalah dua framework paling populer yang tersedia untuk menghasilkan dokumen.
Install library yang kita perlukan dengan menjalankan perintah berikut :
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/swag/http-swagger
go get -u github.com/alecthomas/template
