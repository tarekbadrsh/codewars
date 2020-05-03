
def map_vector(vector, circle1, circle2):
	r = circle2[2] / circle1[2]
	xp1 = circle1[0] - vector[0]
	yp1 = circle1[1] - vector[1]

	x2 = circle2[0] - (r * xp1)
	x2 = round(x2,2)

	y2 = circle2[1] - (r * yp1)
	y2 = round(y2,2)

	return (x2, y2)


### ================ other practices ==================

def map_vector2(vector, circle1, circle2):
    (vx, vy), (x1, y1, r1), (x2, y2, r2) = vector, circle1, circle2
    v = (complex(vx, vy) - complex(x1, y1)) / r1 * r2 + complex(x2, y2)
    return v.real, v.imag