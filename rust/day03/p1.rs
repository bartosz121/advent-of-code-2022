use std::{collections::HashSet, fs};

fn main() {
    let lines = fs::read_to_string("input.txt").unwrap();

    let x: u32 = lines
        .trim()
        .split("\n")
        .map(|rucksack| {
            let (r1, r2) = rucksack.split_at(rucksack.len() / 2);
            let hsr1: HashSet<char> = r1.chars().collect();
            let hsr2: HashSet<char> = r2.chars().collect();
            let r: u32 = hsr1
                .intersection(&hsr2)
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

            r
        })
        .sum();

    println!("{}", x)
}
