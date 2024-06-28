# CIDVerse - Go Utils / Helpers

> A collection of packages that provide utility functions for CIDVerse projects.

## Why split into multiple packages?

Some modules require external dependencies that are not needed in all projects, so we split them into separate packages to keep the dependencies to a minimum.

## Installation

```bash
go get -u github.com/cidverse/cidverseutils/ci@main
go get -u github.com/cidverse/cidverseutils/compress@main
go get -u github.com/cidverse/cidverseutils/containerruntime@main
go get -u github.com/cidverse/cidverseutils/exec@main
go get -u github.com/cidverse/cidverseutils/filesystem@main
go get -u github.com/cidverse/cidverseutils/hash@main
go get -u github.com/cidverse/cidverseutils/network@main
go get -u github.com/cidverse/cidverseutils/redact@main
go get -u github.com/cidverse/cidverseutils/version@main
go get -u github.com/cidverse/cidverseutils/zerologconfig@main
```

## License

Released under the [MIT license](./LICENSE).
