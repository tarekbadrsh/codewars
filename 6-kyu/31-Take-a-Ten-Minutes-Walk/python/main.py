
def is_valid_walk(walk):
    if len(walk) != 10:
        return False
    counter = {'n':0,'s':0,'w':0,'e':0}
    i = 0
    while i<len(walk):
        counter[walk[i]] += 1
        i += 1
    if counter['n']-counter['s'] == 0 and counter['w']-counter['e'] == 0:
        return True
    return False

 
print(is_valid_walk(['n','s','n','s','n','s','n','s','n','s']))

### ================ other practices ==================

def isValidWalkA(walk):
    return len(walk) == 10 and walk.count('n') == walk.count('s') and walk.count('e') == walk.count('w')