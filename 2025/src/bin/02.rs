use std::{fs::read_to_string, time::SystemTime};

const DAY: u8 = 2;
const EXAMPLE: &'static str = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124";

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

fn solve(input: String) -> (usize, usize) {
    let mut first: usize = 0;
    let mut second: usize = 0;

    input.trim().split(",").for_each(|e| {
        let ends = e.split("-").collect::<Vec<&str>>();
        (ends[0].parse::<usize>().unwrap()..(ends[1].parse::<usize>().unwrap() + 1)).for_each(
            |id| {
                if id < 10 {
                    return;
                }
                let id_length = (id.ilog10() + 1) as u8;
                let mut chunk_size = id_length / 2;

                while chunk_size >= 1 && chunk_size >= id_length / 2 {
                    // If not divisible into this size chunk skip loop
                    if id_length % chunk_size != 0 {
                        chunk_size -= 1;
                        continue;
                    }

                    let chunk_divisor = DIVISORS[chunk_size as usize];
                    let check_chunk = id % chunk_divisor;
                    let mut remaining_id = id / chunk_divisor;
                    while remaining_id > 0 {
                        let next_chunk = remaining_id % chunk_divisor;

                        if next_chunk != check_chunk {
                            break;
                        }

                        remaining_id /= chunk_divisor;
                    }
                    if remaining_id == 0 {
                        if chunk_size == id_length / 2 {
                            first += id as usize
                        }
                        second += id as usize;
                        break;
                    }

                    chunk_size -= 1
                }
            },
        )
    });

    (first, second)
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
