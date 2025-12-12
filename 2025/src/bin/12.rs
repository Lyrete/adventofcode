use std::{fs::read_to_string, time::SystemTime};

const DAY: u8 = 12;
const EXAMPLE: &'static str = "0:
###
##.
##.

1:
###
##.
.##

2:
.##
###
##.

3:
##.
###
##.

4:
###
#..
###

5:
###
.#.
###

4x4: 0 0 0 0 2 0
12x5: 1 0 1 0 2 2
12x5: 1 0 1 0 3 2";

#[derive(Debug)]
struct PresentArea {
    width: u16,
    height: u16,
    values: Vec<u16>,
}

fn solve(input: String) -> (u32, u32) {
    let mut present_areas = Vec::new();
    let mut max_area: u16 = 0;

    for chunk in input.trim().split("\n\n") {
        if chunk.contains('x') {
            for line in chunk.split('\n') {
                let (dim_str, value_str) = line.split_once(':').unwrap();
                let (w, h) = dim_str.split_once('x').unwrap();
                let amounts = value_str
                    .trim()
                    .split_ascii_whitespace()
                    .map(|amt| amt.parse().unwrap())
                    .collect();

                present_areas.push(PresentArea {
                    width: w.parse().unwrap(),
                    height: h.parse().unwrap(),
                    values: amounts,
                });
            }
        } else {
            let (_, present_grid) = chunk.split_once(':').unwrap();
            let area = present_grid.chars().filter(|c| c != &'\n').count() as u16 - 1;
            if area > max_area {
                max_area = area;
            }
        }
    }

    // We presume that there is a way to find a pattern where the presents fit in the area regardless of their shape
    let res1 = present_areas
        .iter()
        .map(|area| {
            let max_area_needed: u16 = area.values.iter().map(|v| v * max_area).sum();
            if max_area_needed < area.height * area.width {
                return 1;
            }
            0
        })
        .sum();

    (res1, 0)
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
