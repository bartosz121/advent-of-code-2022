from string import ascii_letters

if __name__ == "__main__":
    result = 0
    with open("input.txt", "r") as f:
        inp = [line.strip() for line in f.readlines()]
        groups = []
        for i in range(0, len(inp), 3):
            group_result = 0
            s1 = set(inp[i])
            s2 = set(inp[i + 1])
            s3 = set(inp[i + 2])

            for c in set.intersection(s1, s2, s3):
                group_result += ascii_letters.index(c) + 1
            result += group_result

        print(result)
