(use judge)

(defn parse [input]
  (filter
    (complement empty?)
    (string/split "\n" input)))

(test (parse "abc\ncde\nefg\n") @["abc" "cde" "efg"])

(def nice-peg
  (peg/compile
    ~{:main (* (look :double) (look :vowels) :no-naughty)
      :double (thru (* (<- 1) (backmatch)))
      :vowels (3 (thru (set "aeiou")))
      :no-naughty (! (thru (+ "ab" "cd" "pq" "xy")))}))

(defn is-nice? [str]
  (not (nil? (peg/match nice-peg str))))

(test (is-nice? "ugknbfddgicrmopn") true)
(test (is-nice? "aaa") true)
(test (is-nice? "jchzalrnumimnmhp") false)
(test (is-nice? "haegwjzuvuyypxyu") false)
(test (is-nice? "dvszwmarrgswjxmb") false)

(def nice-improved-peg
  (peg/compile
    ~{:main (* (look :pair) :repeat)
      :pair (thru (* (<- 2) (thru (backmatch))))
      :repeat (thru (* (<- 1) 1 (backmatch)))}))

(defn is-nice-improved? [str]
  (not (nil? (peg/match nice-improved-peg str))))

(test (is-nice-improved? "qjhvhtzxzqqjkmpb") true)
(test (is-nice-improved? "xxyxx") true)
(test (is-nice-improved? "uurcxstgmygtbstg") false)
(test (is-nice-improved? "ieodomkazucvgmuy") false)

(defn main [&]
  (->>
    (file/read stdin :all)
    parse
    (|{:p1 (count is-nice? $) :p2 (count is-nice-improved? $)})
    pp))
