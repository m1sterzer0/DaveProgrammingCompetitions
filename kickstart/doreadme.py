import os
from pathlib import Path
import yaml

def doHeader(fp) :
    ##p1 = f"# m1sterzer0 Atcoder ABC Solutions ![Language](https://img.shields.io/badge/language-Golang-green.svg) [![License](https://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE) ![GolangProgress](https://img.shields.io/badge/GolangProgress-{goprogress}%20%2F%20{numprobs}-ff69b4.svg)"
    p1 = f"# m1sterzer0 Google Kickstart Solutions ![Language](https://img.shields.io/badge/language-Golang-green.svg) [![License](https://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE)"
    p2 = "These are my solutions for Google's Kickstart problems. Please enjoy!"
    p3 = "`DISCLAIMER`: Most of these file were created/edited after the contest, so many of these solutions were created without the time pressure of the contest and (occasionally) with the benefit of looking at the prepared editorials/solutions.  Note that run-time challenged implementations are indicated in the `Notes` column below."

    print(p1,file=fp)
    print("",file=fp)
    print(p2,file=fp)
    print("",file=fp)
    print(p3,file=fp)

def doTable(fp) :
    with open("kickstart.yaml",'rt') as fp2 :
        xx = yaml.safe_load(fp2)
        contests = xx['contests']
    print('## Contest Shortcuts\n|     |     |     |     |     |\n| --- | --- | --- | --- | --- |',file=fp)
    ptr = 0
    while (ptr < len(contests)) :
        print("|",end='',file=fp)
        for _ in range(5) :
            elem = f" [{contests[ptr]['name']}]({contests[ptr]['roundlink']}); [sol](#{contests[ptr]['name']}-Solutions) |" if ptr < len(contests) else " |"
            ptr += 1
            print(elem,end='',file=fp)
        print("\n",end='',file=fp)
    print("",file=fp)

def doSolutions(fp) :
    with open("kickstart.yaml",'rt') as fp2 :
        xx = yaml.safe_load(fp2)
        contests = xx['contests']
    for contest in contests[::-1] :
        print(f"## {contest['name']} Solutions",file=fp)
        print("| Contest | Problem | Num Correct | Solutions | Notes |",file=fp)
        print("| ------- | ------- | ----------: | --------- | ----- |",file=fp)
        for prob in contest['problems'] :
            sol = ""
            if os.path.exists(f"go/{contest['name']}/{prob['solname']}/{prob['solname']}.go") :
                sol += f" [go](./go/{contest['name']}/{prob['solname']}/{prob['solname']}.go)"
            print(f"| [{contest['name']}]({contest['roundlink']}) | [{prob['name']}]({prob['problink']}) | {prob['correct']} | {sol} | |",file=fp)
        print(f"",file=fp)

if __name__ == "__main__" :
    with open("README.md","wt") as fp :
        doHeader(fp)
        doTable(fp)
        doSolutions(fp)
    Path(f"README.md").touch()
