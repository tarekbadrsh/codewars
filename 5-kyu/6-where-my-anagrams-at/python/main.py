

def checkword(mainword, wordmap, checkword):
    if len(mainword) != len(checkword):
        return False
    tmpmap = {}
    for x in checkword:
        if not wordmap.get(x):
            return False
        if not tmpmap.get(x, None):
            tmpmap[x] = 0
        tmpmap[x] += 1
        if tmpmap[x] > wordmap[x]:
            return False
    return True


def anagrams(mainword, words):
    wordmap = {}
    for v in mainword:
        if not wordmap.get(v, None):
            wordmap[v] = 0
        wordmap[v] += 1
    result = []
    for word in words:
        angrm = checkword(mainword, wordmap, word)
        if angrm:
            result.append(word)
    if len(result) < 1:
        return []
    return result


print(anagrams("abba", ["aabb", "abcd", "bbaa", "dada"]))
