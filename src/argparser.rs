use crate::errs;
use crate::funcs;

#[allow(dead_code)]
pub struct Bus {
    pub prompt_icon: String,
    pub file_type: String,
    pub func: String,
    pub path: String,
    pub help: bool,
}

#[allow(dead_code)]
impl Bus {
    pub fn new() -> Bus {
        return Bus {
            prompt_icon: "".to_owned(),
            file_type: "".to_owned(),
            func: "".to_owned(),
            path: "".to_owned(),
            help: false,
        };
    }

    pub fn set_icon(&mut self, s: &str) {
        self.prompt_icon = s.to_owned();
    }

    pub fn set_path(&mut self, p: &str) {
        self.path = p.to_owned();
    }
}

const VAL_FUNCS: [&'static str; 6] = ["add", "edit", "filter", "list", "update", "remove"];

const VAL_DOCS: [&'static str; 5] = ["events", "ideas", "recurrents", "todos", "wishlist"];

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
                errs::api::crit_err(&e);
            }
        } else if VAL_DOCS.contains(s) {
            if b.file_type.len() == 0 {
                b.file_type = r.to_string();
            } else {
                let e = format!("filetypes '{}' and '{}' conflict", b.file_type, r);
                errs::api::crit_err(&e);
            }
        } else if s.contains("-") {
            println!("gotta do a flag parse here");
        } else if s == &"help" {
            b.help = true;
        } else {
            let s = format!("unrecognized arg '{}'", s);
            errs::api::crit_err(&s);
        }
    } // end argparse loop

    if bus_is_valid(&b) {
        funcs::exec_func(b);
    } else {
        println!("gooch");
    }
}

fn bus_is_valid(b: &Bus) -> bool {
    let boolean = b.func.len() * b.file_type.len();
    if boolean != 0 {
        return true;
    }

    return false;
}
