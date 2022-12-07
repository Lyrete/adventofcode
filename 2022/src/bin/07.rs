use std::collections::BTreeMap;

const MAX_FS_SIZE: i32 = 70000000;
const TARGET_UNUSED_SIZE: i32 = 30000000;

fn main() {
    let input: Vec<String> = include_str!("./data7.txt")
        .lines()
        .map(str::to_string)
        .collect();

    let mut dirstack: Vec<String> = Vec::new();
    let mut fs_sizes: BTreeMap<String, i32> = BTreeMap::new();

    for line in input.iter() {
        let split_line: Vec<String> = line.split(" ").map(str::to_string).collect();

        //Not a command
        if split_line[0] != "$" {
            if split_line[0] == "dir" {
                //We don't care about directories (handled by path)
                continue;
            }

            //Traverse back the path in the stack and add the filesize all the way to root
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
            if "ls" == &split_line[1] {
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

    let unused_space = MAX_FS_SIZE - fs_sizes.get("/").unwrap();

    let smallest_dir: i32 = *fs_sizes
        .iter()
        //Map to tuple of (val, unused space if removed)
        .map(|(_, v)| (v, unused_space - TARGET_UNUSED_SIZE + v))
        .filter(|(_, c)| c > &0) //Any positive val works
        .map(|(v, _)| v)
        .min()
        .unwrap();

    println!("Task 1: {:?}", over100k_sum);
    println!("Task 2: {:?}", smallest_dir);
}
