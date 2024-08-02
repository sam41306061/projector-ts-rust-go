use anyhow::{Ok, Result};
use clap::Parser;
use projector::{config::Config, opts::Opts};
fn main() -> Result<()> {
    let opts: Config = Opts::parse().try_into()?;
    println!("{:?}", opts);

    return Ok(());
}
