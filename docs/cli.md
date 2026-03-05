# Balafetch CLI Documentation

This document serves as a comprehensive guide to using the Balafetch Command Line Interface (CLI). Whether you're a new user or an experienced one, this guide will help you understand how to effectively use the CLI to fetch and display your Balatro profile information.

> [!NOTE]
> Heads up! The balafetch CLI V0.4.0 and above use a different set of exit codes that are not compatible with the previous versions. If you have scripts that rely on the old exit codes, make sure to update them according to the new exit codes listed in the [Balafetch Exit Codes documentation.](errors_codes.md)

## Usage
```bash
balafetch - The stupid balatro flavoured fastfetch wrapper
Usage: balafetch [OPTIONS] [CARD CATEGORY | ALIAS]
Options:
  -h, --help           Show help information
  -t, --timeout int    Set the timeout for API requests in seconds (default 20)
  -v, --version        Show version information
      --version-full   Show detailed version information
Available categories (aliases in parentheses):
 - jokers (joker)
 - tarot cards (tarot, tarots)
 - planet cards (planet, planets)
 - spectral cards (spectral, spectrals)
 - vouchers (voucher)
```
## Options
The balafetch CLI supports the following options:
- `-h`, `--help`: Display help information about the CLI and its usage.
- `-v`, `--version`: Display the current version of balafetch.
- `-t <seconds>`, `--timeout <seconds>`: Set a custom timeout for the CLI operations (default is 20 seconds),
    - Negative values are converted to their absolute value, so `-t -30` is the same as `-t 30`.

## Arguments 
- [CARD CATEGORY](#card-category)
- [ALIAS](#alias)

### CARD CATEGORY
CARD CATEGORY (optional): Specify a card category to fetch. If not provided, a random category will be selected. Valid categories include:
- jokers
- tarot cards
- planet cards
- spectral cards
- vouchers

note that the card category argument is case-insensitive, so you can use any combination of uppercase and lowercase letters (e.g., "Jokers", "jokers", "JOKERS" are all valid).

### ALIAS
ALIAS (optional): Specify an alias for a card category. Aliases are alternative names that can be used to refer to the same card category. The following aliases are supported:
| aliases | category |
| ------- | -------- |
| joker | `jokers` |
| tarot, tarots | `tarot cards` |
| planet, planets | `planet cards` |
| spectral, spectrals | `spectral cards` |
| voucher | `vouchers` |


i.e. `balafetch tarot` will not work, but `balafetch "tarot cards"` will work.

## Additional Information

- For flags that doesn't take any values, chaining is supported, so you can use `balafetch -hv` instead of `balafetch -h -v`.

- However this feature is irrelevant as the `help` flag is prioritized over the `version` flag, so `balafetch -hv` will only display the help information and ignore the version flag.

## FLAGS PRECEDENCE
When multiple flags are provided, balafetch will prioritize them in the following order:
1. `--help` (`-h`): If the help flag is present, balafetch will display the help information and ignore all other flags.
2. `--version` (`-v`): If the version flag is present (and the help flag is not), balafetch will display the version information and ignore all other flags except for `--version-full`.
3. `--version-full`: If this flag is present along with the version flag, balafetch will display detailed version information (including build time and commit hash) instead of just the version number.
4. `--timeout` (`-t`): If the timeout flag is present (and neither the help nor version flags are present), balafetch will use the specified timeout value for API requests.
5. Normal operation: If none of the above flags are present, balafetch will proceed with its normal operation of fetching and displaying a random card from the specified category (or a random category if none is specified).