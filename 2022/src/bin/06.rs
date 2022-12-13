use std::collections::HashSet;
use std::collections::VecDeque;

fn main() {
    let data: Vec<char> = include_str!("./data.txt").trim().chars().collect();

    let mut first_pos = 0;
    let mut q2: VecDeque<char> = VecDeque::new();
    let mut snd_pos = 0;

    for (i, c) in data.iter().enumerate() {
        q2.push_back(*c); //Add to queue
        if q2.len() >= 4 {
            let uniqs: HashSet<&char> = q2.range(q2.len() - 4..).collect(); //Grab last 4 of queue
            if uniqs.len() == 4 && first_pos == 0 {
                first_pos = i + 1;
            }
        }

        if q2.len() == 14 {
            let uniqs: HashSet<&char> = q2.iter().collect();
            if uniqs.len() == 14 && snd_pos == 0 {
                snd_pos = i + 1;
            }
            q2.pop_front(); //Remove from queue once check complete
        }
    }

    println!("First task completed at pos: {}", first_pos);
    println!("Second task completed at pos: {}", snd_pos);
}
