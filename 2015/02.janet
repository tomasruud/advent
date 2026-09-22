(use judge)

(defn parse [input]
  (peg/match
    ~{:main (any (+ :present 1))
      :present (group (* :dim "x" :dim "x" :dim))
      :dim (number :d+)}
    input))

(test (parse "2x3x4\n1x1x10") @[@[2 3 4] @[1 1 10]])

(defn paper-size [dims]
  (let [[a b c] (sort dims)
        surface (+ (* 2 a b) (* 2 b c) (* 2 a c))
        slack (* a b)]
    (+ surface slack)))

(defn solve-1 [presents]
  (sum (map paper-size presents)))

(test (solve-1 (parse "2x3x4")) 58)
(test (solve-1 (parse "1x1x10")) 43)

(defn ribbon-length [dims]
  (let [[a b c] (sort dims)
        wrap (+ (* 2 a) (* 2 b))
        bow (* a b c)]
    (+ wrap bow)))

(defn solve-2 [presents]
  (sum (map ribbon-length presents)))

(test (solve-2 (parse "2x3x4")) 34)
(test (solve-2 (parse "1x1x10")) 14)

(defn main [&]
  (->>
    (file/read stdin :all)
    parse
    (|{:p1 (solve-1 $) :p2 (solve-2 $)})
    pp))
