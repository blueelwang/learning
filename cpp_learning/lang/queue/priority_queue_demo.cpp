#include <queue>
#include <iostream>

using namespace std;

struct Node{
    int x, y;
    Node(int a = 0, int b = 0) : x(a), y(b) {}
};

bool operator< (const Node& a, const Node& b ){
    if (a.x == b.x)
        return a.y < b.y;
    return a.x < b.x;
}

int main() {
    //priority_queue<int, vector<int>, greater<int> > q_g;    // 小顶堆
    //priority_queue<int, vector<int>, less<int> > q_l;       // 大顶堆

    priority_queue<Node> q;
    q.push(Node(1,1));
    q.push(Node(2,2));
    q.push(Node(4,4));
    q.push(Node(3,3));

    while (!q.empty()) {
        cout << q.top().x << q.top().y << endl;
        q.pop();
    }


    return 0;
}