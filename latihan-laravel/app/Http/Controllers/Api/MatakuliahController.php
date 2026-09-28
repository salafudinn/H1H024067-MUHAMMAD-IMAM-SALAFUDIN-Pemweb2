<?php

namespace App\Http\Controllers\Api;

use App\Http\Controllers\Controller;
use App\Models\Matakuliah;
use App\Http\Requests\StoreMatakuliahRequest;
use App\Http\Requests\UpdateMatakuliahRequest;
use App\Http\Resources\MatakuliahResource;

class MatakuliahController extends Controller
{
    public function index()
    {
        return MatakuliahResource::collection(Matakuliah::all());
    }

    public function store(StoreMatakuliahRequest $request)
    {
 
        $matakuliah = Matakuliah::create($request->validated());
        return new MatakuliahResource($matakuliah);
    }

    public function show(Matakuliah $matakuliah)
    {
        return new MatakuliahResource($matakuliah);
    }

    public function update(UpdateMatakuliahRequest $request, Matakuliah $matakuliah)
    {
        $matakuliah->update($request->validated());
        return new MatakuliahResource($matakuliah);
    }

    public function destroy(Matakuliah $matakuliah)
    {
        $matakuliah->delete();
        return response()->json(['message' => 'Matakuliah berhasil dihapus'], 200);
    }
}
