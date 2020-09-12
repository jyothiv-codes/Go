func gradingStudents(grades []int32) []int32 {
    // Write your code here
    var i int
    for i=0;i<len(grades);i++{
        var modval int32
        modval=grades[i]%5
        if modval>=3{
            if grades[i]+modval>=40{
                    grades[i]+=(5-grades[i]%5)
            }
            
        }
    }
    return grades
}
