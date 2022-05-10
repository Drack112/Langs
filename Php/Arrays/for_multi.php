<?php

$cars = array(
    array("Volvo", 22, 18),
    array("BMW", 15, 13),
    array("Saab", 5, 2),
    array("Land Rover", 17, 15)
);

for ($row = 0; $row < 4; $row++) {
    echo "Row Number";
    for ($col = 0; $col < 4; $col++) {
        echo $cars[$row][$col];
    };
}
echo " ";
