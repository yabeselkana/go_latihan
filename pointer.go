package main

import "fmt"

func main() {
    // Mendeklarasikan variabel
    angka := 42

    // Membuat pointer yang menunjuk ke variabel 'angka'
    var pointerAngka *int
    pointerAngka = &angka

    // Menampilkan nilai variabel menggunakan pointer
    fmt.Println("Nilai variabel angka:", angka)
    fmt.Println("Alamat memori variabel angka:", &angka)
    fmt.Println("Nilai yang disimpan di dalam pointer:", *pointerAngka)
}