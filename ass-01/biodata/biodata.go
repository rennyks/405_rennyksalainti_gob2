package main

/*import package*/
import (
	"fmt"     //untuk menampilkan teks pda saat melakukan output
	"os"      // os tersimpan dalam bentuk array dengan pemisahnya adalah spasi
	"strconv" //untuk melakukan konversi
)

//inialisasi struct Students
type Students struct {
	nama 	string
	alamat 	string
	pekerjaan 	string
	alasan 	string
}

func main(){
	
	var argsRaw = os.Args //os.Args untuk mengambil data argument (package os perlu di-import terlebih dahulu)
	//fmt.Printf("%#v\n", argsRaw) //untuk mengecek argument

	var args = argsRaw[1] //pda data arrgsRaw dimasukan array 1 ke args

	num, err := strconv.Atoi(args) //strconv.Atoi() untuk konversi data string ke int
	_ = err //untuk menampung nilai kosong


	fmt.Println(args) //=> untuk mengetahui argument datanya 

	biodata(num) //untuk memanggil func biodata dengan parameter num
}	

	//fun biodata yang memiliki parameter arr dengan tipe datanya integer 
	func biodata(arr int) 	{
	var datapeserta = []Students{
		//datapeserta yang nanti ditampilkan sesuai urutan
		{nama: "renny", alamat: "manado", pekerjaan: "mahasiswa", alasan: "mempelajari hal baru"},
		{nama: "rendy", alamat: "bandung", pekerjaan: "mahasiswa", alasan: "mempelajari hal baru"},
		{nama: "renata", alamat: "jakarta", pekerjaan: "mahasiswa", alasan: "mempelajari hal baru"},
		{nama: "relin", alamat: "purwokerto", pekerjaan: "mahasiswa", alasan: "mempelajari hal baru "},
		{nama: "ryan", alamat: "bali", pekerjaan: "mahasiswa", alasan: "mempelajari hal baru"},
		{nama: "yuni", alamat: "semarang", pekerjaan: "mahasiswa", alasan: "mempelaj hal baru"},
	}
	//untuk menampilkan data nama, alamat, pekerjaan dan alasan yang disimpan dalam variabel datapeserta
	//%s digunakan untuk memformat data string.
	fmt.Printf(" Nama : %s\n Alamat : %s\n Pekerjaan : %s\n Berikan Alasan : %s\n",
	
	//datapesert = nama variabelnya, [arr-1] karena nanti inputan dimulai dari angka/no absen 1
	datapeserta[arr -1].nama, datapeserta[arr -1].alamat, datapeserta[arr -1].pekerjaan, datapeserta[arr -1].alasan)
}