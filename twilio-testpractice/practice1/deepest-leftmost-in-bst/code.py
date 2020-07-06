# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution(object):
    def findBottomLeftValue(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        
        def dfs(root, level, curr):
            if not root:
                return 
            
            
            dfs(root.left, level, curr+1)
            if curr>level[0]:
                level[0] = curr
                ans[0] = root.val
            dfs(root.right, level, curr+1)
            
        ans = [root.val]
        level = [0]
        dfs(root, level, 0)
        return ans[0]