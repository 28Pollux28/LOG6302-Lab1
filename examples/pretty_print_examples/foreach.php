<?php
// 1. Basic foreach Loop (Value Only)
$fruits = ['apple', 'banana', 'cherry'];

echo "Basic foreach Loop (Value Only):\n";
foreach ($fruits as $fruit) {
    echo "Fruit: $fruit\n";
}

echo "\n";

// 2. foreach Loop with Key and Value
$fruitsWithKeys = ['a' => 'apple', 'b' => 'banana', 'c' => 'cherry'];

echo "foreach Loop with Key and Value:\n";
foreach ($fruitsWithKeys as $key => $value) {
    echo "Key: $key; Fruit: $value\n";
}

echo "\n";

// 3. foreach Loop with Objects
class Fruit {
    public $name;
    public $color;

    public function __construct($name, $color) {
        $this->name = $name;
        $this->color = $color;
    }
}

$fruitObjects = [
    new Fruit('apple', 'red'),
    new Fruit('banana', 'yellow'),
    new Fruit('cherry', 'red')
];

echo "foreach Loop with Objects:\n";
foreach ($fruitObjects as $fruit) {
    echo "Fruit: {$fruit->name}, Color: {$fruit->color}\n";
}

echo "\n";

// 4. foreach Loop with References
$fruitsToModify = ['apple', 'banana', 'cherry'];

echo "foreach Loop with References (Before Modification):\n";
print_r($fruitsToModify);

foreach ($fruitsToModify as &$fruit) {
    $fruit = strtoupper($fruit);
}

echo "foreach Loop with References (After Modification):\n";
print_r($fruitsToModify);

echo "\n";

// 5. foreach Loop with Nested Arrays
$nestedFruits = [
    'red' => ['apple', 'cherry'],
    'yellow' => ['banana', 'pineapple']
];

echo "foreach Loop with Nested Arrays:\n";
foreach ($nestedFruits as $color => $fruitList) {
    echo "Color: $color\n";
    foreach ($fruitList as $fruit) {
        echo "  Fruit: $fruit\n";
    }
}

// foreach Loop with Key and Value using Colon Syntax
$fruitsWithKeys = ['a' => 'apple', 'b' => 'banana', 'c' => 'cherry'];

echo "foreach Loop with Key and Value (Colon Syntax):\n";
foreach ($fruitsWithKeys as $key => $value):
    echo "Key: $key; Fruit: $value\n";
endforeach;

?>
