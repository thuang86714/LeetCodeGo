package leetcodego
import (
	"strconv"
)
func compareVersion(version1 string, version2 string) int {
    //TC O(n), SC O(1), two ptr
    idx1, idx2:= 0, 0
    for idx1 < len(version1) || idx2 < len(version2){
        tempStr1, tempStr2 := "", ""
        for idx1 < len(version1) && version1[idx1] != '.'{
            tempStr1 += string(version1[idx1])
            idx1++
        }
        var tempInt1 int
        if tempStr1 == ""{
            tempInt1 = 0
        }else{
            tempInt1, _ = strconv.Atoi(tempStr1)
        }

        for idx2 < len(version2) && version2[idx2] != '.'{
            tempStr2 += string(version2[idx2])
            idx2++
        }
        var tempInt2 int
        if tempStr2 == ""{
            tempInt2 = 0
        }else{
            tempInt2, _ = strconv.Atoi(tempStr2)
        }
        
        if tempInt1 > tempInt2{
            return 1
        }else if tempInt1 < tempInt2 {
            return -1
        }
        if idx1 < len(version1){
            idx1++;
        }
        if(idx2 < len(version2)){
            idx2++;
        }
    }
    return 0
}