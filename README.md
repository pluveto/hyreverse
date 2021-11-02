# HyReverse

```go
|  |      __   ___       ___  __   __   ___ 
|__| \ / |__) |__  \  / |__  |__) /__` |__  
|  |  |  |  \ |___  \/  |___ |  \ .__/ |___ 

```

Simple and handy reverse proxy for http/https services.

## Get started

```console
i@myhost:~$ ./hyreverse --remote https://www.google.com --local http://localhost:23333
```

Now visit `http://localhost:23333`.

## Usage

```console
i@myhost:~$ ./hyreverse.exe -h
usage: hyreverse.exe [-h|--help] -l|--local "<value>" -r|--remote "<value>"
                     [-c|--cert "<value>"] [-k|--key "<value>"]

                     simple and handy reverse proxy for http/https services

Arguments:

  -h  --help    Print help information
  -l  --local   Local address to host on. Example: http://localhost:8080
  -r  --remote  Remote address. Example: https://www.google.com
  -c  --cert    Cert file path
  -k  --key     Key file path

```