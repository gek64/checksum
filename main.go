package main

import (
	"flag"
	"fmt"
	"gek_checksum"
	"gek_path"
	"log"
	"os"
	"path/filepath"
	"strings"
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
	flag.BoolVar(&cliMd5, "md5", false, "use md5")
	flag.BoolVar(&cliCrc32, "crc32", false, "use crc32")
	flag.BoolVar(&cliSha1, "sha1", false, "use sh1 (default)")
	flag.BoolVar(&cliSha256, "sha256", false, "use sh256")
	flag.BoolVar(&cliHelp, "h", false, "show help")
	flag.BoolVar(&cliVersion, "v", false, "show version")
	flag.Parse()

	// 重写显示用法函数
	flag.Usage = func() {
		var helpInfo = `Usage:
  checksum {Mode} [Command] files...

Mode:
  -md5              : use md5
  -crc32            : use crc32
  -sha1             : use sha1 (default)
  -sha256           : use sha256

Command:
  -h                : show help
  -v                : show version

Example:
  1) checksum /root/*.txt
  2) checksum t1.txt t2.txt
  3) checksum -sha1 *
  4) checksum -sha1 t1.txt t2.txt`
		fmt.Println(helpInfo)
	}

	// 哈希函数模式(默认sha1)
	if cliCrc32 {
		cliMode = "crc32"
	} else if cliMd5 {
		cliMode = "md5"
	} else if cliSha1 {
		cliMode = "sha1"
	} else if cliSha256 {
		cliMode = "sha256"
	} else {
		cliMode = "sha1"
	}

	// 打印版本信息
	if cliVersion {
		fmt.Println(`v1.03`)
		os.Exit(0)
	}

	// 打印帮助信息
	if cliHelp || flag.NArg() == 0 {
		flag.Usage()
		os.Exit(0)
	}
}

func showChangelog() {
	var versionInfo = `Changelog:
  1.00:
    - First release
  1.01:
    - Add wildcard(*,?) support in file name
  1.02:
    - Change default from md5 to sha1(faster and safer)
  1.03:
    - Clean code`
	fmt.Println(versionInfo)
}

func main() {
	files, err := WildcardMatchFile(flag.Args())
	if err != nil {
		log.Panicln(err)
	}

	for _, file := range files {
		checksum, err := gek_checksum.Checksum(cliMode, file)
		if err == nil {
			fmt.Println(checksum, file)
		}
	}
}

// WildcardMatchFile 通配符匹配文件,pathList 示例: []string{ "./*", "test/1.txt" }
func WildcardMatchFile(pathList []string) (matchList []string, err error) {
	// 从多地址列表中获取单个地址
	for _, filePath := range pathList {
		// 地址文件名部分
		var base = filepath.Base(filePath)
		// 地址路径部分
		var dir = filepath.Dir(filePath)

		// 如果文件名部分带有通配符
		if strings.Contains(base, "*") || strings.Contains(base, "?") {
			// 匹配文件地址部分中的所有文件
			var fileList []string
			err := gek_path.WalkAll(dir, &fileList, false)
			if err != nil {
				return nil, err
			}

			// 所有文件中按文件名部分进行匹配
			for _, file := range fileList {
				// 正则匹配模式
				matched, err := filepath.Match(base, filepath.Base(file))
				if err != nil {
					return nil, err
				}
				// 匹配成功的加入到匹配列表中
				if matched {
					matchList = append(matchList, file)
				}
			}
		} else {
			// 一般无通配符路径则获取绝对路径存储到匹配列表中
			fullPath, err := filepath.Abs(filePath)
			if err != nil {
				return nil, err
			}
			matchList = append(matchList, fullPath)
		}
	}
	return matchList, nil
}
