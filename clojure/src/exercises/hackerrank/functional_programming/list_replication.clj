(ns exercises.hackerrank.functional-programming.list-replication)

(defn list-replication
  [num lst]
  (flatten (map #(repeat num %1) lst)))
