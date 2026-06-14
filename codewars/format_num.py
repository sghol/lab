#
# def create_phone_number(numbers):
#     s = "".join(str(n) for n in numbers)
#     part1 = "(" + s[0:3] + ")"
#     part2 = " " + s[3:6] + "-"
#     part3 = s[6:]
#     result = part1 + part2 + part3
#     return result


def create_phone_number(numbers):
    return "({}{}{}) {}{}{}-{}{}{}{}".format(*numbers)


# => returns "(123) 456-7890"
x = create_phone_number([1, 2, 3, 4, 5, 6, 7, 8, 9, 0])
print(x)
