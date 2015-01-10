# KellyBackend

[![Build Status](https://secure.travis-ci.org/missdeer/KellyBackend.png)](https://travis-ci.org/missdeer/KellyBackend)
[![Total views](https://sourcegraph.com/api/repos/github.com/missdeer/KellyBackend/counters/views.png)](https://sourcegraph.com/github.com/missdeer/KellyBackend)
[![status](https://sourcegraph.com/api/repos/github.com/missdeer/KellyBackend/.badges/status.png)](https://sourcegraph.com/github.com/missdeer/KellyBackend)
[![Gobuild Download](http://gobuild.io/badge/github.com/missdeer/KellyBackend/downloads.svg)](http://gobuild.io/github.com/missdeer/KellyBackend)

Yiili community which is a part of [Kelly project](https://github.com/missdeer/kelly), it is the server side, AKA backend, of the project. The source code is based on WeTalk project, thanks those guys for their great work.

### Usage

```
go get -u github.com/missdeer/KellyBackend
cd $GOPATH/src/github.com/missdeer/KellyBackend
```

I suggest you [update all Dependencies](#dependencies)

Copy `conf/global/app.ini` to `conf/app.ini` and edit it. All configure has comment in it.

The files in `conf/` can overwrite `conf/global/` in runtime.


**Run KellyBackend**

```
bee run watchall
```

### Dependencies

Contrib

* Beego [https://github.com/astaxie/beego](https://github.com/astaxie/beego) 
* Social-Auth [https://github.com/beego/social-auth](https://github.com/beego/social-auth)
* Compress [https://github.com/beego/compress](https://github.com/beego/compress)
* i18n [https://github.com/beego/i18n](https://github.com/beego/i18n)
* Mysql [https://github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)
* goconfig [https://github.com/Unknwon/goconfig](https://github.com/Unknwon/goconfig)
* fsnotify [https://github.com/howeyc/fsnotify](https://github.com/howeyc/fsnotify)
* resize [https://github.com/nfnt/resize](https://github.com/nfnt/resize)
* blackfriday [https://github.com/slene/blackfriday](https://github.com/slene/blackfriday)

Update all Dependencies

```
go get -u github.com/beego/social-auth
go get -u github.com/beego/compress
go get -u github.com/beego/i18n
go get -u github.com/go-sql-driver/mysql
go get -u github.com/Unknwon/goconfig
go get -u github.com/howeyc/fsnotify
go get -u github.com/nfnt/resize
go get -u github.com/slene/blackfriday
```

### Static Files

KellyBackend use `Google Closure Compile` and `Yui Compressor` compress js and css files.

So you could need Java Runtime. Or close this feature in code by yourself.

### Contact

Maintain by [missdeer](http://minidump.info/)

## License

[Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0.html).

### Prebuilt Binaries Download

Download the pre-built binaries for variant platforms:

[Darwin x86](https://github.com/missdeer/KellyBackend/raw/prebuilt/KellyBackend-darwin-386)

[Darwin amd64](https://github.com/missdeer/KellyBackend/raw/prebuilt/KellyBackend-darwin-amd64)

[OpenBSD x86](https://github.com/missdeer/KellyBackend/raw/prebuilt/KellyBackend-openbsd-386)

[OpenBSD amd64](https://github.com/missdeer/KellyBackend/raw/prebuilt/KellyBackend-openbsd-amd64)

[NetBSD x86](https://github.com/missdeer/KellyBackend/raw/prebuilt/KellyBackend-netbsd-386)

[NetBSD amd64](https://github.com/missdeer/KellyBackend/raw/prebuilt/KellyBackend-netbsd-amd64)

[FreeBSD x86](https://github.com/missdeer/KellyBackend/raw/prebuilt/KellyBackend-freebsd-386)

[FreeBSD amd64](https://github.com/missdeer/KellyBackend/raw/prebuilt/KellyBackend-freebsd-amd64)

[FreeBSD arm6](https://github.com/missdeer/KellyBackend/raw/prebuilt/KellyBackend-freebsd-arm)

[Linux x86](https://github.com/missdeer/KellyBackend/raw/prebuilt/KellyBackend-linux-386)

[Linux amd64](https://github.com/missdeer/KellyBackend/raw/prebuilt/KellyBackend-linux-amd64)

[Linux arm6](https://github.com/missdeer/KellyBackend/raw/prebuilt/KellyBackend-linux-arm)

[Windows x86](https://github.com/missdeer/KellyBackend/raw/prebuilt/KellyBackend-windows-386), append the extension name ".exe" after downloaded.

[Windows amd64](https://github.com/missdeer/KellyBackend/raw/prebuilt/KellyBackend-windows-amd64), append the extension name ".exe" after downloaded.

