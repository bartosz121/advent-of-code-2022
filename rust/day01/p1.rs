use std::fs;

fn main() {
    let lines = fs::read_to_string("input.txt").unwrap();
    let result = lines
        .split("\n\n")
        .map(|x| {
            x.trim()
                .split("\n")
                .map(|n| n.parse::<u32>().unwrap())
                .sum::<u32>()
        })
        .max()
        .unwrap();

    println!("{:?}", result)
}
