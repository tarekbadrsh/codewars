

from itertools import permutations as pm
def permutations2(s): return set(map(''.join, set(pm(s))))


def add_element_to_list(array, lenght, element):
    result = []
    for i in range(lenght):
        for x in array:
            zz = x[:i] + element + x[i:]
            result.append(zz)
    return result


def permutations(input):
    result = input[0]
    for i in range(len(input[1:])):
        result = add_element_to_list(result, i+2, input[i+1])
    return set(result)


# print(permutations('aa'))  # ['aa']
# print(permutations('ab'))  # ['ab', 'ba']
# print(permutations('abc'))  # [cab', 'abc', 'acb', 'bac', 'bca', 'cba']
