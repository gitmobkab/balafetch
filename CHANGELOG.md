# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).



## [Unreleased]


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