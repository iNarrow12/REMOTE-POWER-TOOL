package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/itchyny/volume-go"
)

const (
	AdminToken = "your-super-secure-token"
	Port       = ":8080"
)

func ExecuteAction(action string, value int) error {
	var cmd *exec.Cmd
	switch action {
	case "shutdown":
		if runtime.GOOS == "windows" {
			cmd = exec.Command("shutdown", "/s", "/t", "0")
		} else {
			cmd = exec.Command("shutdown", "now")
		}
	case "restart":
		if runtime.GOOS == "windows" {
			cmd = exec.Command("shutdown", "/r", "/t", "0")
		} else {
			cmd = exec.Command("reboot")
		}
	case "sleep":
		if runtime.GOOS == "windows" {
			cmd = exec.Command("rundll32.exe", "powrprof.dll,SetSuspendState", "0,1,0")
		} else {
			cmd = exec.Command("systemctl", "suspend")
		}
	case "lock":
		if runtime.GOOS == "windows" {
			cmd = exec.Command("rundll32.exe", "user32.dll,LockWorkStation")
		} else {
			cmd = exec.Command("xdg-screensaver", "lock")
		}
	case "find":
		if runtime.GOOS == "windows" {
			sirenCmd := "[Console]::Beep(1500,200); [Console]::Beep(1000,200); [Console]::Beep(1500,200); [Console]::Beep(1000,400)"
			cmd = exec.Command("powershell", "-Command", sirenCmd)
		} else {
			fmt.Print("\a\a\a")
			os.Stdout.Sync()
			cmd = exec.Command("beep", "-f", "2000", "-l", "300")
			if err := cmd.Run(); err != nil {
				cmd = exec.Command("aplay", "/usr/share/sounds/alsa/Front_Center.wav")
			}
			return nil
		}
	case "volume_set":
		return volume.SetVolume(value)
	case "brightness_set":
		if runtime.GOOS == "windows" {
			psCmd := fmt.Sprintf("(Get-WmiObject -Namespace root/WMI -Class WmiMonitorBrightnessMethods).WmiSetBrightness(1, %d)", value)
			cmd = exec.Command("powershell", "-Command", psCmd)
		} else {
			return fmt.Errorf("brightness only supported on windows")
		}
	default:
		return fmt.Errorf("unsupported: %s", action)
	}
	if cmd != nil {
		return cmd.Run()
	}
	return nil
}

const UI_HTML = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no, viewport-fit=cover">
    <title>Node Controller</title>
    <script src="https://cdn.tailwindcss.com"></script>
    <link href="https://fonts.googleapis.com/css2?family=Inter:wght@300;400;600;800&family=Fira+Code:wght@400;600&display=swap" rel="stylesheet">
    <style>
        :root {
            --bg-color: #050505;
            --glass-bg: rgba(20, 20, 25, 0.4);
            --glass-border: rgba(255, 255, 255, 0.08);
            --neon-blue: #00f0ff;
            --neon-red: #ff003c;
            --neon-green: #00ff66;
            --neon-purple: #b000ff;
        }

        body { 
            background-color: var(--bg-color); 
            background-image: 
                radial-gradient(circle at 15% 50%, rgba(176, 0, 255, 0.08), transparent 25%),
                radial-gradient(circle at 85% 30%, rgba(0, 240, 255, 0.08), transparent 25%);
            background-attachment: fixed;
            color: #e0e0e0;
            font-family: 'Inter', sans-serif;
            height: 100dvh;
            display: flex;
            flex-direction: column;
            overflow: hidden;
            animation: bgShift 15s ease-in-out infinite alternate;
        }

        @keyframes bgShift {
            0% { background-position: 0% 0%; }
            100% { background-position: 100% 100%; }
        }

        .glass-panel {
            background: var(--glass-bg);
            backdrop-filter: blur(16px);
            -webkit-backdrop-filter: blur(16px);
            border: 1px solid var(--glass-border);
            box-shadow: 0 4px 30px rgba(0, 0, 0, 0.5);
        }

        .sys-card { 
            transition: all 0.3s cubic-bezier(0.25, 0.8, 0.25, 1);
            position: relative;
            overflow: hidden;
        }
        
        .sys-card::before {
            content: '';
            position: absolute;
            top: 0; left: -100%;
            width: 50%; height: 100%;
            background: linear-gradient(to right, transparent, rgba(255,255,255,0.03), transparent);
            transform: skewX(-20deg);
            transition: 0.5s;
        }

        .sys-card:hover::before {
            left: 150%;
        }

        .sys-card:hover {
            transform: translateY(-4px);
            box-shadow: 0 10px 40px rgba(0, 0, 0, 0.6);
            border-color: rgba(255, 255, 255, 0.15);
        }

        .btn-press:active { 
            transform: scale(0.95); 
        }

        .btn-danger {
            background: linear-gradient(135deg, rgba(255, 0, 60, 0.1), rgba(0,0,0,0));
            border: 1px solid rgba(255, 0, 60, 0.3);
            box-shadow: 0 0 20px rgba(255, 0, 60, 0.1);
        }

        .btn-danger:hover {
            background: linear-gradient(135deg, rgba(255, 0, 60, 0.2), rgba(0,0,0,0));
            border-color: rgba(255, 0, 60, 0.6);
            box-shadow: 0 0 30px rgba(255, 0, 60, 0.2);
        }

        .btn-danger:active { 
            transform: scale(0.97); 
            background: rgba(255, 0, 60, 0.3);
        }

        .terminal {
            font-family: 'Fira Code', monospace;
            background: rgba(0, 0, 0, 0.6);
            border-top: 1px solid var(--glass-border);
        }
        
        .glow-text {
            text-shadow: 0 0 10px currentColor;
        }
        
        /* Custom Sliders */
        input[type=range] {
            -webkit-appearance: none;
            width: 100%;
            background: transparent;
        }
        input[type=range]::-webkit-slider-thumb {
            -webkit-appearance: none;
            height: 16px;
            width: 16px;
            border-radius: 50%;
            background: #fff;
            cursor: pointer;
            margin-top: -6px;
            box-shadow: 0 0 10px rgba(255,255,255,0.8);
            transition: transform 0.1s;
        }
        input[type=range]:active::-webkit-slider-thumb {
            transform: scale(1.3);
        }
        input[type=range]::-webkit-slider-runnable-track {
            width: 100%;
            height: 4px;
            cursor: pointer;
            background: rgba(255, 255, 255, 0.1);
            border-radius: 2px;
        }
        #vol-slider::-webkit-slider-thumb { background: #b000ff; box-shadow: 0 0 10px #b000ff; }
        #vol-slider::-webkit-slider-runnable-track { background: rgba(176, 0, 255, 0.2); }
        #bri-slider::-webkit-slider-thumb { background: #00f0ff; box-shadow: 0 0 10px #00f0ff; }
        #bri-slider::-webkit-slider-runnable-track { background: rgba(0, 240, 255, 0.2); }
    </style>
</head>
<body class="p-4 pb-6 space-y-4 h-screen flex flex-col justify-between"> 
    <div class="flex justify-between items-end pb-2 relative z-10 shrink-0">
        <div>
            <div class="flex items-center gap-2 mb-1">
                <div class="w-2.5 h-2.5 rounded-full bg-[#00ff66] shadow-[0_0_15px_rgba(0,255,102,0.8)] animate-pulse"></div>
                <span class="text-[10px] font-bold text-[#00ff66] uppercase tracking-[0.3em] glow-text">Device Online</span>
            </div>
            <h1 class="text-white text-3xl font-extrabold tracking-tighter uppercase drop-shadow-md">POWER<span class="text-zinc-500 font-light">_TOOL</span></h1>
        </div>
        <div class="text-right flex flex-col items-end pt-1">
            <span class="text-[9px] text-zinc-500 uppercase tracking-widest mb-1">Protocol</span>
            <span class="text-xs text-zinc-300 font-bold tracking-wider">WSS://</span>
        </div>
    </div>

    <div class="grid grid-cols-2 gap-3 grow content-start relative z-10">
        
        <button onclick="trigger('find')" class="glass-panel sys-card btn-press group rounded-2xl py-6 flex flex-col items-center justify-center gap-3">
            <div class="p-3 rounded-full bg-cyan-500/10 text-cyan-400 group-hover:bg-cyan-500/20 group-hover:scale-110 transition-all duration-300">
                <svg class="w-7 h-7 group-active:text-cyan-300 drop-shadow-[0_0_8px_rgba(34,211,238,0.5)]" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M8.288 15.038a5.25 5.25 0 0 1 7.424 0M5.106 11.856c3.807-3.808 9.98-3.808 13.788 0M1.924 8.674c5.565-5.565 14.587-5.565 20.152 0M12.53 18.22l-.53.53-.53-.53a.75.75 0 0 1 1.06 0Z" />
                </svg>
            </div>
            <span class="text-[10px] font-bold tracking-[0.2em] text-zinc-300 uppercase group-hover:text-cyan-100 transition-colors">Ping</span>
        </button>

        <button onclick="trigger('lock')" class="glass-panel sys-card btn-press group rounded-2xl py-6 flex flex-col items-center justify-center gap-3">
            <div class="p-3 rounded-full bg-cyan-500/10 text-cyan-400 group-hover:bg-cyan-500/20 group-hover:scale-110 transition-all duration-300">
                <svg class="w-7 h-7 group-active:text-cyan-300 drop-shadow-[0_0_8px_rgba(34,211,238,0.5)]" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M16.5 10.5V6.75a4.5 4.5 0 1 0-9 0v3.75m-.75 11.25h10.5a2.25 2.25 0 0 0 2.25-2.25v-6.75a2.25 2.25 0 0 0-2.25-2.25H6.75a2.25 2.25 0 0 0-2.25 2.25v6.75a2.25 2.25 0 0 0 2.25 2.25Z" />
                </svg>
            </div>
            <span class="text-[10px] font-bold tracking-[0.2em] text-zinc-300 uppercase group-hover:text-cyan-100 transition-colors">Lock</span>
        </button>

        <button onclick="trigger('sleep')" class="glass-panel sys-card btn-press group rounded-2xl py-6 flex flex-col items-center justify-center gap-3">
            <div class="p-3 rounded-full bg-cyan-500/10 text-cyan-400 group-hover:bg-cyan-500/20 group-hover:scale-110 transition-all duration-300">
                <svg class="w-7 h-7 group-active:text-cyan-300 drop-shadow-[0_0_8px_rgba(34,211,238,0.5)]" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M21.752 15.002A9.72 9.72 0 0 1 18 15.75c-5.385 0-9.75-4.365-9.75-9.75 0-1.33.266-2.597.748-3.752A9.753 9.753 0 0 0 3 11.25C3 16.635 7.365 21 12.75 21a9.753 9.753 0 0 0 9.002-5.998Z" />
                </svg>
            </div>
            <span class="text-[10px] font-bold tracking-[0.2em] text-zinc-300 uppercase group-hover:text-cyan-100 transition-colors">Sleep</span>
        </button>

        <button onclick="trigger('restart')" class="glass-panel sys-card btn-press group rounded-2xl py-6 flex flex-col items-center justify-center gap-3">
            <div class="p-3 rounded-full bg-cyan-500/10 text-cyan-400 group-hover:bg-cyan-500/20 group-hover:scale-110 transition-all duration-300">
                <svg class="w-7 h-7 group-active:text-cyan-300 drop-shadow-[0_0_8px_rgba(34,211,238,0.5)]" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M16.023 9.348h4.992v-.001M2.985 19.644v-4.992m0 0h4.992m-4.993 0 3.181 3.183a8.25 8.25 0 0 0 13.803-3.7M4.031 9.865a8.25 8.25 0 0 1 13.803-3.7l3.181 3.182m0-4.991v4.99" />
                </svg>
            </div>
            <span class="text-[10px] font-bold tracking-[0.2em] text-zinc-300 uppercase group-hover:text-cyan-100 transition-colors">Restart</span>
        </button>

        <div class="glass-panel col-span-2 rounded-2xl p-5 flex flex-col gap-4 mt-1">
            <div class="flex items-center gap-4">
                <svg class="w-5 h-5 text-[#b000ff]" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" d="M19.114 5.636a9 9 0 010 12.728M16.463 8.288a5.25 5.25 0 010 7.424M6.75 8.25l4.72-4.72a.75.75 0 011.28.53v15.88a.75.75 0 01-1.28.53l-4.72-4.72H4.51c-.88 0-1.704-.507-1.938-1.354A9.01 9.01 0 012.25 12c0-.83.112-1.633.322-2.396C2.806 8.756 3.63 8.25 4.51 8.25H6.75z" /></svg>
                <input type="range" id="vol-slider" min="0" max="100" value="50" onchange="triggerVal('volume_set', this.value)">
            </div>
            <div class="flex items-center gap-4">
                <svg class="w-5 h-5 text-[#00f0ff]" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" d="M12 3v2.25m6.364.386l-1.591 1.591M21 12h-2.25m-.386 6.364l-1.591-1.591M12 18.75V21m-4.773-4.227l-1.591 1.591M5.25 12H3m4.227-4.773L5.636 5.636M15.75 12a3.75 3.75 0 11-7.5 0 3.75 3.75 0 017.5 0z" /></svg>
                <input type="range" id="bri-slider" min="0" max="100" value="50" onchange="triggerVal('brightness_set', this.value)">
            </div>
        </div>

        <button onclick="trigger('shutdown')" class="glass-panel sys-card btn-danger col-span-2 rounded-2xl mt-1 flex items-center justify-between p-6 group shrink-0">
            <div class="flex flex-col items-start gap-1">
                <span class="text-[9px] font-bold text-red-500/80 uppercase tracking-[0.4em] group-hover:text-red-400 transition-colors">System Halt</span>
                <span class="text-xl font-black tracking-widest uppercase text-white drop-shadow-[0_0_10px_rgba(255,0,60,0.8)]">Shutdown</span>
            </div>
            <div class="flex items-center justify-center p-3 rounded-full bg-red-500/10 group-hover:bg-red-500/20 group-hover:scale-110 transition-all duration-300">
                <svg class="w-8 h-8 text-red-500 drop-shadow-[0_0_8px_rgba(255,0,60,0.8)]" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M5.636 5.636a9 9 0 1 0 12.728 0M12 3v9" />
                </svg>
            </div>
        </button>
    </div>

    <script>
        async function trigger(action) {
            triggerVal(action, 0);
        }

        async function triggerVal(action, value) {
            try {
                await fetch('/api/v1/power', {
                    method: 'POST',
                    headers: { 
                        'Content-Type': 'application/json',
                        'Authorization': 'Bearer your-super-secure-token' 
                    },
                    body: JSON.stringify({ action: action, value: parseInt(value) })
                });
            } catch (e) {
                console.error("Connection Dropped", e);
            }
        }
    </script>
</body>
</html>
`

func main() {
	http.HandleFunc("/api/v1/power", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			return
		}
		if r.Header.Get("Authorization") != "Bearer "+AdminToken {
			http.Error(w, "Locked", 401)
			return
		}

		var req struct {
			Action string `json:"action"`
			Value  int    `json:"value"`
		}
		json.NewDecoder(r.Body).Decode(&req)

		ExecuteAction(strings.ToLower(req.Action), req.Value)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, UI_HTML)
	})

	fmt.Println("PowerTool Engine v3.4: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
