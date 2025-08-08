(ns exercises.hackerrank.functional-programming.absolute-value)

(defn abs-value
  [lst]
  (map #(Math/abs %1) lst))
