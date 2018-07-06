/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
struct ListNode* swapPairs(struct ListNode* head) {
    
    //遍历链表，每两个节点来一次对换，每次换完把结果挂载在结果链表中，只剩一个节点的时候直接挂载上去
    struct ListNode *current = NULL, *top = NULL, *tmp = NULL, *nexttemp = NULL;
    while(head)
    {
        if (head->next) {
           tmp = head;
           nexttemp = head->next->next;
           head = head->next;
           head->next = tmp;
           //交换完再把头结点挂载到结果链表中
           if(top == NULL){
               current = top = head;
               current = current->next;
           } else {
               current->next = head;
               current = current->next->next;
           }
           
            head->next->next = nexttemp;
            head = head->next->next;
        } else {
            if (top == NULL) {
               return head; 
            } else {
                current->next = head;
            }
            head = head->next;
        }
    }
    
    return top;
}