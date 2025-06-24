Got it! Here’s the updated `README.md` with your `setup.sh` script included:

````markdown
# File Organizer

A simple Go CLI tool to rename files to snake_case and organize them into folders by file type.

## Features

- Rename files to snake_case (e.g., `My File.jpg` → `my_file.jpg`)
- Organize files into folders by extension type (images, documents, videos, audio, archives, others)
- Avoid overwriting by adding a numeric suffix for duplicate filenames
- Easy to extend and customize

## Requirements

- Go 1.18+ installed

## Installation

1. Clone the repository:

```bash
git clone https://github.com/yourusername/file-organizer.git
cd file-organizer
````

2. Build and install the binary globally with the provided script:

```bash
chmod +x setup.sh
./setup.sh
```

This will:

* Build the binary `organize-files`
* Copy it to `/usr/local/bin/` so you can run it anywhere as `organize-files`

## Usage

Run the organizer on a target folder (defaults to current folder if none provided):

```bash
organize-files /path/to/target/folder
```

Example:

```bash
organize-files ~/Downloads
```

## Supported file types and folders

| Folder    | Extensions                                    |
| --------- | --------------------------------------------- |
| images    | jpg, jpeg, png, gif, bmp, svg, webp           |
| documents | pdf, doc, docx, xls, xlsx, ppt, pptx, txt, md |
| videos    | mp4, avi, mov, mkv, flv                       |
| audio     | mp3, wav, aac, flac                           |
| archives  | zip, tar, gz, rar, 7z                         |
| others    | Any other file types                          |

## License

MIT License

