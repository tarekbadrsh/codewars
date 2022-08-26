
def count_by(x, n):
    """
    Return a sequence of numbers counting by `x` `n` times.
    """
    res = []
    i = x
    while i <= (x*n):
        res.append(i)
        i+=x
    return res

