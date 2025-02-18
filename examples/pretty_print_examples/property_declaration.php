<?php
class Example
{
    private bool $modified = false;
    private float $x = 0.0;
    private int $y = 0;


    public string $foo = 'default value' {
        get => $this->foo . ($this->modified ? ' (modified)' : '');

        set {
            $this->foo = strtolower($value);
            $this->modified = true;
        }
    }
}
?>