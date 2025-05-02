' custom_script.vbs
Dim speaker, name
Set speaker = CreateObject("SAPI.SpVoice")
name = "HALLOW USER, THIS TOOL MADE BY I NARROW12. IF YOU HERE ABOUT THIS SOUND OR TEX ON YOUR SREEN. ITS MEAN THE CUSTOM SCRIPT IS WORKING, HAVE NICE DAY"

' Speak the name
speaker.Speak name

' Display the name in a message box
MsgBox "MASGE FROM INARROW12: " & name, vbInformation, "CUSTOM-SCRIPT"
