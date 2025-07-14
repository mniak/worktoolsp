use std::fs;
use std::process;

use clap::{Arg, Command};
use clipboard_rs::Clipboard;
use clipboard_rs::ClipboardContext;
#[derive(serde::Serialize)]
struct QueueMessage {
    #[serde(rename = "MessageBody")]
    message_body: String,
}

fn main() {
    let matches = Command::new("windmill_hsmactions")
        .arg(
            Arg::new("json-list")
                .long("json-list")
                .action(clap::ArgAction::SetTrue)
                .help("Treat input as a single JSON file containing a list of filenames"),
        )
        .arg(
            Arg::new("filenames")
                .required(true)
                .num_args(1..)
                .help("Input file(s) or single JSON file with --json-list"),
        )
        .get_matches();

    let filenames: Vec<_> = matches.get_many::<String>("filenames").unwrap().collect();
    let mut messages = Vec::new();

    for filename in filenames {
        let content = fs::read_to_string(filename).unwrap_or_else(|err| {
            eprintln!("Error reading {}: {}", filename, err);
            process::exit(1);
        });
        if matches.get_flag("json-list") {
            // Parse the content as a JSON array of strings
            let json_msgs: Vec<serde_json::Value> =
                serde_json::from_str(&content).unwrap_or_else(|err| {
                    eprintln!("Error parsing JSON from {}: {}", filename, err);
                    process::exit(1);
                });

            // Read each file in the list and add as queue messages
            for json_msg in json_msgs {
                messages.push(QueueMessage {
                    message_body: serde_json::to_string(&json_msg).unwrap_or_else(|err| {
                        eprintln!("Error serializing JSON message: {}", err);
                        process::exit(1);
                    }),
                });
            }
        } else {
            messages.push(QueueMessage {
                message_body: content,
            });
        }
    }

    let messages_json = serde_json::to_string_pretty(&messages).unwrap();
    println!("{}", messages_json);
    if let Err(e) = ClipboardContext::new().and_then(|ctx| ctx.set_text(messages_json.clone())) {
        eprintln!("Failed to copy to clipboard: {}", e);
    }

    let timestamp = chrono::Utc::now().format("%Y%m%d_%H%M%S");
    let filename = format!("output_{}.json", timestamp);

    if let Err(e) = fs::write(&filename, &messages_json) {
        eprintln!("Failed to write to file {}: {}", filename, e);
    } else {
        println!("Output written to: {}", filename);
    }
}
