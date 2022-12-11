if __name__ == "__main__":
    with open("input.txt", "r") as f:
        inp = [list(map(int, line.strip())) for line in f.readlines()]
        visible = set()
        max_y = len(inp)
        max_x = len(inp[0])

        for y in range(max_y):
            for x in range(max_x):
                if x == 0 or x == max_x - 1 or y == 0 or y == max_y - 1:
                    visible.add((y, x))
                    continue

                n = max((inp[i][x] for i in range(0, y)))
                e = max(inp[y][x + 1 :])
                s = max((inp[i][x] for i in range(y + 1, max_y)))
                w = max(inp[y][:x])

                if any((i < inp[y][x] for i in (n, e, s, w))):
                    visible.add((y, x))

        print(len(visible))
