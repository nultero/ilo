use crate::argparser::Bus;
use crate::colors::col;
use crate::pathutils::paths;

use std::fs;
use std::io::*;

pub fn exec_func(mut b: Bus) {
    let p = paths::get_file_path(&b.path, &b.file_type);
    b.set_path(&p);
    match b.func.as_str() {
        "add" => add(&b),
        "list" => ls(&b),
        _ => println!("blob"),
    }
}

fn add(b: &Bus) {
    println!(">> what to add to {}?", &b.file_type);

    let stdin = stdin();
    let mut inp = "".to_owned();
    stdin.read_line(&mut inp).unwrap();

    // need key capture / possibly lists w/ vim bindings and some sugar over the line writing

    let mut contents = fs::read_to_string(&b.path).unwrap();
    contents.push_str("\n");
    inp = inp.replace("\n", ""); // strip newline for printing / not writing last newline to file
    contents.push_str(&inp);

    fs::write(&b.path, &contents).unwrap();

    let out = col::blue(&b.file_type);
    println!("wrote out `{}` to {}", inp, out);
}

fn ls(b: &Bus) {
    let contents = fs::read_to_string(&b.path).unwrap();
    println!("{}", &contents);
}
