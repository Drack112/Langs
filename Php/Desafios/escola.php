<?php
$nota1 = 10;
$nota2 = 10;

$media = $nota1 / 2 + $nota2 / 2;

echo "Nota atual -> {$media}" . PHP_EOL;

if ($media <= 5) {
    echo "O aluno não passou de ano." . PHP_EOL;
} else if ($media >= 6) {
    echo "O aluno passou de ano." . PHP_EOL;
} else {
    echo "Erro!";
};
