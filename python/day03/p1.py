from string import ascii_letters

if __name__ == "__main__":
    result = 0
    with open("input.txt", "r") as f:
        for line in f.read().split("\n"):
            s1 = set(line[: len(line) // 2])
            s2 = set(line[len(line) // 2 :])

            for c in s1.intersection(s2):
                result += ascii_letters.index(c) + 1
        print(result)
