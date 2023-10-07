# go-rwfs

> Golang read-write filesystem interfaces

![Status](https://img.shields.io/badge/status-active-success.svg)
[![GitHub Issues](https://img.shields.io/github/issues/wwmoraes/go-rwfs.svg)](https://github.com/wwmoraes/go-rwfs/issues)
[![GitHub Pull Requests](https://img.shields.io/github/issues-pr/wwmoraes/go-rwfs.svg)](https://github.com/wwmoraes/go-rwfs/pulls)

[![pre-commit.ci status](https://results.pre-commit.ci/badge/github/wwmoraes/go-rwfs/master.svg)](https://results.pre-commit.ci/latest/github/wwmoraes/go-rwfs/master)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=wwmoraes_go-rwfs&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=wwmoraes_go-rwfs)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=wwmoraes_go-rwfs&metric=coverage)](https://sonarcloud.io/summary/new_code?id=wwmoraes_go-rwfs)

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](/LICENSE)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fwwmoraes%2Fgo-rwfs.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fwwmoraes%2Fgo-rwfs?ref=badge_shield)
[![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/0/badge)](https://bestpractices.coreinfrastructure.org/projects/0)

---

## 📝 Table of Contents

- [About](#-about)
- [Getting Started](#-getting-started)
- [Usage](#-usage)
- [Built Using](#-built-using)
- [TODO](./TODO.md)
- [Contributing](./CONTRIBUTING.md)
- [Authors](#-authors)
- [Acknowledgments](#-acknowledgements)

## 🧐 About

Wrapper around [`io.fs`][std-fs] interfaces to support read-write operations.

Its fully compatible with the `fs` standard package to avoid duplicating
implementations.

### Why?

Golang provides a read-only filesystem and numerous IO abstractions. That's
great to read files without mangling your code with OS details. For some reason
the designers decided to stop there, and provided no read-write interface. If
you need to write files, then you need to fallback and use concrete types such
as `os.File`, which hardcodes the requirement to use a OS-level filesystem.

`go-rwfs` provides:

- a `FS` interface based on `fs.FS` with the extra method `OpenFile`
- a `File` interface based on `fs.File` with writer interfaces from `io`

This means those interfaces are a drop-in replacement for any use-cases of the
`fs` package where you now need write access as well.

## 🏁 Getting Started

Fetch the package:

```shell
go get github.com/wwmoraes/go-rwfs
```

Now you're good to _Go_ 😉

## 🔧 Running the tests

Clone the repository then use `make coverage` to run both unit and integration
tests.

## 🎈 Usage

The package comes with a concrete implementation for the OS filesystem, similar
to how the standard Golang distribution provides `os.DirFS`:

```go
package main

import "github.com/wwmoraes/go-rwfs"

// for brevity's sake there's no error checking, please forgive me ;)
func main() {
  // create a sample folder and file with plain OS methods
  os.Mkdir("foo", 0750)

  osFile, _ := os.Create("foo/bar.txt")
  osFile.Close()

  // create a RWFS on the new folder
  fsys := rwfs.OSDirFS("foo")

  // read-only usage, same as with fs.FS
  entries, _ := fs.ReadDir(fsys, "/")

  for _, entry := range entries {
    fmt.Println("file found:", entry.Name())
  }

  // read-write usage
  fd, _ := fsys.OpenFile("bar.txt", os.O_WRONLY|os.O_TRUNC, 0640)
  defer fd.Close()

  fd.WriteString("hello from rwfs!")
}
```

## 🔧 Built Using

- [Golang](https://go.dev) - Base language

## 🧑‍💻 Authors

- [@wwmoraes](https://github.com/wwmoraes) - Idea & Initial work

## 🎉 Acknowledgements

- Golang team for providing the [`fs`][std-fs] and [`io`][std-io] interfaces

[std-fs]: https://pkg.go.dev/io/fs
[std-io]: https://pkg.go.dev/io
