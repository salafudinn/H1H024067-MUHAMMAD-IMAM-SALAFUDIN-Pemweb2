@extends('layouts.app')

@section('judul', 'Portal Aplikasi')

@section('konten')
<div class="row justify-content-center">
    <div class="col-md-8">
        <div class="text-center mb-5">
            <h1 class="display-4 fw-bold">Selamat Datang</h1>
            <p class="display-7">Nama : Muhammad Imam Salafudin <br>
                NIM : H1H024067 <br>
                Mata Kuliah : Praktikum Pemrograman Web II
            </p>
        </div>

        <div class="row g-4">
            <!-- Modul 1 -->
            <div class="col-md-6">
                <div class="card h-100 shadow-sm border-0">
                    <div class="card-body text-center p-4">
                        <div class="mb-3">
                            <span class="badge bg-primary rounded-pill px-3 py-2">Modul 1</span>
                        </div>
                        <h5 class="card-title fw-bold">Data Mahasiswa Lama</h5>
                        <p class="card-text text-muted mb-4">Menampilkan seluruh data mahasiswa menggunakan controller dasar.</p>
                        <a href="{{ route('mahasiswa.index') }}" class="btn btn-primary px-4 py-2 rounded-pill">Buka Halaman</a>
                    </div>
                </div>
            </div>
            
            <!-- Modul 2 -->
            <div class="col-md-6">
                <div class="card h-100 shadow-sm border-0">
                    <div class="card-body text-center p-4">
                        <div class="mb-3">
                            <span class="badge bg-primary rounded-pill px-3 py-2">Modul 2</span>
                        </div>
                        <h5 class="card-title fw-bold">Data Mata Kuliah</h5>
                        <p class="card-text text-muted mb-4">Menampilkan daftar semua mata kuliah.</p>
                        <a href="{{ route('matakuliah.index') }}" class="btn btn-primary px-4 py-2 rounded-pill">Buka Halaman</a>
                    </div>
                </div>
            </div>  

            <!-- Modul 3: Eloquent & Pagination -->
            <div class="col-md-6">
                <div class="card h-100 shadow-sm border-0">
                    <div class="card-body text-center p-4">
                        <div class="mb-3">
                            <span class="badge bg-primary rounded-pill px-3 py-2">Modul 3</span>
                        </div>
                        <h5 class="card-title fw-bold">Data Mahasiswa (Eloquent)</h5>
                        <p class="card-text text-muted mb-4">Menampilkan daftar mahasiswa dengan relasi Program Studi menggunakan pagination.</p>
                        <a href="{{ route('mahasiswa.data') }}" class="btn btn-primary px-4 py-2 rounded-pill">Buka Halaman</a>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>
@endsection