(defproject exercises "0.1.0-SNAPSHOT"
  :description "Clojure Exercises project"
  :url "https://github.com/alefeans/clojure-exercises"
  :license {:name "MIT License"}
  :dependencies [[org.clojure/clojure "1.10.0"]]
  :main ^:skip-aot exercises.core
  :target-path "target/%s"
  :profiles {:uberjar {:aot :all}})
