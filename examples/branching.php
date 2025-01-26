<?php

// 1. Multiple if branching
function testMultipleIf($value) {
    if ($value > 10) {
        echo "Value is greater than 10\n";
    }
    if ($value % 2 == 0) {
        echo "Value is even\n";
    }
    if ($value < 0) {
        echo "Value is negative\n";
    }
}

// 2. Multiple if-else branching
function testMultipleIfElse($value) {
    if ($value > 10) {
        echo "Value is greater than 10\n";
    } else {
        echo "Value is 10 or less\n";
    }

    if ($value % 2 == 0) {
        echo "Value is even\n";
    } else {
        echo "Value is odd\n";
    }
}

// 3. Multiple if-elseif branching
function testMultipleIfElseIf($value) {
    if ($value > 10) {
        echo "Value is greater than 10\n";
    } elseif ($value > 5) {
        echo "Value is greater than 5 but 10 or less\n";
    } elseif ($value > 0) {
        echo "Value is greater than 0 but 5 or less\n";
    }
}

// 4. Multiple if-elseif-else branching
function testMultipleIfElseIfElse($value) {
    if ($value > 10) {
        echo "Value is greater than 10\n";
    } elseif ($value > 5) {
        echo "Value is greater than 5 but 10 or less\n";
    } elseif ($value > 0) {
        echo "Value is greater than 0 but 5 or less\n";
    } else {
        echo "Value is 0 or negative\n";
    }
}

// 5. Cases with "else if" (alternative to "elseif")
function testElseIf($value) {
    if ($value > 10) {
        echo "Value is greater than 10\n";
    } else if ($value > 5) { // Note the space in "else if"
        echo "Value is greater than 5 but 10 or less\n";
    } else if ($value > 0) {
        echo "Value is greater than 0 but 5 or less\n";
    } else {
        echo "Value is 0 or negative\n";
    }
}

// 6. switch-case branching
function testSwitch($value) {
    switch ($value) {
        case 1:
            echo "One\n";
            break;
        case 2:
            echo "Two\n";
            break;
        default:
            echo "Other number\n";
    }
}

// 7. match branching (PHP 8.0+)
function testMatch($value) {
    echo match (true) {
        $value > 0 => "Positive\n",
        $value < 0 => "Negative\n",
        default => "Zero\n",
    };
}

// 8. while loop
function testWhile($limit) {
    $i = 0;
    while ($i < $limit) {
        echo "While loop iteration: $i\n";
        $i++;
    }
}

// 9. do-while loop
function testDoWhile($limit) {
    $i = 0;
    do {
        echo "Do-while loop iteration: $i\n";
        $i++;
    } while ($i < $limit);
}

// 10. for loop
function testFor($limit) {
    for ($i = 0; $i < $limit; $i++) {
        echo "For loop iteration: $i\n";
    }
}

// 11. foreach loop
function testForeach($array) {
    foreach ($array as $key => $value) {
        echo "Key: $key, Value: $value\n";
    }
}

// 12. Ternary operator
function testTernary($value) {
    echo $value > 0 ? "Positive\n" : "Non-positive\n";
}

// 13. Null coalescing operator (PHP 7.0+)
function testNullCoalescing($value) {
    echo $value ?? "Value is null or undefined\n";
}

// 14. Goto statement
function testGoto() {
    $i = 0;
    label:
    echo "Goto iteration: $i\n";
    $i++;
    if ($i < 3) {
        goto label;
    }
}

// 15. Break and continue in loops
function testBreakAndContinue($limit) {
    for ($i = 0; $i < $limit; $i++) {
        if ($i == 3) {
            echo "Skipping 3\n";
            continue;
        }
        if ($i == 5) {
            echo "Breaking at 5\n";
            break;
        }
        echo "Loop iteration: $i\n";
    }
}

function testBasicTryCatchFinally() {
    try {
        echo "Trying something risky...\n";
        throw new Exception("Something went wrong!");
    } catch (Exception $e) {
        echo "Caught an exception: " . $e->getMessage() . "\n";
    } finally {
        echo "Finally block executed (cleanup or logging).\n";
    }
}

// Function to demonstrate multiple catch blocks
function testMultipleCatch() {
    try {
        echo "Trying something risky...\n";
        throw new InvalidArgumentException("Invalid argument provided!");
    } catch (InvalidArgumentException $e) {
        echo "Caught an InvalidArgumentException: " . $e->getMessage() . "\n";
    } catch (Exception $e) {
        echo "Caught a general exception: " . $e->getMessage() . "\n";
    } finally {
        echo "Finally block executed (cleanup or logging).\n";
    }
}

// Function to demonstrate try-catch without finally
function testTryCatchWithoutFinally() {
    try {
        echo "Trying something risky...\n";
        throw new RuntimeException("A runtime exception occurred!");
    } catch (RuntimeException $e) {
        echo "Caught a runtime exception: " . $e->getMessage() . "\n";
    }
}

// Function to demonstrate try-finally without catch
function testTryFinallyWithoutCatch() {
    try {
        echo "Trying something risky...\n";
        // Simulate some risky operation
        throw new Exception("Something went wrong!");
    } finally {
        echo "Finally block executed even without catch.\n";
    }
}

// Main function to demonstrate all branches
function main() {
    echo "1. Multiple If Branching:\n";
    testMultipleIf(15);
    testMultipleIf(8);
    testMultipleIf(-3);

    echo "\n2. Multiple If-Else Branching:\n";
    testMultipleIfElse(12);
    testMultipleIfElse(7);

    echo "\n3. Multiple If-ElseIf Branching:\n";
    testMultipleIfElseIf(12);
    testMultipleIfElseIf(8);
    testMultipleIfElseIf(3);

    echo "\n4. Multiple If-ElseIf-Else Branching:\n";
    testMultipleIfElseIfElse(12);
    testMultipleIfElseIfElse(8);
    testMultipleIfElseIfElse(3);
    testMultipleIfElseIfElse(-3);

    echo "\n5. Else If Branching:\n";
    testElseIf(12);
    testElseIf(8);
    testElseIf(3);
    testElseIf(-3);


    echo "\n2. Switch-Case Branching:\n";
    testSwitch(1);
    testSwitch(3);

    echo "\n3. Match Branching:\n";
    testMatch(5);
    testMatch(-5);
    testMatch(0);

    echo "\n4. While Loop:\n";
    testWhile(3);

    echo "\n5. Do-While Loop:\n";
    testDoWhile(3);

    echo "\n6. For Loop:\n";
    testFor(3);

    echo "\n7. Foreach Loop:\n";
    testForeach(["a" => 1, "b" => 2, "c" => 3]);

    echo "\n8. Ternary Operator:\n";
    testTernary(5);
    testTernary(0);

    echo "\n9. Null Coalescing Operator:\n";
    testNullCoalescing(null);
    testNullCoalescing(42);

    echo "\n10. Goto Statement:\n";
    testGoto();

    echo "\n11. Break and Continue:\n";
    testBreakAndContinue(10);

    echo "\n12. Basic Try-Catch-Finally:\n";
    testBasicTryCatchFinally();

    echo "\n13. Multiple Catch Blocks:\n";
    testMultipleCatch();

    echo "\n14. Try-Catch Without Finally:\n";
    testTryCatchWithoutFinally();

    echo "\n15. Try-Finally Without Catch:\n";
    testTryFinallyWithoutCatch();

}

// Run the main function
main();
?>
