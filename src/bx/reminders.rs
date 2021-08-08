
use std::fs;
use chrono::{Datelike, Local};
use crate::bx::timeline;
use crate::errs::api;

pub fn run_reminders(path: &str) {
    let now = Local::now();
    let mn = timeline::get_month(now.month());
    let d = now.day();
    let dw = now.weekday();

    println!("{} {}, {:?} // {:?}", mn, d, dw, now);
    println!("{:?}", path);
}

fn get_conf_settings(path: &str) -> String {
    
    // err handle conf
    //parse in another function

}