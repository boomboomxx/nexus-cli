[![CircleCI](https://circleci.com/gh/mlabouardy/nexus-cli.svg?style=svg)](https://circleci.com/gh/mlabouardy/nexus-cli) [![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

<div align="center">
<img src="logo.png" width="60%"/>
</div>

Nexus CLI for Docker Registry

## Usage

<div align="center">
<img src="example.png"/>
</div>

## Build
```shell
# init project
go mod init nexus-cli
# fetch dependency
go mod tidy
# build cli
go build -o nexus-cli .
```


## Download

Below are the available downloads for the latest version of Nexus CLI (1.0.0-beta). Please download the proper package for your operating system and architecture.

### Linux:

```
wget https://nexus.szistech.com/repository/raw-szis-releases/sziscloud/nexus-cli/1.0.1/nexus-cli
```

### Windows:

```
wget https://nexus.szistech.com/repository/raw-szis-releases/sziscloud/nexus-cli/1.0.1/nexus-cli
```

### Mac OS X:

```
wget https://nexus.szistech.com/repository/raw-szis-releases/sziscloud/nexus-cli/1.0.1/nexus-cli
```

### OpenBSD:

```
wget https://nexus.szistech.com/repository/raw-szis-releases/sziscloud/nexus-cli/1.0.1/nexus-cli
```

### FreeBSD:

```
wget https://nexus.szistech.com/repository/raw-szis-releases/sziscloud/nexus-cli/1.0.1/nexus-cli
```

## Available Commands

```
$ nexus-cli configure
```

```
$ nexus-cli image ls
```

```
$ nexus-cli image tags -name mlabouardy/nginx
```

```
$ nexus-cli image info -name mlabouardy/nginx -tag 1.2.0
```

```
$ nexus-cli image delete -name mlabouardy/nginx -tag 1.2.0
```

```
$ nexus-cli image delete -name mlabouardy/nginx -keep 4
```

```
# delete image tags with szis rules
$ nexus-cli image szis delete -name mlabouardy/nginx -k 2
```

```
$ nexus-cli image size -name mlabouardy/nginx
```
## Tutorials

* [Cleanup old Docker images from Nexus Repository](http://www.blog.labouardy.com/cleanup-old-docker-images-from-nexus-repository/)
