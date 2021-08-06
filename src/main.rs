extern crate chrono;
extern crate dirs;

mod colors {
    pub mod col;
}
mod config;
mod errs;
mod funcs;
mod parsers;
mod paths;
mod prompts;

use std::env;
use colors::col;

#[allow(dead_code)]

const PATH: &str = "~/.tailbox/";

fn main() {

    let args: Vec<String> = env::args().skip(1).collect();

    let path = paths::bx_path(PATH);

    if path.is_ok() {

        let path = path.unwrap();

        if args.len() == 0 {
            
            println!("crunk -> {}", path);
            
            
        } else {
            parsers::parse_args(&args, &path);
        }

    } else {

        let s = format!("path '{}' was not able to be expanded or does not exist", col::blue(&PATH));
        errs::basic_err(&s);

        let prmpt = prompts::red_setup_prompt();
        let answer = prompts::get_answer(&prmpt);
        if prompts::confirmed(&answer) {
            config::set_up_conf(PATH);
        }
    }
}
