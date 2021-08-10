const MONTHS: [&'static str; 12] = [
    "Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec",
];

pub fn get_month(i: u32) -> String {
    let i = i as usize;
    return MONTHS[i].to_owned();
}
