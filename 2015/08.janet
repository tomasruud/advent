(use judge)

(def test-input ```
""
"abc"
"aaa\"aaa"
"\x27"
```)

(defn parse [input]
  (filter (complement empty?) (string/split "\n" input)))

(test (parse test-input)
      @[`""` `"abc"` `"aaa\"aaa"` `"\x27"`])

(def interpolate-peg
  (peg/compile
    ~{:main (accumulate :literal)
      :literal (* `"` (any (+ :escaped ':a)) `"`)
      :escaped (+
                 (* `\\` (constant `\`))
                 (* `\"` (constant `"`))
                 (* `\x` (cmt (number (between 1 2 :h) 16) ,string/from-bytes)))}))

(def interpolate
  (comp first (partial peg/match interpolate-peg)))

(test (interpolate `""`) "")
(test (interpolate `"abc"`) "abc")
(test (interpolate `"aaa\"aaa"`) `aaa"aaa`)
(test (interpolate `"\x27"`) "'")

(defn solve-1 [literals]
  (defn length-diff [literal]
    (-
      (length literal)
      (length (interpolate literal))))

  (sum (map length-diff literals)))

(test (solve-1 (parse test-input)) 12)

(def expand-peg
  (peg/compile
    ~{:main (accumulate (* (constant `"`) :literal (constant `"`)))
      :literal (* :quote (any :value) :quote)
      :quote (* `"` (constant `\"`))
      :value (+
               (* `\\` (constant `\\\\`))
               (* `\"` (constant `\\\"`))
               (* (constant `\`) '`\x` '(between 1 2 :h))
               ':a)}))

(def expand
  (comp first (partial peg/match expand-peg)))

(test (expand `""`) `"\"\""`)
(test (expand `"abc"`) `"\"abc\""`)
(test (expand `"aaa\"aaa"`) `"\"aaa\\\"aaa\""`)
(test (expand `"\x27"`) `"\"\\x27\""`)

(defn solve-2 [literals]
  (defn length-diff [literal]
    (-
      (length (expand literal))
      (length literal)))

  (sum (map length-diff literals)))

(test (solve-2 (parse test-input)) 19)

(defn main [&]
  (->>
    (file/read stdin :all)
    parse
    (|{:p1 (solve-1 $) :p2 (solve-2 $)})
    pp))
