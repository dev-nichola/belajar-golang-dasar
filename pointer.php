<?php

$address = [
    "city" => null,
    "province" => null,
    "country" => null
];

$address1 = $address;

$address1["city"] = "Blitar";
$address1["province"] = "Jawa Timur";
$address1["country"] = "Indonesia";

$address2 = $address1;
$address2['city'] = "Malang";

var_dump($address1);
var_dump($address2);

