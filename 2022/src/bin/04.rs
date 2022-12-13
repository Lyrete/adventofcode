use std::collections::HashSet;

fn task1() {
    let subsets = include_str!("./data.txt")
        .lines()
        .map(|n| n.split_once(",").unwrap())
        .map(|(r1, r2)| (r1.split_once("-").unwrap(), r2.split_once("-").unwrap()))
        .map(|(r1, r2)| -> (HashSet<i32>, HashSet<i32>) {
            (
                (HashSet::from_iter(
                    r1.0.parse::<i32>().unwrap()..r1.1.parse::<i32>().unwrap() + 1,
                )),
                (HashSet::from_iter(
                    r2.0.parse::<i32>().unwrap()..r2.1.parse::<i32>().unwrap() + 1,
                )),
            )
        })
        .map(|(r1, r2)| (r1.is_subset(&r2) || r1.is_superset(&r2)))
        .filter(|n| *n)
        .count();

    println!("First task: {:?}", subsets);
}

fn task2() {
    let subsets = include_str!("./data.txt")
        .lines()
        .map(|n| n.split_once(",").unwrap())
        .map(|(r1, r2)| (r1.split_once("-").unwrap(), r2.split_once("-").unwrap()))
        .map(|(r1, r2)| -> (HashSet<i32>, HashSet<i32>) {
            (
                (HashSet::from_iter(
                    r1.0.parse::<i32>().unwrap()..r1.1.parse::<i32>().unwrap() + 1,
                )),
                (HashSet::from_iter(
                    r2.0.parse::<i32>().unwrap()..r2.1.parse::<i32>().unwrap() + 1,
                )),
            )
        })
        .map(|(p1, p2)| (p1.intersection(&p2).count() > 0))
        .filter(|n| *n)
        .count();

    println!("Second task: {:?}", subsets);
}

fn main() {
    task1();
    task2()
}
