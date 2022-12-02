if __name__ == "__main__":
    with open("input.txt", "r") as f:
        print(max([sum(map(int, line.splitlines())) for line in f.read().split("\n\n")]))