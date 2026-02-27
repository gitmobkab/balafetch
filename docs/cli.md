# Balafetch CLI Documentation

This document serves as a comprehensive guide to using the Balafetch Command Line Interface (CLI). Whether you're a new user or an experienced one, this guide will help you understand how to effectively use the CLI to fetch and display your Balatro profile information.

> [!NOTE]
> Heads up! The balafetch CLI V0.4.0 and above use a different set of exit codes that are not compatible with the previous versions. If you have scripts that rely on the old exit codes, make sure to update them according to the new exit codes listed in the [Balafetch Exit Codes documentation.](errors_codes.md)

## Usage
```bash
balafetch [OPTIONS] [CARD_CATEGORY]

Options:
  -h, --help          Show help information
  -t, --timeout int   Set the timeout for API requests in seconds (default 20)
  -v, --version       Show version information
Card Categories:
 - jokers
 - tarot cards
 - planet cards
 - spectral cards
 - vouchers
```
## Options
The balafetch CLI supports the following options:
- `-h`, `--help`: Display help information about the CLI and its usage.
- `-v`, `--version`: Display the current version of balafetch.
- `-t <seconds>`, `--timeout <seconds>`: Set a custom timeout for the CLI operations (default is 20 seconds),
    - Negative values are converted to their absolute value, so `-t -30` is the same as `-t 30`.

## Arguments
card_category (optional): Specify a card category to fetch. If not provided, a random category will be selected. Valid categories include:
- jokers
- tarot cards
- planet cards
- spectral cards
- vouchers

note that the card category argument is case-insensitive, so you can use any combination of uppercase and lowercase letters (e.g., "Jokers", "jokers", "JOKERS" are all valid).

unfortunately, the CLI does not support partial matching for card categories, so you need to provide the full category name. If an invalid category is provided, the CLI will display an error message along with a list of valid categories.

i.e. `balafetch tarot` will not work, but `balafetch "tarot cards"` will work.

## Additional Information

- For flags that doesn't take any values, chaining is supported, so you can use `balafetch -hv` instead of `balafetch -h -v`.

- However this feature is irrelevant as the `help` flag is prioritized over the `version` flag, so `balafetch -hv` will only display the help information and ignore the version flag.
