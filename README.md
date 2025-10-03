# go-app-installer (installexecutable)

A simple Go program to move a file to the `/bin` directory. Useful for system management and quick setup scripts.

---

## Features

- Easy to use
- Moves a file to `/bin`
- Requires root privileges (must run with sudo)

## Requirements

- Go 1.x
- Linux or Unix-based OS
- Root privileges

---

## Installation

1. Clone the repo:
```bash
git clone https://github.com/<username>/installexecutable.git
cd installexecutable
```

2. Build the program:
```bash
go build -o installexecutable main.go
```

3. Usage:
```bash
sudo ./installexecutable <AppName>
```

> ⚠️ Warning: Requires root privileges. It won't work without `sudo`.

---

## Example Usage

```bash
sudo ./installexecutable myscript.sh
```

This command moves the `myscript.sh` file to `/bin` and makes it executable system-wide.

---
