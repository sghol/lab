def bellard_pi_approximation(num_terms):
    """
    Approximates pi using Bellard's formula by summing the first 'num_terms' of the series.
    """
    pi_approx = 0.0

    for n in range(num_terms):
        # Calculate the sign for this term: (-1)^n
        sign = 1 if n % 2 == 0 else -1

        # Calculate the common denominator part: 1 / 2^(10n)
        denominator_2_power = 2 ** (10 * n)

        # Calculate the massive parentheses with all the fractions
        term_in_parentheses = (
            -(2**5) / (4 * n + 1)  # -32/(4n+1)
            - 1 / (4 * n + 3)  # -1/(4n+3)
            + (2**8) / (10 * n + 1)  # +256/(10n+1)
            - (2**6) / (10 * n + 3)  # -64/(10n+3)
            - (2**2) / (10 * n + 5)  # -4/(10n+5)
            - (2**2) / (10 * n + 7)  # -4/(10n+7)
            + 1 / (10 * n + 9)  # +1/(10n+9)
        )

        # Calculate the full term for this n
        term = sign * term_in_parentheses / denominator_2_power
        pi_approx += term

    # Multiply by the final constant: 1 / 2^6 = 1/64
    pi_approx /= 64
    return pi_approx


# How many terms to use for the approximation
number_of_terms = 5

# Calculate and print the result
approximation = bellard_pi_approximation(number_of_terms)
print(f"Pi approximated with {number_of_terms} terms: {approximation}")
print(f"Comparison to math.pi:  {approximation - 3.141592653589793}")
