import os
import shutil

executable = "practica-upo.exe"

os.chdir("../src")
os.system("go build .")

if os.path.exists(executable):
    shutil.move(f"./{executable}", f"../build/{executable}")