<?php

$opts = getopt("f:");
$input = file_get_contents($opts["f"]);
$input = trim($input);
$input = explode("\n", $input);

$hands = [];
foreach ($input as $ln) {
    [$hand, $bet] = explode(" ", $ln);
    $cards = str_split($hand);

    // determine hand type
    $freq = [0];
    $jcount = 0;
    foreach ($cards as $card) {
        if ($card === "J") {
            $jcount++;
        } elseif (!isset($freq[$card])) {
            $freq[$card] = 1;
        } else {
            $freq[$card]++;
        }
    }
    $freq = array_values($freq);
    rsort($freq); // sort values decending

    $type = 0;
    if ($freq[0] + $jcount === 5) {
        $type = 6; // five of a kind
    } elseif ($freq[0] + $jcount === 4) {
        $type = 5; // four of a kind
    } elseif (($freq[0] + $jcount === 3 && $freq[1] === 2) ||
        ($freq[0] === 3 && $freq[1] + $jcount === 2)
    ) {
        $type = 4; // full house
    } elseif ($freq[0] + $jcount === 3) {
        $type = 3; // three of a kind
    } elseif (($freq[0] + $jcount === 2 && $freq[1] === 2) ||
        ($freq[0] === 2 && $freq[1] + $jcount === 2)
    ) {
        $type = 2; // two pairs
    } elseif ($freq[0] + $jcount === 2) {
        $type = 1; // one pair
    }

    // map card values
    $vals = array_map(function ($card) {
        $values = [
            "A" => 14,
            "K" => 13,
            "Q" => 12,
            "T" => 10,
            "9" => 9,
            "8" => 8,
            "7" => 7,
            "6" => 6,
            "5" => 5,
            "4" => 4,
            "3" => 3,
            "2" => 2,
            "1" => 1,
            "J" => 0,
        ];
        return $values[$card];
    }, $cards);

    // add mapped hand
    $hands[] = [$bet, [$type, ...$vals]];
}

usort($hands, function ($a, $b) {
    $av = $a[1];
    $bv = $b[1];

    for ($i = 0; $i < count($av); $i++) {
        if ($av[$i] === $bv[$i]) {
            continue;
        }

        if ($av[$i] < $bv[$i]) {
            return -1;
        }

        if ($av[$i] > $bv[$i]) {
            return 1;
        }
    }
    return 0;
});

$winnings = 0;
for ($i = 0; $i < count($hands); $i++) {
    [$bid] = $hands[$i];
    $winnings += ($i + 1) * $bid;
}

echo "$winnings\n";
