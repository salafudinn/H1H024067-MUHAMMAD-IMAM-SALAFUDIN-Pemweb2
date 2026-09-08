@extends('layouts.app')
@section('judul', 'Daftar Mahasiswa')
<x-kartu-info judul="Informasi">
Data pada halaman ini masih berupa array statis. Pada modul
berikutnya data akan diambil dari basis data.
</x-kartu-info>
@section('konten')
<h1 class="h3 mb-4">Daftar Mahasiswa</h1>
<table class="table table-bordered bg-white">
<thead>
<tr>
<th>NIM</th>
<th>Nama</th>
<th>Angkatan</th>
<th>Aksi</th>
</tr>
</thead>
<tbody>
@forelse ($daftarMahasiswa as $mahasiswa)
<tr>
<td>{{ $mahasiswa['nim'] }}</td>
<td>{{ $mahasiswa['nama'] }}</td>
<td>{{ $mahasiswa['angkatan'] }}</td>
<td>
<a href="{{ route('mahasiswa.show',
$mahasiswa['nim']) }}" class="btn btn-sm btn-primary">
Detail
</a>
</td>
</tr>
@empty
<tr>
<td colspan="4">Data belum tersedia</td>
</tr>
@endforelse
</tbody>
</table>
@endsection