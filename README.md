# HyReverse

```go
          __   ___       ___  __   __   ___ 
|__| \ / |__) |__  \  / |__  |__) /__` |__  
|  |  |  |  \ |___  \/  |___ |  \ .__/ |___ 
                                            
```

Simple and handy reverse proxy for http/https services.

## Get started

```console
i@myhost:~$ ./hyreverse --remote http://www.google.com --local http://localhost:23333
```

Now visit `http://localhost:23333`.