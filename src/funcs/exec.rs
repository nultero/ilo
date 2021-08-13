use crate::argparser::Bus;
use crate::bx::settingparser;
use crate::colors::col;
use crate::errs;
use crate::pathutils::paths;

use std::fs;
use crate::funcs::inputs;
use std::io::*;

const EVENT_STR_CHUNKS: [&str; 3] = ["event", "day", "month"];

pub fn exec_func(mut b: Bus) {
    let conf_path = paths::get_file_path(&b.path, &"conf");
    let conf = fs::read_to_string(&conf_path);

    if conf.is_ok() {
        let conf = conf.unwrap();
        let icon = settingparser::get_icon(conf);
        b.set_icon(&icon);

        let p = paths::get_file_path(&b.path, &b.file_type);
        b.set_path(&p);

        // this WILL have been validated previously but whatever
        let oops = format!("did not match a function to '{}'", col::blue(&b.func));

        match b.func.as_str() {
            "add" => add(&b),
            "list" => ls(&b),
            _ => errs::api::crit_err(&oops),
        }
    } else {
        let fmt_path = col::blue(&format!("{}", &b.path));
        let s = format!("problem accessing bx's config file in '{}'", &fmt_path);
        errs::api::crit_err(&s);
    }
}



fn add(b: &Bus) {
    println!("{} what to add to {}?", &b.prompt_icon, &b.file_type);
    let mut input = "".to_owned();
    let stdin = stdin();

    if &b.file_type == &"events" {
        let prompt = "format: event  @  day  month".to_owned();
        let mut tmp: Vec<String> = vec![];
        for chunk in &EVENT_STR_CHUNKS {
            let mut inp = "".to_owned();
            let s = &prompt.replace(chunk, &col::blue(&chunk));
            println!("{}", s);
            &stdin.read_line(&mut inp).unwrap();
            inp = inp.replace("\n", "");
            tmp.push(inp.to_owned());
        }

        let s = format!("{}  @  {} {}", tmp[0], tmp[1], tmp[2]);
        input.push_str(&s);
    } else {
        // stdin.read_line(&mut input).unwrap();
        input = inputs::read_normal_string().unwrap();
    }

    // need key capture / possibly lists w/ vim bindings and some sugar over the line writing

    if input == "<&&&>" {
        println!("input interrupted");

    } else if input.len() != 0 {

        let mut contents = fs::read_to_string(&b.path).unwrap();
        contents.push_str("\n");
        input = input.replace("\n", ""); // strip newline for printing / not writing last newline to file
        contents.push_str(&input);

        fs::write(&b.path, &contents).unwrap();

        let out = col::blue(&b.file_type);
        println!("wrote out `{}` to {}", input, out);

    } else {
        println!("input empty");
    }
}

fn ls(b: &Bus) {
    let contents = fs::read_to_string(&b.path).unwrap();
    println!("{}", &contents);
}
