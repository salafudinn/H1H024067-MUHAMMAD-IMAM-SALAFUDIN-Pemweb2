<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class MatakuliahController extends Controller
{
    private $matakuliah = [
        ['kode' => 'MK001', 'nama' => 'Pemrograman Web', 'sks' => 3],
        ['kode' => 'MK002', 'nama' => 'Struktur Data', 'sks' => 4],
        ['kode' => 'MK003', 'nama' => 'Pendidikan Pancasila', 'sks' => 2],
        ['kode' => 'MK004', 'nama' => 'Basis Data', 'sks' => 3],
        ['kode' => 'MK005', 'nama' => 'Bahasa Inggris', 'sks' => 2],
    ];

    public function index(Request $request)
    {
        $search = $request->query('q', '');
        $data = $this->matakuliah;

        if (!empty($search)) {
            $data = array_filter($this->matakuliah, function ($item) use ($search) {
                return stripos($item['nama'], $search) !== false || stripos($item['kode'], $search) !== false;
            });
        }

        return view('matakuliah.index', [
            'daftarMatakuliah' => $data,
            'katakunci' => $search
        ]);
    }

    public function show(string $kode)
    {
        $mk = null;
        foreach ($this->matakuliah as $item) {
            if ($item['kode'] === $kode) {
                $mk = $item;
                break;
            }
        }

        if (!$mk) {
            abort(404);
        }

        return view('matakuliah.show', [
            'matakuliah' => $mk,
            'kode' => $kode
        ]);
    }
}