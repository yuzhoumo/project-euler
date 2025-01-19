import math
from functools import reduce


#########
# UTILS #
#########


def fiberator():
    prev, curr = 0, 1
    while True:
        yield curr
        prev, curr = curr, prev + curr


def sieve(n):
    primes = [True] * n
    primes[0] = primes[1] = False

    for i in range(n):
        if not primes[i]: continue

        yield i
        for j in range(i * i, n, i):
            primes[j] = False


def is_prime(n):
    if n < 2: return False
    if n == 2: return True
    for i in range(3, int(math.sqrt(n)) + 1, 2):
        if not (n % i):
            return False
    return True


def is_palindrome_number(n):
    orig, comp = n, 0
    while n:
        comp *= 10
        comp += n % 10
        n //= 10
    return orig == comp


def is_palindrome_binary(n):
    orig, comp = n, 0
    while n:
        comp <<= 1
        comp |= (n & 1)
        n >>= 1
    return orig == comp


def gcd(x, y):
    if x == 0: return y
    return gcd(y % x, x)


def lcm(x, y):
    return x * y // gcd(x, y)


def factorial(n):
    if n <= 1: return 1
    return reduce(lambda a, b: a * b, range(2, n + 1))


def number_to_digits(n):
    res = []
    while n:
        res.append(n % 10)
        n //= 10
    res.reverse()
    return res


def digits_to_number(digits):
    n = 0
    for place, d in enumerate(reversed(digits)):
        n += d * (10 ** place)
    return n


#############
# SOLUTIONS #
#############


def e1():
    return sum(filter(lambda n: not (n % 5 and n % 3), range(1, 1000)))


def e2():
    tot = 0
    for n in fiberator():
        if n > 4_000_000: return tot
        if n & 1: continue
        tot += n


def e3():
    res, target = None, 600851475143
    for n in sieve(int(math.sqrt(target) + 1)):
        if target % n == 0: res = n
    return res


def e4():
    res = 0
    for i in range(100, 1000):
        for j in range(i, 1000):
            n = i * j
            if is_palindrome_number(n):
                res = max(res, n)
    return res


def e5():
    return reduce(lcm, range(1, 21))


def e6():
    s = (100 * (100 - 1)) // 2
    return s * s - sum(map(lambda x: x * x, range(101)))


def e7():
    iter = sieve(1000000)
    for _ in range(10000):
        next(iter)
    return next(iter)


def e8():
    n = "73167176531330624919225119674426574742355349194934" + \
        "96983520312774506326239578318016984801869478851843" + \
        "85861560789112949495459501737958331952853208805511" + \
        "12540698747158523863050715693290963295227443043557" + \
        "66896648950445244523161731856403098711121722383113" + \
        "62229893423380308135336276614282806444486645238749" + \
        "30358907296290491560440772390713810515859307960866" + \
        "70172427121883998797908792274921901699720888093776" + \
        "65727333001053367881220235421809751254540594752243" + \
        "52584907711670556013604839586446706324415722155397" + \
        "53697817977846174064955149290862569321978468622482" + \
        "83972241375657056057490261407972968652414535100474" + \
        "82166370484403199890008895243450658541227588666881" + \
        "16427171479924442928230863465674813919123162824586" + \
        "17866458359124566529476545682848912883142607690042" + \
        "24219022671055626321111109370544217506941658960408" + \
        "07198403850962455444362981230987879927244284909188" + \
        "84580156166097919133875499200524063689912560717606" + \
        "05886116467109405077541002256983155200055935729725" + \
        "71636269561882670428252483600823257530420752963450"

    res, N = 0, len(n)
    for i in range(N-13):
        product = reduce(lambda a, b: a * b, map(int, list(n[i:i+13])))
        res = max(res, product)

    return res


def e9():
    for a in range(1, 1001):
        for b in range(a + 1, 1001 - a):
            c = 1000 - a - b
            aa, bb, cc = a * a, b * b, c * c
            if a + b + c == 1000 and aa + bb == cc:
                return a * b * c


def e10():
    return sum(sieve(2_000_000))


def e30():
    res, n = 0, 2
    while n < 1_000_000:
        sum_of_fifth_powers = sum(map(lambda n: n ** 5, number_to_digits(n)))
        if sum_of_fifth_powers == n:
            res += n
        n += 1
    return res


def e34():
    factorials = {i: factorial(i) for i in range(10)}

    res = 0
    for n in range(3, 10_000_000):
        fac_sum_digits, tmp = 0, n

        while tmp:
            fac_sum_digits += factorials[tmp % 10]
            tmp //= 10

        if fac_sum_digits == n:
            res += n

    return res


def e36():
    res = 0
    for n in range(1_000_000):
        if is_palindrome_binary(n) and is_palindrome_number(n):
            res += n
    return res


if __name__ == "__main__":
    solutions = {
        1:  e1,
        2:  e2,
        3:  e3,
        4:  e4,
        5:  e5,
        6:  e6,
        7:  e7,
        8:  e8,
        9:  e9,
        10: e10,
        30: e30,
        34: e34,
        36: e36,
    }

    for i, f in solutions.items():
        print(f"{i: 5} {f()}")

