(use judge)

(def test-input ```
London to Dublin = 464
London to Belfast = 518
Dublin to Belfast = 141
```)

(defn parse [input]
  (peg/match
    ~{:main (any (* (group :route) (any :s+)))
      :route (* :dest " to " :dest " = " :dist)
      :dest (cmt ':w+ ,keyword)
      :dist (number :d+)}
    input))

(test (parse test-input)
      @[@[:London :Dublin 464]
        @[:London :Belfast 518]
        @[:Dublin :Belfast 141]])

(defn routes [defs]
  (def dists @{})

  (each [a b dist] defs
    (put dists [a b] dist)
    (put dists [b a] dist))

  (def cities (distinct (map first (keys dists))))

  (defn trace [from &opt tos route dist]
    (default tos cities)
    (default route [from])
    (default dist 0)

    (def next-tos (filter |(not= from $) tos))

    (if (empty? next-tos)
      [[route dist]]
      (catseq [to :in next-tos]
        (trace to
               next-tos
               [;route to]
               (+ dist (get dists [from to]))))))

  (->>
    cities
    (mapcat trace)
    from-pairs))

(test (routes (parse test-input))
      @{[:Dublin :London :Belfast] 982
        [:London :Dublin :Belfast] 605
        [:London :Belfast :Dublin] 659
        [:Dublin :Belfast :London] 659
        [:Belfast :Dublin :London] 605
        [:Belfast :London :Dublin] 982})

(defn solve [dists]
  {:p1 (min-of dists) :p2 (max-of dists)})

(test (solve (routes (parse test-input))) {:p1 605 :p2 982})

(defn main [&]
  (->>
    (file/read stdin :all)
    parse
    routes
    solve
    pp))
