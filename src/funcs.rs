use crate::argparser::Bus;

pub fn exec_func(b: Bus) {
    println!("bfunc {}", b.func);
    println!("bfiletype {}", b.file_type);
}
