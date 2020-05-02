
def abbrevName(name):
    for i in range(len(name)):
        if name[i] == " ":
            return name[0].upper() + "." +  name[i+1].upper()

### ================ best practices ==================###

def abbrevNameB(name):
    names = name.split()
    return f"{names[0][0]}.{names[1][0]}".upper()

