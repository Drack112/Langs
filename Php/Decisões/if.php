<?php
$idade = 18;

echo "Você só pode entrar se tiver 18 anos." . PHP_EOL;

if ($idade == 18 or $idade > 18) {
    echo "Idade -> $idade" . PHP_EOL;
    echo "Pó entrar";
}
