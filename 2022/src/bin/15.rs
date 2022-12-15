fn to_tuple_pair(input: &str) -> ((isize, isize), (isize, isize)) {
    let pairs = input.split_once(":").unwrap();

    let start = pairs
        .0
        .split_once(", ")
        .map(|(x, y)| (x.parse::<isize>().unwrap(), y.parse::<isize>().unwrap()))
        .unwrap();

    let end = pairs
        .1
        .split_once(", ")
        .map(|(x, y)| (x.parse::<isize>().unwrap(), y.parse::<isize>().unwrap()))
        .unwrap();

    (start, end)
}

fn dist_between(start: (isize, isize), end: (isize, isize)) -> f64 {
    return f64::sqrt((start.0 - end.0));
}

fn main() {
    let input: Vec<_> = include_str!("../../inputs/ex15.txt")
        .lines()
        .map(|x| {
            x.replace("Sensor at x=", "")
                .replace(" closest beacon is at x=", "")
                .replace("y=", "")
        })
        .map(|x| to_tuple_pair(&x))
        .collect();

    println!("{:?}", input);
}
