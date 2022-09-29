<h2>Nomor peserta : 149368582100-405 </h2>
<h2>Nama : Renny Kristina Salainti </h2>

<h3>sesi 5 </h3>
<h3>Tgl : 28/09/22</h3>

<p>1. Channels ini merupakan sebuah mekanisme untuk Goroutine saling berkomunikasi dengan Goroutine lainnya, komunikasi ini dimaksud adalah
   proses pengiriman dan pertukaran data antara Goroutine satu dengan yg lainnya.
   Untuk membuat sebuah Goroutine, kt memerlukan function make, yg merupakansebuah built in function dr bahasa Go.</p>
   <p>Contoh : <code>func main(){
   c := make(chan string) //variabel c dengan tipe data channel of string
   }</code></p>
   <p>tanda <code><-</code> merupakan operator dr channel untuk mengirim data dr Goroutine satu dengan lainnya.
   <p><code>c <-</code> mengirim data pda channel c.</p>
   <p><code>:= <-c</code> menerima data dr channel c dan data tersebut disimpan di dalam variabel result.
   proses pengiriman dan penerimaan data ini bersifat synchronous.</p>

<p>2. Defer & Exit, defer merupakan keyword pd bahasa Go digunakan untuk memanggil sebuah function
   yg dimana proses eksekusi akan ditahan hingga block sebuah function selesai. Sedangkan, function
   Exit ini berasal dr package os yg berguna untuk menghentikan suatu program secara paksa.</p>

<p>Error, Panic & Recover,</p>
   <p>Error : merupakan tipe data bahasa Go yg digunakan untuk me return sebuah error jika ada error yg terjadi.
   umumnya error akan direturn pda posisi terakhir dr nilai-nilai return sebuah function.</p>
   <p>Panic : digunakan untuk menampilkan stack trace error sekaligus akan menghentikan flow Goroutine
   (karena main() juga merupakan goroutine, maka behavior yg sana dilakukan) Setelah ada panic proses akan terhenti
   apapum itu setelah tidak di eksekusi kecuali sdh di defer sebelumnya, panic akan menampilkan pesan error di consol,
   sama seperti fmt.Println(), bentuknya berupa stack trace error, jdi sangat detail dan heboh/banyak.</p>
   <p>Recover :digunakan untuk menangkap error yang terjadi pd sebuah goroutine.</p>
