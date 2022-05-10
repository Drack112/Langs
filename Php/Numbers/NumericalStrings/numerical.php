<?php
$x = 5985;
echo "$x";
var_dump(is_numeric($x));

$x = "5985";
echo "$x";
var_dump(is_numeric($x));


$x = "59.85" + 100;
echo "$x";
var_dump(is_numeric($x));

$x = "Hello";
echo "$x";
var_dump(is_numeric($x));
