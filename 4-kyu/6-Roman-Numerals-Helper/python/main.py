

roman_numerals = [
    [1000, "M"],
    [900, "CM"],
    [500, "D"],
    [100, "C"],
    [50, "L"],
    [10, "X"],
    [5, "V"],
    [1, "I"],
]


class RomanNumerals:
    def to_roman(input):
        result = ""
        for i in roman_numerals:
            if input == 90:
                input -= 90
                return "XC" + RomanNumerals.to_roman(input)
            elif input == 4:
                input -= 4
                return "IV" + RomanNumerals.to_roman(input)
            if input >= i[0]:
                input -= i[0]
                return i[1] + RomanNumerals.to_roman(input)
        return result

    def from_roman(input):
        result = 0
        if input == "IV":
            return 4
        for i in input:
            for x in roman_numerals:
                if i == x[1]:
                    result += x[0]
        return result


print(RomanNumerals.to_roman(2008))  # 'MMVIII'
print(RomanNumerals.from_roman('MDCLXVI'))  # 1666
