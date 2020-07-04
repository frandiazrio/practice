from collections import OrderedDict
class LRU(OrderedDict):


    def __init__(self, capacity):
        self.capacity = capacity
        pass

    def get(self,key):

        if key not in self:
            return -1
        self.move_to_end(key)
        return self[key]

    def put(self,key,value):

        if key in self:
            self.move_to_end(key)
        self[key]=value
        
        if len(self) > self.capcity:
            self.popitem(last=False)


