
def tops(msg):
    c = 1
    i = 0
    t = ""
    while i < len(msg): 
        i = i + c
        if i < len(msg):
            t = msg[i] + t
        c = c + 4
    return t

print(tops("abcdefghijklmnopqrstuvwxyz12345"))