(use judge)
(use ./build/_md5)

(defn solve [want prefix &opt n]
  (default n 0)
  (if (string/has-prefix?
        want
        (md5 (string prefix n)))
    n
    (solve want prefix (+ n 1))))

(test (solve "00000" "abcdef") 609043)
(test (solve "00000" "pqrstuv") 1048970)

(defn solve-both [prefix]
  (let [p1 (solve "00000" prefix)
        p2 (solve "000000" prefix p1)]
    {:p1 p1 :p2 p2}))

(defn main [&]
  (->>
    (file/read stdin :all)
    string/trim
    solve-both
    pp))
