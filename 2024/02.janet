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
    (reduce
      (fn [acc n]
        (and
          acc
          (< 0 (math/abs (- (get report n) (get report (+ n 1)))) 4)))
      true
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

(defn exclude [arr i]
  (array/join
    (array/slice arr 0 i)
    (array/slice arr (+ i 1))))

(test (exclude [1 3 6 7 9] 0) @[3 6 7 9])
(test (exclude [1 3 6 7 9] 1) @[1 6 7 9])
(test (exclude [1 3 6 7 9] 4) @[1 3 6 7])

(defn is-safe2 [report]
  (or
    (is-safe report)
    (reduce
      (fn [acc i]
        (or acc (is-safe (exclude report i))))
      false
      (range 0 (length report)))))

(test (is-safe2 [7 6 4 2 1]) true)
(test (is-safe2 [1 2 7 8 9]) false)
(test (is-safe2 [9 7 6 2 1]) false)
(test (is-safe2 [1 3 2 4 5]) true)
(test (is-safe2 [8 6 4 4 1]) true)
(test (is-safe2 [1 3 6 7 9]) true)

(defn solve-2 [input]
  (def parsed (parse input))
  (reduce
    (fn [acc report]
      (+ acc (if (is-safe2 report) 1 0)))
    0
    parsed))

(test (solve-2 test-input) 4)

(defn main [&]
  (->>
    (file/read stdin :all)
    (|{:p1 (solve-1 $) :p2 (solve-2 $)})
    pp))
