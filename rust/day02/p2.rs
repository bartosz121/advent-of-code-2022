use std::fs;

fn get_game_score(enemy: &str, me: &str) -> u32 {
    let mut score: u32 = 0;
    match enemy {
        // rock
        "A" => {
            if me == "X" {
                // lose
                score += 3; // scissors
            } else if me == "Y" {
                // draw
                score += 1; // rock
                score += 3; // draw
            } else {
                // win
                score += 2; // paper
                score += 6;
            }
        }
        // paper
        "B" => {
            if me == "X" {
                // lose
                score += 1; // rock
            } else if me == "Y" {
                // draw
                score += 2; // paper
                score += 3; // draw
            } else {
                // win
                score += 3; // scissors
                score += 6;
            }
        }
        // scissors
        "C" => {
            if me == "X" {
                // lose
                score += 2; // paper
            } else if me == "Y" {
                // draw
                score += 3; // scissors
                score += 3; // draw
            } else {
                // win
                score += 1; // rock
                score += 6;
            }
        }
        _ => {}
    }

    score
}

fn main() {
    let lines = fs::read_to_string("input.txt").unwrap();

    let result: u32 = lines
        .trim()
        .split("\n")
        .map(|game| {
            let mut split = game.split_whitespace();
            let enemy_move = split.nth(0).unwrap();
            let my_move = split.nth(0).unwrap();

            get_game_score(enemy_move, my_move)
        })
        .sum();

    println!("{:?}", result)
}
