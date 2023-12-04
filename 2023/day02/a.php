<?php

$opts = getopt("f:");
$input = file_get_contents($opts["f"]);

$rules = [
    "r" => 12,
    "g" => 13,
    "b" => 14,
];

$sum = array_reduce(
    explode("\n", $input),
    function ($acc, $line) use ($rules) {
        $line = trim($line);
        if (strlen($line) === 0) {
            return $acc;
        }

        $game = 0;

        $tokens = explode(" ", $line);
        $it = new ArrayIterator($tokens);
        while ($it->valid()) {
            if ($it->current() === "Game") {
                $it->next(); // advance to number
                $game = filter_var($it->current(), FILTER_SANITIZE_NUMBER_INT); // grab int
                $it->next(); // advance
                continue;
            }

            if (is_numeric($it->current())) {
                $n = intval($it->current());
                $it->next(); // advance to color
                $c = $it->current()[0]; // grab first letter in color
                if ($n > $rules[$c]) {
                    $game = 0;
                    break; // break out if game is impossible
                }

                $it->next(); // advance
                continue;
            }

            $it->next();
        }

        return $acc + $game;
    },
    0
);

echo "$sum\n";
