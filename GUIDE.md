# 실행 가이드

## 요구사항

- Go 1.13 이상
- Docker
- Docker Compose

## 1. 프로젝트 클론

```bash
git clone https://github.com/hyeonjeong33/backend-onboard-task.git
cd backend-onboard-task
```

## 2. 의존성 설치

Go 모듈 의존성을 설치합니다. **처음 프로젝트를 실행하는 경우에** 아래 명령어를 사용하여 모듈을 설치하세요.

```bash
go mod download
```

이 명령어는 `go.mod` 파일에 정의된 모든 의존성을 다운로드하여 프로젝트를 실행할 수 있도록 준비합니다.

## 3. Docker 컨테이너 실행

```bash
docker-compose up -d
```

이 명령어를 실행하면 MySQL 데이터베이스와 애플리케이션이 Docker 컨테이너로 실행됩니다.  
(Docker Compose 파일에서 환경 변수를 자동으로 설정하기 때문에 별도의 `.env` 파일 설정이 필요하지 않습니다.)

## 4. 애플리케이션 실행

```bash
go run main.go
```

애플리케이션이 성공적으로 실행되면, `http://localhost:8080`에서 API를 사용할 수 있습니다.

## 5. API 테스트

Postman 또는 cURL을 통해 API를 테스트할 수 있습니다. 

#### 요청 예시 (bash)

```bash
curl -X GET http://localhost:8080/boards?page=1&limit=10
```

## 6. 종료

애플리케이션과 Docker 컨테이너를 종료하려면 다음 명령어를 사용하세요.

```bash
docker-compose down
```

## 기타

더 자세한 내용은 프로젝트의 [API 명세서](https://github.com/hyeonjeong33/backend-onboard-task/blob/main/API_SPEC.md)를 참고하세요.
