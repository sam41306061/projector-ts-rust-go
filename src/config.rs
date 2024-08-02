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
        let pwd = get_config(value.pwd)?;

        return Ok(Config {
            operation,
            config,
            pwd,
        });
    }
}
#[derive(Debug)]
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
