use std::time::Instant;

//Helper to print a line in ASCII
fn print_line(line: &Vec<char>) {
    println!(
        "{}",
        line.iter()
            .fold("".to_string(), |cur, nxt| cur + &nxt.to_string())
    );
}

fn t1_better(instructions: &Vec<i16>) {
    let mut x: i16 = 1;
    let mut t1 = 0;

    for i in 0..6 {
        let chunk = if i == 0 {
            &instructions[0..19]
        } else {
            &instructions[40 * i - 20 - 1..20 + 40 * i - 1]
        };
        x += chunk.iter().sum::<i16>();
        t1 += (20 + 40 * i) as i16 * x;
    }

    println!("Task 1: {:?}", t1);
}

fn better() {
    let lines: Vec<_> = include_str!("./data10.txt").lines().collect();
    let mut instructions: Vec<i16> = Vec::new();
    let mut reg_pos: i16 = 1;
    for line in lines.iter() {
        instructions.push(0); //Add a no move instruction always
        let cmd: Vec<&str> = line.split(" ").collect();
        //Then if it was a command add an actual move instruction
        if cmd[0] == "addx" {
            instructions.push(cmd[1].parse::<i16>().unwrap())
        }
    }

    t1_better(&instructions);

    let mut grid: Vec<Vec<char>> = vec![vec![' '; 40]; 6];

    for (idx, rule) in instructions.iter().enumerate() {
        let dx: i16 = idx as i16 % 40;
        let dy: i16 = idx as i16 / 40;

        if dx.abs_diff(reg_pos) <= 1 {
            grid[dy as usize][dx as usize] = '█';
        }
        reg_pos += rule;
    }

    println!("Task 2 display:");
    for line in grid {
        print_line(&line);
    }
}

fn main() {
    let now = Instant::now();
    better();
    println!("Elapsed: {:?}", now.elapsed());
}
