
def sum_of_intervals(intervals):
    intervals.sort(key=lambda x:x[0])
    sum = 0
    low = intervals[0][0]
    for v in intervals:
        if v[1] >= low:
            sum = sum + v[1]
            if v[0] > low:
                sum = sum - v[0]
            else:
                sum = sum - low
            low = v[1]
    return sum 

print(sum_of_intervals([(1, 4), (7, 10), (3, 5)]))
### ================ other practices ================== 

def sum_of_intervalsA(intervals):
    s, top = 0, float("-inf")
    for a,b in sorted(intervals):
        if top < a: top    = a
        if top < b: s, top = s+b-top, b
    return s