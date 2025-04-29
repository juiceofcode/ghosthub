# ghosthub-cli

A cross-platform Git profile manager that allows you to easily switch between different identities and SSH keys for different contexts (work, freelance, personal projects, etc).

## Installation

```bash
go install github.com/juiceofcode/ghosthub-cli@latest
```

## Usage

### Add a new profile

```bash
ghosthub profile add my-profile --email="your@email.com" --ssh-key="/path/to/ssh_key" --name="Your Name"
```

The `--name` parameter is optional. If not provided, the profile name will be used.

### List profiles

```bash
ghosthub profile list
```

### Switch to a profile

```bash
ghosthub profile switch my-profile
```

### Remove a profile

```bash
ghosthub profile remove my-profile
```

## Profile Structure

Profiles are stored in `~/.ghosthub/profiles.json` in the following format:

```json
{
  "freelancer": {
    "name": "John Smith",
    "email": "john@freelance.com",
    "sshKeyPath": "/home/user/.ssh/id_freelance"
  },
  "company": {
    "name": "John Corp",
    "email": "john@company.com",
    "sshKeyPath": "/home/user/.ssh/id_company"
  }
}
```

## Features

- ✅ Multiple Git profile management
- ✅ Support for different SSH keys per profile
- ✅ Automatic configuration of user.name and user.email
- ✅ Compatible with Windows, macOS and Linux
- ✅ Support for any Git service (GitHub, GitLab, Bitbucket, etc)

## Contributing

Contributions are welcome! Please feel free to submit pull requests. 