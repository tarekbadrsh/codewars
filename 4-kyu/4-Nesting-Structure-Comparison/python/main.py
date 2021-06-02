
def create_string_arr(arr):
    builder = ""
    if isinstance(arr, list):
        if len(arr) == 0:
            return "[ ]"
        index = 0
        for x in arr:
            if isinstance(x, list):
                builder += create_string_arr(x)
                continue
            if index < len(arr) - 1:
                builder += "x,"
            else:
                builder += "x"
            index += 1
    return "[" + builder + "]"


def same_structure_as(original, other):
    original_build = create_string_arr(original)
    other_build = create_string_arr(other)
    return original_build == other_build


def same_structure_as2(original, other):
    if isinstance(original, list) and isinstance(other, list) and len(original) == len(other):
        for o1, o2 in zip(original, other):
            if not same_structure_as(o1, o2):
                return False
        else:
            return True
    else:
        return not isinstance(original, list) and not isinstance(other, list)


print(same_structure_as([1, 1, 1], [2, 2, 2]))
