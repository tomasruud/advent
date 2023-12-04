<?php

$opts = getopt("f:");

$input = file_get_contents($opts["f"]);
$lines = explode("\n", $input);
$sum = array_reduce($lines, function ($acc, $line) {
    if (strlen($line) === 0) {
        return $acc;
    }

    $ints = filter_var($line, FILTER_SANITIZE_NUMBER_INT);
    $ints = str_replace(["-", "+"], ["", ""], $ints);
    $n = $ints[0] . $ints[strlen($ints) - 1];
    return $acc + intval($n);
}, 0);

echo "$sum\n";
