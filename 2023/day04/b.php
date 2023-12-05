<?php

$opts = getopt("f:");
$input = file_get_contents($opts["f"]);

$results = [];

$lines = explode("\n", $input);
foreach ($lines as $game => $line) {
    if (trim($line) === "") {
        continue;
    }

    $tokens = preg_split("/\s+/", $line);
    $ab = array_slice($tokens, 2);

    $sep = array_search("|", $ab);
    $a = array_slice($ab, 0, $sep);
    $b = array_slice($ab, $sep);

    $wins = array_intersect($a, $b);
    $results[$game] = count($wins);
}

$hits = [];
foreach ($results as $game => $result) {
    $hits[$game] = ($hits[$game] ?? 0) + 1;
    for ($i = 0; $i < $hits[$game]; $i++) {
        for ($j = 1; $j <= $result; $j++) {
            $hits[$game + $j] = ($hits[$game + $j] ?? 0) + 1;
        }
    }
}

$sum = array_sum($hits);
echo "$sum\n";
