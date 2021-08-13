use std::io::*;


#[allow(unused_imports)]
use crossterm::{
    Result,
    execute,
    cursor,
    terminal::{disable_raw_mode, enable_raw_mode, Clear, ClearType},
    event::{read, Event, KeyCode, KeyEvent, KeyModifiers}
};


pub fn read_normal_string() -> Result<String> {
    enable_raw_mode()?;
    let mut stdout = stdout();
    let mut line = String::new();

    while let Event::Key(KeyEvent { modifiers, code }) = read()? {
        match Event::Key(KeyEvent { modifiers, code }) {

            Event::Key(KeyEvent { 
                modifiers: KeyModifiers::CONTROL,
                code: KeyCode::Char('c')
            }) => {
                line.push_str("^C\n");
                execute!(stdout,
                    Clear(ClearType::CurrentLine),
                    cursor::MoveLeft(999),
                ).unwrap();

                let b = line.as_bytes();
                Write::write(&mut stdout, &b).unwrap();
                stdout.flush().unwrap();
                execute!(stdout, cursor::MoveLeft(999)).unwrap();
                disable_raw_mode()?;
                line = "<&&&>".to_string();
                break;
            }

            Event::Key(KeyEvent { 
                modifiers: _,
                code: KeyCode::Enter }) => 
            {
                disable_raw_mode()?;
                break;
            }

            Event::Key(KeyEvent { 
                modifiers: _,
                code: KeyCode::Char(c) }) => 
            {
                line.push(c);
                let b = [c as u8];
                Write::write(&mut stdout, &b).unwrap();
                stdout.flush().unwrap();
            }
            
            Event::Key(KeyEvent { 
                modifiers: _,
                code: KeyCode::Home }) => 
                {
                    line.push('Q');
                    execute!(stdout,
                        Clear(ClearType::CurrentLine),
                        cursor::MoveLeft(999),
                    ).unwrap();
                    stdout.flush().unwrap();
                }



            // KeyCode::Char(c) => {
            
            // }
            // KeyCode::Left   => {
            //     execute!(stdout,
            //         cursor::MoveLeft(1),
            //     ).unwrap();
            //     stdout.flush().unwrap();
            // }
            // KeyCode::Down => {
            //     line.push('z');
            //     let b = [b'z'];
            //     Write::write(&mut stdout, &b).unwrap();
            //     stdout.flush().unwrap();
            // }
            // KeyCode::Home => {
            //     execute!(stdout,
            //         cursor::MoveLeft(999),
            //     ).unwrap();
            //     stdout.flush().unwrap();
            // }
            // Event::Key(KeyEvent { code, modifiers }) => {
            //     if modifiers == (KeyModifiers::ALT | KeyModifiers::SHIFT) {
            //         println!("Alt + Shift {:?}", code);
            //     }
            // }
            // KeyCode::Esc  => {

            // }
            _ => {}
        }
    }

    disable_raw_mode()?;
    Ok(line)
}