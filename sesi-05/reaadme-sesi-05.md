<h1>Nomor peserta : 149368582100-405 </h1>
<h1>Nama : Renny Kristina Salainti </h1>
<h1>sesi 5 </h1>

<h3>Tgl : 28/09/22</h3>

<p>1. Channels ini merupakan sebuah mekanisme untuk Goroutine saling berkomunikasi dengan Goroutine lainnya, komunikasi ini dimaksud adalah
   proses pengiriman dan pertukaran data antara Goroutine satu dengan yg lainnya.
   Untuk membuat sebuah Goroutine, kt memerlukan function make, yg merupakansebuah built in function dr bahasa Go.
   Contoh : <code>func main(){
   c := make(chan string) //variabel c dengan tipe data channel of string
   }</code>
   tanda <- merupakan operator dr channel untuk mengirim data dr Goroutine satu dengan lainnya.
   c <- mengirim data pda channel c.
   := <-c menerima data dr channel c dan data tersebut disimpan di dalam variabel result.
   proses pengiriman dan penerimaan data ini bersifat synchronous.

2. Defer & Exit, defer merupakan keyword pd bahasa Go digunakan untuk memanggil sebuah function
   yg dimana proses eksekusi akan ditahan hingga block sebuah function selesai. Sedangkan, function
   Exit ini berasal dr package os yg berguna untuk menghentikan suatu program secara paksa.

3. Error, Panic & Recover,
   Error : merupakan tipe data bahasa Go yg digunakan untuk me return sebuah error jika ada error yg terjadi.
   umumnya error akan direturn pda posisi terakhir dr nilai-nilai return sebuah function.
   Panic : digunakan untuk menampilkan stack trace error sekaligus akan menghentikan flow Goroutine
   (karena main() juga merupakan goroutine, maka behavior yg sana dilakukan) Setelah ada panic proses akan terhenti
   apapum itu setelah tidak di eksekusi kecuali sdh di defer sebelumnya, panic akan menampilkan pesan error di consol,
   sama seperti fmt.Println(), bentuknya berupa stack trace error, jdi sangat detail dan heboh/banyak.
   Recover :digunakan untuk menangkap error yang terjadi pd sebuah goroutine.