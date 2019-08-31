
func maxPoints(points [][]int) int {
    //用斜率试试
    Kmaps := make(map[string][]int) //key存斜率，value第一个元素存第一次发现该斜率的i值（去重复），第二个元素存出现次数
    
    lens := len(points)
    if lens <=1 {
        return lens
    }
    
    selfMaps := make(map[int]int) //记录每个点重合的次数
    
    for i:=0; i < lens; i++ {
        isFind := false
        for k,_ := range selfMaps{
             if points[k][0] == points[i][0] && points[k][1] == points[i][1] {
                selfMaps[k]++
                isFind = true
                break
            }
        }
      
        if !isFind {
            selfMaps[i]=0
        }
    }
    
    if len(selfMaps)== 1 {
      for _, v := range selfMaps{
        return v+1
      }
    }
    
    
    noKMpas := make(map[int][]int) //斜率不存在的情况，x相同，那就用x作为key, value第一个元素为首次发行斜率不存在时的i，第二个元素是已出现的次数
    
    for i:=0;i<lens-1; i++ {
        for j:=i+1;j < lens; j++ {
            if _,ok :=  selfMaps[j];!ok{
                continue
            }
            x:=points[i][0] - points[j][0]
            y:=points[i][1]-points[j][1]
            
            if x == 0 {
                if y == 0 {
                    //这里因为是去重复的，不会有重复的，不会出现到这里
                    
                } else {
                    //斜率不存在的情况
                     if val,ok := noKMpas[points[i][0]]; ok {
                            if val[0] == i {
                                val[1]++
                                val[1]+=selfMaps[j]
                            }
                        } else {
                            noKMpas[points[i][0]] = []int{i,2}
                            noKMpas[points[i][0]][1]+=selfMaps[i]
                            noKMpas[points[i][0]][1]+=selfMaps[j]
                        }
                }
            } else {
                
                k:= getK(x,y)
                
                if val,ok := Kmaps[k]; ok {
                    if val[0] == i {
                        val[1]++
                        val[1]+=selfMaps[j]
                    }
                } else {
                    Kmaps[k] = []int{i,2}
                    //看当前的点钟是否有重合的点,要查看的点为i和j
                    Kmaps[k][1]+=selfMaps[i]
                    Kmaps[k][1]+=selfMaps[j]
                }
            }
        }
        
    }
    
    max := 0
    for _,v := range Kmaps {
        if v[1] >= max {
            max = v[1]
        }
    }
    
    for _,v := range noKMpas {
        if v[1] >= max {
            max = v[1]
        }
    }
    
    return max
}

func getK(x, y int) string{
    if y == 0 {
        return strconv.Itoa(0)+strconv.Itoa(y)
    }
    divis:= divisor(x,y) 
    return strconv.Itoa(x/divis)+strconv.Itoa(y/divis)
}


//Greatest common divisor
func divisor (a,b int) int {
    if a%b == 0 {
        return b
    } else {
        return divisor(b,a%b)
    }
}




/***

大佬的算法总是这么简洁明了 ╮(╯▽╰)╭

// 暴力算法的基础上优化，解决精度问题
// O(N^2)
// 查找表，键为斜率
// 解决斜率精度问题：用最大公约数，3/6==2/4==1/2
// 整数问题，都要考虑0，正负数
// 数组问题一般要考虑数组长度为0，1这两种情况。
func maxPoints(points [][]int) int {
    if len(points) <= 2 {
		return len(points)
	}

	count := 0
	for i := 0; i < len(points)-1; i++ {
		// repeat表示重复的点
		countMap, repeat, max := make(map[string]int, len(points)>>1), 0, 0
		for j := i + 1; j < len(points); j++ {
			a, b := points[i][1]-points[j][1], points[i][0]-points[j][0]
			if a == 0 && b == 0 {
				repeat++
				continue
			}

			c := gcd(a, b)
			if c != 0 {
				a /= c
				b /= c
			}

			key := strconv.Itoa(a) + "/" + strconv.Itoa(b)
			countMap[key]++
			if max < countMap[key] {
				max = countMap[key]
			}
		}

		if count < repeat+max+1 {
			count = repeat + max + 1
		}
	}

	return count
}

// 求a，b的最大公约数
func gcd(a, b int) int {
    if b == 0 {
        return a
    }
    
    return gcd(b, a%b)
}

****/

