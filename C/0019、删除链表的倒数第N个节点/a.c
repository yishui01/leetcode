/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
struct ListNode* removeNthFromEnd(struct ListNode* head, int n) {
    if(head == NULL)return NULL;
    //只能先从头部遍历到尾部，同时记录每个node的位置
    struct ListNode *place[100]; //用于保存每个node的位置
    int i = 0;
    struct ListNode *current;
    current = head;
    while(current) {
        place[i] = current;
        i++;
        current = current->next;
    }
    if(n == i) {
        //n是头部元素
        if(i == 1) {
            return NULL;
        }
        head = place[1];
    } else if(n == 1) {
        //n是尾部元素
        place[i-2]->next = NULL;
    } else {
        //n是中间元素
        //place[i-n] //移除
        place[i-n-1]->next = place[i-n+1];
    }
    return head;
    
}