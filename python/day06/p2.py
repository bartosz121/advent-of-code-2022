if __name__ == "__main__":
    with open("input.txt", "r") as f:
        inp = f.read()
        for i in range(4, len(inp)):
            if len(set(inp[i - 14 : i])) == 14:
                print(inp[i - 14 : i])
                print(i)
                break
