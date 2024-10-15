# Projector Application

## Overview

This application is designed to manipulate and project configurations through a `Projector` class. The main operations include printing, adding, and removing values from a projected configuration based on user input.

## How It Works

The application parses user options and configuration settings, initializes a projector based on those settings, and performs different operations depending on the specified action (`Print`, `Add`, or `Remove`).

### Operations

1. **Print Operation**:
   - If no arguments are provided, the application prints the entire configuration as a JSON string.
   - If an argument is provided, it retrieves the value associated with the given argument and prints it.

2. **Add Operation**:
   - Adds a new key-value pair to the configuration based on the provided arguments.
   - Saves the updated configuration and prints `"added"` to confirm the action.

3. **Remove Operation**:
   - Removes a key-value pair from the configuration (if it exists).
   - Saves the updated configuration and prints `"added"` as confirmation.

### Command-Line Options

The application uses the following input structure:

- Options are parsed through the `getOpts()` function.
- Configuration is generated based on these options via `getConfig()`.

### Code Breakdown

- **`getOpts()`**: Retrieves the command-line options provided by the user.
- **`getConfig()`**: Loads the application's configuration based on the options.
- **`Projector`**: The core class responsible for getting, setting, and saving configuration values.

### Usage Examples

1. **Print the Entire Configuration**:
   ```
   node index.js --operation print
   ```

2. **Print a Specific Value**:
   ```
   node index.js --operation print --args key
   ```

3. **Add a Key-Value Pair**:
   ```
   node index.js --operation add --args key value
   ```

4. **Remove a Key-Value Pair**:
   ```
   node index.js --operation remove --args key
   ```

### Notes

- Make sure to pass the correct command-line arguments based on the operation you intend to execute.
- The `getOpts()` function will output the parsed options to the console as part of the execution flow.

## License

This project is licensed under the MIT License.

