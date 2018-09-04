#include <vector>
#include <iostream>

using namespace std;


class Solution {
public:
    double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2) {
        vector<int>::size_type len1 = nums1.size();
        vector<int>::size_type len2 = nums2.size();
        vector<int>::size_type len3 = len1 + len2;
        vector<int> allNums(len3);
        int index1 = 0;
        int index2 = 0;
        while (index1 < len1 || index2 < len2) {
            int value1 = (index1 < len1) ? nums1[index1] : 2147483647;
            int value2 = (index2 < len2) ? nums2[index2] : 2147483647;
            if (index1 < len1 && value1 < value2) {
                allNums[index1 + index2] = value1;
                index1++;
            } else {
                allNums[index1 + index2] = value2;
                index2++;
            }
        }
        if (len3 == 0) {
            return 0;
        } else if (len3 <= 1) {
            return allNums[0];
        } else {
            return (len3 % 2 == 0) ? (allNums[len3 / 2] + allNums[len3 / 2 - 1]) / 2.0 : allNums[len3 / 2];
        }
    }
};


int main() {

    Solution s;
    vector<int> nums1 = {1, 3};
    vector<int> nums2 = {2};
    double res = s.findMedianSortedArrays(nums1, nums2);
    cout<<res<<endl;

    return 0;
}