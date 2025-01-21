package main

/*
Select Channel
- Kadang ada kasus dimana kita membuat beberapa channel, dan menjalankan beberapa Goroutine
- Lalu kita ingin mendapatkan data dari semua Channel tersebut
- Untuk melakukan hal tersebut, kita bisa menggunakan select channel di Go
- Dengan select channel, kita bisa memilih data tercepat dari beberapa Channel, jika data datang secara bersamaan di beberapa Channel, maka akan dipilih secara random
*/
