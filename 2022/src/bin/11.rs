use std::collections::VecDeque;
use std::str::FromStr;
use std::time::Instant;

#[derive(Debug, Clone)]
struct Monkey {
    items: VecDeque<usize>,
    testnumber: usize,
    operation: String,
    targets: (u32, u32),
    inspected: usize,
}

impl Monkey {
    fn operate(&self, old: usize) -> usize {
        let op_str = self.operation.replace("old", &old.to_string());

        let op: Vec<&str> = op_str.split(" ").collect();

        if op[1] == "*" {
            op[0].parse::<usize>().unwrap() * op[2].parse::<usize>().unwrap()
        } else {
            op[0].parse::<usize>().unwrap() + op[2].parse::<usize>().unwrap()
        }
    }

    fn throw(&self, item: usize) -> usize {
        if item % self.testnumber == 0 {
            self.targets.0 as usize
        } else {
            self.targets.1 as usize
        }
    }
}

impl FromStr for Monkey {
    type Err = ();

    fn from_str(input: &str) -> Result<Monkey, Self::Err> {
        let lines: Vec<&str> = input.lines().collect();
        let items: VecDeque<usize> = lines[1]
            .split_once(": ")
            .unwrap()
            .1
            .split(", ")
            .flat_map(|x| x.parse::<usize>())
            .collect();

        let testnumber = lines[3]
            .split_once(" divisible by ")
            .unwrap()
            .1
            .parse::<usize>()
            .unwrap();

        let targets: (u32, u32) = (
            lines[4].chars().nth(29).unwrap().to_digit(10).unwrap(),
            lines[5].chars().nth(30).unwrap().to_digit(10).unwrap(),
        );

        let operation = lines[2].split_once("= ").unwrap().1.to_string();

        Ok(Monkey {
            items,
            testnumber,
            operation,
            targets,
            inspected: 0,
        })
    }
}

fn main() {
    let now = Instant::now();
    let monkey_instructions: Vec<_> = include_str!("./data11.txt").split("\n\n").collect();

    let mut monkeys: Vec<Monkey> = Vec::new();

    for ins in monkey_instructions.iter() {
        monkeys.push(Monkey::from_str(ins).unwrap());
    }

    let mut monkeys2 = monkeys.clone();

    for _ in 0..20 {
        for idx in 0..monkeys.len() {
            let current = monkeys[idx].clone();
            current.items.iter().for_each(|x| {
                monkeys[idx].inspected += 1;
                let worry = current.operate(*x) / 3;
                let target_idx = current.throw(worry);
                monkeys[target_idx].items.push_back(worry);
                monkeys[idx].items.pop_front();
            });
        }
    }

    let mut inspections: Vec<usize> = monkeys.iter().map(|x| x.inspected).collect();
    inspections.sort_by(|a, b| b.cmp(a));

    println!("Task 1: {:?}", inspections[0] * inspections[1]);

    let lcm: usize = monkeys.iter().map(|x| x.testnumber).product();

    for _ in 0..10000 {
        for idx in 0..monkeys2.len() {
            let current = monkeys2[idx].clone();
            current.items.iter().for_each(|x| {
                monkeys2[idx].inspected += 1;
                let worry = current.operate(*x) % lcm;
                let target_idx = current.throw(worry);
                monkeys2[target_idx].items.push_back(worry);
                monkeys2[idx].items.pop_front();
            });
        }
    }

    inspections = monkeys2.iter().map(|x| x.inspected).collect();
    inspections.sort_by(|a, b| b.cmp(a));

    println!("Task 2: {:?}", inspections[0] * inspections[1]);
    println!("Elapsed: {:?}", now.elapsed());
}
