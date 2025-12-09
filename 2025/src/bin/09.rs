use std::{fs::read_to_string, time::SystemTime};

const DAY: u8 = 9;
const EXAMPLE: &'static str = "7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3";

fn get_area_between_points(a: (isize, isize), b: (isize, isize)) -> usize {
    let h = if a.0 < b.0 { b.0 - a.0 } else { a.0 - b.0 } + 1;
    let w = if a.1 < b.1 { b.1 - a.1 } else { a.1 - b.1 } + 1;
    (h * w).unsigned_abs()
}

fn rect_contains_point(rect: ((isize, isize), (isize, isize)), point: (isize, isize)) -> bool {
    let (corner1, corner2) = rect;
    if ((corner1.0 < point.0 && point.0 < corner2.0)
        || (corner2.0 < point.0 && point.0 < corner1.0))
        && ((corner1.1 < point.1 && point.1 < corner2.1)
            || (corner2.1 < point.1 && point.1 < corner1.1))
    {
        return true;
    }

    false
}

fn polygon_intersects_rectangle(
    polygon: &Vec<&(isize, isize)>,
    rectangle: ((isize, isize), (isize, isize)),
) -> bool {
    let (a, b) = rectangle;
    let rectangle_min = (a.0.min(b.0), a.1.min(b.1));
    let rectangle_max = (a.0.max(b.0), a.1.max(b.1));

    for polygon_edge in polygon.windows(2) {
        let start = *polygon_edge[0];
        let end = *polygon_edge[1];

        let polyedge_min = (start.0.min(end.0), start.1.min(end.1));
        let polyedge_max = (start.0.max(end.0), start.1.max(end.1));

        // Vertical edges
        if start.0 == end.0 {
            if rectangle_min.0 < start.0 && start.0 < rectangle_max.0 {
                if rectangle_min.1.max(polyedge_min.1) < rectangle_max.1.min(polyedge_max.1) {
                    return true;
                }
            }
        }

        // Horizontal edges
        if start.1 == end.1 {
            if rectangle_min.1 < start.1 && start.1 < rectangle_max.1 {
                if rectangle_min.0.max(polyedge_min.0) < rectangle_max.0.min(polyedge_max.0) {
                    return true;
                }
            }
        }
    }

    false
}

fn solve(input: String) -> (usize, usize) {
    let points = input
        .trim()
        .split('\n')
        .map(|e| {
            let s = e.split_once(',').unwrap();
            (s.0.parse::<isize>().unwrap(), s.1.parse::<isize>().unwrap())
        })
        .collect::<Vec<(isize, isize)>>();

    let mut largest_area = 0;
    let mut largest_contained_area = 0;
    let polygon: Vec<&(isize, isize)> = points.iter().chain([&points[0]]).collect();

    for i in 0..points.len() {
        for j in i + 1..points.len() {
            let area = get_area_between_points(points[i], points[j]);
            if area > largest_area {
                //println!("{:?}, {:?} = {:?}", points[i], points[j], area);
                largest_area = area;
            }
            if area > largest_contained_area
                && !polygon_intersects_rectangle(&polygon, (points[i], points[j]))
            {
                largest_contained_area = get_area_between_points(points[i], points[j]);
            }
        }
    }

    (largest_area, largest_contained_area)
}

fn main() {
    let example_res = solve(EXAMPLE.to_string());
    println!("Example:");
    println!("{:?} {:?}", example_res.0, example_res.1);

    let start = SystemTime::now();
    let actual_res = solve(read_to_string(format!("./inputs/{:02}.txt", DAY)).unwrap());
    let end = SystemTime::now();

    println!("Actual:");
    println!("{:?} {:?}", actual_res.0, actual_res.1);

    println!();
    println!("Execution time: {:?}", end.duration_since(start).unwrap())
}
