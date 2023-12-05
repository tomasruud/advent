<?php

$opts = getopt("f:");
$input = file_get_contents($opts["f"]);

$sum = 0;

$lines = explode("\n", $input);
for ($y = 0; $y < count($lines); $y++) {
    for ($x = 0; $x < strlen($lines[$y]); $x++) {
        if ($lines[$y][$x] === "*") {
            // scan 3x3 and look for ints
            $ints = [];
            for ($y2 = $y - 1; $y2 <= $y + 1; $y2++) {
                for ($x2 = $x - 1; $x2 <= $x + 1; $x2++) {
                    if (is_numeric($lines[$y2][$x2] ?? "")) {
                        // rewind to start of int
                        $x2--;
                        while (is_numeric($lines[$y2][$x2] ?? "")) {
                            $x2--;
                        }
                        $x2++; // go to first digit

                        $n = $lines[$y2][$x2];
                        $x2++;
                        while (is_numeric($lines[$y2][$x2] ?? "")) {
                            $n .= $lines[$y2][$x2];
                            $x2++;
                        }

                        $ints[] = intval($n);
                    }
                }
            }

            if (count($ints) === 2) {
                $sum += $ints[0] * $ints[1];
            }
        }
    }
}

echo "$sum\n";
