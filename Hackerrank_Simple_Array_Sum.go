func simpleArraySum(ar []int32) int32 {
    /*
     * Write your code here.
     */
     var i int
     var sum int32
     for i=0;i<len(ar);i++{
         sum+=ar[i]
     }
    return sum
}
