<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

return new class extends Migration
{
    /**
     * Run the migrations.
     */
    public function up(): void
    {
    Schema::create('mahasiswas', function (Blueprint $table) {
    $table->id();
    $table->foreignId('program_studi_id')->constrained('program_studis')->cascadeOnDelete();
    $table->string('nim', 20)->unique();
    $table->string('nama', 100);
    $table->string('email', 100)->unique();
    $table->year('angkatan');
    $table->decimal('ipk', 4, 2)->default(0);
    $table->boolean('aktif')->default(true);
    $table->timestamps();
    });
    }
    /**
     * Reverse the migrations.
     */
    public function down(): void
    {
        Schema::dropIfExists('mahasiswas');
    }
};
