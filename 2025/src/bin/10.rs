use std::{
    collections::{HashMap, HashSet, VecDeque},
    fs::read_to_string,
    time::SystemTime,
};

use z3::{Solver, ast::Int};

const DAY: u8 = 10;
const EXAMPLE: &'static str = "[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}";

fn print_matrix(matrix: &[Vec<u16>]) {
    for line in matrix {
        for (i, e) in line.iter().enumerate() {
            if i == line.len() - 1 {
                print!("| ");
            }
            print!("{:?} ", e)
        }
        println!();
    }
}

fn find_unique_solutions_z3(buttons: &Vec<Vec<u16>>, solution: &[u16]) -> usize {
    let matrix = solution
        .iter()
        .enumerate()
        .map(|(s_idx, s)| {
            let mut row = buttons
                .iter()
                .map(|bi| if bi.contains(&(s_idx as u16)) { 1 } else { 0 })
                .collect::<Vec<u16>>();

            row.push(*s);
            row
        })
        .collect::<Vec<Vec<u16>>>();

    // print_matrix(&matrix);

    let solver = Solver::new();
    let constraints = (0..buttons.len())
        .map(|i| {
            let label = &format!("x{n}", n = i);
            let press_amt = Int::fresh_const(label);
            solver.assert(press_amt.ge(0));
            press_amt
        })
        .collect::<Vec<Int>>();

    for row in matrix {
        let meaningful_idx = row
            .iter()
            .enumerate()
            .filter(|(_, e)| *e == &1)
            .map(|(i, _)| i)
            .collect::<Vec<usize>>();

        let result = row.last().unwrap();
        solver.assert(
            constraints
                .iter()
                .enumerate()
                .filter_map(|(i, e)| {
                    if meaningful_idx.contains(&i) {
                        return Some(e);
                    } else {
                        None
                    }
                })
                .sum::<Int>()
                .eq(*result),
        );
    }

    let mut best_solution = usize::MAX;

    for solution in solver.solutions(constraints, false) {
        let solution = solution
            .iter()
            .map(Int::as_u64)
            .map(Option::unwrap)
            .sum::<u64>() as usize;
        if best_solution > solution {
            best_solution = solution;
        }
    }

    best_solution
}

fn find_shortest_sequence(buttons: &Vec<Vec<u16>>, desired_lights: &Vec<bool>) -> u32 {
    let mut q = VecDeque::new();

    (0..buttons.len()).for_each(|f| q.push_back(vec![f]));

    while q.len() > 0 {
        let path = &q.pop_front().unwrap();

        let mut state = vec![false; desired_lights.len()];
        for press in path {
            let button = &buttons[*press];
            button
                .iter()
                .for_each(|i| state[*i as usize] = !state[*i as usize]);

            if &state == desired_lights {
                return path.len() as u32;
            }
        }

        let last_press = path[path.len() - 1];

        (0..buttons.len())
            .filter(|next_press| *next_press != last_press)
            .for_each(|n| {
                let mut new_path = path.clone();
                new_path.push(n);
                q.push_back(new_path);
            });
    }

    0
}

fn solve(input: String) -> (u32, usize) {
    if input.len() == 0 {
        panic!("Input has no length");
    }

    let (res1, res2) = input
        .trim()
        .split('\n')
        .map(|line| {
            let parts = line.split_once(']').unwrap(); // Light terminator

            let desired_lights = parts.0[1..]
                .chars()
                .map(|e| match e {
                    '.' => false,
                    '#' => true,
                    _ => false,
                })
                .collect::<Vec<bool>>();

            let (buttonstr, joltage) = parts.1.split_once('{').unwrap(); // Start of joltage

            let desired_joltages = joltage[..joltage.len() - 1]
                .split(',')
                .map(|e| e.parse::<u16>().unwrap())
                .collect::<Vec<u16>>();

            let buttons: Vec<Vec<u16>> = buttonstr
                .split_ascii_whitespace()
                .map(|e| {
                    e[1..e.len() - 1]
                        .split(',')
                        .map(|n| n.parse::<u16>().unwrap())
                        .collect::<Vec<u16>>()
                })
                .collect();
            //let res2 = find_fewest_presses(&buttons, &desired_joltages[..]);
            let res2 = find_unique_solutions_z3(&buttons, &desired_joltages[..]);

            (find_shortest_sequence(&buttons, &desired_lights), res2)
        })
        .reduce(|acc, e| (acc.0 + e.0, acc.1 + e.1))
        .unwrap();

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
