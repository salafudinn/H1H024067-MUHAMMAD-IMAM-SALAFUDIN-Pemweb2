@extends('layouts.app')
@section('judul', 'Detail Matakuliah')

@section('konten')
<h1 class="h3 mb-4">Detail Matakuliah</h1>

<div class="card"> 
    <div class="card-body">
        @if($matakuliah)
            <p class="mb-1">Kode: <strong>{{ $matakuliah['kode'] }}</strong></p>
            <p class="mb-1">Nama Matakuliah: <strong>{{ $matakuliah['nama'] }}</strong></p>
            <p class="mb-0">Beban Studi: <x-badge-sks :sks="$matakuliah['sks']" /></p>
        @else
            <p class="mb-0 text-danger">Data matakuliah dengan kode <strong>{{ $kode }}</strong> tidak ditemukan.</p>
        @endif
    </div>
</div>
<a href="{{ route('matakuliah.index') }}" class="btn btn-secondary mt-3">Kembali</a>
@endsection