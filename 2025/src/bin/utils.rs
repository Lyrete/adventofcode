pub fn get_straight_line_dist(a: (isize, isize, isize), b: (isize, isize, isize)) -> f64 {
    let dist = (((b.0 - a.0).pow(2) + (b.1 - a.1).pow(2) + (b.2 - a.2).pow(2)) as f64).sqrt();

    dist
}

fn main() {}
