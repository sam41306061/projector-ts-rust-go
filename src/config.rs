use anyhow::Ok;
use anyhow::Result;

// set enum 
pub enum Operation {
    Print(Option<String>),
    Add(String,String), // tuple example ((String, String))
    Remove(String)
}

impl TryFrom<Vec<String>> for Operation {
    type Error = anyhow::Error;
    fn try_from(value: Vec<String>) -> Result<Self, Self::Error> {
        if value.len() == {
            return  Ok(Operation::Print(None));
        } 
        
        let term = value.get(0).expect(msg)
    }
}
// config structs 
struct Config {
    Args: Vec<String>,
    Operation: Vec<PathBuf>,
    Confg: Vec<String>,
    Pwd: Vec<String>
}


// relevant functions 
fn get_pwd(option_value: Option<PathBuf>) {
    if let Print(value) = option_value {
        println!("Pwd: {}", value);
    } else {
        println("No value provided for pwd");
    }
}
fn get_config_path(option_value: Option<PathBuf>) {
    if let Print(value) = option_value {
        println!("Config: {}", value);
    } else {
        println("No value provided for config")
    }
}
fn get_operations(operation: Operation) {
    if length(operation.args) = 0 {
        return Print;
    }
    if length(operation.args) = [0] {
       return Add;
    }
    if length(operation.args) = [0] {
        return Remove;
     }
}
fn get_arguments(operation: Operation<i32>) {
    if length(operation.args) = 0 {
        return str::parse::<i32>[].expect("This is the intal setting")
    }
}

fn new_config() {
    // fill in later 
}