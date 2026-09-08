<?php
namespace App\Http\Controllers;
use Illuminate\Http\Request;

class MahasiswaController extends Controller
{
    public function index()
    {
        $daftarMahasiswa = [
            ['nim' => 'H1A123001', 'nama' => 'Andi Prasetyo', 'angkatan' => 2023],
            ['nim' => 'H1A123002', 'nama' => 'Bunga Lestari', 'angkatan' => 2023],
            ['nim' => 'H1A123003', 'nama' => 'Citra Ramadhani', 'angkatan' => 2024],
        ];

        return view('mahasiswa.index', ['daftarMahasiswa' => $daftarMahasiswa]);
    }

    public function show(string $nim)
    {
        return view('mahasiswa.show', ['nim' => $nim]);
    }

    public function cari(Request $request)  {
        $kataKunci = $request->query('q', '');

        return response()->json([
            'kata_kunci' => $kataKunci,
            'metode' => $request->method(),
            'path' => $request->path(),
        ]);
    }
}