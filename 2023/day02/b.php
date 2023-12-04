<?php

$opts = getopt("f:");
$input = file_get_contents($opts["f"]);

$sum = array_reduce(
    explode("\n", $input),
    function ($acc, $line) {
        $line = trim($line);
        if (strlen($line) === 0) {
            return $acc;
        }

        $max = [];

        $tokens = explode(" ", $line);
        $it = new ArrayIterator($tokens);
        while ($it->valid()) {
            if ($it->current() === "Game") {
                $it->next(); // advance to number
                $it->next(); // advance
                continue;
            }

            if (is_numeric($it->current())) {
                $n = intval($it->current());
                $it->next(); // advance to color
                $c = $it->current()[0];
                $max[$c] = max($max[$c] ?? 0, $n); // set max for current set
                $it->next(); // advance
                continue;
            }

            $it->next();
        }

        $power = array_reduce($max, fn ($acc, $n) => $acc * $n, 1); // multiply all max values

        return $acc + $power;
    },
    0
);

echo "$sum\n";
