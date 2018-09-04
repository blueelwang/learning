#include <iostream>
#include <vector>

using namespace std;

/**
给定一个包含 n + 1 个整数的数组 nums，其数字都在 1 到 n 之间，包括 1 和 n ，可知至少存在一个重复的整数。假设只有一个重复的整数，找出这个重复的数。

示例 1:
输入: [1,3,4,2,2]
输出: 2

示例 2:
输入: [3,1,3,4,2]
输出: 3

说明：
不能更改原数组（假设数组是只读的）。
只能使用额外的 O(1) 的空间。
时间复杂度小于 O(n2) 。
数组中只有一个重复的数字，但它可能不止重复出现一次。
 */

class Solution {
public:
    int findDuplicate(vector<int>& nums) {
        int slow = 0;
        int fast = 0;
        // 找到快慢指针相遇的地方
        do{
            slow = nums[slow];
            fast = nums[nums[fast]];
        } while(slow != fast);
        int find = 0;
        // 注意下面这三个时间点
        // 1. 快指针走一圈
        // 2. 慢指针和快指针相遇
        // 3. 慢指针走到1位置
        // 2到3的距离正好等于起点到1的距离
        // 此时处理2时间点，用一个新指针从头开始，和慢指针相遇的位置就是环的起点
        while(find != slow){
            slow = nums[slow];
            find = nums[find];
        }
        return find;
    }
};


int main() {
    Solution s;
    vector<int> data{2,5,9,6,9,3,8,9,7,1};
    int res = s.findDuplicate(data);
    cout << res << endl;
    return 0;
}
