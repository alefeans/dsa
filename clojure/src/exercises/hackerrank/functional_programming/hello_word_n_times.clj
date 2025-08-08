(ns exercises.hackerrank.functional-programming.hello-world-n-times)

(defn hello-word-n-times [n]
  (when-not (= n 0)
    (println "Hello World")
    (recur (dec n))))


(def n (Integer/parseInt (read-line)))
(hello-word-n-times n)
