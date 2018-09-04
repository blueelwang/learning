#include<iostream>
#include<vector>

using namespace std;


/*
给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。画 n 条垂直线，使得垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水，得到最大的容水量

注意：你不能倾斜容器，n 至少是2。
 */


class Solution {
public:
    int maxArea(vector<int>& height) {
        int max = 0;
        int l = 0;
        int r =  height.size() - 1;
        while (l < r) {
            int value = (r - l) * (height[l] < height[r] ? height[l] : height[r]);
            if (value > max) {
                max = value;
            }
            if (height[l] < height[r]) {
                l++;
            } else {
                r--;
            }
        }
        return max;
    }
};

int main() {
    Solution s;
    vector<int> data = {0, 2};
    s.maxArea(data);
}
