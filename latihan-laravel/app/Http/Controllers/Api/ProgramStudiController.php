<?php

namespace App\Http\Controllers\Api;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use App\Models\Mahasiswa;

class ProgramStudiController extends Controller
{
    /**
     * Display a listing of the resource.
     */
    public function mahasiswa(Request $request, $id)
    {
        // Mencari mahasiswa dengan program_studi_id tersebut lalu dipaginasi 
        // Menggunakan parameter 'per_halaman' dari request, defaultnya 10
        $perHalaman = $request->query('per_halaman', 10);
        $mahasiswa = Mahasiswa::where('program_studi_id', $id)->paginate($perHalaman);
        
        // Kembalikan dalam bentuk response JSON
        return response()->json($mahasiswa);
    }
    public function index()
    {
        //
    }

    /**
     * Store a newly created resource in storage.
     */
    public function store(Request $request)
    {
        //
    }

    /**
     * Display the specified resource.
     */
    public function show(string $id)
    {
        //
    }

    /**
     * Update the specified resource in storage.
     */
    public function update(Request $request, string $id)
    {
        //
    }

    /**
     * Remove the specified resource from storage.
     */
    public function destroy(string $id)
    {
        //
    }
}
