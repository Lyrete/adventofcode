use priority_queue::DoublePriorityQueue;
use std::collections::HashMap;
use std::time::Instant;

const CHARS: &str = "abcdefghijklmnopqrstuvwxyz";

fn adjacents(grid: &Vec<Vec<char>>, point: (usize, usize)) -> Vec<(usize, usize)> {
    let mut edges: Vec<(usize, usize)> = Vec::new();

    if point.0 > 0 {
        edges.push((point.0 - 1, point.1));
    }

    if point.0 < grid[0].len() - 1 {
        edges.push((point.0 + 1, point.1));
    }

    if point.1 > 0 {
        edges.push((point.0, point.1 - 1));
    }

    if point.1 < grid.len() - 1 {
        edges.push((point.0, point.1 + 1));
    }

    edges
}

struct TaskResult {
    task1: u16,
    task2: u16,
}

fn find_paths(grid: &Vec<Vec<char>>, start: (usize, usize), end: (usize, usize)) -> TaskResult {
    let mut paths: HashMap<(usize, usize), usize> = HashMap::new();

    let mut result = TaskResult {
        task1: 0,
        task2: 1000,
    };
    let mut pq = DoublePriorityQueue::new();

    paths.insert(end, 0);

    pq.push(end, 0);

    while !(pq.is_empty()) {
        if result.task1 > 0 && result.task2 > 0 {
            return result;
        }

        let (current, priority) = pq.pop_min().unwrap();

        if current == start {
            result.task1 = priority;
        }

        if grid[current.1][current.0] == 'a' && result.task2 > priority {
            result.task2 = priority;
        }

        for neighbor in adjacents(grid, current).iter() {
            let slope: isize = CHARS.find(grid[neighbor.1][neighbor.0]).unwrap() as isize
                - CHARS.find(grid[current.1][current.0]).unwrap() as isize;

            if slope < -1 {
                continue;
            }

            let length = paths.get(&current).unwrap();

            if !paths.contains_key(neighbor) || length - 1 < *paths.get(&neighbor).unwrap() {
                let path = *paths.get(&current).unwrap();
                paths.insert(*neighbor, path + 1);
                pq.push_decrease(*neighbor, (path + 1) as u16);
            }
        }
    }

    print_visited(paths.keys().map(|x| *x).collect(), &grid);
    result
}

fn main() {
    let now = Instant::now();
    let lines: Vec<Vec<char>> = include_str!("./data12.txt")
        .lines()
        .map(|x| x.chars().collect())
        .collect();

    let mut grid: Vec<Vec<char>> = Vec::new();
    let mut start = (100000, 100000);
    let mut end = (0, 0);

    for (idx, line) in lines.iter().enumerate() {
        if line.contains(&'S') {
            start = (line.iter().position(|&c| c == 'S').unwrap(), idx);
        }
        if line.contains(&'E') {
            end = (line.iter().position(|&c| c == 'E').unwrap(), idx);
        }
        grid.push(
            line.clone()
                .iter()
                .map(|&x| {
                    if x == 'S' {
                        'a'
                    } else if x == 'E' {
                        'z'
                    } else {
                        x
                    }
                })
                .collect(),
        );
    }

    let result = find_paths(&grid, start, end);
    println!(
        "Task 1 path: {:?}, Task 2 path: {:?}",
        result.task1, result.task2
    );
    println!("Elapsed: {:?}", now.elapsed());

    let char_map: Vec<_> = include_str!("./data12.txt")
        .lines()
        .enumerate()
        .map(|(y, line)| (line.chars().collect::<Vec<char>>(), y))
        .collect();

    println!("{:?}", char_map);
}

fn print_grid(grid: &Vec<Vec<char>>) {
    for line in grid.iter() {
        println!(
            "{}",
            line.iter().fold(String::new(), |a, b| a + &b.to_string())
        );
    }
}

fn print_visited(visited: Vec<(usize, usize)>, values: &Vec<Vec<char>>) {
    let mut grid = vec![vec![' '; values[0].len()]; values.len()];

    for (x, y) in visited.iter() {
        grid[*y][*x] = values[*y][*x];
    }

    print_grid(&grid);
}
