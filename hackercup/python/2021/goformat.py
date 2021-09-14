import re
import sys

def removeComments(s) :
    return re.sub(r'//.*','',s)
def allwhitespace(s) :
    m = re.match(r'^\s*$',s)
    return bool(m)
def processLine(l2) :
    numtabs = 0
    for c in l2 : 
        if c == '\t' : numtabs += 1
        else : break
    contents = re.sub(r'\s+',' ',re.sub(r'^\t*','',l2))
    return (numtabs,contents)
def isRoom2(s1,s2,linelen) :
    return 4*s1[0] + len(s1[1]) + 2 + len(s2[1]) <= linelen
def combine2(s1,s2) :
    return (s1[0],s1[1] + '; ' + s2[1])
def isRoom3(s1,s2,s3,linelen) :
    return 4*s1[0] + len(s1[1]) + 1 + len(s2[1]) + 1 + len(s3[1]) <= linelen
def combine3(s1,s2,s3) :
    return (s1[0], s1[1] + ' ' + s2[1] + ' ' + s3[1])
def clearStack(st) :
    for xx in st :
        ss = '\t' * xx[0] + xx[1]
        print(ss)
    st.clear()

def reformat(infile,linelen) :
    with open(infile,"rt") as fp1 :
        st = []
        for l in fp1 :
            l2 = l.rstrip()
            l2 = removeComments(l2)
            if allwhitespace(l2) : continue
            (indentlen,contents) = processLine(l2)
            st.append((indentlen,contents))

            ## TODO: deal with import mess.
            ## TODO: think about rules for combining if else if else all in one line
            while(True) :
                if len(st) == 1 : break
                if st[-1][0] > st[-2][0] : break
                if st[-1][0] == st[-2][0] :
                    if re.match(r'^(package |func |type )',st[-1][1]) :
                        xx = st.pop(); clearStack(st); st.append(xx); break
                    #if re.match(r'^(package |func |type |if |for )',st[-2][1]) and st[-2][1].count('}') == st[-2][1].count('{') :
                    #    xx = st.pop(); clearStack(st); st.append(xx); break
                    #if st[-2][1].count('}') > st[-2][1].count('{') :
                    #    xx = st.pop(); clearStack(st); st.append(xx); break
                    if isRoom2(st[-2],st[-1],linelen) :
                        res = combine2(st[-2],st[-1])
                        st.pop(); st.pop(); st.append(res); continue
                    else :
                        xx = st.pop(); clearStack(st); st.append(xx); break
                else :
                    if len(st) >= 3 and st[-1][0] == st[-3][0] and st[-2][0] > st[-1][0] :
                        if isRoom3(st[-3],st[-2],st[-1],linelen) :
                            res = combine3(st[-3],st[-2],st[-1])
                            st.pop(); st.pop(); st.pop(); st.append(res); continue
                        else : xx = st.pop(); clearStack(st); st.append(xx); break
                    else :
                        xx = st.pop(); clearStack(st); st.append(xx); break
        clearStack(st)


if __name__ == "__main__" :
    reformat('C:\\Users\\debr5\\OneDrive\\Documents\\gomod\\oc\\rbtreemap.go',120)
    helpstring = "goformat.py compresses go code for use in competitive programming.\nUsage:\ngoformat.py <filename>"
    good = True
    if len(sys.argv) == 2 : reformat(sys.argv[1],120) 
    else : print(helpstring)
