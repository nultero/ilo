extern crate chrono;
extern crate dirs;

mod config;
mod errs;
mod paths;
mod prompts;
mod colors {
    pub mod col;
}

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
            println!("junk -> {}", path);
            for r in args {
                println!("{}", r);
            }
        }

    } else {

        errs::basic_err();
        print!("path '{}' was not able to be expanded or does not exist \n", col::blue(&PATH));

        let prmpt = prompts::red_setup_prompt();
        let answer = prompts::get_answer(&prmpt);
        if prompts::confirmed(&answer) {
            config::set_up_conf(PATH);
        }
        
        // println!("glop \n {:?}", path.err());
    }
}
