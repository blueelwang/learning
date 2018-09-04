#include <iostream>
#include <vector>

using namespace std;


/*
根据百度百科，生命游戏，简称为生命，是英国数学家约翰·何顿·康威在1970年发明的细胞自动机。

给定一个包含 m × n 个格子的面板，每一个格子都可以看成是一个细胞。每个细胞具有一个初始状态 live（1）即为活细胞， 或 dead（0）即为死细胞。每个细胞与其八个相邻位置（水平，垂直，对角线）的细胞都遵循以下四条生存定律：

1. 如果活细胞周围八个位置的活细胞数少于两个，则该位置活细胞死亡；

2. 如果活细胞周围八个位置有两个或三个活细胞，则该位置活细胞仍然存活；

3. 如果活细胞周围八个位置有超过三个活细胞，则该位置活细胞死亡；

4. 如果死细胞周围正好有三个活细胞，则该位置死细胞复活；

根据当前状态，写一个函数来计算面板上细胞的下一个（一次更新后的）状态。

进阶:

你可以使用原地算法解题吗？请记住，面板上所有格子需要同时被更新：你不能先更新某些格子，然后使用它们的更新后的值更新其他格子。
在此题中，我们使用二维数组来表示面板。原则上，面板是无限的，但当活细胞侵占了面板边界时会造成问题。你将如何解决这些问题？
 */



class Solution {
public:
    void gameOfLife(vector<vector<int>>& board) {
        int rows = board.size();
        if (rows <= 0) {
            return;
        }
        int cols = board[0].size();
        if (cols <= 0) {
            return;
        }

        for (int i = 0; i < rows; i++) {
            for (int j = 0; j < cols; j++) {
                int liveCnt = 0;
                // count live cells around
                for (int m = max(i - 1, 0); m <= min(i + 1, rows - 1); m++) {
                    for (int n = max(j - 1, 0); n <= min(j + 1, cols - 1); n++) {
                        liveCnt += board[m][n] & 1;
                    }
                }
                liveCnt -= board[i][j] & 1;
                if (board[i][j] && (2 == liveCnt || 3 == liveCnt)
                    || !board[i][j] && 3 == liveCnt) {
                    board[i][j] += 2;
                }
            }
        }
        for (int i = 0; i < rows; i++) {
            for (int j = 0; j < cols; j++) {
                board[i][j] >>= 1;
            }
        }
    }
};

int main() {
    Solution s;
    vector<vector<int>> data{{0}};
    s.gameOfLife(data);
    return 0;
}
