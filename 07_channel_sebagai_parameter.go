package main

/*
Channel Sebagai Parameter
- Dalam kenyataan pembuatan palikasi, seringnya kita akan mengirim channel ke function lain via parameter
- Sebelumnya kita tahu bahkan di Go by default, parameter adalah pass by value, artinya value akan diduplikasi lalu dikirim ke function parameter, sehingga jika kita ingin mengirim data asli, kita bisa gunakan pointer (agar pass by reference)
- Berbeda dengan Channel, kita tidak perlu melakukan hal tersebut
*/
