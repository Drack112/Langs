<?php
$idade = 17;

echo "Você só pode entrar se tiver 18 anos." . PHP_EOL;

/*
if ($idade >= 18){
echo "Idade -> $idade" . PHP_EOL;
echo "Pó entrar";
}

if ($idade >= 18 && %nome == "drack"){
echo "Idade -> $idade" . PHP_EOL;
echo "Pó entrar";
}
 */
if ($idade == 18) {
    echo "Idade -> $idade" . PHP_EOL;
    echo "Pó entrar";
} else if ($idade > 18) {
    echo "Idade -> $idade" . PHP_EOL;
    echo "Pó entrar";
} else {
    echo "Num vai entrar >:(" . PHP_EOL;
}
;
