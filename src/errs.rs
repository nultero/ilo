use std::io::Error;
use std::process::exit;


pub fn basic_err(s: &str) {
    print!("\x1b[31;1;4m{}\x1b[0m {}\n", "error:", s);
}

pub fn crit_err(s: &str) {
    print!("\x1b[31;1;4m{}\x1b[0m {}\n", "error:", s);
    println!("exiting ...");
    exit(1);
}

pub fn sys_err(e: Error) {
    println!("\x1b[31;1;4m{}\x1b[0m {}", "error:", e);
    exit(1);
}