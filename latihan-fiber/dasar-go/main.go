package main

import (
    "errors"
    "fmt"
)

type Mahasiswa struct {
    Nim      string
    Nama     string
    Angkatan int
    Ipk      float64
}

func cariMahasiswa(daftar []Mahasiswa, nim string) (Mahasiswa, error) {
    for i := 0; i < len(daftar); i++ {
        if daftar[i].Nim == nim {
            return daftar[i], nil
        }
    }
    return Mahasiswa{}, errors.New("mahasiswa tidak ditemukan")
}

func rerataIpk(daftar []Mahasiswa) float64 {
    if len(daftar) == 0 {
        return 0
    }
    var total float64
    for i := 0; i < len(daftar); i++ {
        total += daftar[i].Ipk
    }
    return total / float64(len(daftar))
}

func main() {
    daftar := []Mahasiswa{
        {Nim: "H1A123001", Nama: "Andi Prasetyo", Angkatan: 2023, Ipk: 3.45},
        {Nim: "H1A123002", Nama: "Bunga Lestari", Angkatan: 2023, Ipk: 3.80},
        {Nim: "H1A124003", Nama: "Citra Ramadhani", Angkatan: 2024, Ipk: 3.20},
    }

    ditemukan, err := cariMahasiswa(daftar, "H1A123002")
    if err != nil {
        fmt.Println("Galat:", err)
        return
    }
    fmt.Println("Ditemukan:", ditemukan.Nama)
    fmt.Printf("Rerata IPK: %.2f\n", rerataIpk(daftar))

    jumlahPerAngkatan := make(map[int]int)
    for i := 0; i < len(daftar); i++ {
        jumlahPerAngkatan[daftar[i].Angkatan] = jumlahPerAngkatan[daftar[i].Angkatan] + 1
    }
    fmt.Println("Jumlah per angkatan:", jumlahPerAngkatan)
}