// use std::result;
// use std::fs;
use dirs::home_dir;
use std::io;
use std::path;

pub fn expand_const_path(p: &str) -> String {
    let mut p = p.to_owned();

    if &p.chars().nth(0).unwrap() == &'~' {
        let home = home_dir().unwrap();
        let hstr = home.to_str().unwrap();
        let mut home = hstr.to_owned();

        let tildeslice = &p[1..];
        home.push_str(tildeslice);

        p = home;
    }

    return p;
}

pub fn bx_path(p: &str) -> Result<String, io::Error> {
    let p = expand_const_path(p);
    let p = path::Path::new(&p).canonicalize()?;
    let p = p.to_str().unwrap();

    return Ok(p.to_owned());
}

fn get_file(filetype: &str) -> &str {
    match filetype {
        "conf" => "/config.txt",
        "cache" => "/cache.txt",
        "events" => "/events.txt",
        "ideas" => "/ideas.txt",
        "recurrents" => "/reminders.txt",
        "todos" => "/todos.txt",
        "wishlist" => "/wishlist.txt",
        _ => "",
    }
}

pub fn get_file_path(path: &str, filetype: &str) -> String {
    let f = get_file(filetype);
    let p = format!("{}{}", path, f);
    return p.to_owned();
}
