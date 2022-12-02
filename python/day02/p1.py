map_ = {
    "A": {
        "win": "C",
        "score": 1,
    },
    "B": {
        "win": "A",
        "score": 2,
    },
    "C": {
        "win": "B",
        "score": 3,
    },
}

if __name__ == "__main__":
    my_points = 0
    with open("input.txt", "r") as f:
        for game in f.readlines():
            game = game.strip()
            enemy, me = game.translate(
                game.maketrans({"Y": "B", "X": "A", "Z": "C"})
            ).split(" ")
            if enemy == me:
                my_points += 3 + map_[enemy]["score"]
            else:
                if map_[me]["win"] == enemy:
                    my_points += 6 + map_[me]["score"]
                else:
                    my_points += map_[me]["score"]

        print(my_points)
