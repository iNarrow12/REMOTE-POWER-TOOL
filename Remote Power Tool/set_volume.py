import sys
import subprocess

def set_volume(percent_str):
    try:
        pct = int(percent_str)
        pct = max(0, min(pct, 100))
    except ValueError:
        print("Invalid volume value:", percent_str)
        return

    # nircmd’s setsysvolume takes 0–65535
    max_val = 65535
    vol = int(pct * max_val / 100)
    subprocess.run(["nircmd.exe", "setsysvolume", str(vol)])

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python set_volume.py <0-100>")
    else:
        set_volume(sys.argv[1])