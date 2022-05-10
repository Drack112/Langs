<?php
class Car
{
    public $color;
    public $model;

    public function __construct($color, $model)
    {
        $this->color = $color;
        $this->model = $model;
    }

    public function message()
    {
        echo "My car is " . $this->color . " " . $this->model . " " . PHP_EOL;
    }
}

$myCar = new Car("red", "BMW");
$myCar->message();
