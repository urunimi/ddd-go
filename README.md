# ddd-go

Go 언어로 작성된 [DDD](https://martinfowler.com/tags/domain%20driven%20design.html) 샘플 프로젝트

### Proejct 구조

[golang-standards/project-layout](https://github.com/golang-standards/project-layout) 를 토대로 구성

| 디렉토리 | 설명 |
| - | - |
| `/cmd` | Main applications for this project. The directory name for each application should match the name of the executable you want to have (e.g., `/cmd/myapp`). |
| `/internal` | Private application and library code. |
| `/internal/app` | Presentation packages are placed here. |
| `/internal/pkg` | Domain packages are placed here. |

### [도메인 정의하기](https://github.com/urunimi/ddd-go/blob/master/internal/pkg/notice)
