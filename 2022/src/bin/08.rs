use std::collections::HashSet;
use std::time::Instant;

fn main() {
    let time = Instant::now();

    let grid: Vec<_> = include_str!("./data8.txt")
        .lines()
        .map(str::to_string)
        .map(|c| {
            c.chars()
                .map(|x| x.to_owned().to_digit(10).unwrap())
                .collect::<Vec<u32>>()
        })
        .collect();

    let mut visible: HashSet<(usize, usize)> = HashSet::new();

    for (line_idx, line) in grid.iter().enumerate() {
        let mut row_largest = 0;

        for (x, tree) in line.iter().enumerate() {
            if x == 0 {
                row_largest = *tree;
                visible.insert((x, line_idx));
                continue;
            }

            if tree <= &row_largest {
                continue;
            }

            row_largest = *tree;
            visible.insert((x, line_idx));
        }

        row_largest = 0;

        for (x, tree) in line.iter().enumerate().rev() {
            if x == line.len() - 1 {
                row_largest = *tree;
                visible.insert((x, line_idx));
                continue;
            }

            if tree <= &row_largest {
                continue;
            }

            row_largest = *tree;
            visible.insert((x, line_idx));
        }
    }

    let size = grid.first().unwrap().len();
    for x in 0..size {
        let mut row_largest = 0;

        for y in 0..size {
            if y == 0 {
                row_largest = grid[y][x];
                visible.insert((x, y));
                continue;
            }

            if grid[y][x] <= row_largest {
                continue;
            }

            row_largest = grid[y][x];
            visible.insert((x, y));
        }

        for y in (0..size).rev() {
            if y == size - 1 {
                row_largest = grid[y][x];
                visible.insert((x, y));
                continue;
            }

            if grid[y][x] <= row_largest {
                continue;
            }

            row_largest = grid[y][x];
            visible.insert((x, y));
        }
    }

    println!("Task 1: {:?}", visible.len());

    let mut highest = 0;
    for (x, y) in visible.iter() {
        if x == &0 || y == &0 || *x == size - 1 || *y == size - 1 {
            continue;
        }
        let score = check_score(*x, *y, grid.to_vec());
        if score > highest {
            highest = score;
        }
    }
    println!("Task 2: {}", highest);
    println!("Took: {:?}", time.elapsed());
}

fn check_score(x: usize, y: usize, grid: Vec<Vec<u32>>) -> usize {
    let tree = grid[y][x];
    let GRID_SIZE = grid.first().unwrap().len();

    let mut score = 0;

    let line = &grid[y];

    let mut left_trees = 0;
    for lx in (0..x).rev() {
        left_trees += 1;
        if line[lx] >= tree {
            break;
        }
    }

    let mut right_trees = 0;
    for rx in x + 1..GRID_SIZE {
        right_trees += 1;
        if line[rx] >= tree {
            break;
        }
    }

    //top
    let mut top_trees = 0;
    for ty in (0..y).rev() {
        top_trees += 1;
        if grid[ty][x] >= tree {
            break;
        }
    }

    let mut bot_trees = 0;
    for dy in y + 1..GRID_SIZE {
        bot_trees += 1;
        if grid[dy][x] >= tree {
            break;
        }
    }

    score += left_trees * right_trees * top_trees * bot_trees;

    return score;
}

//Helper to use for debug
fn printer(trees: HashSet<(usize, usize)>, size: usize) {
    let size = size;

    let mut grid: Vec<Vec<_>> = Vec::with_capacity(size);

    for _ in 0..size {
        grid.push(vec![0; size])
    }

    for (x, y) in trees.iter() {
        grid[*y][*x] = 1;
    }

    for line in grid.iter() {
        println!("{:?}", line);
    }
}
