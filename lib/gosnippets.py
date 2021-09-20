import argparse
import os.path
from pathlib import Path
import shutil
import sys
import re

## Minheap check -- abc164 E

def removeComments(s) :  return re.sub(r'//.*','',s)
def allwhitespace(s) :   m = re.match(r'^\s*$',s); return bool(m)
def processLine(l2) :
    numtabs = 0
    for c in l2 : 
        if c == '\t' : numtabs += 1
        else : break
    contents = re.sub(r'\s+',' ',re.sub(r'^\t*','',l2))
    firstBrace,firstBraceLoc,numOpen,numClose = '',-1,0,0
    for i,c in enumerate(l2) :
        if c == '{' : 
            if not firstBrace : firstBrace = '{'; firstBraceLoc = i
            numOpen += 1
        elif c == '}' :
            if not firstBrace : firstBrace = '}'; firstBraceLoc = i
            numClose += 1
    if numOpen == 0 and numClose == 0 : return (numtabs,"balanced",contents)
    if numOpen == numClose and firstBrace == '{': return (numtabs,"balanced",contents)
    if numOpen == 1 and numClose == 1 and firstBrace == '}': return (numtabs,"closeopen",contents)
    if numOpen == 1 and numClose == 0 : return (numtabs,"open",contents)
    if numClose == 1 and numOpen == 0 : return (numtabs,"close",contents)
    return (numtabs,"other",contents)

def canCombineA(s1,s2,linewidth) :
    if s1[1] != "balanced" or s2[1] != "balanced" : return False
    if s1[0] != s2[0] : return False
    if re.match(r'^(type |func |package )',s1[2]) : return False
    if 4 * s1[0] + len(s1[2]) + 2 + len(s2[2]) > linewidth : return False
    return True

def combineA(s1,s2) : return (s1[0],s1[1],s1[2]+"; "+s2[2])

def canCombineB(s1,s2,s3,linewidth) :
    if s1[1] != "open" or s2[1] != "balanced" or s3[1] != "close": return False
    if s1[0] != s3[0] or s1[0] != s2[0]-1 : return False
    if 4 * s1[0] + len(s1[2]) + 1 + len(s2[2]) + 1 + len(s3[2]) > linewidth : return False
    return True

def combineB(s1,s2,s3) : return (s1[0],"balanced",s1[2]+" "+s2[2]+" "+s3[2])

def canCombineC(s1,s2,s3,s4,s5,linewidth) :
    if s1[1] != "open" or s2[1] != "balanced" or s3[1] != "closeopen": return False
    if s4[1] != "balanced" or s5[1] != "close" : return False
    if s1[0] != s3[0] or s1[0] != s5[0] : return False
    if s1[0] != s2[0]-1 or s1[0] != s4[0]-1 : return False
    if 4 * s1[0] + len(s1[2]) + 1 + len(s2[2]) + 1 + len(s3[2]) + 1 + len(s4[2]) + 1 + len(s5[2]) > linewidth : return False
    return True

def combineC(s1,s2,s3,s4,s5) : return (s1[0],"balanced",s1[2]+" "+s2[2]+" "+s3[2]+" "+s4[2]+" "+s5[2])

def reformat(larr,linewidth) :
    st = []; st2 = []
    for l in larr :
        l2 = removeComments(l)
        if allwhitespace(l2) : continue
        (indentlen,type,contents) = processLine(l2)
        st.append((indentlen,type,contents))
        ## 3 types of combines : none, {open,balanced,close}, and {open,balanced,closeopen,balanced,close}
        while True :
            if len(st) >= 2 and canCombineA(st[-2],st[-1],linewidth) :
                newl = combineA(st[-2],st[-1])
                st.pop(); st.pop(); st.append(newl); continue
            if len(st) >= 3 and canCombineB(st[-3],st[-2],st[-1],linewidth) :
                newl = combineB(st[-3],st[-2],st[-1])
                st.pop(); st.pop(); st.pop(); st.append(newl); continue
            if len(st) >= 5 and canCombineC(st[-5],st[-4],st[-3],st[-2],st[-1],linewidth) :
                newl = combineC(st[-5],st[-4],st[-3],st[-2],st[-1])
                st.pop(); st.pop(); st.pop(); st.pop(); st.pop(); st.append(newl); continue
            break
    return ["\t"*x[0]+x[2] for x in st]
    
def dosubs(larr,subs) :
    def worker(l) :
        l2 = l
        for k,v in subs.items(): l2 = l2.replace(k,v)
        return l2
    return [worker(l) for l in larr]

def dooutput(l3,fn) :
    if fn == "stdout" :
        for l in l3 : print(l)
    else :
        with open(fn,'at') as fp :
            for l in l3 : print(l,file=fp)

def processRequest(fn,subs,outfn) :
    with open(fn,"rt") as fp1 :
        l = [ x.rstrip() for x in fp1 ]
        stidx = -1
        for (i,ll) in enumerate(l) :
            if "START HERE" in ll : stidx = i; break
        if stidx > 0 : l = l[stidx:]
        l2 = reformat(l,120)
        l3 = dosubs(l2,subs)
        dooutput(l3,outfn)

if __name__ == "__main__" :
    helpstring = ("gosnippets.py generates competitive programming code with user-customized classes for certain elements." +
                "USAGE:\n" +
                "    python3 gosnippets.py queue            CLASSNAME DATATYPE FILENAME\n" +
                "    python3 gosnippets.py stack            CLASSNAME DATATYPE FILENAME\n" +
                "    python3 gosnippets.py deque            CLASSNAME DATATYPE FILENAME\n" +
                "    python3 gosnippets.py minheap          CLASSNAME DATATYPE FILENAME\n" +
                "    python3 gosnippets.py segtree          CLASSNAME DATATYPE FILENAME\n" +
                "    python3 gosnippets.py lazysegtree      CLASSNAME DATATYPE FUNCTIONTYPE FILENAME\n" +
                "    python3 gosnippets.py convolver        FILENAME\n" +
                "    python3 gosnippets.py fenwick          FILENAME\n" +
                "    python3 gosnippets.py maxflow          FILENAME\n" +
                "    python3 gosnippets.py matching         FILENAME\n" +
                "    python3 gosnippets.py dsu              FILENAME\n" +
                "    python3 gosnippets.py dsusparse        FILENAME\n" +
                "    python3 gosnippets.py scc              FILENAME\n" +
                "    python3 gosnippets.py twosat           FILENAME\n" +
                "    python3 gosnippets.py bisect           FILENAME\n" +
                ##"    python3 gosnippets.py skiplistset      CLASSNAME DATATYPE FILENAME\n" +
                ##"    python3 gosnippets.py skiplistmultiset CLASSNAME DATATYPE FILENAME\n" +
                ##"    python3 gosnippets.py skiplistmap      CLASSNAME KEYTYPE VALUETYPE FILENAME\n" +
                "    python3 gosnippets.py rbtreeset        CLASSNAME KEYTYPE FILENAME\n" +
                "    python3 gosnippets.py rbtreemultiset   CLASSNAME KEYTYPE FILENAME\n" +
                "    python3 gosnippets.py rbtreemapmap     CLASSNAME KEYTYPE VALUETYPE FILENAME\n" +
                "NOTE: gosnippets.py will APPEND the results to the given filenme.\n" +
                "Use 'stdout' for filename if you just want the codesnippet on stdout.")
    dir = os.path.join(os.path.dirname(os.path.abspath(__file__)),"go") 
    good = True
    if len(sys.argv) <= 1 : good = False
    if len(sys.argv) > 1 :
        if   sys.argv[1] == "queue" and len(sys.argv) == 5 :            processRequest(os.path.join(dir,"queue","queue.go"),{"QUEUE":sys.argv[2], "DATATYPE":sys.argv[3]},sys.argv[4])
        elif sys.argv[1] == "stack" and len(sys.argv) == 5 :            processRequest(os.path.join(dir,"stack","stack.go"),{"STACK":sys.argv[2], "DATATYPE":sys.argv[3]},sys.argv[4])
        elif sys.argv[1] == "deque" and len(sys.argv) == 5 :            processRequest(os.path.join(dir,"deque","deque.go"),{"DEQUE":sys.argv[2], "DATATYPE":sys.argv[3]},sys.argv[4])
        elif sys.argv[1] == "minheap" and len(sys.argv) == 5 :          processRequest(os.path.join(dir,"minheap","minheap.go"),{"MINHEAP":sys.argv[2], "DATATYPE":sys.argv[3]},sys.argv[4])
        elif sys.argv[1] == "segtree" and len(sys.argv) == 5 :          processRequest(os.path.join(dir,"segtree","segtree.go"),{"SEGTREE":sys.argv[2], "DATATYPE":sys.argv[3]},sys.argv[4])
        elif sys.argv[1] == "lazysegtree" and len(sys.argv) == 6 :      processRequest(os.path.join(dir,"lazysegtree","lazysegtree.go"),{"LAZYSEGTREE":sys.argv[2], "DATATYPE":sys.argv[3], "FUNCTYPE":sys.argv[4]},sys.argv[5])
        elif sys.argv[1] == "convolver" and len(sys.argv) == 3 :        processRequest(os.path.join(dir,"convolver","convolver.go"),{},sys.argv[2])
        elif sys.argv[1] == "fenwick" and len(sys.argv) == 3 :          processRequest(os.path.join(dir,"fenwick","fenwick.go"),{},sys.argv[2])
        elif sys.argv[1] == "maxflow" and len(sys.argv) == 3 :          processRequest(os.path.join(dir,"maxflow","maxflow.go"),{},sys.argv[2])
        elif sys.argv[1] == "matching" and len(sys.argv) == 3 :         processRequest(os.path.join(dir,"matching","matching.go"),{},sys.argv[2])
        elif sys.argv[1] == "dsu" and len(sys.argv) == 3 :              processRequest(os.path.join(dir,"dsu","dsu.go"),{},sys.argv[2])
        elif sys.argv[1] == "dsusparse" and len(sys.argv) == 3 :        processRequest(os.path.join(dir,"dsusparse","dsusparse.go"),{},sys.argv[2])
        elif sys.argv[1] == "mincostflow" and len(sys.argv) == 3 :      processRequest(os.path.join(dir,"mincostflow","mincostflow.go"),{},sys.argv[2])
        elif sys.argv[1] == "scc" and len(sys.argv) == 3 :              processRequest(os.path.join(dir,"scc","scc.go"),{},sys.argv[2])
        elif sys.argv[1] == "twosat" and len(sys.argv) == 3 :           processRequest(os.path.join(dir,"twosat","twosat.go"),{},sys.argv[2])
        elif sys.argv[1] == "bisect" and len(sys.argv) == 3 :           processRequest(os.path.join(dir,"bisect","bisect.go"),{},sys.argv[2])
        elif sys.argv[1] == "rbtreemap" and len(sys.argv) == 6 :        processRequest(os.path.join(dir,"rbtreemap","rbtreemap.go"),{"RBTREEMAP":sys.argv[2], "KEYTYPE":sys.argv[3], "VALTYPE":sys.argv[4]},sys.argv[5])
        elif sys.argv[1] == "rbtreeset" and len(sys.argv) == 5 :        processRequest(os.path.join(dir,"rbtreeset","rbtreeset.go"),{"RBTREESET":sys.argv[2], "KEYTYPE":sys.argv[3]}, sys.argv[4])
        elif sys.argv[1] == "rbtreemultiset" and len(sys.argv) == 5 :   processRequest(os.path.join(dir,"rbtreemultiset","rbtreemultiset.go"),{"RBTREEMULTISET":sys.argv[2], "KEYTYPE":sys.argv[3]}, sys.argv[4])
        ##elif sys.argv[1] == "skiplistset" and len(sys.argv) == 5 :      s = getskiplistSet(sys.argv[2],sys.argv[3]); appendStringToFile(s,sys.argv[4])
        ##elif sys.argv[1] == "skiplistmultiset" and len(sys.argv) == 6 : s = getskiplistMultiset(sys.argv[2],sys.argv[3]); appendStringToFile(s,sys.argv[4])
        ##elif sys.argv[1] == "skiplistmap" and len(sys.argv) == 6 :      s = getskiplistMap(sys.argv[2],sys.argv[3],sys.argv[4]); appendStringToFile(s,sys.argv[5])
        else : good = False
    if not good : print(helpstring)




         
