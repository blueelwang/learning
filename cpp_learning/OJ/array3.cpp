#include <iostream>
#include <unordered_map>

using namespace std;


/*
给定四个包含整数的数组列表 A , B , C , D ,计算有多少个元组 (i, j, k, l) ，使得 A[i] + B[j] + C[k] + D[l] = 0。

为了使问题简单化，所有的 A, B, C, D 具有相同的长度 N，且 0 ≤ N ≤ 500 。所有整数的范围在 -228 到 228 - 1 之间，最终结果不会超过 231 - 1 。

例如:

输入:
A = [ 1, 2]
B = [-2,-1]
C = [-1, 2]
D = [ 0, 2]

输出:
2

解释:
两个元组如下:
1. (0, 0, 0, 1) -> A[0] + B[0] + C[0] + D[1] = 1 + (-2) + (-1) + 2 = 0
2. (1, 1, 0, 0) -> A[1] + B[1] + C[0] + D[0] = 2 + (-1) + (-1) + 0 = 0
 */

class Solution {
public:
    int fourSumCount(vector<int>& A, vector<int>& B, vector<int>& C, vector<int>& D) {
        int ans = 0;
        unordered_map<int, int> sums1 = addTwoArray(A, B);
        unordered_map<int, int> sums2 = addTwoArray(C, D);
        for (unordered_map<int, int>::iterator it1 = sums1.begin(); it1 != sums1.end(); it1++) {
            int target = 0 - it1->first;
            unordered_map<int, int>::iterator targetIter = sums2.find(target);
            if (targetIter != sums2.end()) {
                ans += it1->second * targetIter->second;
            }
        }

        return ans;
    }

    unordered_map<int, int> addTwoArray(vector<int> a, vector<int> b) {
        unordered_map<int, int> sums;
        for (int i : a) {
            for (int j : b) {
                sums[i + j]++;
            }
        }
        return sums;
    };


};