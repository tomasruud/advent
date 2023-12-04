<?php

$opts = getopt("f:");

$input = file_get_contents($opts["f"]);
$lines = explode("\n", $input);
$sum = array_reduce($lines, function ($acc, $line) {
    if (strlen($line) === 0) {
        return $acc;
    }

    $ints = str_replace(
        ["one", "two", "three", "four", "five", "six", "seven", "eight", "nine"],
        ["o1e", "t2o", "t3e", "f4r", "f5e", "s6x", "s7n", "e8t", "n9e"],
        $line
    );
    $ints = filter_var($ints, FILTER_SANITIZE_NUMBER_INT);
    $ints = str_replace(["-", "+"], ["", ""], $ints);
    $n = intval($ints[0] . $ints[strlen($ints) - 1]);
    return $acc + $n;
}, 0);

echo "$sum\n";
