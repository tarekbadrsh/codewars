
def get_grade(s1, s2, s3):
	total = s1+ s2+ s3

	score = total / 3

	if 90 <= score and score <= 100 :
		return 'A'
	
	if 80 <= score and score <= 90 :
		return 'B'
	
	if 70 <= score and score <= 80 :
		return 'C'
	
	if 60 <= score and score <= 70 :
		return 'D'
	
	if 0 <= score and score <= 60 :
		return 'F'
	
	return 'A'
