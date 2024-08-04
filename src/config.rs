use std::path::PathBuf;

use anyhow::{anyhow, Context, Ok, Result};

use crate::opts::Opts;
#[derive(Debug)]
pub struct Config {
    pub operation: Operation,
    pub config: PathBuf,
    pub pwd: PathBuf,
}

impl TryFrom<Opts> for Config {
    type Error = anyhow::Error;
    fn try_from(value: Opts) -> Result<Self> {
        let operation = value.args.try_into()?;
        let config = get_config(value.config)?;
        let pwd = get_pwd(value.pwd)?;

        return Ok(Config {
            operation,
            config,
            pwd,
        });
    }
}
#[derive(Debug, PartialEq)]
pub enum Operation {
    Print(Option<String>),
    Add(String, String), // tuple example ((String, String))
    Remove(String),
}

impl TryFrom<Vec<String>> for Operation {
    type Error = anyhow::Error;
    fn try_from(value: Vec<String>) -> Result<Self, Self::Error> {
        if value.len() == 0 {
            return Ok(Operation::Print(None));
        }

        let mut value = value;

        let term = value.get(0).expect("expect to exist!");
        if term == "add" {
            if value.len() != 3 {
                let err = anyhow!("Expects 2 arguments but got:{}", value.len() - 1);
                return Err(err);
            }
            let mut drain = value.drain(1..=2);
            return Ok(Operation::Add(
                drain.next().expect("To Exist 1"),
                drain.next().expect("To Exist 2"),
            ));
        }
        if term == "remove" {
            if value.len() != 2 {
                let err = anyhow!(
                    "operation remove expects 1 arguments but got:{}",
                    value.len() - 1
                );
                return Err(err);
            }
            let arg = value.pop().expect("To exist remove");
            return Ok(Operation::Remove(arg));
        }
        if value.len() > 1 {
            let err = anyhow!(
                "operation print expects 0 or 1 arguments but got:{}",
                value.len()
            );
            return Err(err);
        }
        let arg = value.pop().expect("To exist remove");
        return Ok(Operation::Print(Some(arg)));
    }
}
fn get_config(config: Option<PathBuf>) -> Result<PathBuf> {
    // get the pathbuf , return the pathbuf
    if let Some(v) = config {
        return Ok(v);
    }
    let loc = std::env::var("XDG_CONFIG_HOME").context("unable to get xdg_config_home object")?;
    let mut loc = PathBuf::from(loc);

    loc.push("projector");
    loc.push("projector.json");

    return Ok(loc);
}

fn get_pwd(pwd: Option<PathBuf>) -> Result<PathBuf> {
    if let Some(p) = pwd {
        return Ok(p);
    }

    return std::env::current_dir().context("unable to get std::env::current_dir");
}

#[cfg(test)]
#[cfg(test)]
#[cfg(test)]
mod test {
    use anyhow::Result;

    use crate::{config::Config, config::Operation, opts::Opts};

    use super::get_config;

    #[test]
    fn test_print() -> Result<()> {
        let opts: Config = Opts {
            args: vec![],
            pwd: None,
            config: None,
        }
        .try_into()?;

        assert_eq!(opts.operation, Operation::Print(None));

        return Ok(());
    }
    #[test]
    fn test_print_key() -> Result<()> {
        let opts: Config = Opts {
            args: vec!["foo".to_string()],
            pwd: None,
            config: None,
        }
        .try_into()?;

        assert_eq!(opts.operation, Operation::Print(Some(String::from("foo"))));

        return Ok(());
    }

    #[test]
    fn test_add_key_value() -> Result<()> {
        let opts: Config = Opts {
            args: vec![
                String::from("add"),
                String::from("foo"),
                String::from("bar"),
            ],
            pwd: None,
            config: None,
        }
        .try_into()?;

        assert_eq!(
            opts.operation,
            Operation::Add(String::from("foo"), String::from("bar"))
        );

        return Ok(());
    }
    #[test]
    fn test_remove_key() -> Result<()> {
        let opts: Config = Opts {
            args: vec![String::from("remove"), String::from("foo")],
            pwd: None,
            config: None,
        }
        .try_into()?;

        assert_eq!(opts.operation, Operation::Remove(String::from("foo")));

        return Ok(());
    }

    //#[test]
    // fn test_add() -> Result<()> {
    //     let opts: Config = Opts {
    //         config: None,
    //         pwd: None,
    //         args: vec!["add".to_string(), "foo".into(), String::from("bar")],
    //     };

    //     let config = get_config(opts)?;

    //     assert_eq!(
    //         config.operation,
    //         Operation::Add((String::from("foo"), String::from("bar")))
    //     );

    //     return Ok(());
    // }
}
