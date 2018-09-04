#include <iostream>
#include <vector>
#include<queue>

using namespace std;


/**
 * Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity
 * Example:
 * Input:
 * [
 *  1->4->5,
 *  1->3->4,
 *  2->6
 * ]
 * Output: 1->1->2->3->4->4->5->6
 */

struct ListNode {
    int val;
    ListNode *next;

    ListNode(int x) : val(x), next(NULL) {}
};

struct HeapNode {
    int val;
    int ListNum;
    HeapNode(int val, int ListNum) : val(val), ListNum(ListNum) {}
    bool operator<(const HeapNode& rhs) const { return val > rhs.val; }
};

class Solution {
public:
    /*ListNode* mergeKLists(vector<ListNode*>& lists) {
        ListNode node(0);
        ListNode *head = &node;
        ListNode *tail = head;


        while (true) {
            bool isFinished = true;

            ListNode *minNode = NULL;
            vector<ListNode*>::size_type minNodeIndex = 0;

            vector<ListNode*>::size_type vecLen = lists.size();
            for (vector<ListNode *>::size_type i = 0; i < vecLen; i++) {
                if (NULL == lists[i]) {
                    continue;
                }
                isFinished = false;

                if (!minNode || lists[i]->val < minNode->val) {
                    minNode = lists[i];
                    minNodeIndex = i;
                }
            }
            if (isFinished) {
                break;
            }

            tail->next = new ListNode(minNode->val);
            tail = tail->next;

            lists[minNodeIndex] = lists[minNodeIndex]->next;

        }

        return head->next;
    }*/
    ListNode* mergeKLists(vector<ListNode*>& lists) {
        ListNode* dummy = new ListNode(0);
        ListNode* now = dummy;
        priority_queue<HeapNode> q;
        for (int i = 0; i < lists.size(); i++)
            if (lists[i])
                q.push(HeapNode(lists[i]->val, i));
        while (!q.empty()) {
            auto ListNum = q.top().ListNum;
            q.pop();
            now = now->next = lists[ListNum];
            lists[ListNum] = lists[ListNum]->next;
            if (lists[ListNum])
                q.push(HeapNode(lists[ListNum]->val, ListNum));
        }
        now->next = NULL;
        ListNode* ans = dummy->next;
        delete dummy;
        return ans;
    }
};

int main() {
    ListNode n11(1);
    ListNode n12(1);
    ListNode n2(2);
    ListNode n3(3);
    ListNode n41(4);
    ListNode n42(4);
    ListNode n5(5);
    ListNode n6(6);

    n11.next = &n41;
    n41.next = &n5;

    n12.next = &n3;
    n3.next = &n42;

    n2.next = &n6;

    vector<ListNode*> vec = {&n11, &n12, &n2};


    Solution s;
    ListNode *res = s.mergeKLists(vec);
    while (res) {
        cout<<res->val<<endl;
        res = res->next;
    }
    return 0;
}