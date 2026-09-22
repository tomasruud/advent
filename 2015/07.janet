(use judge)

(def test-input ```
123 -> x
456 -> y
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> i
```)

(defn parse [input]
  (peg/match
    ~{:main (any (* (group :instruction) (any :s+)))
      :instruction (* (group :op) " -> " :wire)

      :op (+
            (* (constant :and) :op-and)
            (* (constant :or) :op-or)
            (* (constant :lsft) :op-lsft)
            (* (constant :rsft) :op-rsft)
            (* (constant :not) :op-not)
            (* (constant :plain) :op-plain))

      :op-and (* :value " AND " :value)
      :op-or (* :value " OR " :value)
      :op-lsft (* :value " LSHIFT " :value)
      :op-rsft (* :value " RSHIFT " :value)
      :op-not (* "NOT " :value)
      :op-plain :value

      :value (+ :number :wire)

      :wire (cmt (<- :a+) ,keyword)
      :number (number :d+)}
    input))

(test (parse test-input)
      @[@[@[:plain 123] :x]
        @[@[:plain 456] :y]
        @[@[:and :x :y] :d]
        @[@[:or :x :y] :e]
        @[@[:lsft :x 2] :f]
        @[@[:rsft :y 2] :g]
        @[@[:not :x] :h]
        @[@[:not :y] :i]])

(defn probe [signals wire &opt overrides]
  (default overrides {})

  (def ops (merge
             (from-pairs (map reverse signals))
             overrides))

  (defn uint16 [n] (band n 0xffff))

  (defn op-thunk [[op & args] thunk-for]
    (defn arg [x] ((thunk-for (args x))))
    (var cache nil)

    (fn []
      (if-not (nil? cache)
        cache
        (set cache (case op
                     :plain (arg 0)
                     :and (band (arg 0) (arg 1))
                     :or (bor (arg 0) (arg 1))
                     :lsft (uint16 (blshift (arg 0) (arg 1)))
                     :rsft (uint16 (brshift (arg 0) (arg 1)))
                     :not (uint16 (bnot (arg 0))))))))

  (def thunks @{})

  (defn thunk-for [value]
    (cond
      (number? value) (fn [] value)
      (has-key? thunks value) (thunks value)
      (let [thunk (op-thunk (ops value) thunk-for)]
        (put thunks value thunk)
        thunk)))

  ((thunk-for wire)))

(deftest "probe"
  (let [sut (partial probe (parse test-input))]
    (test (sut :d) 72)
    (test (sut :e) 507)
    (test (sut :f) 492)
    (test (sut :g) 114)
    (test (sut :h) 65412)
    (test (sut :i) 65079)
    (test (sut :x) 123)
    (test (sut :y) 456)))

(defn solve-1 [signals]
  (probe signals :a))

(defn solve-2 [p1 signals]
  (probe signals :a {:b [:plain p1]}))

(defn main [&]
  (let [signals (parse (file/read stdin :all))
        p1 (solve-1 signals)
        p2 (solve-2 p1 signals)]
    (pp {:p1 p1 :p2 p2})))
