# Checksum
- support crc32 md5 sha1 sha256

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

## License

**GPL-3.0 License**

See `LICENSE` for details
