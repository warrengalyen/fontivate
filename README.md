# fontivate

![release](https://github.com/warrengalyen/fontivate/actions/workflows/release.yml/badge.svg)
[![GoDoc](https://godoc.org/github.com/warrengalyen/fontivate?status.svg)](https://godoc.org/github.com/warrengalyen/fontivate)
![Supported Version](https://img.shields.io/badge/go%20version-%3E%3D1.19-turquoise)
[![Go Report Card](https://goreportcard.com/badge/github.com/warrengalyen/fontivate)](https://goreportcard.com/report/github.com/warrengalyen/fontivate)
[![License](https://img.shields.io/github/license/warrengalyen/fontivate)](https://github.com/warrengalyen/fontivate/blob/master/LICENSE)

`fontivate` is a command line font management tool used with my upcoming desktop Electron-based Font Manager.

## Installation

```sh
go install github.com/warrengalyen/fontivate@latest
```

Fontivate has been tested and supports Go versions >=1.20.x.

## Usage

Run `fontivate -h` to print help instructions.

---

### Installing Fonts

```sh
fontivate install "Font-Name.otf"
```

### Uninstalling Fonts

```sh
fontivate uninstall "Font-Name.otf"
```

### Multiple Fonts

Installing multiple fonts is supported by separating paths by a space.

```sh
fontivate install "Font-Name.otf" "Font-Name Bold.otf" "Font-Name Bold Italic.otf"
```

### Temporary Installation

Windows supports installing fonts temporarily. Fonts are removed after a system reboot.

```sh
fontivate install --temporary=true "Font-Name.otf"
```

### Temp Uninstallation

```sh
fontivate uninstall --temporary=true "Font-Name.otf"
```

---

## Font Management

Included utility commands to help manage system fonts.

### Finding Fonts

Finds and prints fonts including sub folders.

```sh
fontivate fonts find --root "C:\Fonts"
```

### Copying Fonts

Copies fonts from source to destination folder.

```sh
fontivate fonts copy --source "C:\Fonts" --destination "C:\Dest"
```

---

## License

See [LICENSE](https://github.com/warrengalyen/fontivate/blob/master/LICENSE).
