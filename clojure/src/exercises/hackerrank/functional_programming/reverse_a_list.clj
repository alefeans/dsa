(ns exercises.hackerrank.functional-programming.reverse-a-list)

(defn reverse-a-list [lst]
  (reduce conj '() lst))
