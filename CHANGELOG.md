# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.0.5] - 2023-11-26

* Update dependencies.

## [0.0.4] - 2023-11-17

* Allow for output functions to accept `nil` as `output.Output` without rising an error (panic).

## [0.0.3] - 2023-10-15

* Introduce `output.Printer` interface with `output.OptionPrinter` to set it with and **Writer Printer** as its first
  implementation.

## [0.0.2] - 2023-10-14

* Include current `output.Output` instance in `output.PrintFunc`.

## [0.0.1] - 2023-10-10

* Release to GitHub

[0.0.5]: https://github.com/SevenOfSpades/go-aggregated-output/releases/tag/v0.0.5
[0.0.4]: https://github.com/SevenOfSpades/go-aggregated-output/releases/tag/v0.0.4
[0.0.3]: https://github.com/SevenOfSpades/go-aggregated-output/releases/tag/v0.0.3
[0.0.2]: https://github.com/SevenOfSpades/go-aggregated-output/releases/tag/v0.0.2
[0.0.1]: https://github.com/SevenOfSpades/go-aggregated-output/releases/tag/v0.0.1