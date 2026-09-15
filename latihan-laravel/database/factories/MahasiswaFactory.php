<?php

namespace Database\Factories;

use App\Models\Mahasiswa;
use Illuminate\Database\Eloquent\Factories\Factory;

/**
 * @extends Factory<Mahasiswa>
 */
class MahasiswaFactory extends Factory
{
    /**
     * Define the model's default state.
     *
     * @return array<string, mixed>
     */
    public function definition(): array
    {
    return [
    'program_studi_id' => 1,
    'nim' => 'H1A' . fake()->unique()->numberBetween(100000, 999999),
    'nama' => fake('id_ID')->name(),
    'email' => fake()->unique()->safeEmail(),
    'angkatan' => fake()->numberBetween(2021, 2025),
    'ipk' => fake()->randomFloat(2, 2.50, 4.00),
    'aktif' => fake()->boolean(85),
    ];
    }
}
