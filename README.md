# ORMEncrption

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

A command-line tool for encrypting and decrypting `orms_core_config.xml` files used in oppo and realme phones.

## Table of Contents

1. [ORMEncrption](#ormencrption)
   1. [Table of Contents](#table-of-contents)
   2. [Installation](#installation)
   3. [Usage](#usage)
      1. [Encrypt Command](#encrypt-command)
      2. [Decrypt Command](#decrypt-command)
   4. [Examples](#examples)
      1. [Encrypt Example](#encrypt-example)
      2. [Decrypt Example](#decrypt-example)
      3. [Print Decrypted Content to Terminal](#print-decrypted-content-to-terminal)
   5. [Contributing](#contributing)
   6. [License](#license)

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/ORMEncrption.git
   ```

2. Navigate to the project directory:
   ```sh
   cd ORMEncrption
   ```

3. Build and install the tool:
   ```sh
   go build -o ORMEncrption main.go
   ```

4. Optionally, move the binary to a location in your PATH if you want it available globally:
   ```sh
   sudo mv ORMEncrption /usr/local/bin/
   ```

## Usage

### Encrypt Command

The `encrypt` command is used to encrypt an XML configuration file (`orms_core_config.xml`) using AES/GCM and Base64 encoding. The encrypted content will be saved in a specified output file.

```sh
ORMEncrption encrypt -i <input-file-path> -o <output-file-path>
```

- `-i, --input`: Path to the input file (required).
- `-o, --output`: Path to the output file. If not specified, the default is `./orms_core_config_encrypted.xml`.

### Decrypt Command

The `decrypt` command is used to decrypt an AES/GCM encrypted and Base64 encoded XML configuration file (`orms_core_config.xml`). The decrypted content can be saved in a specified output file or printed directly into the terminal if no output file is provided.

```sh
ORMEncrption decrypt -i <input-file-path> [-o <output-file-path>] [--no-output]
```

- `-i, --input`: Path to the input file (required).
- `-o, --output`: Path to the output file. If not specified, the default is `./orms_core_config_decrypted.xml`.
- `--no-output`: Do not write to a file and print the decrypted content in the terminal instead.

## Examples

### Encrypt Example
```sh
$ ORMEncrption encrypt -i /path/to/orms_core_config.xml -o ./encrypted.xml
```

### Decrypt Example
```sh
$ ORMEncrption decrypt -i /path/to/encrypted.xml -o ./decrypted.xml
```

### Print Decrypted Content to Terminal
```sh
$ ORMEncrption decrypt -i /path/to/encrypted.xml --no-output
```

## Contributing

Contributions are welcome! Please read the [CONTRIBUTING.md](CONTRIBUTING.md) file for details on how to contribute.

## License

This project is licensed under the Apache-2.0 License - see the [LICENSE](LICENSE) file for details.
