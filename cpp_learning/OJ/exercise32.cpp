#include<iostream>

using namespace std;

/*
Given a string containing just the characters '(' and ')', find the length of the longest valid (well-formed) parentheses substring.

Example 1:

Input: "(()"
Output: 2
Explanation: The longest valid parentheses substring is "()"
Example 2:

Input: ")()())"
Output: 4
Explanation: The longest valid parentheses substring is "()()"
 */


class Solution {
public:
    int longestValidParentheses(string s) {
        int len = s.length();
        if (len < 2) {
            return 0;
        }

        int result = 0;

        int cnt = 0;
        int l = 0;
        for (; l < len; l++) {
            if ('(' == s[l]) {
                cnt = 1;
                break;
            }
        }
        if (l >= len - 1) {
            return result;
        }

        for (int r = l + 1; r < len; r++) {
            if ('(' == s[r]) {
                cnt++;
                continue;
            }
            cnt--;
            if (0 == cnt) {
                int rangeLen = r - l + 1;
                if (rangeLen > result) {
                    result = rangeLen;
                }
            } else if (cnt < 0) {
                r++;
                l = r;
                cnt = 0;
            }
        }
        if (cnt > 0) {
            for (int r = len - 1; l < r; l++) {
                if (')' == s[l]) {
                    cnt++;
                    continue;
                }
                cnt--;
                if (0 == cnt) {
                    int rangeLen = r - l;
                    if (rangeLen > result) {
                        result = rangeLen;
                    }
                    break;
                }
            }
        }
        return result;
    }
};


int main() {
    Solution s;
    int res = s.longestValidParentheses("(()");
    cout<<res<<endl;
    return 0;
}