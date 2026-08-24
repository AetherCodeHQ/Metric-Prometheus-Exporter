# 🌐 Metric Prometheus Exporter

![Go](https://img.shields.io/badge/Go-1.21%2B-00ADD8?style=flat-square&logo=go&logoColor=white)
![Version](https://img.shields.io/badge/Version-v2.1.0-00ADD8?style=flat-square)
![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)
![PRs](https://img.shields.io/badge/PRs-Welcome-brightgreen?style=flat-square)

> Infrastructure tool by [AetherCodeHQ](https://github.com/AetherCodeHQ)

`infrastructure` `devops` `cli` `golang` `http`

---

## What is Metric-Prometheus-Exporter?

**Metric-Prometheus-Exporter** is an infrastructure tool for monitoring, inspecting, and managing systems and services.

## Features

- ✅ Network operations
- 🚀 **Zero dependencies** — only Go standard library
- 📦 **Single binary** — compile and run anywhere
- 🔄 **Offline capable** — no internet required

## Installation

```bash
# Clone
git clone https://github.com/AetherCodeHQ/Metric-Prometheus-Exporter.git
cd Metric-Prometheus-Exporter

# Build
go build -o metric-prometheus-exporter .

# Run
./metric-prometheus-exporter [file]
```

### Or directly with `go run`:
```bash
go run main.go [file]
```

## Usage

```bash
# Basic usage
./metric-prometheus-exporter [file]
```

### Example Output

```
$ ./metric-prometheus-exporter [file]
prometheus exporter on :%s/metrics\n
```

## Project Structure

```
Metric-Prometheus-Exporter/
  main.go          # Entry point (37 lines)
  go.mod            # Go module definition
  go.sum            # Dependency checksums
  README.md         # This file
  LICENSE           # MIT License
  CHANGELOG.md      # Version history
```

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

MIT License - see [LICENSE](LICENSE) for details.

---

Built with ❤️ by [AetherCodeHQ](https://github.com/AetherCodeHQ)
