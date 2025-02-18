<?php
// Standard syntax examples

// Example of a while loop
echo "While Loop (Standard Syntax):\n";
$i = 0;
while ($i < 5) {
    echo "i is $i\n";
    $i++;
}

// Example of a for loop
echo "\nFor Loop (Standard Syntax):\n";
for ($j = 0; $j < 5; $j++) {
    echo "j is $j\n";
}

// Example of an if statement
echo "\nIf Statement (Standard Syntax):\n";
$x = 10;
if ($x < 20) {
    echo "x is less than 20\n";
} elseif ($x == 20) {
    echo "x is equal to 20\n";
} else {
    echo "x is greater than 20\n";
}

// Example of a switch statement
echo "\nSwitch Statement (Standard Syntax):\n";
$fruit = "apple";
switch ($fruit) {
    case "banana":
        echo "Banana is good!\n";
        break;
    case "apple":
        echo "Apple is good too!\n";
        break;
    default:
        echo "I have never tried that fruit.\n";
}

// Example of a try-catch block
echo "\nTry-Catch Block (Standard Syntax):\n";
try {
    // This will throw an exception
    throw new Exception("An error occurred!");
} catch (Exception $e) {
    echo "Caught exception: " . $e->getMessage() . "\n";
} finally {
    echo "This will always execute.\n";
}

// Alternative syntax examples

// Example of a while loop with alternative syntax
echo "\nWhile Loop (Alternative Syntax):\n";
$i = 0;
while ($i < 5):
    echo "i is $i\n";
    $i++;
endwhile;

// Example of a for loop with alternative syntax
echo "\nFor Loop (Alternative Syntax):\n";
for ($j = 0; $j < 5; $j++):
    echo "j is $j\n";
endfor;

// Example of an if statement with alternative syntax
echo "\nIf Statement (Alternative Syntax):\n";
$x = 10;
if ($x < 20):
    echo "x is less than 20\n";
elseif ($x == 20):
    echo "x is equal to 20\n";
else:
    echo "x is greater than 20\n";
endif;

// Example of a switch statement with alternative syntax
echo "\nSwitch Statement (Alternative Syntax):\n";
$fruit = "apple";
switch ($fruit):
    case "banana":
        echo "Banana is good!\n";
        break;
    case "apple":
        echo "Apple is good too!\n";
        break;
    default:
        echo "I have never tried that fruit.\n";
endswitch;

// Example of a try-catch block with alternative syntax
echo "\nTry-Catch Block (Alternative Syntax):\n";
try {
    // This will throw an exception
    throw new Exception("An error occurred!");
} catch (FirstExceptionType | SecondExceptionType $e) {
    echo "Caught exception: " . $e->getMessage() . "\n";
} finally {
    echo "This will always execute.\n";
}
?>
