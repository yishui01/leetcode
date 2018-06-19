/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */

/*
每个链表头结点依次比对，找出最小值，再重新比对一遍，找出和最小值相等的链表，保存到结果链表中，并将原来的链表指向next，无限循环，直到所有链表都指向NULL
*/
struct ListNode* mergeKLists(struct ListNode** lists, int listsSize) {
    int min, i = 0,j = 0, flag = 0, isend = 1;//flag标识当前最小的链表是第几个, isend标识是否全部链表以及到头了
    struct ListNode * head, *current;
    current = head = NULL;
        while(1) {
            isend = 1;
            min = INT_MAX;
            for(i = 0; i < listsSize;i++) {
                if(lists[i]){
                    isend = 0; //代表还有链表没完成
                    if(i == 0) {
                        min = lists[i]->val;
                        flag = i;
                    } else {
                        if(lists[i]->val < min) {
                            min = lists[i]->val;
                            flag = i;
                        }
                    } 
                }

            }

          //把所有值等于最小值的链表+1
            for(j = 0; j < listsSize; j++) {
                if(lists[j] && (lists[j]->val == min)){
                   
                    if (head == NULL) {
                        head = current = lists[j];
                    } else {
                        current->next = lists[j];
                        current = current->next;
                    }
                     lists[j] = lists[j]->next;

                }
            }
            if(isend == 1)break;
        }
    
    return head;
    
    
}