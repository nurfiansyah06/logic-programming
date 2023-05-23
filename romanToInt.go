package logicprogramming

func romanToInt(s string) int {
    symbolMap := map[byte]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }
    
    var result int

    for i := 0; i < len(s); i++ {
        if i != len(s)-1 && symbolMap[s[i]] < symbolMap[s[i+1]] {
            result -= symbolMap[s[i]]
        } else {
            result += symbolMap[s[i]]
        }
    }

    return result
}