(use judge)

(def test-input ```
3   4
4   3
2   5
1   3
3   9
3   3
```)

(defn parse [input]
  (peg/match
    '{:main (any (* :line (any "\n")))
      :line (group (* :digit :s+ :digit))
      :digit (number :d+)}
    input))

(test (parse test-input)
      @[@[3 4]
        @[4 3]
        @[2 5]
        @[1 3]
        @[3 9]
        @[3 3]])

(defn solve-1 [input]
  (def parsed (parse input))
  (def lhs (-> (map |(get $ 0) parsed) sort))
  (def rhs (-> (map |(get $ 1) parsed) sort))

  (reduce
    (fn [acc i]
      (def diff (math/abs (- (get lhs i) (get rhs i))))
      (+ acc diff))
    0
    (range 0 (length lhs))))

(test (solve-1 test-input) 11)

(defn solve-2 [input]
  (def parsed (parse input))
  (def lhs (map |(get $ 0) parsed))
  (def rhs (map |(get $ 1) parsed))

  (def fqs (frequencies rhs))
  (reduce
    (fn [acc n]
      (+ acc (* n (or (get fqs n) 0))))
    0
    lhs))

(test (solve-2 test-input) 31)

(defn main [&]
  (->>
    (file/read stdin :all)
    (|{:p1 (solve-1 $) :p2 (solve-2 $)})
    pp))
