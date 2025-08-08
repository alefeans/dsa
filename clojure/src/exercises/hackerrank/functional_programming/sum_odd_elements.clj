(ns exercises.hackerrank.functional-programming.sum-odd-elements)

(defn sum-odd-elements
  [lst]
  (reduce + (filter odd? lst)))
