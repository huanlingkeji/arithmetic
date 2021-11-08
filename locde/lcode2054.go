package locde

/*
class Solution {
public:
    int maxTwoEvents(vector<vector<int>>& e) {
        sort(e.begin(), e.end());
        int n = e.size(), res = 0;
        vector<int> f(n);
        f[n - 1] = e[n - 1][2];
        for(int i = n - 2; i >= 0; i --){
            f[i] = max(f[i + 1], e[i][2]);
        }
        for(int i = 0; i < n; i ++){
            int t = e[i][2];
            int l = 0, r = n - 1;
            while(l < r){
                int mid = l + r >> 1;
                if(e[mid][0] > e[i][1]) r = mid;
                else l = mid + 1;
            }
            if(e[l][0] > e[i][1]) t += f[r];
            res = max(res, t);
        }
        return res;
    }
};
*/

func maxTwoEvents(events [][]int) int {
	maxR := -1
	lth := len(events)
	for i, arr := range events {
		if arr[2] > maxR {
			maxR = arr[2]
		}
		for j := i + 1; j < lth; j++ {
			f1, f2 := 0, 0
			if arr[0] < events[j][0] {
				f1 = i
				f2 = j
			} else {
				f1 = j
				f2 = i
			}
			if events[f1][1] < events[f2][1] && events[f1][2]+events[f2][2] > maxR {
				maxR = events[f1][2] + events[f2][2]
			}
		}
	}
	return maxR
}
