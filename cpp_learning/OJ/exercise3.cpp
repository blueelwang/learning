/*
Given a string, find the length of the longest substring without repeating characters.

Examples:

Given "abcabcbb", the answer is "abc", which the length is 3.

Given "bbbbb", the answer is "b", with the length of 1.

Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
 */


#include <iostream>
#include <string>
#include <map>

using namespace std;

class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        int len = s.length();
        if (len < 2) {
            return len;
        }
        int maxSubStringLen = 1;
        for (int i = 0, j = 1; j < len; j++) {
            for (int index = j - 1; index >= i; index--) {
                int subStringLen = 0;
                if (s[index] == s[j]) {
                    i = index + 1;
                    subStringLen = j - index;
                } else {
                    subStringLen = j - index + 1;
                }
                if (subStringLen > maxSubStringLen) {
                    maxSubStringLen = subStringLen;
                }
            }
        }
        return maxSubStringLen;
    }
};

int main() {
    Solution s;
    int maxSubStrLen = s.lengthOfLongestSubstring("DannyPlaysWithHisFriends");
    cout<<maxSubStrLen<<endl;

    return 0;
}