def is_even(n):
    return n % 2 == 0


def test_is_even():
    test_cases = [(4, True), (3, False), (3, False), (5, False), (12, True)]
    for number, result in test_cases:
        assert is_even(number) == result, f"({number}, {result})"
        print("Test passed.")


test_is_even()
