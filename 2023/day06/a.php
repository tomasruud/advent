<?php

$opts = getopt("f:");
$input = file_get_contents($opts["f"]);
$input = trim($input);
$input = explode("\n", $input);

$times = preg_split('/\s+/', $input[0]);
$times = array_slice($times, 1);

$dists = preg_split('/\s+/', $input[1]);
$dists = array_slice($dists, 1);

$product = 1;
foreach ($times as $race => $time) {
	$dist = $dists[$race];

	$ways = 0;
	for ($hold = 0; $hold <= $time; $hold++) {
		$speed = $hold;
		$rest = $time - $hold;
		$travel = $speed * $rest;

		if ($travel > $dist) {
			$ways++;
		}
	}

	if ($ways > 0) {
		$product *= $ways;
	}
}

echo "$product\n";
