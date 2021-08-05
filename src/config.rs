
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

pub fn set_up_conf(path: &str) {

    let p = paths::exp_const_path(path);
    let res = fs::create_dir(&p); 

    if res.is_ok() {

        res.unwrap();
        println!("{} created dir {}", col::bx_print(), &p.to_string());

        for f in FILES.iter() {
            let mut fp = p.to_owned();
            fp.push_str(f);
            let fp = Path::new(&fp);
            let w = fs::write(fp, "");
            
            if w.is_ok() {
                println!("{} created file {}", col::bx_print(), fp.to_str().unwrap());

            } else {
                errs::sys_err(w.err().unwrap())
            }
        }

    } else {
        errs::sys_err(res.err().unwrap());
    }
}