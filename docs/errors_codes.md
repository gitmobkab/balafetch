# Balafetch Exit Codes

To this day, balafetch registers 6 exit codes.

You might refer to these if you plan to write a script with balafetch or when debugging balafetch.

> [!NOTE]
> The balafetch CLI V0.4.0 and above use a different set of exit codes that are not compatible with the previous versions. If you have scripts that rely on the old exit codes, make sure to update them according to the new exit codes listed in this documentation.

## Exit Codes Table

| Code | Type | Reason | Category |
|------|------|--------|----------|
| 50 | Setup Failure | Happens when balafetch can't access the balafetch.log file (think of it as startup unsuccessful) | Fatal |
| 0 | Success | Happens when balafetch successfully completes (no Fatal/Fallible errors encountered) | Success |
| 1 | Internet Related Failures | Any internet-related error such as no internet connection, temporary failure in DNS resolution, or timeout | Failable |
| 2 | Command Not Found, Invalid Usage | Happens when there's no command named 'fastfetch' in your system | Fatal |
| 3 | Unexpected API Response | Happens when the Balatro API responds to a request in a format that balafetch doesn't know how to handle | Failable |
| 4 | File Related Failures | Happens when the OS refuses to grant balafetch file access. **This is different from Error 50 because it only triggers when balafetch tries to download a card image and can't access your temporary folder** | Failable |

## Fixes

### Code 50: Setup Failure
- Make sure that the `balafetch.log` file and `balafetch/` folder have the right permissions
- `balafetch/` needs **read**, **write**, and **execute** permissions for the user
- `balafetch/balafetch.log` needs at least **read** and **write** permissions for the user

**Linux/macOS/FreeBSD:**
```bash
chmod u=rwx ~/balafetch/
chmod u=rw ~/balafetch/balafetch.log
```

### Code 0: Success
Why would you fix something that works? 🎉

### Code 1: Internet Related Failures
- Check your internet connection
- Retry later
- Verify DNS is working properly

### Code 2: Command Not Found,Invalid Usage
- Install the **fastfetch** binary on your system/OS.
- Make sure the binary is named `fastfetch` and is in your system's PATH. You can verify this by running `which fastfetch` (Linux/macOS/FreeBSD) or `where fastfetch` (Windows) in your terminal.
- Make sure you're using the correct command syntax. You can refer to the [cli documentation for the correct usage.](cli.md)


### Code 3: Unexpected API Response
Open an issue on the [official balafetch repository](https://github.com/gitmobkab/balafetch/issues). This is a task for the developers.


### Code 4: File Related Failures
- Make sure your temporary folder has **write** permissions for the user
- Rare but possible: not enough storage space for balafetch to download the image. The images are very lightweight, but in such cases, try to free up some space or add more storage

**Check temporary folder permissions:**

**Linux/macOS/FreeBSD:**
```bash
# Check /tmp permissions
ls -ld /tmp
```

**Windows:**
```powershell
# Check %TEMP% permissions
icacls %TEMP%
```


**Installation options:**
- Package manager (recommended): `apt install fastfetch`, `brew install fastfetch`, `winget install fastfetch`, etc.
- Manual download from [fastfetch releases](https://github.com/fastfetch-cli/fastfetch/releases)
