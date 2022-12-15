use std::cmp;
use std::collections::HashSet;

#[derive(Debug, Hash, Eq, PartialEq, Clone, Copy)]
struct Point {
    x: usize,
    y: usize,
}

impl From<(usize, usize)> for Point {
    fn from((x, y): (usize, usize)) -> Point {
        return Point { x, y };
    }
}

impl From<Point> for (usize, usize) {
    fn from(point: Point) -> (usize, usize) {
        return (point.x, point.y);
    }
}

#[derive(Debug)]
struct Sand {
    location: Point,
}

impl Sand {
    fn fall(&mut self, blocked: &HashSet<Point>) {
        let (x, y) = (self.location.x, self.location.y);
        if !blocked.contains(&Point { x, y: y + 1 }) {
            self.location.y += 1;
            return;
        }

        if !blocked.contains(&Point { x: x - 1, y: y + 1 }) {
            self.location.y += 1;
            self.location.x -= 1;
            return;
        }

        if !blocked.contains(&Point { x: x + 1, y: y + 1 }) {
            self.location.y += 1;
            self.location.x += 1;
            return;
        }
    }
}

const FALLPOINT: Point = Point { x: 500, y: 0 };

fn main() {
    let mut walls: HashSet<Point> = HashSet::new();

    let wallpoints: Vec<Vec<(usize, usize)>> = include_str!("../../inputs/data14.txt")
        .lines()
        .map(|x| {
            x.split(" -> ")
                .map(|x| x.split_once(",").unwrap())
                .map(|(x, y)| (x.parse::<usize>().unwrap(), y.parse::<usize>().unwrap()))
                .collect()
        })
        .collect();

    let mut lowest_wall = 0;

    for wall in wallpoints {
        for i in 1..wall.len() {
            let start = Point::from(wall[i - 1]);
            let end = Point::from(wall[i]);

            let max_y = cmp::max(start.y, end.y);
            if lowest_wall < max_y {
                lowest_wall = max_y;
            }

            // Vertical wall
            if start.x == end.x {
                let range = if start.y < end.y {
                    (start.y, end.y)
                } else {
                    (end.y, start.y)
                };

                for y in range.0..range.1 + 1 {
                    walls.insert(Point::from((start.x, y)));
                }
            }

            // Horizontal wall
            if start.y == end.y {
                let range = if start.x < end.x {
                    (start.x, end.x)
                } else {
                    (end.x, start.x)
                };
                for x in range.0..range.1 + 1 {
                    walls.insert(Point::from((x, start.y)));
                }
            }
        }
    }

    let mut sand_stopped = 0;
    let mut current = Sand {
        location: FALLPOINT,
    };

    loop {
        current.location = FALLPOINT;

        while current.location.y < lowest_wall {
            let previous = current.location.clone();
            current.fall(&walls);
            //println!("{:?}", current);

            //didn't fall anymore
            if previous == current.location {
                sand_stopped += 1;
                walls.insert(current.location);
                break;
            }
        }

        if current.location.y >= lowest_wall {
            break;
        }
    }

    let mut sand2 = 0;

    loop {
        current.location = FALLPOINT;

        loop {
            let previous = current.location.clone();
            current.fall(&walls);

            //didn't fall anymore
            if previous == current.location || current.location.y == lowest_wall + 1 {
                sand2 += 1;
                walls.insert(current.location);
                break;
            }
        }

        if current.location.y == 0 {
            break;
        }
    }

    println!("Sand stopped: {}", sand_stopped);
    println!("Sand 2 stopped: {}", sand2 + sand_stopped);
}
