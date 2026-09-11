<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class MatakuliahController extends Controller
{
    // Data dummy array minimal 5 matakuliah
    private $matakuliah = [
        ['kode' => 'MK001', 'nama' => 'Pemrograman Web', 'sks' => 3],
        ['kode' => 'MK002', 'nama' => 'Struktur Data', 'sks' => 4],
        ['kode' => 'MK003', 'nama' => 'Pendidikan Pancasila', 'sks' => 2],
        ['kode' => 'MK004', 'nama' => 'Basis Data', 'sks' => 3],
        ['kode' => 'MK005', 'nama' => 'Bahasa Inggris', 'sks' => 2],
    ];

    public function index(Request $request)
    {
        $data = $this->matakuliah;
        
        // Fitur pencarian sederhana dari query string
        $search = $request->query('search');
        if ($search) {
            $data = array_filter($data, function($item) use ($search) {
                return stripos($item['nama'], $search) !== false;
            });
        }

        return view('matakuliah.index', ['matakuliah' => $data, 'search' => $search]);
    }

    public function show($kode)
    {
        // Mencari matakuliah berdasarkan kode
        $mk = collect($this->matakuliah)->firstWhere('kode', $kode);
        
        if (!$mk) {
            abort(404);
        }

        return view('matakuliah.show', ['mk' => $mk]);
    }
}