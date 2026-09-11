<?php

namespace App\View\Components;

use Closure;
use Illuminate\Contracts\View\View;
use Illuminate\View\Component;

class BadgeSks extends Component
{
    public int $sks;
    public function __construct(int $sks) 
    {
        $this->sks = $sks;
    }

    public function render(): View|Closure|string
    {
        return view('components.badge-sks');
    }
}