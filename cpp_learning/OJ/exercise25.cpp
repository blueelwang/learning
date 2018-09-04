#include <iostream>

using namespace std;

/**
 * Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.
 * k is a positive integer and is less than or equal to the length of the linked list.
 * If the number of nodes is not a multiple of k then left-out nodes in the end should remain as it is.
 *
 * Example:
 * Given this linked list: 1->2->3->4->5
 * For k = 2, you should return: 2->1->4->3->5
 * For k = 3, you should return: 3->2->1->4->5
 *
 * Note:
 * Only constant extra memory is allowed.
 * You may not alter the values in the list's nodes, only nodes itself may be changed.
 */


struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

class Solution {
public:
    ListNode* reverseKGroup(ListNode* head, int k) {
        if (k < 2) {
            return head;
        }
        ListNode dummy(0);
        ListNode *pdummy = &dummy;
        pdummy->next = head;

        ListNode *lastTail  = pdummy;
        ListNode *start     = head;
        ListNode *cur       = head;

        bool isFinished = false;
        while (!isFinished) {
            for (int i = 1; i <= k - 1; i++) {
                if (!cur || !cur->next) {
                    isFinished = true;
                    break;
                }
                cur = cur->next;

                start->next = cur->next;
                cur->next = lastTail->next;
                lastTail->next = cur;

                cur = start;
            }
            if (isFinished) {
                break;
            }

            lastTail = start;
            start = start->next;
            cur = start;
        }

        if (lastTail->next != start) {
            start = lastTail->next;
            cur = start;
            while (true) {
                if (!cur->next) {
                    break;
                }
                cur = cur->next;

                start->next = cur->next;
                cur->next = lastTail->next;
                lastTail->next = cur;

                cur = start;
            }
        }

        return pdummy->next;
    }
};


int main() {
    ListNode n1(1);
    ListNode n2(2);
    ListNode n3(3);
    ListNode n4(4);
    ListNode n5(5);
    ListNode n6(6);
    ListNode n7(7);
    ListNode n8(8);
    ListNode n9(9);
    ListNode n10(10);
    ListNode n11(11);


    n1.next = &n2;
    n2.next = &n3;
    n3.next = &n4;
    n4.next = &n5;
    n5.next = &n6;
    n6.next = &n7;
    n7.next = &n8;
    n8.next = &n9;
    n9.next = &n10;
    n10.next = &n11;

    int k = 11;
    Solution s;
    ListNode *res = s.reverseKGroup(&n1, k);
    while (res) {
        cout<<res->val<<endl;
        res = res->next;
    }

    return 0;
}