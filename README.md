# Metric Prometheus Exporter

![CI](https://github.com/Qyroxen/Metric-Prometheus-Exporter/actions/workflows/ci.yml/badge.svg)
![CodeQL](https://github.com/Qyroxen/Metric-Prometheus-Exporter/actions/workflows/codeql.yml/badge.svg)
![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/License-MIT-yellow.svg)
![Stars](https://img.shields.io/github/stars/Qyroxen/Metric-Prometheus-Exporter?style=social)
![Issues](https://img.shields.io/github/issues/Qyroxen/Metric-Prometheus-Exporter)
![PRs](https://img.shields.io/github/issues-pr/Qyroxen/Metric-Prometheus-Exporter)

> A production-ready CLI tool built with Go

[![Star Badge](https://img.shields.io/github/stars/Qyroxen/Metric-Prometheus-Exporter?style=social)](https://github.com/Qyroxen/Metric-Prometheus-Exporter/stargazers)

## What is it?

Metric Prometheus Exporter is a production-ready CLI tool built with Go. It provides powerful functionality with a beautiful terminal interface.

## Features

- Fast and efficient (written in Go)
- Beautiful CLI with colored output
- Comprehensive documentation
- GitHub Actions CI/CD
- CodeQL security analysis
- Dependabot for dependency updates
- MIT Licensed
- Fully offline - zero cloud dependency

## Quick Start

```bash
# Install
git clone https://github.com/Qyroxen/Metric-Prometheus-Exporter.git
cd Metric-Prometheus-Exporter
go build -o metricprometheusexporter .

# Run
./metricprometheusexporter --help
```

## CLI Usage

```bash
# Basic usage
./metricprometheusexporter

# With flags
./metricprometheusexporter --verbose --output json

# Get help
./metricprometheusexporter --help
```

## Examples

```bash
# Example 1
./metricprometheusexporter example1

# Example 2
./metricprometheusexporter example2 --flag value
```

## Development

```bash
# Run tests
go test ./...

# Build
go build -o metricprometheusexporter .

# Lint
golangci-lint run

# Security scan
codeql analyze
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## Security

For security vulnerabilities, please see [SECURITY.md](SECURITY.md).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<p align="center">
  <a href="https://github.com/Qyroxen/Metric-Prometheus-Exporter/stargazers">
    <img src="https://img.shields.io/github/stars/Qyroxen/Metric-Prometheus-Exporter?style=social" alt="Star this repo">
  </a>
  <a href="https://github.com/Qyroxen/Metric-Prometheus-Exporter/forks">
    <img src="https://img.shields.io/github/forks/Qyroxen/Metric-Prometheus-Exporter?style=social" alt="Fork this repo">
  </a>
  <a href="https://github.com/Qyroxen/Metric-Prometheus-Exporter/issues">
    <img src="https://img.shields.io/github/issues/Qyroxen/Metric-Prometheus-Exporter" alt="Issues">
  </a>
  <a href="https://github.com/Qyroxen/Metric-Prometheus-Exporter/pulls">
    <img src="https://img.shields.io/github/issues-pr/Qyroxen/Metric-Prometheus-Exporter" alt="Pull Requests">
  </a>
</p>
