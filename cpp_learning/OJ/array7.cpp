#include <iostream>
#include <vector>
#include <map>

using namespace std;

/*
给定一个未排序的整数数组，找出最长连续序列的长度。

要求算法的时间复杂度为 O(n)。

示例:

输入: [100, 4, 200, 1, 3, 2]
输出: 4
解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。
 */


class Solution {
public:
    int longestConsecutive(vector<int>& nums) {
        vector<int>::size_type size = nums.size();
        if (size < 2) {
            return size;
        }
        map<int, int> data;
        for (int i : nums) {
            data[i]++;
        }
        int max = 1;
        int len = 0;
        int last;
        for (map<int, int>::iterator it = data.begin(); it != data.end(); it++) {
            if (it == data.begin() || it->first == last + 1) {
                last = it->first;
                len++;
                if (len > max) {
                    max = len;
                }
                continue;
            }
            last = it->first;
            len = 1;
        }
        return max;
    }

    // https://www.youtube.com/watch?v=rc2QdQ7U78I
    int longestConsecutive(vector<int>& nums) {
        unordered_map<int, int> m; // m[k] : length of the seq bounded by k

        int res = 0;
        for(auto k : nums) {
            if(m[k]!=0) {
                continue; // dup, k is visited already
            }

            int leftLen = m[k-1], rightLen = m[k+1]; // lengths of seqs on left and right
            if(leftLen==0 && rightLen==0) {
                // k is a single elem seq
                m[k] = 1;
            } else if(leftLen==0) {
                // has a seq on right side. New seq is [k ... k+right]
                // update m[k] and m[k+right]
                m[k] = m[k+rightLen] = rightLen+1;
            } else if(rightLen==0) {
                m[k] = m[k-leftLen] = leftLen+1;
            } else {
                // has seq on both sides. need to update m[k] to make it non-zero
                m[k] = m[k-leftLen] = m[k+rightLen] = leftLen + rightLen +1;
            }
            res = max(res, m[k]);
        }
        return res;
    }
};

int main() {
    Solution s;
    vector<int> data{9,1,4,7,3,-1,0,5,8,-1,6};
    s.longestConsecutive(data);
    return 0;
}