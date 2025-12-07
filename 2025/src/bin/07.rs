use std::{fs::read_to_string, time::SystemTime};

const DAY: u8 = 7;
const EXAMPLE: &'static str = ".......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............";
const START: char = 'S';
const SPLITTER: char = '^';

fn solve(input: String) -> (usize, usize) {
    let mut start = 0;
    let mut splitters = Vec::new();
    let mut height = 0;
    let mut width = 0;

    input.trim().split('\n').enumerate().for_each(|(y, line)| {
        line.char_indices().for_each(|(x, c)| {
            match c {
                START => start = x,
                SPLITTER => splitters.push((x, y)),
                _ => {}
            };
        });

        width = line.len();
        height = y;
    });

    let mut beams = vec![0; width];
    beams[start] = 1;
    let mut res1 = 0;

    let mut y = 0;

    while y < height {
        let mut x = 0;
        while x < width {
            let curr = beams[x];
            if curr > 0 && splitters.contains(&(x, y)) {
                res1 += 1;
                beams[x] = 0;
                beams[x - 1] += curr;
                beams[x + 1] += curr;
            }

            x += 1;
        }

        y += 1;
    }

    let res2 = beams.iter().sum::<usize>();
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
