package main

/*
sync.Mutex

Mutex (Mutual Exclusion)
- Untuk mengatasi masalah race condition tersebut, di Go teradapat sebuah struct bernama sync.Mutex
- Mutex bisa digunakan untuk melakukan locking dan unlocking, dimana ketika kita melakukan locking terhadap Mutex, maka tidak ada yang bisa melakukan locking lagi sampai kita melakukan unlock
- Dengan demikian, jika ada beberapa Goroutine melakukan lock terhadap Mutex, maka hanya 1 Goroutine yang diperbolehkan, setelah Goroutine tersebut melakukan unlock, baru Goroutine selanjutnya diperbolehkan melakukan lock lagi
- Ini sangat cocok sebagai solusi ketika ada masalah race condition yang sebelumnya kita hadapi
*/
