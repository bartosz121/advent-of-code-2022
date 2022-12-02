if __name__ == "__main__":
    with open("input.txt", "r") as f:
        x = [sum(map(int, line.splitlines())) for line in f.read().split("\n\n")]
        x.sort(reverse=True)
        print(sum(x[:3]))