use std::{collections::HashSet, fs::read_to_string, time::SystemTime};

const DAY: u8 = 4;
const EXAMPLE: &'static str = "..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.";

static PAPER_ROLL: char = '@';
const MODIFIERS: &'static [(isize, isize)] = &[
    (1, 0),
    (1, 1),
    (0, 1),
    (-1, 1),
    (-1, 0),
    (-1, -1),
    (0, -1),
    (1, -1),
];

fn get_neighbours(point: &(isize, isize)) -> Vec<(isize, isize)> {
    MODIFIERS
        .iter()
        .map(|(dx, dy)| {
            let x2 = point.0 + dx;
            let y2 = point.1 + dy;

            (x2, y2)
        })
        .collect()
}

fn solve(input: String) -> (usize, usize) {
    let mut paper_rolls: HashSet<(isize, isize)> = input
        .trim()
        .split_ascii_whitespace()
        .enumerate()
        .flat_map(|(y, line)| {
            line.chars()
                .enumerate()
                .filter_map(|(x, c)| {
                    if c == PAPER_ROLL {
                        return Some((x as isize, y as isize));
                    }
                    None
                })
                .collect::<HashSet<(isize, isize)>>()
        })
        .collect();

    let mut removable: HashSet<(isize, isize)> = paper_rolls
        .iter()
        .filter(|point| {
            get_neighbours(point)
                .iter()
                .filter(|point| paper_rolls.contains(point))
                .count()
                < 4
        })
        .map(|point| *point)
        .collect();

    let res1 = removable.len();
    let mut res2 = res1;

    while removable.len() > 0 {
        removable.iter().for_each(|point| {
            paper_rolls.remove(point);
        });

        removable = removable
            .iter()
            .flat_map(|point| {
                get_neighbours(point)
                    .iter()
                    .filter(|point| paper_rolls.contains(point))
                    .map(|p| *p)
                    .collect::<Vec<(isize, isize)>>()
            })
            .filter(|point| {
                get_neighbours(point)
                    .iter()
                    .filter(|point| paper_rolls.contains(point))
                    .count()
                    < 4
            })
            .collect();

        res2 += removable.len()
    }

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
    println!("Execution time: {:?}", end.duration_since(start).unwrap());
}
