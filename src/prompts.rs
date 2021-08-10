use std::io::Write;
use std::io::{stdin, stdout};

use crate::colors::col;

pub fn confirmed(answer: &str) -> bool {
    let a = &answer.to_lowercase();
    let a = a.as_str();
    return !a.contains("n");
}

pub fn get_answer(prompt: &str) -> String {
    let mut out = stdout();
    out.write(&prompt.as_bytes()).unwrap();
    out.flush().unwrap();

    let mut s: String = "".to_owned();
    stdin().read_line(&mut s).unwrap();
    return s;
}

pub fn red_setup_prompt() -> String {
    let mut slash = col::red_no_underline(&"   \\");
    slash.push_str(&" Set up bx folder and config? [ Y / n ] ");
    return slash;
}
