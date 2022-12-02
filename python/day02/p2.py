map_ = {
    "A": {
        "win": "C",
        "lose": "B",
        "score": 1,
    },
    "B": {
        "win": "A",
        "lose": "C",
        "score": 2,
    },
    "C": {
        "win": "B",
        "lose": "A",
        "score": 3,
    },
}

if __name__ == "__main__":
    with open("input.txt", "r") as f:
        my_points = 0
        for game in f.readlines():
            game = game.strip()
            enemy, me = game.translate(
                game.maketrans({"Y": "B", "X": "A", "Z": "C"})
            ).split(" ")
            if me == "B":
                my_points += 3 + map_[enemy]["score"]
            elif me == "C":
                my_points += 6 + map_[map_[enemy]["lose"]]["score"]
            elif me == "A":
                my_points += map_[map_[enemy]["win"]]["score"]

        print(my_points)
