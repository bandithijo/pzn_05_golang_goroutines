package main

/*
Buffered Channel
- Seperti yang dijelaskan sebelumnya, bahwa secara default Channel itu hanya bisa menerima 1 data
- Artinya jika kita menambah data ke-2, maka kita akan diminta menunggu sampai data ke-1 ada yang mengambil
- Kadang-kadang ada kasus dimana pengirim lebih cepat dibanding penerima, dalama hal ini jika kita menggunakan Channel, maka otomatis pengirim akan ikut lambat juga
- Untungnya ada Buffered Channel, yaitu buffer yang bisa digunakan untuk menampung data antrian di Channel

Buffer Capacity
- Kita bebas memasukkan berapa jumlah kapasitas antrian di dalam buffer
- Jika kita set misal 5, artinya kita bisa menerima 5 data di buffer
- Jika kita mengirim data ke-6, maka kita diminta untuk menunggu sampai buffer ada yang kosong
- Ini cocok sekali ketika memang Goroutine yang menerima data lebih lambat dari yang menerima data
*/
