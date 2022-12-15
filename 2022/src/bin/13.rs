use std::collections::VecDeque;

fn find_parentheses(input: &str) -> VecDeque<(usize, usize)> {
    let mut opens = Vec::new();
    let mut pairs = VecDeque::new();

    input.chars().enumerate().for_each(|(i, c)| {
        if c == '[' {
            opens.push(i)
        } else if c == ']' {
            pairs.push_back((opens.pop().unwrap(), i))
        }
    });

    pairs
}

fn main() {
    let pairs: Vec<_> = include_str!("./ex13.txt")
        .split("\n\n")
        .map(|x| x.trim().split_once("\n").unwrap())
        .collect();

    let mut goodpairs: Vec<usize> = Vec::new();
    for (idx, pair) in pairs.iter().enumerate().map(|(i, x)| (i + 1, x)) {
        let mut left_str: String = pair.0.to_string();
        let mut right_str: String = pair.1.to_string();
        let (mut left_p, mut right_p) = (find_parentheses(pair.0), find_parentheses(pair.1));

        let checks: Vec<bool> = Vec::new();

        while left_p.len() > 0 && right_p.len() > 0 {
            let (left, right) = (left_p.pop_front().unwrap(), right_p.pop_front().unwrap());

            let left_values: Vec<isize> = left_str[left.0 + 1..left.1]
                .split(",")
                .map(|x| x.parse::<isize>().unwrap_or(-1))
                .filter(|x| *x >= 0)
                .collect();

            let right_values: Vec<isize> = right_str[right.0 + 1..right.1]
                .split(",")
                .map(|x| x.parse::<isize>().unwrap_or(-1))
                .filter(|x| *x >= 0)
                .collect();

            println!("{:?} vs {:?}", left_values, right_values);
        }

        println!("{:?} - {:?}", left_p, right_p);
        println!();
    }

    println!("Result: {:?}", goodpairs);
}
