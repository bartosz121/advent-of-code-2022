use std::{collections::HashSet, fs};

fn main() {
    let lines = fs::read_to_string("input.txt").unwrap();

    let x: Vec<_> = lines.trim().split("\n").collect();

    let result: u32 = x
        .chunks(3)
        .map(|games| {
            let gsr1: HashSet<char> = games.get(0).unwrap().chars().collect();
            let gsr2: HashSet<char> = games.get(1).unwrap().chars().collect();
            let gsr3: HashSet<char> = games.get(2).unwrap().chars().collect();

            let number: u32 = gsr1
                .intersection(&gsr2)
                .cloned()
                .collect::<HashSet<_>>()
                .intersection(&gsr3)
                .map(|c| {
                    let mut ord = *c as u32;
                    if ord > 96 {
                        ord -= 96
                    } else {
                        ord -= 38
                    }
                    ord
                })
                .sum();
            number
        })
        .sum();

    println!("{}", result)
}
