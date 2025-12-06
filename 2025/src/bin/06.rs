use std::{fs::read_to_string, time::SystemTime};

const DAY: u8 = 6;
const EXAMPLE: &'static str = "123 328  51 64
 45 64  387 23
  6 98  215 314
*   +   *   +  ";

fn solve(input: String) -> (usize, usize) {
    let mut lines = input.trim().split('\n').collect::<Vec<&str>>();

    let operands = lines
        .pop()
        .unwrap()
        .split_ascii_whitespace()
        .map(|e| e.chars().next().unwrap())
        .collect::<Vec<char>>();

    let res1 = lines
        .iter()
        .map(|line| {
            line.split_ascii_whitespace()
                .map(|e| e.parse::<usize>().unwrap())
                .collect::<Vec<usize>>()
        })
        .reduce(|acc, line_vec| {
            acc.iter()
                .enumerate()
                .map(|(idx, acc_elem)| {
                    let next = line_vec[idx];

                    match operands[idx] {
                        '*' => return *acc_elem * next,
                        '+' => return *acc_elem + next,
                        _ => return *acc_elem,
                    }
                })
                .collect::<Vec<usize>>()
        })
        .unwrap()
        .iter()
        .sum();

    let mut numbers = vec![0; lines[0].len() + 1];

    lines.iter().for_each(|line| {
        line.char_indices().for_each(|(i, c)| {
            if c.is_digit(10) {
                let digit = c.to_digit(10).unwrap() as usize;
                numbers[i] = numbers[i] * 10 + digit;
            }
        });
    });

    let chunked = numbers.chunk_by(|a, _b| a != &0);

    let res3: usize = operands
        .iter()
        .zip(chunked)
        .map(|(op, chunk)| {
            //println!("{:?}", chunk);
            let filtered = chunk.iter().filter(|&e| *e != 0);

            match op {
                '*' => filtered.product(),
                '+' => filtered.sum(),
                _ => 0,
            }
        })
        .sum();

    (res1, res3)
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
