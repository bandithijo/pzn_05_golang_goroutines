package main

/*
sync.Map

Map
- Go memiliki sebuah struct bernama sync.Map
- Map ini mirip Go map, namun yang membedakan, Map ini aman untuk menggunakan concurrent menggunakan Goroutine
- Ada beberapa function yang bisa kita gunakan di Map:
  - Store(key, value) | untuk menyimpan data ke Map
  - Load(key) | untuk mengambil data dari Map menggunakan key
  - Delete(key) | untuk menghapus data di Map menggunakan key
  - Range(function(key, value)) | digunakan untuk melakukan iterasi seluruh data di Map
*/
