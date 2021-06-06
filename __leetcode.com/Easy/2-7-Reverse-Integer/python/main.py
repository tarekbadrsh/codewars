def reverse(x: int) -> int:
    res = 0
    if x < 0:
        res = -int(str(str(x)[1:])[::-1])
        if res < -2147483648:
            return 0
    else:
        res = int(str(x)[::-1])
        if res > 2147483648:
            return 0
    return res


print(reverse(123))
print(reverse(-123))
print(reverse(120))
print(reverse(0))
print(reverse(1534236469))
