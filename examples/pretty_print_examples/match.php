<?php
// Example of using match expression in PHP

function getFruitColor(string $fruit): string {
    return match ($fruit) {
        'apple' => 'red',
        'banana' => 'yellow',
        'cherry' => 'red',
        'grape' => 'purple',
        'orange' => 'orange',
        default => 'unknown color',
    };
}

// Test cases
$fruits = ['apple', 'banana', 'cherry', 'grape', 'orange', 'kiwi'];

foreach ($fruits as $fruit) {
    echo "The color of $fruit is " . getFruitColor($fruit) . ".\n";
}
?>
