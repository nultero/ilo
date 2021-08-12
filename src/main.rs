extern crate chrono;
extern crate crossterm;
extern crate dirs;

mod bx {
    pub mod reminders;
    pub mod settingparser;
    pub mod timeline;
}
mod colors {
    pub mod col;
}
mod errs {
    pub mod api;
}
mod pathutils {
    pub mod paths;
}
mod argparser;
mod config;
mod funcs;
mod prompts;

use bx::reminders;
use colors::col;
use std::env;

const PATH: &'static str = "~/.tailbox/";

fn main() {
    let args: Vec<String> = env::args().skip(1).collect();

    let path = pathutils::paths::bx_path(PATH);

    if path.is_ok() {
        let path = path.unwrap();

        if args.len() == 0 {
            reminders::run_reminders(&path);
        } else {
            argparser::parse_args(&args, &path);
        }
    } else {
        let s = format!(
            "path '{}' was not able to be expanded or does not exist",
            col::blue(&PATH)
        );
        errs::api::basic_err(&s);

        let prmpt = prompts::red_setup_prompt();
        let answer = prompts::get_answer(&prmpt);
        if prompts::confirmed(&answer) {
            config::set_up_conf(PATH);
        }
    }
}
