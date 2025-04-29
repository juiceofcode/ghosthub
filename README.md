# GhostHub CLI

A CLI tool to manage multiple Git profiles with SSH keys.

## Installation

### Option 1: Using Go Install (Recommended for developers)

If you have Go installed, you can install GhostHub directly:

```bash
go install github.com/juiceofcode/ghosthub-cli@latest
```

### Option 2: Manual Installation

#### Linux/macOS

1. Download the latest release for your system:
   - For Linux: `ghosthub-linux-amd64`
   - For macOS (Intel): `ghosthub-darwin-amd64`
   - For macOS (Apple Silicon): `ghosthub-darwin-arm64`

2. Make the file executable and move it to your PATH:
```bash
chmod +x ghosthub-<your-system>
sudo mv ghosthub-<your-system> /usr/local/bin/ghosthub
```

#### Windows

1. Download the latest release: `ghosthub-windows-amd64.exe`
2. Rename it to `ghosthub.exe`
3. Move it to a directory in your PATH (e.g., `%USERPROFILE%\AppData\Local\Microsoft\WindowsApps`)

### Option 3: Using Installation Scripts

#### Linux/macOS

1. Clone the repository:
```bash
git clone https://github.com/juiceofcode/ghosthub-cli.git
cd ghosthub-cli
```

2. Build and install:
```bash
./build.sh
./install.sh
```

#### Windows

1. Clone the repository:
```bash
git clone https://github.com/juiceofcode/ghosthub-cli.git
cd ghosthub-cli
```

2. Build and install:
```bash
build.bat
install.bat
```

## Usage

### Add a new Git profile

```bash
ghosthub add [profile] --name "Your Name" --email "your.email@example.com" --keygen "ed25519"
```

Options:
- `--name`: Your Git user name
- `--email`: Your Git email
- `--keygen`: SSH key type (ed25519 or rsa-4096)

### Delete a Git profile

```bash
ghosthub delete [profile]
```

### List all Git profiles

```bash
ghosthub list
```

### Switch to a Git profile

```bash
ghosthub use [profile]
```

### Show current Git profile information

```bash
ghosthub info
```

### Generate a new SSH key

```bash
ghosthub key [profile] --type "ed25519"
```

Options:
- `--type`: SSH key type (ed25519 or rsa-4096)

## Examples

### Add a new profile

```bash
ghosthub add work --name "John Doe" --email "john@company.com" --keygen "ed25519"
```

### Switch to a profile

```bash
ghosthub use work
```

### List all profiles

```bash
ghosthub list
```

### Delete a profile

```bash
ghosthub delete work
```

### Generate a new SSH key

```bash
ghosthub key work --type "ed25519"
```

## System Requirements

- Go 1.16 or higher (for development)
- Git
- OpenSSH (for SSH key management)
- Windows 10/11, macOS 10.15+, or Linux

## License

MIT 