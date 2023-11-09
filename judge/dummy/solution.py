def generate_fibonacci(n):
    if n == 0:
        return []
    if n == 1:
        return [0]
    fib_series = [0, 1]
    while len(fib_series) < n:
        fib_series.append(fib_series[-1] + fib_series[-2])
    return fib_series
