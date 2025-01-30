<?php
class DBWrapper {
    public $mysql;

    public function __construct($dsn, $user, $pass) {
        $this->mysql = new PDO($dsn, $user, $pass);
    }
}

// Create the object
$object = new DBWrapper("mysql:host=localhost;dbname=testdb;charset=utf8", "root", "password");

// Execute a query
$object->mysql->exec("DELETE FROM users WHERE last_login < NOW() - INTERVAL 1 YEAR");
?>