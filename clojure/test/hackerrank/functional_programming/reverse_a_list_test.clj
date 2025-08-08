(ns exercises.hackerrank.functional-programming.reverse-a-list-test
  (:require [clojure.test :refer :all]
            [exercises.hackerrank.functional-programming.reverse-a-list :as r]))

(deftest reverse-a-list
  (testing "If returns the reversed list"
    (is (= '(4 3 2 1) (r/reverse-a-list [1 2 3 4])))))
