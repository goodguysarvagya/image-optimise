# PNG to JPG Converter

This repository contains a Golang script that converts PNG images to JPG format in bulk using the [libvips](https://www.libvips.org/) image processing library. This script is designed to handle large numbers of images efficiently, taking advantage of libvips' high-performance capabilities.

## Why?

- Needed to convert lot of wallpapers to another format to save space.
- Why not PNG simply because Wallpapers do not need a transparency layer and i am fine with small artifacts.
- The savings in disk space are worth it. The jpegs are literally 1/10th of the size of png files sometimes.

## Features

- **Batch Processing**: Convert multiple PNG files to JPG in a single run.
- **High Performance**: Utilizes libvips for fast and memory-efficient image processing.

## Requirements

- **Go**: Version 1.20 or higher.
- **libvips**: You must have libvips binaries available in the script folder. Installation instructions for libvips can be found [here](https://www.libvips.org/install.html).

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/goodguysarvagya/image-optimise.git
   cd image-optimise
   ```

2. Install the necessary Go packages:

   ```bash
   go mod tidy
   ```

3. Ensure libvips is placed in proper location on your system.

## Usage

To convert all PNG files in a directory to JPG format:

```bash
go run main.go
```

This command will convert all PNG files in the `./original` directory to JPG format, saving them to the `./compressed` directory.

## Plans

- Make this with run with arguments like input, output, quality, special options like delete original on conversion.

## Contribution

Contributions are welcome! Please submit a pull request or open an issue if you encounter any bugs or have suggestions for improvements.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

---
