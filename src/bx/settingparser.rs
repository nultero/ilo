use crate::colors::col;
use crate::errs::api;

fn parse(file_contents: &str, search: &str) -> String {
    let lines = file_contents.split("\n");

    for line in lines {
        let line = line.trim();
        if line.chars().nth(0).unwrap() != '#' {
            let spl: Vec<&str> = line.split("=").collect();

            let setting = &spl[0].trim();
            if &search == setting {
                return spl[1].to_owned();
            }
        }
    }

    return "".to_owned();
}

pub fn get_icon(conf: String) -> String {
    let res = parse(&conf, &"prompt_icon");
    if res.len() == 0 {
        let set = col::blue("prompt_icon");
        let s = format!("could not find '{}' setting in bx's config", set);
        api::basic_err(&s);
    }

    return format!("{} ", res);
}
