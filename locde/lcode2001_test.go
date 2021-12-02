package locde

func interchangeableRectangles(rectangles [][]int) int64 {
	recordsMap := make(map[int]map[int]int64)
	for _, v := range rectangles {
		g := gcd(v[0], v[1])
		r1 := v[0] / g
		r2 := v[1] / g
		if _, ok := recordsMap[r1]; !ok {
			recordsMap[r1] = make(map[int]int64)
		}
		recordsMap[r1][r2]++
	}
	var ans int64
	for _, v := range recordsMap {
		for _, vv := range v {
			ans += vv * (vv - 1) / 2
		}
	}
	return ans
}

// 最大公约数
func gcd(a, b int) int {
	if a < b {
		return gcd(b, a)
	}
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}

// class Solution {
// public:
// int gcd(int a, int b) {
// if (a < b) return gcd(b, a);
// if (a % b == 0) return b;
// return gcd(b, a%b);
// }

// unordered_map<int, unordered_map<int, long long>> cnt;

// long long interchangeableRectangles(vector<vector<int>>& rectangles) {
// for (auto r : rectangles) {
// int c = gcd(r[0], r[1]);
// r[0] /= c;
// r[1] /= c;

// cnt[r[0]][r[1]]++;
// }

// long long ans = 0;

// for (auto iter = cnt.begin(); iter != cnt.end(); iter++) {
// for (auto i = iter->second.begin(); i != iter->second.end(); i++) {
// ans += i->second * (i->second - 1) / 2;
// }
// }

// return ans;
// }
// };
