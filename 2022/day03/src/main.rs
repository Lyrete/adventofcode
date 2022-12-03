use std::collections::HashSet;

const CHARS: &str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ";

fn task1() {
    let lines: Vec<_> = include_str!("./data.txt")
        .lines()
        .map(|n| n)
        .map(|n| n.split_at(n.len() / 2))
        .map(|(c1, c2)| {
            (
                HashSet::from_iter(c1.chars()),
                HashSet::from_iter(c2.chars()),
            )
        })
        .collect::<Vec<(HashSet<char>, HashSet<char>)>>();

    //println!("{:?}", lines);
    let mut found: Vec<&char> = Vec::new();

    for line in lines.iter() {
        let mut uniq: Vec<&char> = line.0.iter().filter(|c| line.1.contains(&c)).collect();

        found.append(&mut uniq)
    }

    let mut sum = 0;
    for &c in found.iter() {
        //println!("{:?} - {:?}", c, *c as u32);
        //println!("{:?} - {:?}", c, letters.find(*c).unwrap() + 1);
        sum += CHARS.find(*c).unwrap() + 1;
    }

    println!("Task 1 result: {:?}", sum);
}

fn task2() {
    let lines: Vec<_> = include_str!("./data.txt")
        .lines()
        .map(|n| HashSet::from_iter(n.chars()))
        .collect::<Vec<HashSet<char>>>();

    let mut i = 0;
    let mut score = 0;
    while i < lines.len() {
        let shared: Vec<&char> = lines[i]
            .intersection(&lines[i + 1])
            .filter(|n| lines[i + 2].contains(n))
            .collect();

        if shared.len() > 0 {
            score += CHARS.find(*shared[0]).unwrap() + 1
        }

        i += 3;
    }

    println!("Task 2 result: {:?}", score)
}

fn main() {
    task1();
    task2()
}
