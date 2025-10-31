(use judge)

(def test-input ```
7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
```)

(defn parse [input]
  (peg/match
    '{:main (any (* :report (any "\n")))
      :report (group (* :level (any (* " " :level))))
      :level (number :d+)}
    input))

(test (parse test-input)
      @[@[7 6 4 2 1]
        @[1 2 7 8 9]
        @[9 7 6 2 1]
        @[1 3 2 4 5]
        @[8 6 4 4 1]
        @[1 3 6 7 9]])

(defn is-safe [report]
  (and
    (or (< ;report) (> ;report))
    (reduce2
      (fn [acc n]
        (and
          acc
          (< 0 (math/abs (- (get report n) (get report (+ n 1)))) 4)))
      (range 0 (- (length report) 1)))))

(test (is-safe [7 6 4 2 1]) true)
(test (is-safe [1 2 7 8 9]) false)
(test (is-safe [9 7 6 2 1]) false)
(test (is-safe [1 3 2 4 5]) false)
(test (is-safe [8 6 4 4 1]) false)
(test (is-safe [1 3 6 7 9]) true)

(defn solve-1 [input]
  (def parsed (parse input))
  (reduce
    (fn [acc report]
      (+ acc (if (is-safe report) 1 0)))
    0
    parsed))

(test (solve-1 test-input) 2)

(defn main [&]
  (->>
    (file/read stdin :all)
    (|{:p1 (solve-1 $)})
    pp))
