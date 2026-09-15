<?php

namespace Database\Seeders;
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
    $this->call(ProgramStudiSeeder::class);
    Mahasiswa::factory()->count(30)->create();
    }
}
