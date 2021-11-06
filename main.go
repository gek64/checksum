package main

import (
	"flag"
	"fmt"
	"gek_checksum"
	"os"
)

var (
	cliHelp    bool
	cliVersion bool
	cliMode    string
	cliMd5     bool
	cliCrc32   bool
	cliSha1    bool
	cliSha256  bool
)

func init() {
	flag.BoolVar(&cliMd5, "md5", false, "use md5 (default)")
	flag.BoolVar(&cliCrc32, "crc32", false, "use crc32")
	flag.BoolVar(&cliSha1, "sha1", false, "use sh1")
	flag.BoolVar(&cliSha256, "sha256", false, "use sh256")
	flag.BoolVar(&cliHelp, "h", false, "show help")
	flag.BoolVar(&cliVersion, "v", false, "show version")
	flag.Parse()

	// 重写显示用法函数
	flag.Usage = func() {
		var helpInfo = `Version:
  1.00

Usage:
  checksum [Command] {files...}

Command:
  -md5              : use md5 (default)
  -crc32            : use crc32
  -sha1             : use sha1
  -sha256           : use sha256
  -h                : show help
  -v                : show version

Example:
  1) checksum text.txt photo.jpg
  1) checksum -md5 text.txt photo.jpg
  1) checksum -crc32 text.txt photo.jpg
  2) checksum -sha1 text.txt photo.jpg
  3) checksum -sha256 text.txt photo.jpg
  2) checksum -h
  3) checksum -v`
		fmt.Println(helpInfo)
	}

	// 哈希函数模式(默认md5)
	if cliCrc32 {
		cliMode = "crc32"
	} else if cliMd5 {
		cliMode = "md5"
	} else if cliSha1 {
		cliMode = "sha1"
	} else if cliSha256 {
		cliMode = "sha256"
	} else {
		cliMode = "md5"
	}

	// 打印帮助信息
	if cliHelp || flag.NArg() == 0 {
		flag.Usage()
		os.Exit(0)
	}

	// 打印版本信息
	if cliVersion {
		showVersion()
		os.Exit(0)
	}
}

func showVersion() {
	var versionInfo = `Changelog:
  1.00:
    - First release`

	fmt.Println(versionInfo)
}

func main() {
	for i := 0; i < flag.NArg(); i++ {
		checksum, err := gek_checksum.Checksum(cliMode, flag.Args()[i])
		if err == nil {
			fmt.Println(flag.Args()[i], checksum)
		}
	}
}
