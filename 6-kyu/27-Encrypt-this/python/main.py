
def encrypt_this(text):

	words = text.split()
	result = []
	for w in words:
		nw = ""
		i = 0
		for letter in w:
			if i == 0:
				nw += str(ord(letter))
			elif i == 1 :
				nw += w[len(w)-1]
			elif i == (len(w) - 1):
				nw += w[1]
			else:
				nw += w[i]
			i += 1
		result.append(nw)
	return " ".join(result)
