
use std::fs;
use std::path::Path;
use crate::colors::col;
use crate::paths;
use crate::errs;

const FILES: [&str; 7] = [
    "cache.txt",
    "config.txt",
    "events.txt",
    "ideas.txt",
    "reminders.txt",
    "todos.txt",
    "wishlist.txt",
];

const DEFAULT_CONF_LINES: [&'static str; 4]  = [
    "todoSymbol = ○",
    "prompt_icon = ❯➤",
    "# days ahead to check for",
    "days = 3",
];

pub fn set_up_conf(path: &str) {

    let p = paths::expand_const_path(path);
    let res = fs::create_dir(&p); 

    if res.is_ok() {

        res.unwrap();
        println!("{} created dir {}", col::bx_print(), &p.to_string());

        for f in FILES.iter() {

            let mut fp = p.to_owned();
            fp.push_str(f);
            let fp = Path::new(&fp);

            let mut output = String::from("");

            if f == &"config.txt" {
                for line in &DEFAULT_CONF_LINES {
                    output.push_str(&line);
                    output.push_str(&"\n");
                }
                // slice last newline off
                let tmp = &output[..output.len() - 1];
                output = tmp.to_owned();
            }
            
            let w = fs::write(fp, output);
            
            if w.is_ok() {
                println!("{} created file {}", col::bx_print(), fp.to_str().unwrap());

            } else {
                errs::api::sys_err(w.err().unwrap())
            }
        }

    } else {
        errs::api::sys_err(res.err().unwrap());
    }
}