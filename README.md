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
