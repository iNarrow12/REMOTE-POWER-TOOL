from flask import Flask, request, send_from_directory
import os
import subprocess
import threading
import sys

app = Flask(__name__)
PORT = 5000

@app.route('/')
def index():
    return send_from_directory('static', 'index.html')

@app.route('/action', methods=['POST'])
def handle_action():
    action = request.form.get('action')
    print(f"[+] Action received: {action}")

    if action == 'shutdown':
        subprocess.run(['shutdown', '/s', '/t', '0'])
    elif action == 'restart':
        subprocess.run(['shutdown', '/r', '/t', '0'])
    elif action == 'lock':
        subprocess.run('rundll32.exe user32.dll,LockWorkStation')
    elif action == 'logoff':
        subprocess.run(['shutdown', '/l'])
    elif action == 'sleep':
        subprocess.run('rundll32.exe powrprof.dll,SetSuspendState 0,1,0')
    elif action == 'screenoff':
        subprocess.run(['nircmd.exe', 'monitor', 'off'])
    elif action == 'set_volume':
        volume = int(request.form.get('volume'))
        subprocess.run(['nircmd.exe', 'setsysvolume', str(int(volume * 65535 / 100))])
    elif action == 'set_brightness':
        brightness = int(request.form.get('brightness'))
        subprocess.run([
            'powershell', '-Command',
            f'(Get-WmiObject -Namespace root/wmi -Class WmiMonitorBrightnessMethods).WmiSetBrightness(1,{brightness})'
        ])
    elif action == 'custom_fixed':
        script_path = os.path.join(os.getcwd(), 'custom_script.vbs')
        if os.path.exists(script_path):
            print(f"[+] Running fixed script: {script_path}")
            subprocess.run([script_path], shell=True)
        else:
            print("[!] Fixed custom script not found.")
    elif action == 'custom':
        path = request.form.get('script_path')
        if path and os.path.exists(path):
            print(f"[+] Running user script: {path}")
            subprocess.run([path], shell=True)
        else:
            print("[!] Invalid script path.")
    elif action == 'stop':
        print("[x] Stopping remote tool...")
        threading.Thread(target=graceful_exit).start()
    return '', 204

def graceful_exit():
    os._exit(0)

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=PORT)