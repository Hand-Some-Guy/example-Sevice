# 디바이스 관리 API

이 프로젝트는 Go와 Gin 프레임워크를 사용해 디바이스와 관련 펌웨어를 관리하는 RESTful API입니다. 계층 구조 아키텍처를 따르며, 컨트롤러, 서비스, 레포지토리, 도메인 모델로 책임을 분리했습니다. API는 디바이스 생성, 조회, 상태 업데이트를 지원하며, 서비스 타입에 기반한 펌웨어 정보를 활용합니다.

## 목차
- [프로젝트 구조](#프로젝트-구조)
- [사전 요구사항](#사전-요구사항)
- [설치](#설치)
- [프로젝트 빌드](#프로젝트-빌드)
- [애플리케이션 실행](#애플리케이션-실행)
- [API 엔드포인트](#api-엔드포인트)
- [테스트](#테스트)
- [기여](#기여)
- [라이선스](#라이선스)

## 프로젝트 구조

이 프로젝트는 유지보수성과 확장성을 위해 모듈화된 계층 구조 아키텍처를 따릅니다. 아래는 주요 디렉터리와 파일의 구조 및 설명입니다:

```
project/
├── Dockerfile                # 애플리케이션 빌드 및 실행을 위한 Docker 설정
├── go.mod                    # Go 모듈 의존성 정의
├── go.sum                    # 의존성 체크섬
├── main.go                   # 애플리케이션 진입점, Gin 서버 및 의존성 초기화
├── Jenkinsfile               # 애플리케이션 CI 파이프라인 실행을 위한 Jenkins 설정 
├── controllers/              # HTTP 요청 처리 (프레젠테이션 계층)
│   └── device_controller.go  # 디바이스 관련 API 요청 처리
├── services/                 # 비즈니스 로직 (서비스 계층)
│   └── device_service.go     # 디바이스 및 펌웨어 작업 구현
├── repositories/             # 데이터 접근 계층 (레포지토리 계층)
│   ├── device_repository.go  # 디바이스 메모리 저장소
│   └── firmware_repository.go # 펌웨어 메모리 저장소
├── models/                   # 도메인 모델 및 로직 (도메인 계층)
│   ├── device.go             # 디바이스 구조체 및 관련 메서드
│   └── firmware.go           # 펌웨어 구조체 및 관련 메서드
└── README.md                 # 프로젝트 문서 (이 파일)
```

### 파일 설명
- **`Dockerfile`**: `golang:1.23`으로 빌드하고 `alpine:latest`로 실행하는 멀티스테이지 빌드 설정. 8080 포트를 노출.
- **`Jenkinsfile`** : 
- **`go.mod` 및 `go.sum`**: `github.com/gin-gonic/gin` 등 Go 의존성을 관리.
- **`main.go`**: Gin 라우터 설정, 의존성 주입, 초기 펌웨어 데이터 설정.


## 사전 요구사항
- **Go**: 버전 1.23 이상 (`go version`으로 확인).
- **Docker**: 컨테이너 빌드 및 배포를 위해 필요 (`docker --version`으로 확인).
- **curl** 또는 **Postman**: API 엔드포인트 테스트용.
- **Git**: 리포지토리 클론 시 필요 (선택).

## 설치
1. 리포지토리를 클론합니다:
   ```bash
   git clone https://github.com/Hand-Some-Guy/example-Sevice.git
   cd project
   ```
2. Go 의존성을 설치합니다:
   ```bash
   go mod download
   ```

## 프로젝트 빌드
프로젝트는 로컬 또는 Docker를 사용해 빌드할 수 있습니다.

### 로컬 빌드
Go 애플리케이션을 컴파일합니다:
```bash
go build -o main .
```

### Docker 빌드
제공된 `Dockerfile`을 사용해 Docker 이미지를 빌드합니다:
```bash
docker build -t my-device-api:latest .
```
- 경량화된 이미지(~15MB)를 생성.

## 애플리케이션 실행
애플리케이션은 로컬 또는 Docker 컨테이너에서 실행할 수 있습니다.

### 로컬 실행
컴파일된 바이너리를 실행합니다:
```bash
./main
```
- 서버는 `http://localhost:8080`에서 시작.

### Docker 실행
Docker 컨테이너를 실행합니다:
```bash
docker run -d -p {Host Port}:8080 --name device-api my-device-api:latest
```
- 호스트의 8080 포트를 컨테이너의 8080 포트에 매핑.
- 컨테이너 실행 상태 확인:
  ```bash
  docker ps
  ```

서버 로그 확인:
```bash
docker logs device-api
```

## API 엔드포인트
API는 디바이스 관리를 위한 엔드포인트를 제공하며, 모든 엔드포인트는 `/devices`로 시작.

| 메서드 | 엔드포인트             | 설명                           | 요청 바디 예시                                         |
|--------|------------------------|-------------------------------|-------------------------------------------------------|
| POST   | `/devices`             | 새로운 디바이스 생성           | `{"id":"dev1","service_type":"NEW"}`                  |
| GET    | `/devices/:id`         | ID로 디바이스 조회            | 해당 없음                                             |
| PUT    | `/devices/:id/state`   | 디바이스 상태 업데이트         | `{"state":"UPDATE"}`                                  |

### 요청 예시
1. **디바이스 생성**:
   ```bash
   curl -X POST http://localhost:8080/devices \
   -H "Content-Type: application/json" \
   -d '{"id":"dev1","service_type":"NEW"}'
   ```
   응답:
   ```json
   {"id":"dev1","service":"NEW","firmwareID":"fw1","firmwareVersion":"1.0.0","state":"WAIT"}
   ```

2. **디바이스 조회**:
   ```bash
   curl http://localhost:8080/devices/dev1
   ```

3. **디바이스 상태 업데이트**:
   ```bash
   curl -X PUT http://localhost:8080/devices/dev1/state \
   -H "Content-Type: application/json" \
   -d '{"state":"UPDATE"}'
   ```

## 테스트
서버가 정상 작동하는지 확인하려면:
1. 애플리케이션을 실행 (로컬 또는 Docker).
2. `curl` 또는 Postman으로 API 엔드포인트를 테스트 (위 예시 참조).
3. 서버 로그에서 오류 확인:
   ```bash
   docker logs device-api
   ```
4. Go 테스트로 자동화된 테스트 실행:
   ```bash
   go test ./...
   ```

간단한 테스트 스크립트(`test.sh`) 예시:
```bash
#!/bin/bash
docker build -t my-device-api:latest .
docker run -d -p 8080:8080 --name device-api my-device-api:latest
sleep 2
curl -X POST http://localhost:8080/devices -H "Content-Type: application/json" -d '{"id":"dev1","service_type":"NEW"}'
docker stop device-api
docker rm device-api
```

## 라이선스
이 프로젝트는 MIT 라이선스 하에 배포됩니다. 자세한 내용은 [LICENSE](LICENSE) 파일을 참조하세요.