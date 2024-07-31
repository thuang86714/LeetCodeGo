package leetcodego

//credit to P_A_R_U_S
type SmallestInfiniteSet struct {
    removed map[int]bool
    curMin int
}


func Constructor2() SmallestInfiniteSet {
    return SmallestInfiniteSet{
        removed: map[int]bool{},
        curMin: 1,
    }
}


func (this *SmallestInfiniteSet) PopSmallest() int {
    ans := this.curMin
    this.removed[this.curMin] = true
    this.curMin++
    for this.removed[this.curMin] {
        this.curMin++
    }
    return ans
}


func (this *SmallestInfiniteSet) AddBack(num int)  {
    if _, exits := this.removed[num]; exits {
        this.removed[num] = false
    }
    if num < this.curMin{
        this.curMin = num
    }
}


/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */