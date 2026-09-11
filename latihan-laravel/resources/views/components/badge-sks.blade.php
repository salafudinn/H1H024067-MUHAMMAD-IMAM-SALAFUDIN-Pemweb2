
@php
    $warna = $sks >= 3 ? 'bg-success' : 'bg-warning text-dark';
@endphp
<span class="badge {{ $warna }}">
    {{ $sks }} SKS
</span>