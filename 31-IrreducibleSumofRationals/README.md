# [![Irreducible Sum of Rationals](https://www.codewars.com/kata/5517fcb0236c8826940003c9)](https://www.codewars.com/kata/5517fcb0236c8826940003c9)


Description:

You will have a list of rationals in the form
```
lst = [ [numer_1, denom_1] , ... , [numer_n, denom_n] ]
```
or
```
lst = [ (numer_1, denom_1) , ... , (numer_n, denom_n) ]
```
where all numbers are positive integers. You have to produce their sum `N / D` in an irreducible form: this means that `N` and `D` have only `1` as a common divisor.

Return the result in the form:

* `[N, D]` in Ruby, Python, Clojure, JS, CS, PHP, Julia
* `Just "N D"` in Haskell
* `"[N, D]"` in Java, CSharp, TS, Scala, PowerShell
* `{N, D}` in C++, Elixir
* `Some((N, D))` in Rust
* `Some "N D"` in F#, Ocaml
* `c(N, D)` in R

If the result is an integer (`D` evenly divides `N`) return:

* an integer in Ruby, Elixir, Clojure, Python, JS, CS, PHP, R, Julia
* `Just "n"` (Haskell)
* `"n"` Java, CSharp, TS, Scala, PowerShell
* `{n, 1}` in C++, C
* `Some((n, 1))` in Rust
* `Some "n"` in F#, Ocaml,
* `(n, 1)` in Swift

If the input list is empty, return `nil/None/null/Nothing` (or `{0, 1}` in C++, C) (or `"0"` in Scala, PowerShell)

## Example:

```
[ [1, 2], [1, 3], [1, 4] ]  -->  [13, 12]

    1/2  +  1/3  +  1/4     =      13/12
```

See sample tests for more examples.
