<?php

$opts = getopt("f:");
$input = file_get_contents($opts["f"]);

$sum = 0;

function isSymbol($c)
{
    return $c !== "." && !is_numeric($c);
}

$lines = explode("\n", $input);
for ($y = 0; $y < count($lines); $y++) {
    for ($x = 0; $x < strlen($lines[$y]); $x++) {
        if (is_numeric($lines[$y][$x])) {
            $isPart = false;

            // start of number
            $n = $lines[$y][$x];

            // check left and vertical
            if (
                isSymbol($lines[$y][$x - 1] ?? ".") ||
                isSymbol($lines[$y - 1][$x - 1] ?? ".") ||
                isSymbol($lines[$y + 1][$x - 1] ?? ".") ||
                isSymbol($lines[$y - 1][$x] ?? ".") ||
                isSymbol($lines[$y + 1][$x] ?? ".")
            ) {
                $isPart = true;
            }

            $x++; // next token
            while ($x < strlen($lines[$y]) && is_numeric($lines[$y][$x])) {
                $n .= $lines[$y][$x];

                // check vertical
                if (
                    isSymbol($lines[$y - 1][$x] ?? ".") ||
                    isSymbol($lines[$y + 1][$x] ?? ".")
                ) {
                    $isPart = true;
                }

                $x++; // next token
            }
            // end of number

            // check right
            if (
                isSymbol($lines[$y][$x] ?? ".") ||
                isSymbol($lines[$y - 1][$x] ?? ".") ||
                isSymbol($lines[$y + 1][$x] ?? ".")
            ) {
                $isPart = true;
            }

            if ($isPart) {
                $sum += intval($n);
            }
        }
    }
}

echo "$sum\n";
