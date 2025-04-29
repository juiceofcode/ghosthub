# GhostHub CLI

A CLI tool to manage multiple Git profiles with SSH keys.

## Installation

```bash
go install github.com/juiceofcode/ghosthub-cli
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

## License

MIT 