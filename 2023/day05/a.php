<?php

$opts = getopt("f:");
$input = file_get_contents($opts["f"]);

final class node
{
	public function __construct(
		public string $type,
		public int $value,
		public ?node $next = null,
	) {
	}
}

$seeds = [];
$low = null;

$lines = explode("\n", $input);
$z = 0;
for ($y = 0; $y < count($lines); $y++) {
	$tokens = explode(" ", $lines[$y]);

	// skip empty lines
	if (trim($tokens[0]) === "") {
		continue;
	}

	// gather seed info
	if ($tokens[0] === "seeds:") {
		foreach (array_slice($tokens, 1) as $token) {
			$seeds[] = new node("seed", intval($token));
		}
		continue;
	}

	// handle maps
	if ($tokens[1] === "map:") {
		$type = explode("-", $tokens[0])[2];

		$lookups = $seeds;
		for ($zd = 0; $zd < $z; $zd++) {
			$lookups = array_map(fn (node $n) => $n->next, $lookups);
		}

		$y++;
		while (trim($lines[$y]) !== "") {
			$tokens = explode(" ", $lines[$y]);
			$dst = intval($tokens[0]);
			$src = intval($tokens[1]);
			$n = intval($tokens[2]);

			foreach ($lookups as $lookup) {
				if ($lookup->value >= $src && $lookup->value < $src + $n) {
					// map range
					$lookup->next = new node($type, $dst + ($lookup->value - $src));
				}
			}

			$y++;
		}

		// fill any values not set
		foreach ($lookups as $lookup) {
			if ($lookup->next === null) {
				$lookup->next = new node($type, $lookup->value);
			}

			if ($type === "location") {
				if ($low === null) {
					$low = $lookup->next->value;
				} else {
					$low = min($low, $lookup->next->value);
				}
			}
		}

		$z++;
	}
}

echo "$low\n";
