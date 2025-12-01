use std::fs::read_to_string;

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

fn task1(input: String) -> u16 {
    let instructions = input
        .split_ascii_whitespace()
        .map(|e| e.replace("L", "-").replace("R", "").parse::<i16>().unwrap())
        .collect::<Vec<i16>>();

    let mut zero_hits: u16 = 0;
    instructions.iter().fold(START, |acc, elem| {
        let curr = (acc + elem) % 100;
        if curr == 0 {
            zero_hits += 1
        }
        curr
    });
    zero_hits
}

fn task2(input: String) -> u16 {
    let instructions = input
        .split_ascii_whitespace()
        .map(|e| e.replace("L", "-").replace("R", "").parse::<i16>().unwrap())
        .collect::<Vec<i16>>();

    let mut zero_hits: u16 = 0;
    instructions.iter().fold(START, |acc, elem| {
        let new_pos: i16 = acc + elem;

        let mut curr: i16 = new_pos % 100;
        if curr < 0 {
            curr = 100 + curr
        }

        let mut extra_hits = (elem / 100).abs();

        if curr == 0 {
            extra_hits += 1;
        } else if acc != 0 && elem < &0 && curr > acc || elem > &0 && curr < acc {
            extra_hits += 1;
        }
        zero_hits += extra_hits as u16;
        curr
    });
    zero_hits
}

fn main() {
    println!("Example part 1: {:?}", task1(EXAMPLE.to_string()));
    println!(
        "Real input part 1: {:?}",
        task1(read_to_string("./inputs/01.txt").unwrap())
    );
    println!();

    println!("Example part 2: {:?}", task2(EXAMPLE.to_string()));
    println!(
        "Real input part 2: {:?}",
        task2(read_to_string("./inputs/01.txt").unwrap())
    );
}
