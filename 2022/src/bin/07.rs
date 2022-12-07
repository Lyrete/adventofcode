use std::collections::BTreeMap;

fn main() {
    let mut input = include_str!("./data7.txt").lines();

    let mut dirstack: Vec<String> = Vec::new();

    let mut fs_sizes: BTreeMap<String, i32> = BTreeMap::new();

    loop {
        let split_line: Vec<String> = input
            .next()
            .unwrap_or("")
            .split(" ")
            .map(str::to_string)
            .collect();

        if split_line.len() <= 1 {
            //Exit out on last line, crude but works.
            break;
        }

        if split_line[0] != "$" {
            if split_line[0] == "dir" {
                continue;
            }

            for i in (1..dirstack.len() + 1).rev() {
                let c_path = &dirstack[0..i]
                    .iter()
                    .map(|s| s.to_string())
                    .reduce(|a, b| a + "/" + &b)
                    .unwrap()
                    .replace("//", "/");

                fs_sizes
                    .entry(c_path.to_string())
                    .and_modify(|curr| *curr += split_line[0].parse::<i32>().unwrap())
                    .or_insert(split_line[0].parse().unwrap());
            }
        } else {
            let cmd = &split_line[1];
            if cmd == "ls" {
                continue;
            }

            let dest = &split_line[2];
            if dest == ".." {
                dirstack.pop();
            } else {
                dirstack.push(dest.to_string());
            }
        }
    }

    let over100k_sum: i32 = fs_sizes
        .iter()
        .map(|(_, v)| v)
        .filter(|v| *v <= &100000)
        .sum::<i32>();

    let unused_space = 70000000 - fs_sizes.get("/").unwrap();

    let smallest_dir: i32 = *fs_sizes
        .iter()
        .map(|(_, v)| v)
        .map(|v| (v, unused_space - 30000000 + v))
        .filter(|(_, c)| c > &0)
        .map(|(v, _)| v)
        .min()
        .unwrap();

    println!("Task 1: {:?}", over100k_sum);
    println!("Task 2: {:?}", smallest_dir);
}
