# API 명세

## 1. 회원가입 API

- **URL**: `/signup`
- **Method**: `POST`
- **Description**: 새로운 사용자를 등록하는 API. 사용자는 이메일과 비밀번호를 입력하여 회원가입할 수 있습니다.

### 요청 예시 (bash)

```bash
curl -X POST http://localhost:8080/signup \
-H "Content-Type: application/json" \
-d '{"email": "user@example.com", "password": "test1234"}'
```

### 요청 예시 (json)

```json
{
  "email": "user@example.com",
  "password": "test1234"
}
```

### 요청 파라미터

| Parameter | Type   | Description              | Required |
| --------- | ------ | ------------------------ | -------- |
| email     | string | 회원 이메일 (@ 포함)     | Yes      |
| password  | string | 회원 비밀번호 (8자 이상) | Yes      |

### 응답 예시 (성공 시)

```json
{
  "message": "회원가입이 완료되었습니다."
}
```

### 응답 예시 (에러 시)

```json
{
  "error": "이미 등록된 이메일입니다."
}
```

---

## 2. 로그인 API

- **URL**: `/login`
- **Method**: `POST`
- **Description**: 등록된 이메일과 비밀번호를 통해 사용자가 로그인하고 JWT 토큰을 발급받는 API.

### 요청 예시 (bash)

```bash
curl -X POST http://localhost:8080/login \
-H "Content-Type: application/json" \
-d '{"email": "user@example.com", "password": "test1234"}'
```

### 요청 예시 (json)

```json
{
  "email": "user@example.com",
  "password": "test1234"
}
```

### 요청 파라미터

| Parameter | Type   | Description   | Required |
| --------- | ------ | ------------- | -------- |
| email     | string | 회원 이메일   | Yes      |
| password  | string | 회원 비밀번호 | Yes      |

### 응답 예시 (성공 시)

```json
{
  "message": "로그인 성공",
  "token": "eyJhbGciOiJIUzI1NiIsInR..."
}
```

### 응답 예시 (에러 시)

```json
{
  "error": "존재하지 않는 사용자입니다."
}
```

---

## 3. 게시글 생성 API (JWT 인증 필요)

- **URL**: `/boards`
- **Method**: `POST`
- **Description**: 인증된 사용자가 게시글을 작성하는 API. 제목과 내용을 입력해 게시글을 생성합니다.

### 요청 예시 (bash)

```bash
curl -X POST http://localhost:8080/boards \
-H "Authorization: Bearer {JWT_TOKEN}" \
-H "Content-Type: application/json" \
-d '{"title": "새로운 게시글", "content": "게시글 내용"}'
```

### 요청 예시 (json)

```json
{
  "title": "새로운 게시글",
  "content": "게시글 내용"
}
```

### 요청 파라미터

| Parameter | Type   | Description | Required |
| --------- | ------ | ----------- | -------- |
| title     | string | 게시글 제목 | Yes      |
| content   | string | 게시글 내용 | Yes      |

### 헤더

| Header        | Description                                      | Required |
| ------------- | ------------------------------------------------ | -------- |
| Authorization | `Bearer {JWT_TOKEN}` 형태로 JWT 인증 토큰을 전달 | Yes      |
| Content-Type  | `application/json`                               | Yes      |

### 응답 예시 (성공 시)

```json
{
  "message": "게시글이 성공적으로 생성되었습니다."
}
```

---

## 4. 게시글 목록 조회 API

- **URL**: `/boards`
- **Method**: `GET`
- **Description**: 게시글 목록을 조회하는 API. `page`와 `limit` 쿼리 파라미터를 사용해 페이지네이션을 적용할 수 있습니다.

### 요청 예시 (bash)

```bash
curl -X GET http://localhost:8080/boards?page=1&limit=10
```

### 요청 파라미터

| Parameter | Type | Description                       | Required |
| --------- | ---- | --------------------------------- | -------- |
| page      | int  | 페이지 번호 (기본값 1)            | No       |
| limit     | int  | 한 페이지의 게시글 수 (기본값 10) | No       |

### 응답 예시 (성공 시)

```json
{
  "boards": [
    {
      "id": 1,
      "title": "첫 번째 게시글",
      "content": "내용",
      "views": 10,
      "createdAt": "2024-09-07T12:34:56Z"
    },
    {
      "id": 2,
      "title": "두 번째 게시글",
      "content": "내용",
      "views": 5,
      "createdAt": "2024-09-07T12:35:00Z"
    }
  ],
  "total": 50,
  "page": 1,
  "limit": 10
}
```

---

## 5. 특정 게시글 조회 API

- **URL**: `/boards/:id`
- **Method**: `GET`
- **Description**: 특정 게시글을 조회하는 API. 게시글의 ID를 받아 해당 게시글의 상세 정보를 반환하며, 조회 시 조회수가 증가합니다.

### 요청 예시 (bash)

```bash
curl -X GET http://localhost:8080/boards/1
```

### 응답 예시 (성공 시)

```json
{
  "id": 1,
  "title": "첫 번째 게시글",
  "content": "이것은 첫 번째 게시글의 내용입니다.",
  "views": 11,
  "createdAt": "2024-09-07T12:34:56Z"
}
```

### 응답 예시 (에러 시)

```json
{
  "error": "해당 게시글을 찾을 수 없습니다"
}
```

---

## 6. 게시글 수정 API (JWT 인증 필요)

- **URL**: `/boards/:id`
- **Method**: `PUT`
- **Description**: 인증된 사용자가 게시글을 수정하는 API. 해당 게시글을 작성한 사용자만 수정할 수 있습니다.

### 요청 예시 (bash)

```bash
curl -X PUT http://localhost:8080/boards/1 \
-H "Authorization: Bearer {JWT_TOKEN}" \
-H "Content-Type: application/json" \
-d '{"title": "수정된 제목", "content": "수정된 내용"}'
```

### 요청 예시 (json)

```json
{
  "title": "수정된 제목",
  "content": "수정된 내용"
}
```

### 요청 파라미터

| Parameter | Type   | Description | Required |
| --------- | ------ | ----------- | -------- |
| title     | string | 게시글 제목 | No       |
| content   | string | 게시글 내용 | No       |

### 헤더

| Header        | Description                                      | Required |
| ------------- | ------------------------------------------------ | -------- |
| Authorization | `Bearer {JWT_TOKEN}` 형태로 JWT 인증 토큰을 전달 | Yes      |
| Content-Type  | `application/json`                               | Yes      |

### 응답 예시 (성공 시)

```json
{
  "message": "게시글이 성공적으로 수정되었습니다."
}
```

---

## 7. 게시글 삭제 API (JWT 인증 필요)

- **URL**: `/boards/:id`
- **Method**: `DELETE`
- **Description**: 인증된 사용자가 게시글을 삭제하는 API. 해당 게시글을 작성한 사용자만 삭제할 수 있습니다.

### 요청 예시 (bash)

```bash
curl -X DELETE http://localhost:8080/boards/1 \
-H "Authorization: Bearer {JWT_TOKEN}"
```

### 헤더

| Header        | Description                                      | Required |
| ------------- | ------------------------------------------------ | -------- |
| Authorization | `Bearer {JWT_TOKEN}` 형태로 JWT 인증 토큰을 전달 | Yes      |

### 응답 예시 (성공 시)

```json
{
  "message": "게시글이 성공적으로 삭제되었습니다."
}
```

### 응답 예시 (에러 시)

```json
{
  "error": "해당 게시글을 삭제할 권한이 없습니다."
}
```
