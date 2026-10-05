<?php

namespace App\Http\Middleware;

use Closure;
use Illuminate\Http\Request;
use Symfony\Component\HttpFoundation\Response;

class PeranAdmin
{
    /**
     * Handle an incoming request.
     */
    public function handle(Request $request, Closure $next): Response
    {
        // Pastikan pengguna sudah login dan punya peran 'admin'
        if ($request->user() && $request->user()->peran === 'admin') {
            return $next($request);
        }

        // Jika bukan admin, tolak akses
        return response()->json([
            'pesan' => 'Akses ditolak. Anda bukan admin.',
        ], 403);
    }
}
