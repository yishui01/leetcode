#include <stdio.h>
#include <string.h>
#include <stdlib.h>

//递归函数
void recursion(int target, int now, char *tmp, char *res, int size) 
{
    int i = 0, now_count = 0, now_index = 0;
    char now_char = '\0';
    if (now != target) {
        memset(tmp, '\0', size);
        for(i = 0; i < strlen(res); i++) {
			
            //遍历当前的序列，将新序列记录到tmp中，最后把tmp的序列复制到res中
            if(res[i] != now_char) {
                //把之前记录的填充到tmp中
                if(now_count != 0) {
                    //now_count个now_char
                    tmp[now_index++] = now_count + '0'; //数字转字符
                    tmp[now_index++] = now_char;
                }
                
                now_char = res[i]; //当前的值
                now_count = 1; //当前的次数
            } else {
                now_count++;
            }
        }
		//把最后一段加进去
		 tmp[now_index++] = now_count + '0';
		 tmp[now_index++] = now_char;
		 
		 
        //这一轮刷完了吧tmp同步到res中，进行下一轮同步
        strcpy(res,tmp);
		now++;
        recursion(target, now, tmp, res, size);
    }
}

char* countAndSay(int n) {
    if(n == 1)return "1";
    int i = 0,size = 1000;
    char *res = (char *)malloc(sizeof(char)*size);
    char *tmp = (char *)malloc(sizeof(char)*size);
    memset(tmp, '\0', size);
    memset(res, '\0', size);
    res[0] = '1'; //起始值
    recursion(n, 1, tmp, res, size); 
    
    return res;
}



int main(int argc, char **argv)
{
	int n = 0;
	printf("请输入您的目标数:__");
	scanf("%d", &n);
	if(n <= 0) {
		printf("输入错误");
		return 0;
	}
	printf("当前目标为%d\n", n);
	char *s;
	s = countAndSay(n);
	printf("值为%s", s);
	getchar();
	return 0;
}






