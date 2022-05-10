<?php

declare(strict_types=1); // Permitir tipagem nos valores

function addNumbers(float $f1, float $f2): float
{
    return $f1 + $f2;
};

echo addNumbers(1.2, 1.54);

function addNumber(float $a, float $b): int
{
    return (int)($a + $b);
}
echo addNumber(1.2, 5.2);
