
def get_middle(s):
    n = len(s)
    if n %2 == 1:
        return s[int(n/2)]
    return str(s[int(n/2-1) : int(n/2+1)])
