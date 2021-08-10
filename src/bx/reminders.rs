use crate::bx::{settingparser, timeline};
use crate::colors::col;
use crate::errs::api;
use crate::pathutils::paths;
use chrono::{Datelike, Local};
use std::fs;
use std::io::Error;

pub fn run_reminders(path: &str) {
    let now = Local::now();
    let mn = timeline::get_month(now.month());
    let d = now.day();
    let wkday = now.weekday();

    let conf = get_conf_settings(&path);

    if conf.is_ok() {
        let conf = conf.unwrap();
        let icon = settingparser::get_icon(conf);
        let now_str = format!("{} {} {} {:?}", icon, mn, d, wkday);

        println!("{}", now_str);

        let cache = read_cache(&path);

        if cache.is_ok() {
            let cache = cache.unwrap();
            let cache: Vec<&str> = cache.split("\n").collect();
            if cache_is_old(&cache[0], &now_str) {
                write_cache(&path, &now_str);
            } else {
                println!("cache is not old");
                //  printcache
            }
        }
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

fn read_cache(path: &str) -> Result<String, Error> {
    let ext = paths::get_file(&"cache");
    let path = format!("{}{}", path, ext);

    let cache = fs::read_to_string(path)?;
    return Ok(cache);
}

fn cache_is_old(cached_str: &str, now_str: &str) -> bool {
    return !(cached_str == now_str);
}

fn write_cache(path: &str, now_str: &str) {
    let ext = paths::get_file(&"cache");
    let path = format!("{}{}", path, ext);
    let cache = fs::write(&path, &now_str);
    if cache.is_err() {
        let p = col::blue(&path);
        let s = format!("bx encountered problem accessing cache @ {}", &p);
        api::crit_err(&s);
    } else {
        cache.unwrap();
    }
}
