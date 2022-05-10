<?php
$x = "Hello World!";
$y = 2;

var_dump(is_int($y)) . PHP_EOL;
var_dump(is_nan($y)) . PHP_EOL;
var_dump(is_file($x)) . PHP_EOL;
var_dump(is_long($x)) . PHP_EOL;
var_dump(is_null($x)) . PHP_EOL;
var_dump(is_dir($x)) . PHP_EOL;
var_dump(is_array($x)) . PHP_EOL;
var_dump(is_float($x)) . PHP_EOL;
var_dump(is_double($x)) . PHP_EOL;
var_dump(is_finite($y)) . PHP_EOL;
