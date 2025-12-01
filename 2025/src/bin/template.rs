use std::fs::read_to_string;

const DAY: u8 = 0;
const EXAMPLE: &'static str = "";

fn solve(input: String) -> (u32, u32) {
    (0, 0)
}

fn main() {
    let example_res = solve(EXAMPLE.to_string());
    println!("Example:");
    println!("{:?} {:?}", example_res.0, example_res.1);

    let actual_res = solve(read_to_string(format!("./inputs/{:02}.txt", DAY)).unwrap());
    println!("Actual:");
    println!("{:?} {:?}", actual_res.0, actual_res.1);
}
