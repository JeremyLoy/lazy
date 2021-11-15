# lazy
[![PkgGoDev](https://pkg.go.dev/badge/github.com/JeremyLoy/lazy)](https://pkg.go.dev/github.com/JeremyLoy/lazy)
[![GitHub issues](https://img.shields.io/github/issues/JeremyLoy/lazy.svg)](https://github.com/JeremyLoy/lazy/issues)
[![license](https://img.shields.io/github/license/JeremyLoy/lazy.svg?maxAge=2592000)](https://github.com/JeremyLoy/lazy/LICENSE)
[![Release](https://img.shields.io/github/release/JeremyLoy/lazy.svg?label=Release)](https://github.com/JeremyLoy/lazy/releases)

Package lazy is a light wrapper around sync.Once providing support for return values.
It removes the burden of capturing return values via closures from the caller.

```golang
// server.go
type server struct {
  DB func() *sql.DB
}

func (s *server) someHttpHandler(w http.ResponseWriter, r *http.Request) {
  db := s.DB()
  _ = db // use db throughout the handler
}

// db.go
func newDB() *sql.DB {
  // ommitted
}

// main.go
s := server {
  DB: lazy.Lazy(newDB),
}

s.ListenAndServe()
```
