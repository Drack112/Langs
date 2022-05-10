<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Valor PHP</title>
</head>

<body>
    <?php
    $valor = $_GET["v"];
    $raiz = sqrt($valor);
    echo "<h1>{$raiz}</h1>";
    ?>
    <a href="./form.php">Voltar</a>
</body>

</html>
