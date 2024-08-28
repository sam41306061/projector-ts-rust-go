use std::{collections::HashMap, path::PathBuf};

use serde::{Deserialize, Serialize};

use crate::config::Config;

type HM = HashMap<String, String>;
#[derive(Debug, Serialize, Deserialize)]
struct Data {
    pub projector: HashMap<PathBuf, HM>,
}

pub struct Projector {
    config: Config,
    data: Data,
}

fn default_data() -> Data {
    return Data {
        projector: HashMap::new(),
    };
}

impl Projector {
    pub fn from_config(config: Config) -> Self {
        if std::fs::metadata(config.config.clone()).is_ok() {
            let contents = std::fs::read_to_string(config.config.clone());
            let contents = contents.unwrap_or(String::from("{\"projector\":{}}"));
            let data = serde_json::from_str(&contents);
            let data = data.unwrap_or(default_data());
            return Projector { config, data };
        }

        return Projector {
            config,
            data: default_data(),
        };
    }
}
