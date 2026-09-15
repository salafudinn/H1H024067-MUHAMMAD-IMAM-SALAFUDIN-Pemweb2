<?php

namespace Database\Seeders;
use App\Models\ProgramStudi;
use Illuminate\Database\Console\Seeds\WithoutModelEvents;
use Illuminate\Database\Seeder;

class ProgramStudiSeeder extends Seeder
{
    /**
     * Run the database seeds.
     */
    public function run(): void
    {
    $daftar = [
    ['kode' => 'TK', 'nama' => 'Teknik Komputer', 'jenjang' => 'S1'],
    ['kode' => 'IF', 'nama' => 'Informatika', 'jenjang' => 'S1'],
    ['kode' => 'TE', 'nama' => 'Teknik Elektro', 'jenjang' => 'S1'],
    ];
    foreach ($daftar as $item) {
    ProgramStudi::create($item);
    }
    }
}
