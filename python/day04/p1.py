with open("input.txt", "r") as f:
    result = 0
    for line in f.read().split("\n"):
        p1, p2 = map(
            lambda y: map(int, y), map(lambda x: x.split("-"), line.split(","))
        )

        # Add +1 to make it inclusive
        s1, e1 = p1
        e1 += 1
        s2, e2 = p2
        e2 += 1
        range1 = set(range(s1, e1))
        range2 = set(range(s2, e2))

        if range1.issubset(range2) or range2.issubset(range1):
            result += 1

    print(result)
