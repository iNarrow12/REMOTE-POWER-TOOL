<h1 align="center">
  <br>
  <img src="https://img.shields.io/badge/POWER__TOOL-Engine%20v3.4-00f0ff?style=for-the-badge&logo=go&logoColor=white&labelColor=050505" alt="PowerTool">
</h1>

<p align="center">
  <a href="#-features">Features</a> •
  <a href="#-tech-stack">Tech Stack</a> •
  <a href="#-api-endpoints">API Endpoints</a> •
  <a href="#-installation--usage">Installation</a> •
  <a href="#-security">Security</a>
</p>

---

A lightweight, high-performance cross-platform system utility written in Go. It provisions a sleek, mobile-optimized cyberpunk web GUI to remotely monitor, find, and execute hardware power functions on your machine.

<p align="center">
  <img src="https://img.shields.io/badge/Platform-Windows%20%7C%20Linux-lightgrey?style=flat-square" alt="Platform Supported">
  <img src="https://img.shields.io/badge/Go-1.18+-00ADD8?style=flat-square&logo=go&logoColor=white" alt="Go Version">
  <img src="https://img.shields.io/badge/Tailwind_CSS-3.0+-38B2AC?style=flat-square&logo=tailwind-css&logoColor=white" alt="Tailwind Version">
</p>

## ⚡ Features

- **🎯 Device Finder:** High-frequency, aggressive high-low siren sequencer (PowerShell) or hardware alert bell/audio fallbacks (`aplay`/`beep`) on Linux to find your misplaced machine instantly.
- **🖥️ Hardware Sliders:** Smooth inline adjustments for target operating system volume and screen brightness metrics.
- **🔒 Secure Architecture:** Bearer authorization tokens natively validating downstream layout manipulation payloads.
- **🔋 Power State Control:** Low-latency triggers for `Shutdown`, `Restart`, `Sleep`, and desktop console environment locking state sessions.
- **📱 Responsive UI:** Micro-engineered, auto-scaling `100dvh` flex-grid layout intentionally tailored to match viewport dimensions perfectly without scrolling container overflow.

## 🛠️ Tech Stack

- **Backend core logic:** [Go](https://go.dev/) (Native `net/http` multiplexer context handlers)
- **Frontend styling components:** HTML5, [Tailwind CSS CDN](https://tailwindcss.com), Custom Glassmorphism Panels
- **OS Interfacing bindings:** [volume-go](https://github.com/itchyny/volume-go) framework wrapper, Windows Win32 API extensions via PowerShell wrappers

## 🛣️ API Endpoints

All downstream executions are bound under token protection patterns:

| Method | Route | Authorization Header | Payload Definition (`JSON`) | Description |
| :--- | :--- | :--- | :--- | :--- |
| **POST** | `/api/v1/power` | `Bearer <your-super-secure-token>` | `{"action": "find"\|"shutdown"\|"volume_set", "value": 0}` | Fires underlying host automation scripts. |
| **GET** | `/` | *None* | *None* | Generates the responsive web UI asset panel. |

## 🚀 Installation & Usage

### Prerequisites
Make sure Go is installed and configured along with necessary platform libraries:
```bash
# Clone or move into your project development environment
cd power_tool

# Ensure required libraries are resolved
go mod download
