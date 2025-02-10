package main

/*
GOMAXPROCS
- Sebelumnya di awal, kita sudah bahas bahwa Goroutine itu sebenarnya dijalankan di dalam Thread
- Pertanyannya, seberapa banyak Thread yang ada di Go ketika aplikasi kita berjalan?
- Untuk mengetahui berapa jumlah Thread, kita bisa menggunakan GOMAXPROCS, yaitu sebuah function di package runtime yang bisa kita gunakan untuk mengubah jumlah Thread atau mengambil jumlah Thread
- Secara default, jumlah thread di Go itu sebanyak jumlah CPU di komputer kita
- Kita juga bisa melihat berapa jumlah CPU kita dengan menggunakan function runtime.NumCpu()
*/
