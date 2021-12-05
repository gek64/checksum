# Checksum
- support crc32 md5 sha1 sha256
- Written in golang

## Usage
```
Version:
  1.00

Usage:
  checksum {Mode} [Command] files...

Mode:
  -md5              : use md5 (default)
  -crc32            : use crc32
  -sha1             : use sha1
  -sha256           : use sha256

Command:
  -h                : show help
  -v                : show version

Example:
  1) checksum text.txt photo.jpg
  1) checksum -md5 text.txt photo.jpg
  1) checksum -crc32 text.txt photo.jpg
  2) checksum -sha1 text.txt photo.jpg
  3) checksum -sha256 text.txt photo.jpg
  2) checksum -h
  3) checksum -v
```

## Build
```sh
# dependence
git clone https://github.com/gek64/gek.git

git clone https://github.com/gek64/checksum.git

cd checksum

go build -v -trimpath -ldflags "-s -w"
```

## QA

### Q: Windows Security detect `.exe` as `Trojan:Win32/Wacatac.B!ml`
A: This report occurred after `Windows 10 21h2`. This application does not contain any malware, backdoors, and advertisements, all released files are build by github actions. For more information, see https://go.dev/doc/faq#virus

### Q: Why should I clone `https://github.com/gek64/gek.git` before building
A: I don’t want the project to depend on a certain cloud service provider, and this is also a good way to avoid network problems.


## License

**GNU Lesser General Public License v2.1**

See `LICENSE` for details
