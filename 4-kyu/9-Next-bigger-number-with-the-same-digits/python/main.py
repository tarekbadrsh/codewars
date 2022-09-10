
def next_bigger(n):
    if n < 12 :
        return -1
    can_be = False
    i = 0
    str_base = str(n)
    while i < len(str_base)-1:
        if str_base[i] < str_base[i+1]:
            can_be = True
            break
        i += 1
    if not can_be:
        return -1
    n_plus = n
    while True:
        n_plus +=1
        str_plus = str(n_plus)
        valid = True
        for v in str_base:
            if str_base.count(v) != str_plus.count(v):
                valid = False
                break
        if not valid:
            continue
        return n_plus

### ================ other practices ================== 


import itertools
def next_biggerA(n):
    s = list(str(n))
    for i in range(len(s)-2,-1,-1):
        if s[i] < s[i+1]:
            t = s[i:]
            m = min(filter(lambda x: x>t[0], t))
            t.remove(m)
            t.sort()
            s[i:] = [m] + t
            return int("".join(s))
    return -1
