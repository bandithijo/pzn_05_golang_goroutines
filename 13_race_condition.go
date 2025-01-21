package main

/*
Race Condition

Masalah Dengan Goroutine
- Saat kita menggunakan Goroutine, dia tidak hanya berjalan secara concurrent, tapi bisa parallel juga, karena bisa ada beberapa thread yang berjalan secara parallel
- Hal ini sangat berbahaya ketika kita melakukan manipulasi data variable yang sama oleh beberapa Goroutine secara bersamaan
- Hal ini bisa menyebabkan masalah yang namanya Race Condition
*/
