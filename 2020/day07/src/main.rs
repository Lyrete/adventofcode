use std::collections::HashMap;

type Bags<'a> = HashMap<String, [(u32, &'a str)]>;

fn main() {
    let lines: Vec<&str> = include_str!("./ex.txt").lines().map(|n| n).collect();

    let mut bags = Bags::new();

    for line in lines.iter() {
        let split: Vec<&str> = line.split(" contain ").collect();
        if split[1] == "no other bags." {
            //Can skip this line as no bags can be inside the colour
            continue;
        }
        let contains: Vec<(u32, &str)> = split[1]
            .split(", ")
            .map(|n| n.replace(" bags", ""))
            .map(|n| n.replace(" bag", ""))
            .map(|n| n.replace(".", ""))
            .map(|n| {
                (
                    n.chars().nth(0).unwrap().to_digit(10).unwrap(),
                    n[2..n.len()],
                )
            })
            .collect();
        println!("{:?}", contains);
        bags.insert(split[0].replace(" bags", ""), contains);
    }

    println!("{:?}", bags);
}
