<?php

namespace Database\Seeders;
use App\Models\Matakuliah;
use App\Models\Mahasiswa;
use App\Models\User;
use Illuminate\Database\Console\Seeds\WithoutModelEvents;
use Illuminate\Database\Seeder;

class DatabaseSeeder extends Seeder
{
    use WithoutModelEvents;

    /**
     * Seed the application's database.
     */
    public function run(): void
{
    $this->call([
        ProgramStudiSeeder::class,
        MatakuliahSeeder::class,
    ]);

    Mahasiswa::factory()->count(30)->create();

    // Attach courses to some students as an example for the pivot table
    $mahasiswa = Mahasiswa::first();
    if ($mahasiswa) {
        $matakuliahs = Matakuliah::take(3)->get();
        foreach ($matakuliahs as $mk) {
            $mahasiswa->matakuliah()->attach($mk->id, ['nilai' => rand(70, 100) + rand(0, 99) / 100]);
        }
    }
}
    
}
