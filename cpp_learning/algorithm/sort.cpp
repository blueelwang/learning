#include <iostream>
#include <vector>

using namespace std;

vector<int> bubble(vector<int> nums) {
    for (vector<int>::size_type i = nums.size() - 1; i > 0; i--) {
        for (vector<int>::size_type j = 0; j < i; j++) {
            if (nums[j] > nums[j + 1]) {
                int tmp = nums[j + 1];
                nums[j + 1] = nums[j];
                nums[j] = tmp;
            }
        }
    }
    return nums;
}

vector<int> bubble_ex(vector<int> nums) {
    vector<int>::size_type boundary = nums.size() - 1;
    while (boundary > 0) {
        vector<int>::size_type b = 0;
        for (vector<int>::size_type i = 0; i < boundary; i++) {
            if (nums[i] > nums[i + 1]) {
                int tmp = nums[i + 1];
                nums[i + 1] = nums[i];
                nums[i] = tmp;
                b = i;
            }
        }
        boundary = b;
    }
    return nums;
}

vector<int> insert(vector<int> nums) {
    int len = nums.size();
    if (len <= 1) {
        return nums;
    }
    for (vector<int>::size_type i = 1; i < len; i++) {
        int curNum = nums[i];
        for (vector<int>::size_type j = i - 1; j >= 0; j--) {
            if (nums[j] > curNum) {
                nums[j + 1] = nums[j];
                if (j == 0) {
                    nums[j] = curNum;
                    break;
                }
            } else {
                nums[j + 1] = curNum;
                break;
            }
        }
    }
    return nums;
}

vector<int> fast(vector<int> nums) {

    return nums;
}

int main() {
    vector<int> data = {823, 3, 22, 29, 228, 29, 983, 2, 37, 883, 20, 49, 18};
    data = insert(data);
    for (int d : data) {
        cout << d << endl;
    }
    return 0;
}
