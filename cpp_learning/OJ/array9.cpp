#include <iostream>
#include <string>
#include <vector>
#include <stack>

using namespace std;


/**
实现一个基本的计算器来计算一个简单的字符串表达式的值。

字符串表达式仅包含非负整数，+， - ，*，/ 四种运算符和空格 。 整数除法仅保留整数部分。

示例 1:

输入: "3+2*2"
输出: 7
示例 2:

输入: " 3/2 "
输出: 1
示例 3:

输入: " 3+5 / 2 "
输出: 5
说明：

你可以假设所给定的表达式都是有效的。
请不要使用内置的库函数 eval。
 */



class Solution {
public:
    int calculate(string s) {
        int res = 0, num = 0, n = s.size();
        char op = '+';
        stack<int> st;
        for (int i = 0; i < n; ++i) {
            if (s[i] >= '0') {
                num = num * 10 + s[i] - '0';
            }
            if ((s[i] < '0' && s[i] != ' ') || i == n - 1) {
                if (op == '+') st.push(num);
                if (op == '-') st.push(-num);
                if (op == '*' || op == '/') {
                    int tmp = (op == '*') ? st.top() * num : st.top() / num;
                    st.pop();
                    st.push(tmp);
                }
                op = s[i];
                num = 0;
            }
        }
        while (!st.empty()) {
            res += st.top();
            st.pop();
        }
        return res;
    }

    vector<string> split(string s, char c) {
        vector<string> result;
        if (!s.size()) {
            return result;
        }

        int start = -1;
        for (int i = 0; i <= s.size(); i++) {
            if (i == s.size() || s[i] == c) {
                string sub = s.substr(start + 1, i - 1 - start);
                result.push_back(sub);
                start = i;
            }
        }
        return result;
    }

    string trim(string s) {
        if (!s.size()) {
            return "";
        }
        int start = 0;
        while (s[start] == ' ' && start < s.size()) {
            start++;
        }
        int end = s.size() - 1;
        while (s[end] == ' ' && end >= 0) {
            end--;
        }
        if (start > end) {
            return "";
        }
        return s.substr(start, end - start + 1);
    }
};


int main() {
    Solution s;
    vector<string> data = s.split("abc def g", ' ');
    for (auto d : data) {
        cout << d << endl;
    }
    cout << s.trim("   ABCD  ") << endl;

    return 0;
}