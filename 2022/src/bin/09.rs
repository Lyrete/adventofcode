use std::collections::HashSet;
use std::str::FromStr;

#[derive(Debug, PartialEq)]
enum Dir {
    Right,
    Left,
    Up,
    Down,
}

impl FromStr for Dir {
    type Err = ();

    fn from_str(input: &str) -> Result<Dir, Self::Err> {
        match input {
            "R" => Ok(Dir::Right),
            "L" => Ok(Dir::Left),
            "U" => Ok(Dir::Up),
            "D" => Ok(Dir::Down),
            _ => Err(()),
        }
    }
}

#[derive(Debug, PartialEq)]
struct Move {
    direction: Dir,
    amount: isize,
}

//Move tail towards the head
fn move_towards(head: &(isize, isize), tail: &mut (isize, isize), dir: &Dir) {
    let c_dist = (tail.0 - head.0).abs().max((tail.1 - head.1).abs());

    if c_dist <= 1 {
        return;
    }

    if head.0 < tail.0 && *dir == Dir::Left {
        tail.0 = head.0 + 1;
    }

    if head.0 > tail.0 && *dir == Dir::Right {
        tail.0 = head.0 - 1;
    }

    if head.1 < tail.1 && *dir == Dir::Down {
        tail.1 = head.1 + 1;
    }

    if head.1 > tail.1 && *dir == Dir::Up {
        tail.1 = head.1 - 1;
    }
}

fn main() {
    let input: Vec<_> = include_str!("./data9.txt")
        .lines()
        .map(|c| {
            c.split_once(" ")
                .map(|s| Move {
                    direction: Dir::from_str(s.0).unwrap(),
                    amount: s.1.parse::<isize>().unwrap(),
                })
                .unwrap()
        })
        .collect();

    let mut head = (0, 0);
    let mut tail = (0, 0);

    let mut visited: HashSet<(isize, isize)> = HashSet::new();
    visited.insert(tail);

    for move_daddy in input.iter() {
        if move_daddy.direction == Dir::Right {
            head.0 += move_daddy.amount;
            for _ in 0..move_daddy.amount {
                let c_dist = (tail.0 - head.0).abs().max((tail.1 - head.1).abs());
                if c_dist > 1 {
                    tail.1 = head.1;
                    tail.0 += 1;
                    visited.insert(tail);
                }
            }
        }

        if move_daddy.direction == Dir::Left {
            head.0 -= move_daddy.amount;
            for _ in 0..move_daddy.amount {
                let c_dist = (tail.0 - head.0).abs().max((tail.1 - head.1).abs());
                if c_dist > 1 {
                    tail.1 = head.1;
                    tail.0 -= 1;
                    visited.insert(tail);
                }
            }
        }

        if move_daddy.direction == Dir::Up {
            head.1 += move_daddy.amount;
            for _ in 0..move_daddy.amount {
                let c_dist = (tail.0 - head.0).abs().max((tail.1 - head.1).abs());
                if c_dist > 1 {
                    tail.0 = head.0;
                    tail.1 += 1;
                    visited.insert(tail);
                }
            }
        }

        if move_daddy.direction == Dir::Down {
            head.1 -= move_daddy.amount;
            for _ in 0..move_daddy.amount {
                let c_dist = (tail.0 - head.0).abs().max((tail.1 - head.1).abs());
                if c_dist > 1 {
                    tail.0 = head.0;
                    tail.1 -= 1;
                    visited.insert(tail);
                }
            }
        }
    }

    println!("Task 1: {:?}", visited.len());
}
