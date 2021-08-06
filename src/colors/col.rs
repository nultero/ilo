
pub fn bx_print() -> String {
    let s = format!("\x1b[32;1m{}\x1b[0m", " |>  bx: ");
    return s;
}

pub fn blue(s: &str) -> String {
    let s = format!("\x1b[32;1;4m{}\x1b[0m", s);
    return s;
}

pub fn red_no_underline(s: &str) -> String {
    let s = format!("\x1b[31;1m{}\x1b[0m", s);
    return s;
}

// pub fn red_underline(s: &str) -> String {
//     let s = format!("\x1b[31;1;4m{}\x1b[0m", s);
//     return s;
// }