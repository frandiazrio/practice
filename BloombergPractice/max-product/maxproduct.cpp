#include <iostream>



int maxProduct(vector<int>& nums) {  
        if(nums.size()==1)
          return nums[0];
        int curr_max=nums[0];
        int curr_min=nums[0]; 
        int prev_min=nums[0]; 
        int prev_max=nums[0]; 
        int ans=nums[0];
        for(int i=1;i<nums.size();i++)
        {
          curr_max=max(nums[i]*prev_max,max(nums[i]*prev_min,nums[i]));
          curr_min=min(nums[i]*prev_min,min(nums[i]*prev_max,nums[i]));
          ans=max(ans,curr_max); 
          prev_min=curr_min; 
          prev_max=curr_max;
        } 
        return ans;  
}
