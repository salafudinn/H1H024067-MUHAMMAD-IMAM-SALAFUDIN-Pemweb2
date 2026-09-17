@extends('layouts.app')
@section('judul', 'Top 10 IPK Teknik Komputer')
@section('konten')
<h1 class="h3 mb-4">Top 10 IPK Teknik Komputer</h1>
<table class="table table-striped bg-white">
<thead>
<tr>
<th>NIM</th>
<th>Nama</th>
<th>Program Studi</th>
<th>Angkatan</th>
<th>IPK</th>
</tr>
</thead>
<tbody>
@forelse ($daftarMahasiswa as $mahasiswa)
<tr>
<td>{{ $mahasiswa->nim }}</td>
<td>{{ $mahasiswa->nama }}</td>
<td>{{ $mahasiswa->programStudi->nama }}</td>
<td>{{ $mahasiswa->angkatan }}</td>
<td>{{ $mahasiswa->ipk }}</td>
</tr>
@empty
<tr>
<td colspan="5" class="text-center">Belum ada data mahasiswa</td>
</tr>
@endforelse
</tbody>
</table>
@endsection
