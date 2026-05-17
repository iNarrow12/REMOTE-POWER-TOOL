<div align="center">

![header](https://capsule-render.vercel.app/api?type=waving&height=300&text=POWERTOOL&textBg=false&fontColor=ffffff&fontAlignY=42)

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![HTML5](https://img.shields.io/badge/HTML5-Tailwind-FF69B4?style=for-the-badge&logo=tailwindcss&logoColor=white)
![Status](https://img.shields.io/badge/Status-Active-34A853?style=for-the-badge)

</div>

---

## `$ Overview`

**PowerTool** is a sleek, lightweight, and beautiful remote system control panel built in **Go**. It turns any browser into a powerful command center to control your PC — Shutdown, Restart, Volume, Brightness, Lock, and even a loud **Find My Device** siren.

> **Aesthetic**: Cyber-Neon Glassmorphism  
> **Philosophy**: Maximum control with minimum friction.

---

## `$ Screenshots`

![PowerTool Dashboard]()


---

## `$ Tree Overview`

```
.
├── main.go                    # Complete single-file backend + embedded UI
├── README.md
├── go.mod
├── go.sum
└── screenshots/

```

---

## `$ Features`

| Feature              | Description |
|----------------------|-----------|
| **Power Actions**    | Shutdown, Restart, Sleep, Lock |
| **Find Device**      | Loud siren / beep locator |
| **Volume Control**   | Real-time system volume slider |
| **Brightness Control** | Screen brightness adjustment (Windows) |
| **Modern UI**        | Glassmorphic neon cyber interface |
| **Secure API**       | Bearer Token authentication |
| **Cross Platform**   | Works on Windows and Linux |

---

## `$ Supported Platforms`

| Action             | Windows | Linux |
|--------------------|---------|-------|
| Shutdown           | ✅      | ✅    |
| Restart            | ✅      | ✅    |
| Sleep              | ✅      | ✅    |
| Lock               | ✅      | ✅    |
| Find (Siren)       | ✅      | ✅    |
| Volume Control     | ✅      | ✅    |
| Brightness         | ✅      | ❌    |

---

## `$ Configuration`

Edit the constants at the top of `main.go`:

```go
const (
    AdminToken = "your-super-secure-token"   // ← CHANGE THIS!
    Port       = ":8080"
)
```

---

## `$ Installation`

```bash
git clone https://github.com/iNarrow12/PowerTool.git
cd PowerTool
```

### Build & Run

```bash
# Run directly
go run main.go

# Or build binary
go build -o powertool main.go
./powertool
```

Open browser → `http://localhost:8080`

---

## `$ Usage`

| URL                    | Description |
|------------------------|-----------|
| `http://localhost:8080` | Main Control Panel |
| `/api/v1/power`        | Protected JSON API endpoint |

**Quick Actions Available:**
- **Ping** → Trigger loud find-my-device siren
- **Lock** → Instantly lock the screen
- **Sleep** → Put system to sleep
- **Restart** / **Shutdown** → Power commands

---

## `$ Attack Surface / Security`

- Protected with **Bearer Token** authentication
- Intended for **local network** use only
- **Do not expose directly to the internet** without reverse proxy + strong auth

---

## `$ Tech Stack`

- **Language**: Go
- **UI**: Tailwind CSS + Custom Neon Glassmorphism
- **Volume**: [volume-go](https://github.com/itchyny/volume-go)
- **Single Binary** — No dependencies after build

---

## `$ Future Plans`

- [ ] WebSocket real-time status
- [ ] Docker support
- [ ] Multi-device dashboard
- [ ] Linux brightness support (`brightnessctl`)
- [ ] Authentication UI page
- [ ] Logging & history

---

## `$ License`

MIT — Free to use for personal and educational purposes.

<div align="center">

![footer](https://capsule-render.vercel.app/api?type=waving&color=gradient&customColorList=12,20,24&height=100&section=footer)

**Made with 🔥 by iNarrow12**

</div>
