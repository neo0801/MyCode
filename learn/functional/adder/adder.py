def adder():
    sum = 0

    def f(value):
        nonlocal sum
        sum += value
        return sum
    return f


ii = adder()
for i in range(11):
    print(f"1 + 2 + ... + {i} = {ii(i)}")
