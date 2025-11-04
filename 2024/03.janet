(use judge)

(def test-input "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))")

(defn solve-1 [input]
  ((peg/match
     ~{:main (cmt (any (thru :mul)) ,+)
       :mul (cmt (* "mul(" :num "," :num ")") ,*)
       :num (number (between 1 3 :d))}
     input) 0))

(test (solve-1 test-input) 161)

(def test-input-2 "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))")

(defn solve-2 [input]
  (->>
    (peg/match
      ~{:main (any (thru (+ :do :dont :mul)))
        :mul (cmt (* "mul(" :num "," :num ")") ,*)
        :num (number (between 1 3 :d))
        :do (/ "do()" true)
        :dont (/ "don't()" false)}
      input)

    (reduce
      (fn [[sum enabled] v]
        (case (type v)
          :boolean [sum v]
          :number [(+ sum (if enabled v 0)) enabled]))
      [0 true])
    0))

(test (solve-2 test-input-2) 48)

(defn main [&]
  (->>
    (file/read stdin :all)
    (|{:p1 (solve-1 $) :p2 (solve-2 $)})
    pp))
