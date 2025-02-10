package main

/*
time.Ticker

- Ticker adalah representasi kejadian yang berulang
- Ketika waktu Ticker sudah expire, maka event akan dikirim ke dalam channel
- Untuk membuat Ticker, kita bisa menggunakan time.NewTicker(duration)
- Untuk menghentikan Ticker, kita bisa menggunakan Ticker.Stop()

time.Tick()
- Kadang kita tidak butuh data Ticker nya, kita hanya butuh data channel nya saja
- Jika demikian, kita bisa menggunakan function timer.Tick(duration), function ini tidak akan mengembalikan Ticker, hanya mengembalikan channel timer nya saja
*/
