# lets-go-chat

![Go](https://img.shields.io/badge/Go-1.20-blue.svg?logo=go&longCache=true&logoColor=white&style=flat-square&colorA=4c566a&colorB=5e81ac)
[![GitHub Issues](https://img.shields.io/github/issues/if1bonacci/lets-go-chat.svg?style=flat-square&colorA=4c566a&colorB=ebcb8b)](https://github.com/if1bonacci/lets-go-chat/issues)
[![GitHub Stars](https://img.shields.io/github/stars/if1bonacci/lets-go-chat.svg?style=flat-square&colorB=ebcb8b&colorA=4c566a)](https://github.com/if1bonacci/lets-go-chat/stargazers)
[![GitHub Forks](https://img.shields.io/github/forks/if1bonacci/lets-go-chat.svg?style=flat-square&colorA=4c566a&colorB=ebcb8b)](https://github.com/if1bonacci/lets-go-chat/network)

## Run

```
go run github.com/if1bonacci/lets-go-chat/cmd/chat
```

## Function

*func* [HashPassword](https://github.com/if1bonacci/lets-go-chat/blob/master/pkg/hasher/hasher.go#L9)
```go
HashPassword(password string) (string, error)
```

*func* [CheckPasswordHash](https://github.com/if1bonacci/lets-go-chat/blob/master/pkg/hasher/hasher.go#L19)
```go
CheckPasswordHash(password, hash string) bool
```