mod utils;
use std::{collections::HashSet, fs::read_to_string, time::SystemTime};

use crate::utils::get_straight_line_dist;

const DAY: u8 = 8;
const EXAMPLE: &'static str = "162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689";

fn connect_points(points: &Vec<(isize, isize, isize)>, n: Option<usize>) -> (usize, isize) {
    let mut circuits: Vec<HashSet<(isize, isize, isize)>> = Vec::new();
    let mut found = Vec::new();

    // HashMap<(isize, isize, isize), Vec<f64, (isize, isize, isize)>>
    // let mut direct_connections = HashMap::new();

    for i in 0..points.len() {
        for j in i + 1..points.len() {
            let a = points[i];
            let b = points[j];

            found.push((a, b));
        }
    }

    found.sort_by(|a, b| {
        let dist1 = get_straight_line_dist(a.0, a.1) as usize;
        let dist2 = get_straight_line_dist(b.0, b.1) as usize;

        dist2.cmp(&dist1)
    });

    let oob_idx = found.len();
    let mut last_connection = ((0, 0, 0), (0, 0, 0));

    for _ in 0..n.unwrap_or(found.len()) {
        if circuits.len() > 0 && circuits[0].len() == points.len() {
            break;
        }

        let (a, b) = found.pop().unwrap();

        let mut a_idx = oob_idx;
        let mut b_idx = oob_idx;

        for i in 0..circuits.len() {
            if a_idx != oob_idx && b_idx != oob_idx {
                break;
            }

            if circuits[i].contains(&a) {
                a_idx = i;
            }

            if circuits[i].contains(&b) {
                b_idx = i;
            }
        }
        last_connection = (a, b);

        if a_idx == oob_idx && b_idx == oob_idx {
            let mut new_set = HashSet::new();
            new_set.insert(a);
            new_set.insert(b);
            circuits.push(new_set);
            continue;
        }

        if a_idx == oob_idx && b_idx != oob_idx {
            circuits[b_idx].insert(a);
            continue;
        }

        if b_idx == oob_idx && a_idx != oob_idx {
            circuits[a_idx].insert(b);
            continue;
        }

        if a_idx != b_idx {
            let removed = circuits.remove(b_idx);
            a_idx = if b_idx < a_idx { a_idx - 1 } else { a_idx };
            circuits[a_idx].extend(removed);
        }
    }

    let mut sizes = circuits.iter().map(|e| e.len()).collect::<Vec<usize>>();
    sizes.sort();
    let three_biggest_prod = sizes[sizes.len().saturating_sub(3)..].iter().product();

    (
        three_biggest_prod,
        last_connection.0.0 * last_connection.1.0,
    )
}

fn solve(input: String, n: usize) -> (usize, isize) {
    //let mut circuits: Vec<Vec<(isize, isize, isize)>> = Vec::new();

    let points = input
        .trim()
        .split('\n')
        .map(|e| {
            let mut split = e.split(',').map(|e| e.parse::<isize>().unwrap());

            (
                split.next().unwrap(),
                split.next().unwrap(),
                split.next().unwrap(),
            )
        })
        .collect::<Vec<(isize, isize, isize)>>();

    let (res1, _) = connect_points(&points, Some(n));
    let (_, res2) = connect_points(&points, None);

    (res1, res2)
}

fn main() {
    let example_res = solve(EXAMPLE.to_string(), 10);
    println!("Example:");
    println!("{:?} {:?}", example_res.0, example_res.1);

    let start = SystemTime::now();
    let actual_res = solve(
        read_to_string(format!("./inputs/{:02}.txt", DAY)).unwrap(),
        1000,
    );
    let end = SystemTime::now();

    println!("Actual:");
    println!("{:?} {:?}", actual_res.0, actual_res.1);

    println!();
    println!("Execution time: {:?}", end.duration_since(start).unwrap())
}
