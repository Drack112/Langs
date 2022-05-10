<?php
function divice($divi, $divo)
{
    if ($divo == 0) {
        throw new Exception("Division by Zero");
    }
    return $divi / $divo;
};

echo divice(1, 0);
