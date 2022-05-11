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

$newFruit = new Fruit();
$newFruit->set_name("Banana", "Yellow");
echo $newFruit->get_name();
