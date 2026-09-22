(use judge)

(def deltas
  {(chr "(") 1
   (chr ")") -1})

(defn parse [input]
  (map deltas (filter deltas input)))

(defn solve-1 [moves]
  (sum moves))

(test (solve-1 (parse "(())")) 0)
(test (solve-1 (parse "()()")) 0)
(test (solve-1 (parse "(()(()(")) 3)
(test (solve-1 (parse "))(")) -1)

(defn solve-2 [moves]
  (first
    (reduce
      (fn [[i floor] delta]
        (if (= floor -1)
          [i -1]
          [(+ 1 i) (+ floor delta)]))
      [0 0]
      moves)))

(test (solve-2 (parse ")")) 1)
(test (solve-2 (parse "()())")) 5)

(defn main [&]
  (->>
    (file/read stdin :all)
    parse
    (|{:p1 (solve-1 $) :p2 (solve-2 $)})
    pp))
