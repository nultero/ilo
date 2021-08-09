
use std::io::Error;
use std::fs;
use chrono::{Datelike, Local};
use crate::bx::{timeline, settingparser};
use crate::colors::col;
use crate::errs::api;
use crate::pathutils::paths;



pub fn run_reminders(path: &str) {
    let now = Local::now();
    let mn = timeline::get_month(now.month());
    let d = now.day();
    let wkday = now.weekday();

    
    let conf = get_conf_settings(&path);
    
    if conf.is_ok() {
        let conf = conf.unwrap();
        let icon = settingparser::get_icon(conf);
        
        println!("{} {} {} {:?}", icon, mn, d, wkday);


        // if cache_is_old() {
        //   do checks
        //} else {
        //  printcache
        //}

        

    } else {
        let fmt_path = col::blue(&path);
        let s = format!("problem accessing bx's config file in '{}'", &fmt_path);
        api::crit_err(&s);
    }
}

fn get_conf_settings(path: &str) -> Result<String, Error> {

    let ext = paths::get_file(&"conf");
    let path = format!("{}{}", path, ext);

    let conf = fs::read_to_string(path)?;
    return Ok(conf);
}