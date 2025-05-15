# 디바이스 관리 API

이 프로젝트는 Go와 Gin 프레임워크를 사용해 디바이스, 펌웨어, 사용자 인증을 관리하는 RESTful API입니다. 계층 구조 아키텍처를 따르며, 컨트롤러, 서비스, 레포지토리, 도메인 모델로 책임을 분리했습니다. API는 디바이스 생성/조회/상태 업데이트, 펌웨어 생성/삭제/조회, 사용자 로그인 기능을 지원하며, 서비스 타입에 기반한 펌웨어 정보를 활용합니다.

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
│   ├── device_controller.go  # 디바이스 관련 API 요청 처리
│   ├── firmware_controller.go # 펌웨어 관련 API 요청 처리
│   └── user_controller.go    # 사용자 인증 관련 API 요청 처리
├── services/                 # 비즈니스 로직 (서비스 계층)
│   ├── device_service.go     # 디바이스 작업 구현
│   ├── firmware_service.go   # 펌웨어 작업 구현
│   └── user_service.go       # 사용자 인증 작업 구현
├── repositories/             # 데이터 접근 계층 (레포지토리 계층)
│   ├── device_repository.go  # 디바이스 메모리 저장소
│   ├── firmware_repository.go # 펌웨어 메모리 저장소
│   └── user_repository.go    # 사용자 메모리 저장소
├── models/                   # 도메인 모델 및 로직 (도메인 계층)
│   ├── device.go             # 디바이스 구조체 및 관련 메서드
│   ├── firmware.go           # 펌웨어 구조체 및 관련 메서드
│   └── user.go               # 사용자 구조체 및 관련 메서드
└── README.md                 # 프로젝트 문서 (이 파일)
```

### 파일 설명
- **`Dockerfile`**: `golang:1.23`으로 빌드하고 `alpine:latest`로 실행하는 멀티스테이지 빌드 설정. 8080 포트를 노출.
- **`Jenkinsfile`**: Jenkins를 사용한 CI 파이프라인 설정으로, 빌드와 테스트를 자동화.
- **`go.mod` 및 `go.sum`**: `github.com/gin-gonic/gin`, `github.com/Masterminds/semver/v3`, `golang.org/x/crypto`, `github.com/golang-jwt/jwt/v5` 등 Go 의존성을 관리.
- **`main.go`**: Gin 라우터 설정, 의존성 주입, 초기 펌웨어 및 사용자 데이터 설정.

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
docker run -d -p 8083:8080 --name device-api my-device-api:latest
```
- 호스트의 8083 포트를 컨테이너의 8080 포트에 매핑 (Nginx 프록시 설정 반영).
- 컨테이너 실행 상태 확인:
  ```bash
  docker ps
  ```

서버 로그 확인:
```bash
docker logs device-api
```

## API 엔드포인트

API는 디바이스, 펌웨어, 사용자 인증을 관리하기 위한 엔드포인트를 제공합니다. 아래는 주요 엔드포인트와 설명입니다.

### 디바이스 엔드포인트
디바이스 관리 API는 디바이스를 생성, 조회, 상태를 업데이트합니다. 디바이스는 서비스 타입(`NEW`, `OLD`), 펌웨어 정보, 상태(`PENDING`, `COMPLETED`, `FAILED`)를 저장합니다.

| 메서드 | 엔드포인트             | 설명                           | 요청 바디 예시                                         |
|--------|------------------------|-------------------------------|-------------------------------------------------------|
| POST   | `/devices`             | 새로운 디바이스 생성           | `{"id":"dev1","service_type":"NEW"}`                  |
| GET    | `/devices/:id`         | ID로 디바이스 조회            | 해당 없음                                             |
| PUT    | `/devices/:id/state`   | 디바이스 상태 업데이트         | `{"state":"COMPLETED"}`                               |

### 펌웨어 엔드포인트
펌웨어 관리 API는 펌웨어를 생성, 삭제, 서비스 타입별 최신 펌웨어를 조회합니다. 펌웨어는 서비스 타입(`NEW`, `OLD`), 버전, 파일 경로를 저장합니다.

| 메서드 | 엔드포인트                    | 설명                           | 요청 바디 예시                                         |
|--------|-------------------------------|-------------------------------|-------------------------------------------------------|
| POST   | `/firmwares`                  | 새로운 펌웨어 생성             | `{"id":"fw4","service_type":"NEW","version":"1.0.2","path":"/firmware_test.bin"}` |
| DELETE | `/firmwares/:id`              | ID로 펌웨어 삭제              | 해당 없음                                             |
| GET    | `/firmwares/service/:service` | 서비스 타입으로 최신 펌웨어 조회 | 해당 없음                                             |

### 사용자 엔드포인트
사용자 인증 API는 로그인을 처리하고 JWT 토큰을 반환합니다. 사용자는 ID와 비밀번호로 인증되며, 계정 잠금 상태를 관리합니다.

| 메서드 | 엔드포인트 | 설명               | 요청 바디 예시                                         |
|--------|------------|-------------------|-------------------------------------------------------|
| POST   | `/login`   | 사용자 로그인 및 JWT 토큰 발급 | `{"id":"testuser","password":"password123"}`          |

## 테스트
서버가 정상 작동하는지 확인하려면 아래 단계를 따르세요.

1. **애플리케이션 실행**:
   - 로컬 또는 Docker로 서버를 실행 (위 [애플리케이션 실행](#애플리케이션-실행) 참조).
   - Nginx 프록시를 통해 `http://localhost:8083`으로 요청.

2. **테스트용 파일 생성**:
   - 펌웨어 경로 유효성 검증을 위해 테스트용 파일 생성:
     ```bash
     mkdir testdata
     touch testdata/firmware.bin
     touch testdata/firmware_new.bin
     touch testdata/firmware_old.bin
     touch testdata/firmware_test.bin
     ```

3. **curl 명령어로 테스트**:
   아래는 각 도메인의 주요 엔드포인트를 테스트하는 `curl` 명령어입니다.

   ### 사용자 테스트
   - **로그인 성공**:
     ```bash
     curl -v -X POST http://localhost:8083/login \
     -H "Content-Type: application/json" \
     -d '{"id":"testuser","password":"password123"}'
     ```
     **기대 응답**:
     ```json
     {"message":"로그인 성공","token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."}
     ```

   - **로그인 실패 (잘못된 비밀번호)**:
     ```bash
     curl -v -X POST http://localhost:8083/login \
     -H "Content-Type: application/json" \
     -d '{"id":"testuser","password":"wrongpassword"}'
     ```
     **기대 응답**:
     ```json
     {"error":"잘못된 비밀번호 또는 계정이 잠겨 있습니다"}
     ```

   ### 펌웨어 테스트
   - **펌웨어 생성**:
     ```bash
     curl -v -X POST http://localhost:8083/firmwares \
     -H "Content-Type: application/json" \
     -d '{"id":"fw4","service_type":"NEW","version":"1.0.2","path":"testdata/firmware_test.bin"}'
     ```
     **기대 응답**:
     ```json
     {"id":"fw4","service":"NEW","version":"1.0.2","path":"testdata/firmware_test.bin"}
     ```

   - **최신 펌웨어 조회**:
     ```bash
     curl -v http://localhost:8083/firmwares/service/NEW
     ```
     **기대 응답**:
     ```json
     {"id":"fw2","service":"NEW","version":"1.0.1","path":"testdata/firmware_new.bin"}
     ```

   - **펌웨어 삭제**:
     ```bash
     curl -v -X DELETE http://localhost:8083/firmwares/fw4
     ```
     **기대 응답**:
     ```json
     {"message":"펌웨어가 삭제되었습니다"}
     ```

   ### 디바이스 테스트
   - **디바이스 생성**:
     ```bash
     curl -v -X POST http://localhost:8083/devices \
     -H "Content-Type: application/json" \
     -d '{"id":"dev1","service_type":"NEW"}'
     ```
     **기대 응답**:
     ```json
     {
         "id": "dev1",
         "service": "NEW",
         "firmwareID": "fw2",
         "firmwareVersion": "1.0.1",
         "status": "PENDING"
     }
     ```

   - **디바이스 조회**:
     ```bash
     curl -v http://localhost:8083/devices/dev1
     ```
     **기대 응답**:
     ```json
     {
         "id": "dev1",
         "service": "NEW",
         "firmwareID": "fw2",
         "firmwareVersion": "1.0.1",
         "status": "PENDING"
     }
     ```

   - **디바이스 상태 업데이트**:
     ```bash
     curl -v -X PUT http://localhost:8083/devices/dev1/state \
     -H "Content-Type: application/json" \
     -d '{"state":"COMPLETED"}'
     ```
     **기대 응답**:
     ```json
     {"message":"state changed"}
     ```

4. **로그 확인**:
   - 서버 로그에서 요청 처리 및 오류 확인:
     ```bash
     docker logs device-api
     ```

## 기여
1. 리포지토리를 포크합니다.
2. 기능 브랜치를 생성합니다 (`git checkout -b feature/기능-이름`).
3. 변경 사항을 커밋합니다 (`git commit -m '기능 추가'`).
4. 브랜치를 푸시합니다 (`git push origin feature/기능-이름`).
5. 풀 리퀘스트를 엽니다.



## 라이선스
이 프로젝트는 MIT 라이선스 하에 배포됩니다. 자세한 내용은 [LICENSE](LICENSE) 파일을 참조하세요.