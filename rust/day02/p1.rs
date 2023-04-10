use std::fs;

fn get_game_score(enemy: &str, me: &str) -> u32 {
    let mut score: u32 = 0;
    match enemy {
        "A" => {
            if me == "Y" {
                score += 2; // paper
                score += 6; // win
            } else if me == "X" {
                score += 1; // rock
                score += 3; // draw
            } else {
                score += 3; // scissors
            }
        }
        "B" => {
            if me == "Z" {
                score += 3; // scissors
                score += 6; // win
            } else if me == "Y" {
                score += 2; // paper
                score += 3; // draw
            } else {
                score += 1; // rock
            }
        }
        "C" => {
            if me == "X" {
                score += 1; // rock
                score += 6; // win
            } else if me == "Z" {
                score += 3; // scissors
                score += 3; // draw
            } else {
                score += 2; // paper
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
