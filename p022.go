package euler

/*
Names scores
Using names.txt (right click and 'Save Link/Target As...'), a 46K text file containing over five-thousand first names, begin by sorting it into alphabetical order. Then working out the alphabetical value for each name, multiply this value by its alphabetical position in the list to obtain a name score.

For example, when the list is sorted into alphabetical order, COLIN, which is worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list. So, COLIN would obtain a score of 938 × 53 = 49714.

What is the total of all the name scores in the file?
*/

import (
	"io/ioutil"
	"strings"
	"sort"
)

func P22() int {
    data, err := ioutil.ReadFile("p022_names.txt")
    if err != nil {
    	panic(err);
    }
    names := strings.Split(strings.Replace(string(data), "\"", "", -1), ",");
    sort.Strings(names);

    sum, score := 0, 0;
    for i := 0; i < len(names); i++ {
    	score = 0;
    	for j := 0; j < len(names[i]); j++ {
    		score += int(names[i][j]) - 64; 
    	}
    	sum += score * (i + 1);
    }
    return sum;
}