Numbers = list[int]


def binary_search1(elements: Numbers, target: int):
    """If the target be in the elements return the index of the element otherwise -1."""
    for index, element in enumerate(elements):
        if element == target:
            return index
    return -1


def binary_search(elements: Numbers, target: int):
    """If the target be in the elements return the index of the element otherwise -1."""
    i = 0
    found = False
    while i < len(elements):
        if elements[i] == target:
            found = True
            break
        i += 1
    if found:
        return i
    else:
        return -1


def test_binary_search():

    test_cases = [
        ([1, 2, 3, 4, 5], 3, 2),
        ([10, 20, 1], 5, -1),
        ([], 5, -1),
        ([2, 7, 1], 7, 1),
        ([1], 1, 0),
    ]

    x = 0
    for numbers, target, result in test_cases:
        print(f"Exemening Test #{x}")
        assert binary_search(numbers, target) == result
        print(f"Test #{x} passed.")
        x += 1


def bubble_sort(elements: Numbers):
    """sort elements and return sorted arry."""
    i = 0
    j = 1
    while i < len(elements):
        print(elements)
        current_element = elements[i]
        while j < len(elements) - i:
            next_element = elements[j]
            if next_element < current_element:
                elements[i] = next_element
                elements[j] = current_element
            j += 1
        i += 1
    print(elements)
    return elements


def test_bubble_sort():
    test_cases = [
        ([4, 3, 2, 1], [1, 2, 3, 4]),
        ([7, 3, 9, 1], [1, 3, 7, 9]),
        ([], []),
        ([0], [0]),
    ]

    x = 0
    for source, result in test_cases:
        print(f"Exemening Test #{x}")
        assert bubble_sort(source) == result
        print(f"Test #{x} passed.")
        x += 1

test_bubble_sort()
