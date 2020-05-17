import time
import subprocess
import os
def main(args):
    name = args.get("name", "stranger")
    greeting = "Hello " + name + "!"
    time.sleep(5)
    os.chdir("/opt/intel/sgxsdk/SampleCode/SampleEnclave")
    process = subprocess.Popen(['./app'], stdout=subprocess.PIPE, stderr=subprocess.STDOUT)
    res = str(process.communicate()[0])
    greeting = res
    print(greeting)
    return {"greeting": greeting}
