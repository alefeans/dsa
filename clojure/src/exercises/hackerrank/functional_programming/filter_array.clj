(ns exercises.hackerrank.functional-programming.filter-array)

(defn filter-array [delim lst]
  (filter #(> delim %1) lst))
