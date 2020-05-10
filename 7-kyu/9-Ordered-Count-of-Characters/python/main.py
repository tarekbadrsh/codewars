
def ordered_count(inp):
	dic = {}
	order = []
	ret = []

	for c in inp : 
		if  dic.get(c,0) == 0 :
			order.append(c)
		dic[c] = dic.get(c, 0) + 1

	for c in order :
		ret.append((c, dic[c]))
	
	return ret 
