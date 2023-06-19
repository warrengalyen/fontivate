# fontivate

![release](https://github.com/warrengalyen/fontivate/actions/workflows/release.yml/badge.svg)
[![GoDoc](https://godoc.org/github.com/warrengalyen/fontivate?status.svg)](https://godoc.org/github.com/warrengalyen/fontivate)
![Supported Version](https://img.shields.io/badge/go%20version-%3E%3D1.20-turquoise)
[![License](https://img.shields.io/github/license/warrengalyen/fontivate)](https://github.com/warrengalyen/fontivate/blob/master/LICENSE)
![Go version](https://img.shields.io/github/go-mod/go-version/warrengalyen/fontivate)

Fontivate is a command line font management tool to be used with my upcoming desktop Electron-based Font Manager.

## Installation

```sh
go install github.com/warrengalyen/fontivate@latest
```

Fontivate has been tested and supports Go versions >=1.20.x.

## Usage

Run `fontivate -h` to print help instructions.

---

## Font Installation

> The following commands must be used with elevated or administrative privileges.

### Adding Fonts

```sh
fontivate install "<path-source-folder>/Font-Name.otf"
```

### Removing Fonts

```sh
fontivate uninstall "<path-system-folder>/Font-Name.otf"
```

### Temporary Font Installation

Windows supports installing fonts temporarily. Fonts are automatically removed after a system reboot.

```sh
fontivate install --temporary=true "Font-Name.otf"
```

### Temporary Font Uninstallation

```sh
fontivate uninstall --temporary=true "Font-Name.otf"
```

### Multiple Fonts

Working with arrays of fonts is supported by separating each font path by a space.

```sh
fontivate install "Font-Name.otf" "Font-Name Bold.otf" "Font-Name Bold Italic.otf"
```

---

## Copy Commands

Utility commands to help manage system fonts.

### Finding Fonts

Finds and prints fonts including sub folders.

```sh
fontivate fonts find --root "C:\Fonts"
```

### Copying Fonts

Copies fonts from `source` to `destination` folder.

```sh
fontivate fonts copy --source "C:\Fonts" --destination "C:\Dest"
```

---

## License

See [LICENSE](https://github.com/warrengalyen/fontivate/blob/master/LICENSE).
