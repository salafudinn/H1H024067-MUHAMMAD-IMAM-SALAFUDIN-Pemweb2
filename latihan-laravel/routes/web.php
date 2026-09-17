<?php
use App\Http\Controllers\MahasiswaController;
use Illuminate\Support\Facades\Route;
use App\Http\Controllers\MatakuliahController;
use App\Http\Controllers\MahasiswaWebController;

Route::get('/', function () {
    return view('welcome');
});

Route::get('/salam', function () {
    return 'Selamat datang di Pemrograman Web II';
});

route::get('/mahasiswa/{nim}', function(string $nim){
    return 'data mahasiswa dengan nim ' . $nim;
});

route::get('/matakuliah/{kode}', function(?string $kode=null){
    if ($kode == null){
        return 'menampilkan seluruh matakuliah';
    }
    return 'menampilkan matakuliah kode ' . $kode;
});

Route::get('/semester/{angka}', function (int $angka) {
    return 'Semester ke ' . $angka;
})->whereNumber('angka');

Route::get('/data-mahasiswa', [MahasiswaController::class,'index'])->name('mahasiswa.index');
Route::get('/data-mahasiswa/{nim}', [MahasiswaController::class,'show'])->name('mahasiswa.show');
Route::get('/cari-mahasiswa', [MahasiswaController::class, 'cari']);

Route::get('/matakuliah', [MatakuliahController::class, 'index'])->name('matakuliah.index');
Route::get('/matakuliah/{kode}', [MatakuliahController::class, 'show'])->name('matakuliah.show');
Route::get('/mahasiswa-data', [MahasiswaWebController::class, 'index'])->name('mahasiswa.data');
Route::get('/mahasiswa-data/{mahasiswa}', [MahasiswaWebController::class, 'show'])->name('mahasiswa.detail');
Route::get('/mahasiswa-top-ipk', [MahasiswaWebController::class, 'topIpk'])->name('mahasiswa.top-ipk');
Route::resource('mahasiswa', MahasiswaWebController::class);

