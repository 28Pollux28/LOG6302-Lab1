<?php
/*
    Test
*/
declare(strict_types=1);

namespace \Vendor\Package;

require 'path/to/app/Helpers/array_helpers.php';
use some\namespace\{ClassA, ClassB, ClassC as C};
use function App\Helpers\array_flatten as flatten, function App\Helpers\array_first as first;
use parent as parentAlias;
use self as selfAlias;
use static as staticAlias;

class ReturnTypeVariations
{
    private int|float $foo;
    public function functionName(int $arg1, $arg2): string
    {
        return 'foo';
    }

    public function anotherFunction(
        string $foo,
        string $bar,
        int $baz
    ): string {
        return 'foo';
    }
}