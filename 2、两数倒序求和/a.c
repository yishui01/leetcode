/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */

struct ListNode* addTwoNumbers(struct ListNode* l1, struct ListNode* l2) {
    int v1 = 0, v2 = 0, i = 0, j = 0;
    int num = 0; //迭代变量
    int isadd = 0; //两数相加是否进位
    struct ListNode * head = NULL; //头指针
    struct ListNode * p1 = NULL;   
    struct ListNode * p2 = NULL; 
    while(l1 || l2 || isadd){
       v1 = (l1 == NULL) ? 0 : l1->val;
       v2 = (l2 == NULL) ? 0 : l2->val;
       num = v1 + v2;
       if(isadd != 0){
           num += 1; //进位
       }
       if(num >= 10){
           isadd = 1;
       } else{
           isadd = 0;
       }
       num = num % 10; 
      p1 = (struct ListNode *) malloc (sizeof(struct ListNode));
         if(p1 == NULL) {
             printf("Cann't create node");
             return NULL;
         }else{
             p1->val = num;
             p1->next = NULL;
         }
        if(head == NULL){
            p2 = head = p1;
        }else{
            p2->next = p1;
            p2 = p2->next;
        }
        if(l1){
            l1 = l1->next;
        }
        if(l2){
            l2 = l2->next;
        }
       
        
    }
    
    return head;
}