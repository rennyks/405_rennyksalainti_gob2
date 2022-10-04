Nomor peserta : 149368582100-405 
Nama : Renny Kristina Salainti

Tgl : 05/10/22

A. Database
1) PostgreSQL 
sebuah relational database manajemen system (RDBMS) yang dikembangkan oleh tim relawan yang ada di seluruh dunia yang
bersifat open source, artinya siapa saja bisa mengembangkannya. PostgreSQL tidak dikelola oleh perusahaan atau badan swasta
lainnya, sehingga source code (kode program) yang tersedia bisa di dapatkan secara gratis.
- installing SQL driver "go get -u guthub.com/lib/"
- creating database with query
- connecting to database
- CRUD

B. GORM
Gorm adalah ORM yang cukup populer untuk bahasa Go, yang dimana Gorm telah menyediakan berbagai
macam fitur seperti auto migration, eager loading, association, query method, dan lain-lain.
Dengan menggunakan Gorm, maka akan dapat mempercepat pengembangan aplikasi kita karena fitur-fitur
yang telah disediakan oleh Gorm. Untuk menginstall Gorm, maka kita perlu menjalankan 
perintah: go get -u gorm.io/gorm.
- declaring models
- association
- connecting to database and table migration "go get gorm.io/driver/Postgre"
- CRUD, HOOKS, Eager loading