use std::io::Error;
use std::process::exit;


pub fn basic_err() {
    print!("\x1b[31;1;4m{}\x1b[0m ", "error:");
}

pub fn sys_err(e: Error) {
    println!("\x1b[31;1;4m{}\x1b[0m {}", "error:", e);
    exit(1);
}