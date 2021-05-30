
def are_opposite(e1, e2): 
	if e1 == "NORTH" and e2 == "SOUTH":
			return True
	if e1 == "SOUTH" and e2 == "NORTH":
			return True
	if e1 == "EAST" and e2 == "WEST":
			return True
	if e1 == "WEST" and e2 == "EAST": 
			return True 
	return False 


def dirReduc(arr):
    if len(arr) > 1:
        for i in range(len(arr)-1):
            if are_opposite(arr[i], arr[i+1]):
                newarr = arr[:i] + arr[i+2:]
                return dirReduc(newarr)
    return arr

print(dirReduc(["NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "WEST"]))
