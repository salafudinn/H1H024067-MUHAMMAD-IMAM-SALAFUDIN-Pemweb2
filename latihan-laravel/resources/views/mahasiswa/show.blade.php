@extends('layouts.app')
@section('judul', 'Detail Mahasiswa')
@section('konten')
<h1 class="h3 mb-4">Detail Mahasiswa</h1>
<div class="card">
<div class="card-body">
<p class="mb-0">NIM yang diminta: <strong>{{ $nim }}</strong></p>
</div>
</div>
<a href="{{ route('mahasiswa.index') }}" class="btn btn-secondary
mt-3">Kembali</a>
@endsection