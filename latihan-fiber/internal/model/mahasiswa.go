package model

type Mahasiswa struct {
    ID       int     `json:"id"`
    Nim      string  `json:"nim"`
    Nama     string  `json:"nama"`
    Angkatan int     `json:"angkatan"`
    Ipk      float64 `json:"ipk"`
}

var Daftar = []Mahasiswa{
    {ID: 1, Nim: "H1A123001", Nama: "Andi Prasetyo", Angkatan: 2023, Ipk: 3.45},
    {ID: 2, Nim: "H1A123002", Nama: "Bunga Lestari", Angkatan: 2023, Ipk: 3.80},
    {ID: 3, Nim: "H1A124003", Nama: "Citra Ramadhani", Angkatan: 2024, Ipk: 3.20},
}