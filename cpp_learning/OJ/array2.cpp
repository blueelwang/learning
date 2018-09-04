#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
    vector<int> spiralOrder(vector<vector<int>>& matrix) {
        int m = matrix.size();
        if (m < 1) {
            return {};
        }
        int n = matrix[0].size();
        if (n < 1) {
            return {};
        }

        vector<int> result(m * n);
        int len[4] = {n, m, n, m};

        int dir = 0;
        int i = 0;
        int j = -1;
        int move = 0;
        for (int index = 0; index < m * n; index++) {
            if (++move <= len[dir]) {
                switch (dir) {
                    case 0:
                        j++;
                        break;
                    case 1:
                        i++;
                        break;
                    case 2:
                        j--;
                        break;
                    default:
                        i--;
                }
                result[index] = matrix[i][j];
            }
            if (move >= len[dir]) {
                len[(dir + 1) % 4]--;
                len[(dir + 3) % 4]--;
                dir = (dir + 1) % 4;
                move = 0;
            }
        }

        return result;
    }
};

int main() {

    Solution s;
    vector<vector<int>> input = {{1, 2, 3}, {4, 5, 6}, {7, 8, 9}};
    vector<int> res = s.spiralOrder(input);
    for (int i = 0; i < res.size(); i++) {
        cout<<res[i]<<endl;
    }

    return 0;
}