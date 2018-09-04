#include<iostream>
#include<vector>
#include<map>
using namespace std;

class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        int size = nums.size();
        for (int i = 0; i < size - 1; i++) {
            for (int j = i+1; j < size; j++) {
                if (nums[i] + nums[j] == target) {
                    return {i, j};
                }
            }
        }
    }
};

int main() {
    vector<int> nums = {3,2,4};
    int target = 6;

    Solution s;
    vector<int> res = s.twoSum(nums, target);
    vector<int>::iterator iter;
    for (iter = res.begin(); iter != res.end(); iter++) {
        cout<<*iter<<"\t";
    }
    return 0;
}
