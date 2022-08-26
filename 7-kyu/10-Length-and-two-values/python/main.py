
def alternate(n, first_value, second_value):
    i = 0
    slice = []
    for i in range(n):
        if i%2==0:
            slice.append(first_value)
        elif i%2==1:
            slice.append(second_value)
        i+=1
    return slice


