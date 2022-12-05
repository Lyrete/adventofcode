use std::collections::BTreeMap;

fn main() {
    let input = include_str!("./data.txt").split_once("\n\n").map(|(p1, p2)| (p1.lines().map(String::from).collect::<Vec<String>>(), p2.lines().map(String::from).collect::<Vec<_>>())).unwrap();
    
    let moves: Vec<_> = input.1.iter().map(|n| n.replace("move ", "")).map(|n| n.replace("from", "to")).map(|n| n.split(" to ").map(|n| n.parse().unwrap()).collect::<Vec<i8>>()).collect();
   
    let crate_input: Vec<String> = input.0.iter().map(String::from).collect::<Vec<String>>();
    let mut crate_stacks: BTreeMap<i8, Vec<_>> = BTreeMap::new();

    for crate_line in crate_input.iter() {
        let length = crate_line.len();
        
        let mut i = 0;
        while i < length {
            let slice = crate_line[i..i+3].replace(" ", "").replace("[", "").replace("]", "");
            let idx: i8 = (i / 4 + 1).try_into().unwrap(); //Get idx (aka stack of the crate)
            if slice.len() > 0 && !slice.parse::<i32>().is_ok() {
                if !crate_stacks.contains_key(&idx){
                    crate_stacks.insert(idx, Vec::new());
                }
                let mut stack: Vec<String> = crate_stacks.get(&idx).expect("INPUT BROKEN").to_vec();
                stack.push(slice);
                crate_stacks.insert(idx, stack);
            }
            i += 4;
        }        
    }
    
    for (_, val) in crate_stacks.iter_mut(){
        val.reverse();
    }

    let mut crate_stacks2 = crate_stacks.clone();

    for move_input in moves.iter(){
        let amount: i8 = move_input[0];
        let out_stack = &mut crate_stacks.get_mut(&move_input[1]).unwrap().to_owned();
        let in_stack = &mut crate_stacks.get_mut(&move_input[2]).unwrap().to_owned();

        for _i in 0..amount {
            let move_crate = out_stack.pop().unwrap();
            crate_stacks.insert(move_input[1], out_stack.to_vec());
            in_stack.push(move_crate);
            crate_stacks.insert(move_input[2], in_stack.to_vec());
        }
    }


    

    println!("{:?}", crate_stacks2); 
    for move_input in moves.iter(){
        let amount: i8 = move_input[0];
        let out_stack = &mut crate_stacks2.get_mut(&move_input[1]).unwrap().to_owned();
        let in_stack = &mut crate_stacks2.get_mut(&move_input[2]).unwrap().to_owned();

        let mut moved_crates = Vec::new();
        for _i in 0..amount {
            moved_crates.push(out_stack.pop().unwrap());
        }

        moved_crates.reverse();
        crate_stacks2.insert(move_input[1], out_stack.to_vec());
        in_stack.append(&mut moved_crates);
        crate_stacks2.insert(move_input[2], in_stack.to_vec());

        println!("{:?}", crate_stacks2);
    }

    println!("{:?}", crate_stacks2);

    println!("Task 1 result: {:?}", crate_stacks.iter().map(|(_, v)| v.last()).fold("".to_string(), |acc, value| acc + value.unwrap()));
    println!("Task 2 result: {:?}", crate_stacks2.iter().map(|(_, v)| v.last()).fold("".to_string(), |acc, value| acc + value.unwrap())); 
}
