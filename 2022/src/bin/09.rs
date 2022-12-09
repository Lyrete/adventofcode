use std::collections::HashSet;
use std::str::FromStr;
use std::time::Instant;

#[derive(Debug, PartialEq, Clone, Copy)]
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

#[derive(Debug, PartialEq, Clone, Copy)]
struct Move {
    direction: Dir,
    amount: isize,
}

//Move tail towards the head
fn move_towards(head: (isize, isize), tail: &mut (isize, isize)) {
    let dy = head.0 - tail.0;
    let dx = tail.1 - head.1;

    let c_dist = dy.abs().max(dx.abs());

    if c_dist <= 1 {
        return;
    }

    let edges = vec![
        //edges
        (head.0, head.1 - 1),
        (head.0, head.1 + 1),
        (head.0 + 1, head.1),
        (head.0 - 1, head.1),
        //corners
        (head.0 - 1, head.1 - 1),
        (head.0 - 1, head.1 + 1),
        (head.0 + 1, head.1 + 1),
        (head.0 + 1, head.1 - 1),
    ];

    for edge in edges.iter() {
        let dy = edge.0 - tail.0;
        let dx = edge.1 - tail.1;

        let c_dist = dy.abs().max(dx.abs());
        if c_dist == 1 {
            *tail = *edge;
            return;
        }
    }
}

fn move_head(head: &mut (isize, isize), dir: Dir) {
    let amt = if dir == Dir::Left || dir == Dir::Down {
        -1
    } else {
        1
    };

    if dir == Dir::Left || dir == Dir::Right {
        head.0 += amt;
    }

    if dir == Dir::Up || dir == Dir::Down {
        head.1 += amt;
    }
}

fn main() {
    let start = Instant::now();

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

    let mut visited: HashSet<(isize, isize)> = HashSet::new();

    let mut knots = vec![(0, 0); 10];
    let mut visited2: HashSet<(isize, isize)> = HashSet::new();

    for move_daddy in input.iter() {
        for _ in 0..move_daddy.amount {
            move_head(&mut knots[0], move_daddy.direction);
            for idx in 1..knots.len() {
                move_towards(knots[idx - 1], &mut knots[idx]);
            }
            visited.insert(knots[1]);
            visited2.insert(*knots.last().unwrap());
        }
    }

    println!("Task 1: {:?}", visited.len());
    println!("Task 2: {:?}", visited2.len());
    println!("Elapsed: {:.2?}", start.elapsed());
}
