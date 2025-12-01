use std::fs::read_to_string;

const DAY: u8 = 1;
const EXAMPLE: &'static str = "
L68
L30
R48
L5
R60
L55
L1
L99
R14
L82
";
const START: i16 = 50;

fn solve(input: String) -> (u16, u16) {
    let instructions = input
        .split_ascii_whitespace()
        .map(|e| e.replace("L", "-").replace("R", "").parse::<i16>().unwrap())
        .collect::<Vec<i16>>();

    let mut zero_hits: u16 = 0;
    let mut passes: u16 = 0;
    instructions.iter().fold(START, |acc, elem| {
        let mut curr = (acc + elem) % 100;
        if curr < 0 {
            curr = 100 + curr
        }

        passes += (elem / 100).abs() as u16;

        if curr == 0 {
            zero_hits += 1
        } else if acc != 0 && elem < &0 && curr > acc || elem > &0 && curr < acc {
            // Really dumb check but i cba
            passes += 1;
        }
        curr
    });

    (zero_hits, zero_hits + passes)
}

fn main() {
    let example_res = solve(EXAMPLE.to_string());
    println!("Example:");
    println!("{:?} {:?}", example_res.0, example_res.1);

    let actual_res = solve(read_to_string(format!("./inputs/{:02}.txt", DAY)).unwrap());
    println!("Actual:");
    println!("{:?} {:?}", actual_res.0, actual_res.1);
}
