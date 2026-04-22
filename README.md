(AI GEN SUMMARY)
## File Organizer (Go) – README

This project is a simple command-line file organizer written in Go. It scans a
target directory (default: `~/Downloads/`) and automatically sorts files into
folders based on their file extensions. This is for personal use and is not 
intended for you.

### How it works

* The program reads all files in the specified directory.
* It extracts each file’s extension.
* It matches the extension against a predefined map (`extensionMap`) that categorizes files into groups such as:

  * `image` (jpg, png, gif, etc.)
  * `video` (mp4, mov, mkv, etc.)
  * `audio` (mp3)
  * `code` (go, py, cpp, etc.)
  * `document` (pdf, docx, txt, etc.)
  * `archive` (zip, rar, 7z)
  * `web`, `notebook`, `iso`, and `other`
* It creates the target folder if it doesn’t exist.
* It moves the file into the appropriate folder.

### Key functions

* `getFileType(filename string)`
  Extracts the file extension from a filename.

* `moveFileToFolder(root, filename, folder string)`
  Ensures the destination folder exists and moves the file into it.

* `main()`
  Iterates over all files in the target directory and organizes them.

### Usage

Update the `path` variable in `main()` to point to your desired directory, then run:

```bash
go run main.go
```

## NEW FEATURES
*   CLI Flags: Use the flag package to accept the target directory as a command-line argument instead of hardcoding it in
main.go.
*   Configuration File: Move the extensionMap to an external config.json or config.yaml file. This allows you to update
categories without recompiling the binary.

Robustness & Safety
*   Dry Run Mode: Add a -dry-run flag that prints which files would be moved without performing the actual move operations.
*   Conflict Resolution: Implement logic to handle cases where a file with the same name already exists in the destination
folder (e.g., append a timestamp or a counter to the filename).
*   Error Logging: Instead of just failing, log errors to a file so you can debug why certain files were not moved (e.g.,
permission issues).

Features
*   Ignore List: Allow users to define a list of files or folders to explicitly ignore (e.g., .DS_Store, specific system
folders).
*   Recursive Sorting: Add an option to scan subdirectories within the target directory rather than just the top level.


