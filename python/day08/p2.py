from functools import reduce
from operator import mul

if __name__ == "__main__":
    with open("input.txt", "r") as f:

        def score(n: int, l: list[int]) -> int:
            for i in range(len(l)):
                if l[i] >= n:
                    return i + 1
            return i + 1

        inp = [list(map(int, line.strip())) for line in f.readlines()]
        max_y = len(inp)
        max_x = len(inp[0])
        max_score = -1

        for y in range(max_y - 1):
            for x in range(max_x - 1):
                if y == 0 or x == 0:
                    continue

                num = inp[y][x]

                n_score = score(num, [inp[i][x] for i in range(0, y)][::-1])
                e_score = score(num, inp[y][x + 1 :])
                s_score = score(num, [inp[i][x] for i in range(y + 1, max_y)])
                w_score = score(num, inp[y][:x][::-1])

                s = reduce(mul, (n_score, e_score, s_score, w_score))
                if s > max_score:
                    max_score = s

        print(max_score)
