use std::{fs::read_to_string, time::SystemTime};

const DAY: u8 = 3;
const EXAMPLE: &'static str = "987654321111111
811111111111119
234234234234278
818181911112111";

static DIVISORS: &'static [usize] = &[
    1,
    10,
    100,
    1000,
    10000,
    100000,
    1000000,
    10000000,
    100000000,
    1000000000,
    10000000000,
    100000000000,
];

fn int_length(input: usize) -> u32 {
    if input == 0 {
        return 0;
    }
    input.ilog10() + 1
}

fn get_max_jolt(bank: &str, jolt_length: u8) -> usize {
    let mut result: usize = 0;
    let mut n = 0;

    for c in bank.chars() {
        let digit = c.to_digit(10).unwrap() as usize;
        let curr_length = int_length(result) as usize;

        let trunc = (0..curr_length).rev().find_map(|i| {
            let divisor = DIVISORS[i];
            let truncated = result / divisor;
            let search_digit = truncated % 10;
            let trunc_length = curr_length - i;

            if search_digit < digit && jolt_length as usize - trunc_length < bank.len() - n {
                return Some((truncated / 10) * 10 + digit);
            }

            None
        });

        if trunc.is_none() && curr_length < jolt_length as usize {
            result = result * 10 + digit
        } else if trunc.is_some() {
            result = trunc.unwrap();
        }

        n += 1;
    }

    return result;
}

fn solve(input: String) -> (usize, usize) {
    let banks = input.trim().split_ascii_whitespace().collect::<Vec<&str>>();

    let res1: usize = banks.iter().map(|bank| get_max_jolt(bank, 2)).sum();
    let res2: usize = banks.iter().map(|bank| get_max_jolt(bank, 12)).sum();
    // let res2 = 0;
    // let res1 = 0;

    (res1, res2)
}

fn main() {
    let example_res = solve(EXAMPLE.to_string());
    println!("Example:");
    println!("{:?} {:?}", example_res.0, example_res.1);

    let start = SystemTime::now();
    let actual_res = solve(read_to_string(format!("./inputs/{:02}.txt", DAY)).unwrap());
    let end = SystemTime::now();

    println!("Actual:");
    println!("{:?} {:?}", actual_res.0, actual_res.1);

    println!();
    println!("Execution time: {:?}", end.duration_since(start).unwrap())
}
