
def decipher_this(string):
    words = string.split()
    result = []

    for w in words:
        nw = "" 
        ll = "" 
        i = 0 
        lenOfNumber = 0
        for letter in w :
            if letter.isdigit():
                ll += letter
                lenOfNumber +=1
            else:
                if ll != "":
                    nw += chr(int(ll))
                    ll = ""
                if i == lenOfNumber:
                    nw += w[len(w)-1]
                elif i == len(w)-1:
                    nw += w[lenOfNumber]
                else:
                    nw += w[i]

            if i == len(w)-1 and ll != "":
                nw += chr(int(ll))
            i += 1
        result.append(nw)

    return " ".join(result)



print(decipher_this("65 119esi 111dl 111lw 108dvei 105n 97n 111ka"))