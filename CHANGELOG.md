# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).



## [Unreleased]

## [0.8.0] - 2026-06-05

### Added
- `--update` / `-u` flag for self-updating to the latest release from GitHub
  - Downloads and verifies binary integrity via SHA256 checksums
  - Atomic binary replacement (safe on Windows, Linux, macOS, FreeBSD)
  - Cross-platform support: automatically detects and fetches the right binary for your OS/architecture
  - Returns exit code 5 on update failure; exit code 0 if already up-to-date
- `compile.bat`: Windows batch script for building binaries on Windows (equivalent to `compile.sh`)
  - Cross-platform builds for all supported OS/architecture combinations
  - SHA256 checksum generation via built-in `certutil` (no extra dependencies)

### Changed
- Release build process now generates checksums in `sha256sum` text-mode format (two spaces) for better tool compatibility
- `compile.sh`: removed `-b` flag from `sha256sum` to produce text-mode checksums

### Fixed
- Fixed checksum file format to be compatible with `go-selfupdate` validator (two-space separator instead of asterisk)

## [0.7.0] - 2026-03-28

### Added
- Live progress display during network operations showing elapsed time against timeout (e.g. `fetching... [1.4/20 seconds]`)
- `-t 0` / `--timeout 0` restored to mean no timeout (infinite)
- Negative timeout values are silently corrected to their absolute value for backward compatibility

### Changed
- `--timeout` help text updated to document 0 behavior

### Improved
- Minor performance improvement on network operations via shared HTTP client
## [0.6.1] - 2026-03-05

### Fixed
- Fixed a bug where balafetch would corrupt fastfetch output if any operations failed when using the inlineStreamer

## [0.6.0] - 2026-03-05

### Added
- Alias support for card categories (e.g. `joker`, `tarots`, `tarot`, `planet`, `planets`, `spectral`, `spectrals`, `voucher`)
- `--version-full` flag to display detailed build information (build time, commit hash)
- Live single-line feedback during network operations

### Changed
- `--version` output shortened to `V{version} ({os}/{arch})`
- Improved error message for invalid category — includes quote hint and available categories
- Help output now displays aliases next to each canonical category

### Fixed
- Image download failure incorrectly returned wrong exit code

## [0.5.1] - 2026-02-27

### Fixed
- Fixed bug where balafetch would print "fastfetch not installed" when `card_category` is provided but invalid, even if fastfetch is installed.

## [0.5.0] - 2026-02-27

### Changed
- Updated `balafetch --version` output to include build time and commit hash information, in addition to the version number.
- Added optional 'card_category' argument to support case-insensitive matching and added error handling for invalid categories, providing a user-friendly message listing valid categories.

### Removed
- Go-figure dependency for version output (replaced with simple formatted text output)


## [0.4.0] - 2026-02-12

### Added
- SHA256 checksums for all binaries in the release notes (available in the "Assets" section of the release and since v0.4.0)
- GNU/UNIX style long and short flags support
    - `-h`, `--help`: display help message
    - `-v` ,`--version`: display version information

- `-t`, `--timeout` flag to specify a timeout for fetching data (e.g., `balafetch -t 5s` to set a 5-second timeout)

### Changed
- Updated installation instructions to reflect the new versioning scheme and the addition of checksums in the release notes.
- cli Exit codes

## [0.3.0] - 2026-01-27

### Added

- Command-line flags support [DEPREACATED]
    - `-h`: display help message
    - `-v`: display version information

> [!NOTE]
> The flags in this version doesn't follow the GNU/UNIX convention.
> only the short flags are supported. 

## [0.2.0] - 2026-01-16

- **First actual release** (v0.1.0 was tagged but never released - missing Windows binaries)
- Added Windows binaries
- Initial feature set

## [0.1.0] - Never released
- Tagged but incomplete (missing Windows builds)