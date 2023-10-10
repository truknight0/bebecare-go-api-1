# Bebecare api with Golang

이 api 는 go 1.20.6 version 으로 개발되었습니다.

이 api 는 http 통신을 위해 go.gin web framework 를 사용하였습니다.

이 api 는 local 과 dev 환경을 제공합니다.

## Install go

go 최신버전 설치를 위해 다음의 절차를 따르십시오.

https://go.dev/dl/ 로 이동하여 자신의 개발환경에 맞는 버전을 받아 설치하십시오.

설치 후 다음을 수행하여 정상적으로 설치되었는지 확인하십시오.

```bash
go version #go version go1.20.6 darwin/arm64
go env #golang evironment variables list
```

운영체제에 따른 `GOPATH` 설정은 운영체제별 `GOPATH 설정` 을 검색하십시오. (ex : mac gopath 설정)

## API Startup

go 소스를 빌드하여 바이너리 파일을 생성하고 실행하기 위해 다음을 수행하십시오. (윈도우의 경우 bash 프로그램 혹은 cmd 를 활용하십시오.)

```bash
cd bebecare-go-api-1 #project location
go get
go build server.go
```

바이너리 파일명을 변경할 경우 다음의 명령어를 사용하십시오.

```bash
go build -o {binary file name} server.go
#ex: go build -o bin/apiserver server.go
```

## Etc
* 다음의 쉘 스크립트 파일 apistart.sh 은 실행된 바이너리 파일을 자동으로 종료, 실행하는 기능을 수행합니다. 
* 배포시 자동으로 바이너리 파일의 실행 프로세스를 종료하고 재시작 하기위해 다음의 명령어를 이용하십시오.
```bash
sh apistart.sh
```

END