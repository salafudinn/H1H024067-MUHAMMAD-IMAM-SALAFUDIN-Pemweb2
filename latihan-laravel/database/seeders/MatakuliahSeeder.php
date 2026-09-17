<?php

namespace Database\Seeders;

use App\Models\Matakuliah;
use Illuminate\Database\Seeder;

class MatakuliahSeeder extends Seeder
{
    public function run(): void
    {
        $daftar = [
            ['kode' => 'TK101', 'nama' => 'Algoritma dan Pemrograman', 'sks' => 3, 'semester' => 1],
            ['kode' => 'TK102', 'nama' => 'Struktur Data', 'sks' => 3, 'semester' => 2],
            ['kode' => 'TK201', 'nama' => 'Basis Data', 'sks' => 3, 'semester' => 3],
            ['kode' => 'TK202', 'nama' => 'Pemrograman Web II', 'sks' => 3, 'semester' => 4],
            ['kode' => 'TK301', 'nama' => 'Sistem Mikrokontroller', 'sks' => 3, 'semester' => 5],
        ];

        foreach ($daftar as $item) {
            Matakuliah::create($item);
        }
    }
}