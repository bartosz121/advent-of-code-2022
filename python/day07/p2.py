from collections import defaultdict

if __name__ == "__main__":
    with open("input.txt", "r") as f:
        fs = defaultdict(int)
        cwd = []
        for x in f.read().split("\n"):
            if x.startswith("$ cd"):
                dst: str = x[x.rindex(" ") + 1 :]
                if dst == "..":
                    cwd.pop()
                elif dst == "/":
                    cwd = []
                else:
                    cwd.append(dst)
            elif x.startswith("$ ls"):
                continue
            elif x.startswith("dir"):
                continue
            else:
                size, name = x.split(" ")
                for i in range(len(cwd) + 1):
                    fs["/".join(cwd[0:i])] += int(size)

    TOTAL_SIZE = 70_000_000
    TARGET = 30_000_000
    unused_space = TOTAL_SIZE - fs[""]
    needed = TARGET - unused_space

    smallest_directory_to_delete = TOTAL_SIZE

    for v in fs.values():
        if v >= needed and v <= smallest_directory_to_delete:
            smallest_directory_to_delete = v

    print(smallest_directory_to_delete)
