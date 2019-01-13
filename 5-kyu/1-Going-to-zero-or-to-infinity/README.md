
# [![Going to zero or to infinity?](https://www.codewars.com/kata/55a29405bc7d2efaff00007c/go)](https://www.codewars.com/kata/55a29405bc7d2efaff00007c/go)

Description:

Consider the following numbers (where n! is factorial(n)):
```
u1 = (1 / 1!) * (1!)
u2 = (1 / 2!) * (1! + 2!)
u3 = (1 / 3!) * (1! + 2! + 3!)
un = (1 / n!) * (1! + 2! + 3! + ... + n!)
```
Which will win: `1 / n!` or `(1! + 2! + 3! + ... + n!)`?

Are these numbers going to `0` because of `1/n!` or to infinity due to the sum of factorials?
Task

Calculate `(1 / n!) * (1! + 2! + 3! + ... + n!)` for a given `n`, where `n` is an integer greater or equal to `1`.

To avoid discussions about rounding, return the result truncated to 6 decimal places, for example:

```
1.0000989217538616 will be truncated to 1.000098
1.2125000000000001 will be truncated to 1.2125
```

Remark

Keep in mind that factorials grow rather rapidly, and you need to handle large inputs.
