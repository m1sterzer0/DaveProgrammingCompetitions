
import random
import sys
import collections


class maxHeap :
    v = []
    def __init__(self) : self.v = [0]
    def len(self) : return len(self.v)-1
    def isempty(self) : return len(self.v) == 1
    def top(self) : return self.v[1]
    def push(self,val) :
        self.v.append(val)
        self._bubbleup(len(self.v)-1)
    def pop(self) :
        ans = self.v[1]
        xx = self.v.pop()
        if len(self.v) > 1 :
            self.v[1] = xx
            self._bubbledown(1)
        return ans
    def _bubbleup(self,idx) :
        if idx == 1 : return
        j = idx >> 1
        if self.v[j] < self.v[idx] :
            self.v[j],self.v[idx] = self.v[idx],self.v[j]
            self._bubbleup(j)
    def _bubbledown(self,idx) :
        l = idx << 1; r = l + 1
        ll = len(self.v)
        res1 = l >= ll or self.v[idx] >= self.v[l]
        res2 = r >= ll or self.v[idx] >= self.v[r]
        if res1 and res2 : return
        if res1 : self.v[idx],self.v[r] = self.v[r],self.v[idx]; self._bubbledown(r); return
        if res2 : self.v[idx],self.v[l] = self.v[l],self.v[idx]; self._bubbledown(l); return
        if self.v[l] >= self.v[r] : self.v[idx],self.v[l] = self.v[l],self.v[idx]; self._bubbledown(l); return
        self.v[idx],self.v[r] = self.v[r],self.v[idx]; self._bubbledown(r)

class minHeap :
    v = []
    def __init__(self) : self.v = [0]
    def len(self) : return len(self.v)-1
    def isempty(self) : return len(self.v) == 1
    def top(self) : return self.v[1]
    def push(self,val) :
        self.v.append(val)
        self._bubbleup(len(self.v)-1)
    def pop(self) :
        ans = self.v[1]
        xx = self.v.pop()
        if len(self.v) > 1 :
            self.v[1] = xx
            self._bubbledown(1)
        return ans
    def _bubbleup(self,idx) :
        if idx == 1 : return
        j = idx >> 1
        if self.v[j] > self.v[idx] :
            self.v[j],self.v[idx] = self.v[idx],self.v[j]
            self._bubbleup(j)
    def _bubbledown(self,idx) :
        l = idx << 1; r = l + 1
        ll = len(self.v)
        res1 = l >= ll or self.v[idx] <= self.v[l]
        res2 = r >= ll or self.v[idx] <= self.v[r]
        if res1 and res2 : return
        if res1 : self.v[idx],self.v[r] = self.v[r],self.v[idx]; self._bubbledown(r); return
        if res2 : self.v[idx],self.v[l] = self.v[l],self.v[idx]; self._bubbledown(l); return
        if self.v[l] <= self.v[r] : self.v[idx],self.v[l] = self.v[l],self.v[idx]; self._bubbledown(l); return
        self.v[idx],self.v[r] = self.v[r],self.v[idx]; self._bubbledown(r)

def testit(numop,maxval) :
    wrchance = random.uniform(0.45,0.55)
    mh = maxHeap()
    for i in range(numop) : 
        if mh.isempty() or random.random() < wrchance :
            mh.push(random.randrange(1,maxval+1))
        else :
            mh.pop()
        if not mh.isempty() and max(mh.v[1:]) != mh.top() :
            print(f"ERROR: numop:{numop} maximum:{max(mh.v[1:])} top:{mh.top()}")

def testit2(numop,maxval) :
    wrchance = random.uniform(0.45,0.55)
    mh = minHeap()
    for i in range(numop) : 
        if mh.isempty() or random.random() < wrchance :
            mh.push(random.randrange(1,maxval+1))
        else :
            mh.pop()
        if not mh.isempty() and min(mh.v[1:]) != mh.top() :
            print(f"ERROR: numop:{numop} minimum:{min(mh.v[1:])} top:{mh.top()}")


if __name__ == '__main__' :
    random.seed(19006492568)
    for i in range(10) :
        print(f"iter:{i}")
        testit(10,10)
        testit(100,100)
        testit(1000,1000)
        testit(1000,1000)
    for i in range(10) :
        print(f"iter:{i}")
        testit2(10,10)
        testit2(100,100)
        testit2(1000,1000)
        testit2(1000,1000)

