# 🔌 REMOTE POWER TOOL - BY INARROW12 (FULL GUIDE)

🔌 REMOTE POWER TOOL - BY INARROW12 (FULL GUIDE)

Hi, I'm inarrow12 – the creator of this tool.

This guide will help you set up and use the REMOTE POWER TOOL to control your Windows PC from your mobile browser.

-------------------------------------------------------------
📌 What Does This Tool Do?
-------------------------------------------------------------
This tool allows you to:

✔️ Remotely Restart your PC

✔️ Remotely Shutdown your PC

✔️ Remotely Turn Off your pc Screen

✔️ Remotely Lock your pc

✔️ Remotely Run Custom Script On your pc

✔️ Remotely sleep Your pc

✔️ Remotely Log Off your pc

❌ This is NOT full remote access (you can’t see or control screen)

⭐ This tool Starts automatically on system boot(So You Didn't Wan To Run The Tool Agin)

-------------------------------------------------------------
🧰 Requirements Before You Start
-------------------------------------------------------------
1. A Windows PC
2. A mobile phone on the SAME Wi-Fi network
3. Python (latest version) installed
4. Flask installed using pip

-------------------------------------------------------------
🧪 STEP-BY-STEP SETUP INSTRUCTIONS
-------------------------------------------------------------

🟢 STEP 1: Install Python (if not already installed)
---------------------------------------------------
1. Go to: https://www.python.org/downloads/
2. Download the latest version for Windows
3. Install it and MAKE SURE to check the box:
   👉 “Add Python to PATH”
4. Complete the installation.

🟢 STEP 2: Install Flask
-------------------------------
1. Open Command Prompt or Terminal
2. Type the following command:

   pip install flask

3. Hit ENTER and wait for installation to complete.

🟢 STEP 3: Prepare the Tool Folder
-----------------------------------------
1. Copy the entire folder named: `Remote Power Tool` to your DESKTOP
2. Also copy the `RESTART.bat` file and place it inside that same folder

🟢 STEP 4: Run Setup
-------------------------------
1. Open the `Remote Power Tool` folder on your Desktop
2. Double-click on `setup.bat`
3. Wait for it to complete setup

🟢 STEP 5: Start the Flask Server
---------------------------------
1. Double-click on the file `RESTART.bat`
2. A terminal window will open and Flask will start the local web server

-------------------------------------------------------------
📱 HOW TO CONTROL FROM YOUR PHONE
-------------------------------------------------------------

1. On your PC, check your IP address:
   - Open Command Prompt
   - Type: ipconfig
   - Look for: `IPv4 Address` (e.g., 192.168.1.12)

2. On your phone, open a web browser (Chrome, Firefox, etc.)

3. Type the following in the address bar:

   http://(YOUR-IP):5000

   For example:
   http://192.168.1.12:5000

4. You’ll see the Remote Power Tool interface.

⚠️ NOTE:
• Your phone must be connected to the same Wi-Fi network as your PC.
• The default port is: 5000

-------------------------------------------------------------
🛠️ TROUBLESHOOTING
-------------------------------------------------------------

❌ Flask Not Found?
✔️ Make sure you typed `pip install flask` correctly and that Python is installed.

❌ Website not loading on phone?
✔️ Check your IP address again
✔️ Ensure port 5000 is not blocked by firewall
✔️ Try turning off Windows Firewall temporarily (for testing)

-------------------------------------------------------------
📌 ABOUT
-------------------------------------------------------------
• Made with ❤️ using Python + Flask
• Lightweight and beginner-friendly
• For personal, educational use only

-------------------------------------------------------------
✅ You're now ready to use REMOTE POWER TOOL!
-------------------------------------------------------------
   Enjoy the power – from your pocket 😉
   – inarrow12
===================================================================
