use crate::errs;
use crate::funcs;

#[allow(dead_code)]
struct Bus {
    prompt_icon:    String,
	file_type:      String,
	func:           String,
	path:           String,
	help:           bool,
}

#[allow(dead_code)]
impl Bus {
    pub fn new() -> Bus {
        return Bus {
            prompt_icon:    "".to_owned(),
            file_type:      "".to_owned(),
            func:           "".to_owned(),
            path:           "".to_owned(),
            help:           false,
        }
    }

    pub fn set_icon(&mut self, s: &str) {
        self.prompt_icon = s.to_owned();
    }

    pub fn set_path(&mut self, p: &str) {
        self.path = p.to_owned();
    }
}

const VAL_FUNCS: [&str; 6] = [
    "add",
    "edit",
    "filter",
    "list",
    "update",
    "remove",
];

const VAL_DOCS: [&str; 5] = [
	"events",
	"ideas",
	"recurrents",
	"todos",
	"wishlist",
];

pub fn parse_args(args: &Vec<String>, path: &str) {

    let mut b = Bus::new();
    b.set_path(path);

    for r in args {

        let s = &r.as_str();

        if VAL_FUNCS.contains(s) {
            if b.func.len() == 0 {
                b.func = r.to_string();
            } else {
                let e = format!("args '{}' and '{}' conflict", b.func, r);
                errs::crit_err(&e);
            }
            continue;
        }

        if VAL_DOCS.contains(s) {
            if b.file_type.len() == 0 {
                b.file_type = r.to_string();
            } else {
                let e = format!("filetypes '{}' and '{}' conflict", b.file_type, r);
                errs::crit_err(&e);
            }
            continue;
        }

        if s.contains("-") {
            println!("gotta do a flag parse here");
        }
    } // end argparse loop


    funcs::exec_func();

}