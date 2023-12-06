<?php

$opts = getopt("f:");
$input = file_get_contents($opts["f"]);
$input = trim($input);
$input = explode("\n", $input);

$times = preg_split('/\s+/', $input[0]);
$times = array_slice($times, 1);
$time = (int) implode("", $times);

$dists = preg_split('/\s+/', $input[1]);
$dists = array_slice($dists, 1);
$dist = (int) implode("", $dists);

$ways = 0;
for ($hold = 0; $hold <= $time; $hold++) {
	$speed = $hold;
	$rest = $time - $hold;
	$travel = $speed * $rest;

	if ($travel > $dist) {
		$ways++;
	}
}

echo "$ways\n";
