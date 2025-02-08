package main

/*
sync.Once
- Once adalah fitur di Go yang bisa kita gunakan untuk memastikan bahwa sebuah function di eksekusi hanya sekali
- Jadi berapa banyak pun Goroutine yang mengakses, bisa dipastikan bahwa Goroutine yang pertama yang bisa mengeksekusi function nya
- Goroutine yang lain akan di hiraukan, artinya function tidak akan dieksekusi lagi
*/
