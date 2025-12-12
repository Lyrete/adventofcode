use std::{
    collections::{HashMap, VecDeque},
    fs::read_to_string,
    time::SystemTime,
};

const DAY: u8 = 11;
const EXAMPLE: &'static str = "svr: aaa bbb
aaa: fft
fft: ccc
bbb: tty
tty: ccc
ccc: ddd eee
ddd: hub
hub: fff
eee: dac
dac: fff
fff: ggg hhh
ggg: out
hhh: out";

fn count_possible_paths<'a>(
    wires: &HashMap<&'a str, Vec<&'a str>>,
    paths: &mut HashMap<&'a str, usize>, // Keep track of checked paths
    start: &'a str,
    end: &str,
) -> usize {
    // If we've counted the paths for this node before, no need to do it again.
    if let Some(p) = paths.get(start) {
        return *p;
    }

    // Reached end and found a path
    if start == end {
        return 1;
    }

    // Stop panics on wires that arent inputs to anything
    let total = match wires.get(start) {
        Some(vv) => vv
            .iter()
            .map(|next| count_possible_paths(wires, paths, next, end))
            .sum(),
        None => 0,
    };
    paths.insert(start, total);

    total
}

fn solve(input: String) -> (usize, usize) {
    let mut wires: HashMap<&str, Vec<&str>> = HashMap::new();

    input.trim().split('\n').for_each(|line| {
        let (input, output) = line.split_once(": ").unwrap();
        let outputs = output.split_ascii_whitespace().collect();
        wires.insert(input, outputs);
    });

    let res1 = count_possible_paths(&wires, &mut HashMap::new(), "you", "out");

    let res2 = count_possible_paths(&wires, &mut HashMap::new(), "svr", "fft")
        * count_possible_paths(&wires, &mut HashMap::new(), "fft", "dac")
        * count_possible_paths(&wires, &mut HashMap::new(), "dac", "out")
        + count_possible_paths(&wires, &mut HashMap::new(), "svr", "dac")
            * count_possible_paths(&wires, &mut HashMap::new(), "dac", "fft")
            * count_possible_paths(&wires, &mut HashMap::new(), "fft", "out");

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
