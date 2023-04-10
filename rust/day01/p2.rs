use std::fs;

fn main() {
    let lines = fs::read_to_string("input.txt").unwrap();
    let mut vec: Vec<u32> = lines
        .split("\n\n")
        .map(|x| {
            x.trim()
                .split("\n")
                .map(|n| n.parse::<u32>().unwrap())
                .sum::<u32>()
        })
        .collect();

    vec.sort_by(|a, b| b.cmp(a));

    let result: u32 = vec.iter().take(3).sum();

    println!("{:?}", result)
}
