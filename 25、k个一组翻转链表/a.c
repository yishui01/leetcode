/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
struct ListNode* reverseKGroup(struct ListNode* head, int k) {
    //每三个节点来一次对换，换完之后保存到数组，再把数组组成新的一小段链表挂载到结果链表中，不够三个节点则不交换直接挂载上去
    struct ListNode *current = NULL, *top = NULL, *tmp = NULL, *check = NULL, *last = NULL;
    struct ListNode *arr[k];
    int i = 0, j = 0, z = 0,  isend = 0;
   
    while (head) {
        check = head; 
        for(i = 0; i < k-1; i++) { //看剩余节点够不够k个
            if(!check->next){
                isend = 1;
                break;
            }
            check = check->next;
        }

        if (isend == 1) {
            //如果剩余节点不够K个了，直接挂载
            if(top == NULL) {
                return head;
            } else {
                current->next = head;
            }
            return top;
        } else {
            //如果没结束，交换k个节点，将交换后的顺序保存到数组中，再把数组组成链表挂载到结果链表中
            i = 0;
            j = k-1;
            
            while(i<k) {
                tmp = head;
                for(z = 0;z<i; z++) {
                    tmp = tmp->next;
                }
                arr[j] = tmp;
                tmp = head;
                for(z = 0;z<j; z++) {
                    tmp = tmp->next;
                }
                arr[i] = tmp;
                i++;
                j--;
            }
            
             head = arr[0]->next; //保存最原始的下一轮head节点

            //再把arr里的节点按顺序挂载到current节点中
            for(i = 0; i < k; i++) {
                if(current == NULL) {
                    top=current = arr[i];
                } else {
                     current->next = arr[i];
                     current = current->next;
                     current->next = NULL;
                    
                }

            }

        }
 
    }
    return top;
}