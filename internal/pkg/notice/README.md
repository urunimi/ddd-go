# 도메인 정의하기

```bash
├── entity.go   // 도메인에서 사용할 Entity 모음 
├── repo
│   ├── db_model.go // DB ORM Model 정의
│   ├── mapper.go   // DB Model 을 Entity 로 변경
│   ├── repo.go     // Repository 구현체
│   └── repo_test.go
├── repository.go   // 도메인에서 사용할 Repository interface
├── usecase.go      // 도메인에서 제공할 UseCase interface
└── usecase_test.go
```
