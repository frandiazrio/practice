class Solution:
    def knightDialer(self, N: int) -> int:
        
        d = {
            1: [6,8],
            2: [7,9],
            3: [4,8],
            4: [3,9,0],
            5: [],
            6: [1,7,0],
            7: [2,6],
            8: [1,3],
            9: [2,4],
            0: [4,6]
        }
        
        m = {}
        
        def kd(k, n):
            if (k,n) in m:
                return m[(k,n)]
            
            if n == 1:
                return 1
            
            total = 0
            for i in d[k]:
                total += kd(i, n-1)
            
            m[(k,n)] = total
            
            return total
        
        r = 0
        for i in range(10):
            r += kd(i, N)
            
        return r % (10**9 + 7)
