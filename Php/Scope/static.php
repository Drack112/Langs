<?php
function myTest()
{
    // static --> Não vai ser deletado
    static $x = 0;
    echo $x;
    $x++;
}

myTest();
myTest();
myTest();
