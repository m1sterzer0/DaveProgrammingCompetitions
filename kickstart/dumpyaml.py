from bs4 import BeautifulSoup
import re
import os
import sys

def parse(contest) :
    if not os.path.exists(f'html/{contest}.html') :
        raise Exception (f"Could not find the html for contest:{contest}")
    with open(f"html/{contest}.html") as fp :
        myhtml = fp.read()
    soup = BeautifulSoup(myhtml,'html.parser')
    a1 = soup.find('section','problems-nav')
    a2 = a1.find_all('div',re.compile(r'problems-nav-selector-item-container section-row-column.*'))
    probnames = [x.find('p').getText().rstrip().lstrip() for x in a2]
    links = [x.find('a','body-5 problems-nav-problem-link').attrs['href'] for x in a2]
    xx = soup.find('p','ranking-table-page-number-total-pages').getText()
    idx = 0
    while idx < len(xx) and xx[idx] not in '123456789' : idx += 1
    numpeople = 0
    if idx < len(xx) : numpeople = 50 * int(xx[idx:])
    return probnames,links,numpeople

def printyaml(contest,problems,links,participants) :
    roundlink = "https://codingcompetitions.withgoogle.com/kickstart"
    m = re.match(r'.*.kickstart.round.[^/]*',links[0])
    if m : roundlink = m.group(0)
    print(f'- name: "{contest}"')
    print(f'  roundlink: "{roundlink}"')
    print(f'  participants: {participants}')
    print(f'  problems :')
    for (p,l) in zip(problems,links) :
        solname = "".join(x for x in p if x in "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
        print(f'  - name : "{p}"')
        print(f'    problink : "{l}"')
        print(f'    correct : 0')
        print(f'    solname : "{solname}"')

if __name__ == "__main__" :
    ## Do command line parsing
    if len(sys.argv) <= 1 or len(sys.argv) > 2 : raise Exception("Wrong number of args")
    contest = sys.argv[1]
    problems,links,participants = parse(contest)
    printyaml(contest,problems,links,participants)



