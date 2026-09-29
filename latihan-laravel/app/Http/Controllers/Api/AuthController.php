<?php

namespace App\Http\Controllers\Api;

use App\Http\Controllers\Controller;
use App\Models\User;
use Illuminate\Http\JsonResponse;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Hash;
use Illuminate\Validation\Rules\Password;

class AuthController extends Controller
{
    public function register(Request $request): JsonResponse
    {
        $data = $request->validate([
            'name' => ['required', 'string', 'max:100'],
            'email' => ['required', 'email', 'max:100', 'unique:users,email'],
            'password' => ['required', 'confirmed', Password::min(8)->letters()->numbers()],
        ]);

        $data['password'] = Hash::make($data['password']);
        $data['peran'] = 'mahasiswa';

        $pengguna = User::create($data);

        $token = $pengguna->createToken('token-perangkat')->plainTextToken;

        return response()->json([
            'sukses' => true,
            'pesan' => 'Registrasi berhasil',
            'data' => [
                'pengguna' => [
                    'id' => $pengguna->id,
                    'nama' => $pengguna->name,
                    'email' => $pengguna->email,
                    'peran' => $pengguna->peran,
                ],
                'token' => $token,
            ],
        ], 201);
    }

    public function login(Request $request): JsonResponse
    {
        $data = $request->validate([
            'email' => ['required', 'email'],
            'password' => ['required', 'string'],
        ]);

        $pengguna = User::where('email', $data['email'])->first();

        if ($pengguna === null || Hash::check($data['password'], $pengguna->password) === false) {
            return response()->json([
                'sukses' => false,
                'pesan' => 'Email atau kata sandi tidak sesuai',
            ], 401);
        }

        $kemampuan = $pengguna->peran === 'admin'
            ? ['mahasiswa:baca', 'mahasiswa:tulis']
            : ['mahasiswa:baca'];

        $token = $pengguna->createToken('token-perangkat', $kemampuan)->plainTextToken;

        return response()->json([
            'sukses' => true,
            'pesan' => 'Login berhasil',
            'data' => [
                'pengguna' => [
                    'id' => $pengguna->id,
                    'nama' => $pengguna->name,
                    'email' => $pengguna->email,
                    'peran' => $pengguna->peran,
                ],
                'token' => $token,
            ],
        ]);
    }

    public function profil(Request $request): JsonResponse
    {
        $pengguna = $request->user();

        return response()->json([
            'sukses' => true,
            'data' => [
                'id' => $pengguna->id,
                'nama' => $pengguna->name,
                'email' => $pengguna->email,
                'peran' => $pengguna->peran,
                'kemampuan' => $pengguna->currentAccessToken()->abilities,
            ],
        ]);
    }

    public function logout(Request $request): JsonResponse
    {
        $request->user()->currentAccessToken()->delete();

        return response()->json([
            'sukses' => true,
            'pesan' => 'Logout berhasil',
        ]);
    }

    public function logoutSemua(Request $request): JsonResponse
    {
        $request->user()->tokens()->delete();

        return response()->json([
            'sukses' => true,
            'pesan' => 'Seluruh sesi perangkat telah diakhiri',
        ]);
    }
}
