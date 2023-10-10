# Bebecare api with Golang

이 api 는 go 1.20.6 version 으로 개발 되었습니다.

이 api 는 http 통신을 위해 go.gin web framework 를 사용하였습니다.

## Install go

go 최신버전 설치를 위해 다음의 절차를 따르십시오.

https://go.dev/dl/ 로 이동하여 자신의 개발환경에 맞는 버전을 받아 설치하십시오.
(본 예제 에서는 개발에 사용된 1.20.6 버전 설치를 예문으로 사용합니다. 최신 버전 확인은 다음의 링크에서 확인하십시오. https://go.dev/dl/)

```bash
cd bebecare-mobile-web #project location
npm install
npm run local
```

api 서비스를 binary 로 build하여 실행하기 위해 다음을 수행하십시오.

(이 서비스는 golang 설치가 필수적으로 요구됩니다. 다음의 링크로 이동해 node.js 설치 후 하단의 절차를 따르십시오. https://nodejs.org/ko)


```bash
cd bebecare-mobile-web #project location
npm install
npm run local
```

## Dev server service up

dev server 환경에서 서비스를 실행하기 위해 하단의 절차를 따르십시오.

(해당 서비스는 Ubuntu 18.04 및 20.04 에서 테스트 되었습니다.)

```bash
cd bebecare-mobile-web #project location
npm install
npm run build-dev
```

npm background 실행을 위해 빌드 전 다음을 수행하십시오.

```bash
(nohup npm run build-dev &)
```

## Etc...

* 각 서비스 환경에 따라 .env 파일을 제공합니다. 해당 환경에 대한 설정은 rollup.config.js 파일을 확인하십시오.
* npm 패키지 버전에 대한 상세한 내용은 지원하지 않습니다. 호환성 오류는 npm update 를 참고하십시오.
* nohup 에 의한 process 실행을 중단할 경우 다음을 참고하십시오.
```bash
ps -ef | grep node* #search for node process to this project directory
kill -9 {PID}
```

END