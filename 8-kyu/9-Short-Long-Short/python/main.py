
def solution(a, b):
    if len(a)<len(b):
        return f"{a}{b}{a}"
    return f"{b}{a}{b}"


