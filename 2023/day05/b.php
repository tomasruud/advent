<?php

$opts = getopt("f:");

$input = file_get_contents($opts["f"]);
$input = trim($input);
$input = explode("\n\n", $input);

// resolve pairs of seeds
$ns = array_shift($input); // shifts off first element
$ns = explode(" ", $ns);
$ns = array_slice($ns, 1);
$ns = array_chunk($ns, 2);
$seeds = [];
foreach ($ns as [$lo, $len]) {
	$seeds[] = [$lo, $lo + $len];
}

// iterate over map chunks
foreach ($input as $chunk) {
	// parse maps in chunk
	$lns = explode("\n", $chunk);
	$lns = array_slice($lns, 1); // skip title
	$maps = array_map(fn ($ln) => explode(" ", $ln), $lns);

	// iterate seeds
	$next = [];
	foreach ($seeds as [$lo, $hi]) {
		$found = false;
		foreach ($maps as [$dst, $src, $len]) {
			// <[>]
			if ($lo < $src && $hi > $src && $hi <= $src + $len) {
				$found = true;
				$seeds[] = [$lo, $src];
				$next[] = [$dst, $dst + ($hi - $src)];
				continue;
			}

			// <[]>
			if ($lo < $src && $hi > $src + $len) {
				$found = true;
				$seeds[] = [$lo, $src];
				$next[] = [$dst, $dst + $len];
				$seeds[] = [$src + $len, $hi];
				continue;
			}

			// [<]>
			if ($lo >= $src && $lo < $src + $len && $hi > $src + $len) {
				$found = true;
				$next[] = [$dst + ($lo - $src), $dst + $len];
				$seeds[] = [$src + $len, $hi];
				continue;
			}

			// [<>]
			if ($lo >= $src && $lo < $src + $len && $hi > $src && $hi <= $src + $len) {
				$found = true;
				$next[] = [$dst + ($lo - $src), $dst + ($hi - $src)];
				continue;
			}
		}

		// if no overlap
		if (!$found) {
			$next[] = [$lo, $hi];
		}
	}

	// set up new seed list
	$seeds = $next;
}

// loop over seeds, find lowest
[$min] = $seeds[0];
foreach ($seeds as [$lo]) {
	if ($lo < $min) {
		$min = $lo;
	}
}

echo "$min\n";
