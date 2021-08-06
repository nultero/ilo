#[allow(dead_code)]

// use std::result;
// use std::fs;
use dirs::home_dir;
use std::path;
use std::io;

pub fn exp_const_path(p: &str) -> String {
    
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

    let p = exp_const_path(p);
    let p = path::Path::new(&p).canonicalize()?;
    let p = p.to_str().unwrap();

    return Ok(p.to_owned());
}
