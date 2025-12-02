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
                let num_str = id.to_string();
                let top_end = num_str.len();
                let mut chunk_size = top_end / 2;
                while chunk_size > 0 {
                    let chunk = &num_str[0..chunk_size];
                    let mut check_start = chunk_size;
                    while check_start + chunk_size <= top_end {
                        let check_chunk = &num_str[check_start..check_start + chunk_size];

                        if check_chunk != chunk {
                            break;
                        }

                        check_start += chunk_size
                    }
                    if check_start >= top_end {
                        if chunk_size == top_end / 2 {
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
