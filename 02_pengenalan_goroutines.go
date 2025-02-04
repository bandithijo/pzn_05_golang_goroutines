package main

/*
Pengenalan Goroutine
- Goroutine adalah sebuah thread ringan yang dikelola ole Go Runtime
- Ukuran Goroutine sangat kecil, sekitar 2kb, jauh lebih kecil dibandingkan Thread yang bisa sampai 1mb atau 1000kb
- Namun tidak seperti Thread yang berjalan parallel, goroutine berjalan secara concurrent

Cara Kerja Goroutine
- Sebenarnya, Goroutine dijalankan oleh Go Scheduler dalam Thread, dimana jumlah thread nya sebanyak GOMAXPROCS (biasanya sejumlah core CPU)
- Jadi sebenarnya tidak bisa dibilang Goroutine itu pengganti Thread, karena Goroutine sendiri berjalan di atas Thread
- Namun yang mempermudah kita adalah, kita tidak perlu melakukan manajemen Thread secara manual, semua sudah diatur oleh Go Scheduler

Cara Kerja Go Scheduler
Dalam Go Scheduler, kita akan mengenal beberapa teknologi:
- G:Goroutine
- M:Thread (Machine)
- P:Processor
*/
