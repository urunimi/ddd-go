# 도메인 정의하기

Domain package layout 예

```bash
├─- repo
│   └─ ...              // Repository 구현
├── entity.go           // 도메인에서 사용할 Entity 모음
├── repository.go       // 도메인에서 사용할 Repository interface
├── usecase.go          // 도메인에서 제공할 UseCase interface
└── usecase_test.go
```

## 도메인 UseCase 정의하기

`entity.go` 에 entity 를 추가하고 이를 사용하서 도메인에서 제공할 API 명세를 `usecase.go`에 정의한다.

```go
//entity.go
type Notice struct {
    ID        int64     `json:"id"`
    Title     string    `json:"title"`
    Content   string    `json:"content"`
    UserTypes *string   `json:"-"`
    Lang      string    `json:"-"`
    CreatedAt time.Time `json:"-"`
    UpdatedAt time.Time `json:"-"`
}

// usecase.go
type UseCase interface {
    First(title string) (*Notice, error)
    GetBy(lang string, userType string, lastID *int64) (Notices, error)
    Save(notice *Notice) error
}
```

## 도메인 Repository 정의하기

위 UseCase 를 구현하면서 Data layer 에서 필요한 정보들을 `repository.go` 에 정의하고 UseCase 의 비즈니스 로직을 구현한다.

```go
type Repository interface {
    First(title string) (*Notice, error)
    Find(query string, args ...interface{}) (Notices, error)
    Save(notice *Notice) error
}

```

## Repository 구현하기

위 `Repository` 의 구현체는 도메인 로직과 별도의 package 로 구별해서 작성한다.

```bash
├── repo
│   ├── db_model.go     // DB ORM Model 정의
│   ├── mapper.go       // DB Model 을 Entity 로 변경
│   ├── repo.go         // Repository 구현체
│   └── repo_test.go
```
