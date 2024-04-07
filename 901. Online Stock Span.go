package leetcodego

type Pair struct{
    First int
    Second int
}

type StockSpanner struct {
    priceSlice []Pair
}


func Constructor() StockSpanner {
    return StockSpanner{priceSlice: []Pair{}}
}


func (this *StockSpanner) Next(price int) int {
    span := 1
    for i := len(this.priceSlice) - 1; i >= 0; i--{
        if this.priceSlice[i].First <= price{
            span += this.priceSlice[i].Second
            this.priceSlice = this.priceSlice[:len(this.priceSlice) - 1]
        }else{
            break
        }
    }
    this.priceSlice = append(this.priceSlice, Pair{First: price, Second: span})
    return span
}


/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */