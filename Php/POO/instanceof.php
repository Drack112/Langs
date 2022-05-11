<?php
class Fruit
{
    public string $name;
    public string $color;

    function set_name($name, $color)
    {
        $this->name = $name;
        $this->color = $color;
    }

    function get_name()
    {
        return $this->name;
    }
};

$apple = new Fruit();
var_dump($apple instanceof Fruit);
