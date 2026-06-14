



a = [2, 4, 3]
b = [5, 6, 4]

a = int(''.join(str(d) for d in a[::-1]))
b = int(''.join(str(d) for d in b[::-1]))
result = [int(d) for d in str(a+b)[::-1]]
print(result)

