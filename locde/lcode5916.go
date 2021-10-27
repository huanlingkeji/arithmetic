package locde

/*
class Solution {
public:
    typedef pair<int, int> pii;
    int minimumOperations(vector<int>& nums, int start, int goal) {
        queue<pii> q;// x, 操作次数
        q.push({start, 0});
        unordered_set<int> st;
        st.insert(start);
        while(q.size()) {
            pii t = q.front(); q.pop();
            for(auto& y : nums) {
                if((t.first + y == goal) || (t.first - y == goal) || ((t.first ^ y) == goal)) return t.second + 1;
                if(0 <= t.first + y && t.first + y <= 1000 && !st.count(t.first + y)) {
                    st.insert(t.first + y);
                    q.push({t.first + y, t.second + 1});
                }
                if(0 <= t.first - y && t.first - y <= 1000 && !st.count(t.first - y)) {
                    st.insert(t.first - y);
                    q.push({t.first - y, t.second + 1});
                }
                if(0 <= (t.first ^ y) && (t.first ^ y) <= 1000 && !st.count(t.first ^ y)) {
                    st.insert(t.first ^ y);
                    q.push({t.first ^ y, t.second + 1});
                }
            }
        }
        return -1;
    }
};
*/

type Pair struct {
	First  int
	Second int
}

func minimumOperations(nums []int, start int, goal int) int {
	var q = []Pair{{
		First:  start,
		Second: 0,
	}}
	var usedMap = map[int]bool{
		start: true,
	} // 关键的去重操作
	for len(q) != 0 {
		// pop的操作曾经出现错误  使用arr实现pop应该是去除第一元素
		t := q[0]
		q = q[1:]
		for _, v := range nums {
			if t.First^v == goal || t.First-v == goal || t.First+v == goal {
				return t.Second + 1
			}
			if 0 <= t.First+v && t.First+v <= 1000 && !usedMap[t.First+v] {
				q = append(q, Pair{First: t.First + v, Second: t.Second + 1})
				usedMap[t.First+v] = true
			}
			if 0 <= t.First-v && t.First-v <= 1000 && !usedMap[t.First-v] {
				q = append(q, Pair{First: t.First - v, Second: t.Second + 1})
				usedMap[t.First-v] = true
			}
			if 0 <= t.First^v && t.First^v <= 1000 && !usedMap[t.First^v] {
				q = append(q, Pair{First: t.First ^ v, Second: t.Second + 1})
				usedMap[t.First^v] = true
			}
		}
	}
	return -1
}
