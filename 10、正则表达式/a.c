bool isMatch(char* s, char* p) {
    char star;
    int isstar = 0, isspot = 0, i = 0, j = 0, bt = 0; //bt为*的差值
    bool res = true;
    while(*(s+i))
    {  
        
        if(*(p+j) == '*'){
          isstar = 1; 
          star = *(p+j-1);
        }
        
          printf("j是%d--重复字符是--%c--s是--%c\n", j,star, *(s+i));
        if(*(s+i) != *(p+j)){
         res = false;
         if(*(p+j) == '.'){
             res = true;
         }else{
            if(*(p+j) == '*' ){
                if(star == '.'){
                    //之前的是个点
                    res = true;
                }else if(*(p+j+bt) == *(s+i) ||  *(p+j+bt) == '.'){
                     printf("抵消重复\n ");
                 res = true;
                 j+=bt;
                 isstar = 0;
                 star = ' ';
                }else{
                    printf("靠重复\n");
                    //之前的不是个点
                    if(star == *(s+i)){
                        res = true;
                    }else{
                        res = false;
                    }
                }
            }else{
                if(*(p+j+1) == '*'){
                    res = true;
                    j++;
                    i--;
                }else{
                    return false; 
                }
               
            }
         }
        }
      
        if(!res){
            return false;
        }
        
     

        
        if(isstar == 0 || (star != *(s+i) && star != '.')){
            j++;
            isstar = 0;
            star = ' ';
            bt = 0;
        }else{
            bt++;
        }
        
        
        i++;
    }
    
    if(strlen(s) < strlen(p)){
         while(*(p+j+bt)){
        if(*(p+j+bt) == '.'){
            res = true;
        }else if(*(p+j+bt) == '*'){
            if(*(p+j+bt-1)=='.'){
                res = true;
            }else if(*(p+j+bt-1)==*(s+i-1)){
                res = true;
            }
        }else{
            return false;
        }

            j++;
        }
    }
    
    
    return res;
}