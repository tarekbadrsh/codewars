

def strip_comments(input, markers):
    lines = input.split("\n")
    result = []
    for line in lines:
        line_result = line
        for mark in markers:
            if mark in line:
                tmp_result = line.split(mark)[0]
                if len(tmp_result) < len(line_result):
                    line_result = tmp_result
        result.append(line_result.strip())
    x = "\n".join(result)
    return x


def strip_comments2(input, markers):
    parts = input.split('\n')
    for s in markers:
        parts = [v.split(s)[0].rstrip() for v in parts]
    return '\n'.join(parts)


print(strip_comments("#my \napples, pears # and bananas\ngrapes\nbananas !apples",
                     ["#", "!"]))  # "\napples, pears\ngrapes\nbananas"

print(strip_comments("a #b\nc\nd $e f g", ["#", "$"]))  # "a\nc\nd"

print(strip_comments("@\n= - oranges watermelons watermelons =\nstrawberries",
      ["'", '-', '#', '!', '^', '@']))  # "\n=\nstrawberries"
