(use judge)

(def test-input ```
turn on 1,2 through 2,3
turn off 2,4 through 6,7
toggle 4,3 through 2,1
```)

(defn parse [input]
  (peg/match
    ~{:main (any (* (group :instruction) (any :s+)))
      :instruction (* :action " " :pos " through " :pos)
      :action (+ (* "turn on" (constant :on))
                 (* "turn off" (constant :off))
                 (* "toggle" (constant :toggle)))
      :pos (group (* (number :d+) "," (number :d+)))}
    input))

(test (parse test-input)
      @[@[:on @[1 2] @[2 3]]
        @[:off @[2 4] @[6 7]]
        @[:toggle @[4 3] @[2 1]]])

(defn solve-1 [instructions &opt size]
  (default size 1000)

  (def lights (array/new-filled (* size size) false))

  (each [action [x1 y1] [x2 y2]] instructions
    (loop [y :range-to [y1 y2]]
      (def start (+ (* y size) x1))
      (def end (+ (* y size) x2))

      (loop [i :range-to [start end]]
        (put lights i (case action
                        :on true
                        :off false
                        :toggle (not (get lights i)))))))

  (count identity lights))

(test (solve-1 (parse "turn on 0,0 through 999,999")) 1_000_000)
(test (solve-1 (parse "toggle 0,0 through 999,0")) 1_000)
(test (solve-1 (parse "turn on 0,0 through 999,999\nturn off 499,499 through 500,500")) 999_996)

(defn solve-2 [instructions &opt size]
  (default size 1000)

  (def lights (array/new-filled (* size size) 0))

  (each [action [x1 y1] [x2 y2]] instructions
    (loop [y :range-to [y1 y2]]
      (def start (+ (* y size) x1))
      (def end (+ (* y size) x2))

      (loop [i :range-to [start end]]
        (def light (get lights i))
        (put lights i (case action
                        :on (+ light 1)
                        :off (max 0 (- light 1))
                        :toggle (+ light 2))))))

  (sum lights))

(test (solve-2 (parse "turn on 0,0 through 0,0")) 1)
(test (solve-2 (parse "toggle 0,0 through 999,999")) 2_000_000)

(defn main [&]
  (->>
    (file/read stdin :all)
    parse
    (|{:p1 (solve-1 $) :p2 (solve-2 $)})
    pp))
