<?php
$idade = 17;

echo "Você só pode entrar se tiver 18 anos." . PHP_EOL;

/*
if ($idade >= 18){
echo "Idade -> $idade" . PHP_EOL;
echo "Pó entrar";
}
 */
if ($idade == 18 or $idade > 18) {
    echo "Idade -> $idade" . PHP_EOL;
    echo "Pó entrar";
} else {
    echo "Tu não tem autorização" . PHP_EOL;
}
