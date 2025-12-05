use std::{fs::read_to_string, time::SystemTime};

const DAY: u8 = 5;
const EXAMPLE: &'static str = "3-5
10-14
16-20
4-18

1
5
8
11
17
32";

fn solve(input: String) -> (usize, usize) {
    let parts = input.trim().split_once("\n\n").unwrap();

    let ranges = parts
        .0
        .split('\n')
        .map(|line| line.split_once('-').unwrap())
        .collect::<Vec<(&str, &str)>>();
    let mut valid_ranges = Vec::new();

    ranges.iter().for_each(|range| {
        let start = range.0.parse::<usize>().unwrap();
        let end = range.1.parse::<usize>().unwrap();

        let end_contained_idx = valid_ranges
            .iter()
            .position(|(valid_start, valid_end)| *valid_start <= end && *valid_end >= end);

        let start_contained_idx = valid_ranges
            .iter()
            .position(|(valid_start, valid_end)| *valid_start <= start && *valid_end >= start);

        let mut new_range = (start, end);

        if end_contained_idx.is_some() && start_contained_idx.is_some() {
            let end_idx = end_contained_idx.unwrap();
            let start_idx = start_contained_idx.unwrap();

            //idx are the same so the new range was fully contained in an existing one
            if end_idx == start_idx {
                return;
            }

            let (_, curr_end) = valid_ranges[end_idx];
            let (curr_start, _) = valid_ranges[start_idx];
            new_range = (curr_start, curr_end);
            valid_ranges.remove(start_idx);
            valid_ranges.remove(if start_idx < end_idx {
                end_idx - 1
            } else {
                end_idx
            }); // idx might change by one because of prev removal
        } else if end_contained_idx.is_some() {
            let end_idx = end_contained_idx.unwrap();
            let (_, curr_end) = valid_ranges[end_idx];
            new_range = (start, curr_end);
            valid_ranges.remove(end_idx);
        } else if start_contained_idx.is_some() {
            let start_idx = start_contained_idx.unwrap();
            let (curr_start, _) = valid_ranges[start_idx];
            new_range = (curr_start, end);
            valid_ranges.remove(start_idx);
        }

        valid_ranges = valid_ranges
            .iter()
            .filter(|(comp_start, comp_end)| {
                !(comp_start >= &new_range.0 && comp_end <= &new_range.1)
            })
            .map(|p| *p)
            .collect();

        valid_ranges.push(new_range);
    });

    let res1 = parts
        .1
        .split('\n')
        .map(|e| e.parse::<usize>().unwrap())
        .filter(|item| {
            valid_ranges
                .iter()
                .any(|(start, end)| start <= item && end >= item)
        })
        .count();

    //println!("{:?}", valid_ranges);

    let res2 = valid_ranges
        .iter()
        .map(|(start, end)| end - start + 1)
        .sum();

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
