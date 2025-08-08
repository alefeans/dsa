(ns exercises.hackerrank.functional-programming.filter-positions)

(defn filter-positions
  [lst]
  (take-nth 2 (rest lst)))
