#include <iostream>
using namespace std;

struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
    int carry = 0;
    struct ListNode *head = NULL;
    struct ListNode *cur = NULL;
    while (NULL != l1 && NULL != l2) {
        int sum = l1->val + l2->val + carry;
        carry = (int) sum / 10;
        struct ListNode *tmp = new ListNode(sum % 10);

        if (NULL == cur) {
            head = tmp;
            cur = tmp;
        } else {
            cur->next = tmp;
            cur = cur->next;
        }
        l1 = l1->next;
        l2 = l2->next;
    }
    while (NULL != l1) {
        int sum = l1->val + carry;
        carry = (int) sum / 10;
        struct ListNode *tmp = new ListNode(sum % 10);

        if (NULL == cur) {
            head = tmp;
            cur = tmp;
        } else {
            cur->next = tmp;
            cur = cur->next;
        }
        l1 = l1->next;
    }
    while (NULL != l2) {
        int sum = l2->val + carry;
        carry = (int) sum / 10;
        struct ListNode *tmp = new ListNode(sum % 10);

        if (NULL == cur) {
            head = tmp;
            cur = tmp;
        } else {
            cur->next = tmp;
            cur = cur->next;
        }
        l2 = l2->next;
    }
    if (carry > 0) {
        struct ListNode *tmp = new ListNode(carry);
        cur->next = tmp;
    }

    return head;
}

int main() {
    ListNode t11(2);
    ListNode t12(4);
    ListNode t13(3);
    ListNode t21(5);
    ListNode t22(6);
    ListNode t23(4);

    t11.next = &t12;
    t12.next = &t13;

    t21.next = &t22;
    t22.next = &t23;

    ListNode *res = addTwoNumbers(&t11, &t21);

    do {
        cout<<res->val<<endl;
        res = res->next;
    } while(NULL != res);
}