import array, math
# Given integer x, this returns the integer floor(sqrt(x)).
def sqrt(x):
    assert x >= 0
    i = 1
    while i * i <= x:
        i *= 2
    y = 0
    while i > 0:
        if (y + i)**2 <= x:
            y += i
        i //= 2
    return y

# Returns a list of True and False indicating whether each number is prime.
# For 0 <= i <= n, result[i] is True if i is a prime number, False otherwise.
def list_primality(n):
    # Sieve of Eratosthenes
    result = [True] * (n + 1)
    result[0] = result[1] = False
    for i in range(sqrt(n) + 1):
        if result[i]:
            for j in range(i * i, len(result), i):
                result[j] = False
    return result


# Returns all the prime numbers less than or equal to n, in ascending order.
# For example: listPrimes(97) = [2, 3, 5, 7, 11, 13, 17, 19, 23, 29, ..., 83, 89, 97].
def list_primes(n):
    return [i for (i, isprime) in enumerate(list_primality(n)) if isprime]

def compute():
    isprime = list_primality(1000000)
    for i in range(len(isprime)):
        if not isprime[i]:
            continue

        n = [int(c) for c in str(i)]
        for mask in range(1 << len(n)):
            digits = do_mask(n, mask)
            count = 0
            for j in range(10):
                if digits[0] != 0 and isprime[to_number(digits)]:
                    count += 1
                digits = add_mask(digits, mask)

            if count == 8:
                digits = do_mask(n, mask)
                for j in range(10):
                    if digits[0] != 0 and isprime[to_number(digits)]:
                        return str(to_number(digits))
                    digits = add_mask(digits, mask)
    raise AssertionError("Not found")


def do_mask(digits, mask):
    return [d * ((~mask >> i) & 1) for (i, d) in enumerate(digits)]


def add_mask(digits, mask):
    return [d + ((mask >> i) & 1) for (i, d) in enumerate(digits)]


def to_number(digits):
    result = 0
    for d in digits:
        result = result * 10 + d
    return result


if __name__ == "__main__":
    print(compute())