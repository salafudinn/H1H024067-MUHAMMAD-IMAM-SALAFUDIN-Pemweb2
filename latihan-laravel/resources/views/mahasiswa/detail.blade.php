@extends('layouts.app')

@section('judul', 'Detail Mahasiswa')

@section('konten')
<h1 class="h3 mb-4">Detail Mahasiswa</h1>

<table class="table table-bordered bg-white w-50">
    <tr><th>NIM</th><td>{{ $mahasiswa->nim }}</td></tr>
    <tr><th>Nama</th><td>{{ $mahasiswa->nama }}</td></tr>
    <tr><th>Program Studi</th><td>{{ $mahasiswa->programStudi->nama }}</td></tr>
    <tr><th>Angkatan</th><td>{{ $mahasiswa->angkatan }}</td></tr>
    <tr><th>IPK</th><td>{{ $mahasiswa->ipk }}</td></tr>
</table>

<h2 class="h4 mt-4 mb-3">Matakuliah yang Diambil</h2>
<table class="table table-striped bg-white">
    <thead>
        <tr>
            <th>Kode</th>
            <th>Nama Matakuliah</th>
            <th>SKS</th>
            <th>Nilai</th>
        </tr>
    </thead>
    <tbody>
        @forelse ($mahasiswa->matakuliah as $matakuliah)
        <tr>
            <td>{{ $matakuliah->kode }}</td>
            <td>{{ $matakuliah->nama }}</td>
            <td>{{ $matakuliah->sks }}</td>
            <td>{{ $matakuliah->pivot->nilai ?? '-' }}</td>
        </tr>
        @empty
        <tr><td colspan="4">Belum ada matakuliah yang diambil</td></tr>
        @endforelse
    </tbody>
</table>
@endsection