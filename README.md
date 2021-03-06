# Checksum
- support crc32 md5 sha1 sha256

## Usage
```
Usage:                                  
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
  4) checksum -sha1 t1.txt t2.txt
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
