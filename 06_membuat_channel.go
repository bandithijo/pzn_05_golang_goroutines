package main

/*
Membuat Channel
- Channel di Go direpresentasikan dengan tipe data chan
- Untuk membuat Channel sangat mudah, kita bisa menggunakan make(), mirip ketika membuat map
- Namun saat pembuatan Channel, kita harus tentukan tipe data apa yang bisa dimasukkan ke dalam Channel tersebut

Mengirim dan Menerima Data dari Channel
- Seperti yang sudah dibahas sebelumnya, channel bisa digunakan untuk mengirim dan menerima data
- Untuk mengirim data, kita bisa gunakan kode: channel <- data
- Sedangkan untuk menerima data, bisa gunakan kode: data <- channel
- Jika selesai, jangan lupa untuk menutup channel menggunakan function close()
*/
