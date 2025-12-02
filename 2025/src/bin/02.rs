use std::fs::read_to_string;

const DAY: u8 = 2;
const EXAMPLE: &'static str = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124";

fn solve(input: String) -> (usize, usize) {
    let mut first = 0;
    let mut second = 0;

    input.trim().split(",").for_each(|e| {
        let ends = e.split("-").collect::<Vec<&str>>();
        (ends[0].parse::<usize>().unwrap()..(ends[1].parse::<usize>().unwrap() + 1)).for_each(
            |id| {
                let id_length = (id as f64).log10().floor() as u32 + 1;
                let mut chunk_size: u32 = (id_length / 2) as u32;
                while chunk_size > 0 {
                    // If not divisible into this size chunk skip loop
                    if id_length % chunk_size != 0 {
                        chunk_size -= 1;
                        continue;
                    }

                    let chunk_divisor = 10_usize.pow(chunk_size);
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
                            first += id
                        }
                        second += id;
                        break;
                    }

                    chunk_size -= 1;
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

    let actual_res = solve(read_to_string(format!("./inputs/{:02}.txt", DAY)).unwrap());
    println!("Actual:");
    println!("{:?} {:?}", actual_res.0, actual_res.1);
}
