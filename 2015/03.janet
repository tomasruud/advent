(use judge)

(def deltas
  {(chr ">") [1 0]
   (chr "v") [0 -1]
   (chr "<") [-1 0]
   (chr "^") [0 1]})

(defn parse [input]
  (map deltas (filter deltas input)))

(defn visits [moves]
  (def all @[[0 0]])
  (var pos [0 0])

  (each [dx dy] moves
    (set pos [(+ (pos 0) dx) (+ (pos 1) dy)])
    (array/push all pos))

  all)

(defn solve-1 [moves]
  (->>
    (visits moves)
    distinct
    length))

(test (solve-1 (parse ">")) 2)
(test (solve-1 (parse "^>v<")) 4)
(test (solve-1 (parse "^v^v^v^v^v")) 2)

(defn split [moves]
  (def santa @[])
  (def robot @[])

  (eachk i moves
    (array/push
      (if (even? i) santa robot)
      (moves i)))

  [santa robot])

(defn solve-2 [moves]
  (->>
    (split moves)
    (mapcat visits)
    distinct
    length))

(test (solve-2 (parse "^v")) 3)
(test (solve-2 (parse "^>v<")) 3)
(test (solve-2 (parse "^v^v^v^v^v")) 11)

(defn main [&]
  (->>
    (file/read stdin :all)
    parse
    (|{:p1 (solve-1 $) :p2 (solve-2 $)})
    pp))
