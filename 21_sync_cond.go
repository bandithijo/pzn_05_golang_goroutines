package main

/*
sync.Cond

Cond
- Cond adalah implementasi locking berbasis kondisi
- Cond membutuhkan Locker (bisa menggunakan Mutex atau RWMutex) untuk implementasi locking nya, namun berbeda dengan Locker, biasanya di Cond terdapat function Wait() untuk menunggu apakah perlu menunggu atau tidak
- Function Signal() bisa digunakan untuk memberi tahu sebuah Goroutine agar tidak perlu menunggu lagi, sedangkan function Broadcast() digunakan untuk memberi tahu semua Goroutine agar tidak peru menunggu lagi
- Untuk membuat Cond, kita bisa menggunakan function sync.NewCond(Locker)
*/
