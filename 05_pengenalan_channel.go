package main

/*
Pengenalan Channel
- Channel adalah tempat komunikasi secara synchronous yang bisa dilakukan oleh Goroutine
- Di Channel terdapat pengirim dan penerima, biasanya pengirim dan penerima adalah Goroutine yang berbeda
- Saat melakukan pengiriman data ke Channel, Goroutine akan ter-block, sampai ada yang menerima data tersebut
- Maka dari itu, Channel disebut sebagai alat komunikasi synchronous (blocking)
- Channel cocok sekali sebagai alternatif seperti mekanisme async await yang terdapat di beberapa bahasa pemrograman lain

Karakteristik Channel
- Secara default Channel hanya bisa menampung satu data, jika kita ingin menambahkan data lagi, harus menunggu data yang ada di Channel diambil
- Channel hanya bisa menerima satu jenis data.
- Channel bisa diambil dari lebih dari satu Goroutine
- Channel harus di close jika tidak digunakan, atau bisa menyebabkan memory leak
*/
