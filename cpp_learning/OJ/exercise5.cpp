#include <iostream>
#include <string>


using namespace std;


class Solution {
public:
    string longestPalindrome(string s) {
        int len = s.length();
        if (len < 2) {
            return s;
        }
        int subStrMaxLen = 0;
        int subStrStart = 0;
        int subStrEnd = 0;
        for (int mid = 0; mid < len - 1; mid++) {
            for (int l = mid, r = mid; l >= 0 && r <= len; l--, r++) {
                if (s[l] == s[r]) {
                    int tmpLen = r - l + 1;
                    if (tmpLen > subStrMaxLen) {
                        subStrMaxLen = tmpLen;
                        subStrStart = l;
                        subStrEnd = r;
                    }
                    continue;
                }
                break;
            }
            for (int l = mid, r = mid + 1; l >= 0 && r <= len; l--, r++) {
                if (s[l] == s[r]) {
                    int tmpLen = r - l + 1;
                    if (tmpLen > subStrMaxLen) {
                        subStrMaxLen = tmpLen;
                        subStrStart = l;
                        subStrEnd = r;
                    }
                    continue;
                }
                break;
            }
        }
        return s.substr(subStrStart, subStrMaxLen);
    }

};

int main() {

    string tmp = "0123456";
    cout << tmp.substr(0, 3)<< endl;

    Solution s;
    string str = "babad";
    string res = s.longestPalindrome(str);
    cout<<res<<endl;

    return 0;
}