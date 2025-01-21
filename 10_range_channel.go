package main

/*
Range Channel
- Kadang-kadang ada kasus sebuah Channel dikirim data secara terus menerus oleh pengirim
- Dan kadang tidak jelas kapan Channel tersebut akan berhenti menerima data
- Salah satu yang bisa kita lakukan adalah dengan menggunakan perulangan range ketika menerima data dari Channel
- Ketika sebuah Channel di close(), maka secara otomatis perulangan tersebut akan berhenti
- Ini lebih sederhana dari pada kita melakukan pengecekan Channel secara manual
*/
