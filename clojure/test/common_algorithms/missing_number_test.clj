(ns exercises.common-algorithms.missing-number-test
  (:require [clojure.test :refer :all]
            [exercises.common-algorithms.missing-number :as m]))

(deftest missing-number-test
  (testing "If returns the missing number in an array"
    (is (= 3
           (m/missing-number [1 2 0 4])))
    (is (= 5
           (m/missing-number [1 2 3 4 0 6 7])))))
