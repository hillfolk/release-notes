# Github 버전 관리하기
1. changelog 정책 잡기 (.github/release.yml)
2. tag 지정 및 release 하기
3. github action으로 자동화하기
4. git hook으로 자동화하기

## release.yml 
```
# .github/release.yml
changelog:
  categories:
    - title: 배포 주의 PR 🛠
      labels:
        - deploy # deploy 라벨이 있는 이슈들을 모아서 배포 주의 카테고리로 만듭니다.
    - title: 기능 추가 ✨
      labels:
        - feature
    - title: 버그 수정 🐛
      labels:
        - bug
    - title: 성능 개선 🚀
      labels:
        - performance
    - title: 빌드 시스템 📦
      labels:
        - build
    - title: 리팩토링 ♻️
      labels:
        - refactor
    - title: 테스트 🧪
      labels:
        - test
    - title: 문서 📚
      labels:
        - docs
    - title: 기타 변경사항 📄
      labels:
        - "*" # 모든 라벨
```

## github 릴리즈 기능으로 버전 관리 할때 장점
- github에서 제공하는 릴리즈 기능을 사용하면, changelog를 자동으로 만들어준다.
- 변경 사항을 라벨 별로 분류해서 볼 수 있다.
- 상위 및 하위 버전의 변경사항을 한눈에 볼 수 있다.
- 커밋 기준이 아닌 PR 기준으로 변경사항을 볼 수 있다.
- 마지막 릴리즈한 버전을 기준으로 변경사항을 정리해 준다.
- 릴리즈 수행 후 자동으로 추가 작업을 수행할 수 있다. (ex. slack 알림 등등 )

## PR시 추가 되어야할 내용 
- 이슈 번호 & 링크 
- PR 내용 요약
- 리뷰 요청 내용
- PR 관련된 이슈가 있다면, 이슈 번호 & 링크
- 배포 관련된 내용이 있다면, 배포 관련된 내용(배포 특이사항) 
- 라벨 지정 (deploy)

## 주의 사항
- PR은 지정된 라벨은 우선 순위에 따라서 릴리즈 노트에 1개만 분류되어 표시 된다.
- PR은 가능하면 작은 단위로 나누어서 작성한다.

## 배포 순서 
1. 배포할 tag 지정
2. Release 노트 자동 작성 버튼 클릭
3. pre-release 수행 및 수정
4. release 버튼 클릭
5. jenkins 또는 github action으로 배포 진행
6. 후 작업 진행

## 참고 
- https://docs.github.com/en/repositories/releasing-projects-on-github/automatically-generated-release-notes