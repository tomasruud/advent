<?php

$opts = getopt("f:");
$input = file_get_contents($opts["f"]);

$sum = 0;
$lines = explode("\n", $input);
foreach ($lines as $line) {
    $tokens = preg_split("/\s+/", $line); // split by space
    $game = array_slice($tokens, 2); // remove game info

    $sep = array_search("|", $game);
    $a = array_slice($game, 0, $sep);
    $b = array_slice($game, $sep);

    $wins = array_intersect($a, $b);

    if (count($wins) > 0) {
        $sum += pow(2, count($wins) - 1);
    }
}

echo "$sum\n";
