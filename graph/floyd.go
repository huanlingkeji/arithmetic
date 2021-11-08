package graph

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

// EdgeType 边的权值类型
type EdgeType int

//INF 无穷大
const INF = 0xfffff

//Graph 图
type Graph struct {
	VNum, ENum int          //顶点、边的个数
	G          [][]EdgeType //邻接矩阵
}

//CreateGraph 创建一个图并且初始化邻接矩阵
func CreateGraph(n int) (graph Graph) {
	graph = Graph{VNum: n}
	graph.G = make([][]EdgeType, n)
	for i := 0; i < n; i++ {
		graph.G[i] = make([]EdgeType, n)
		for j := 0; j < n; j++ {
			graph.G[i][j] = INF
			if i == j {
				graph.G[i][j] = 0
			}
		}
	}
	return graph
}

//BuildGraph 创建图
func BuildGraph(path string) (graph Graph) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	buf := bufio.NewReader(f)

	i := 0
	//边的数目
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return graph
			}
			panic(err)
		}
		line = strings.TrimSpace(line)
		data := strings.Split(line, " ")
		if i == 0 {
			n, err := strconv.Atoi(data[0])
			if err != nil {
				panic(err)
			}
			graph = CreateGraph(n)

			graph.ENum, err = strconv.Atoi(data[1])
			if err != nil {
				panic(err)
			}
		} else if i <= graph.ENum {
			s, err := strconv.Atoi(data[0])
			if err != nil {
				panic(err)
			}
			t, err := strconv.Atoi(data[1])
			if err != nil {
				panic(err)
			}
			weight, err := strconv.Atoi(data[2])
			if err != nil {
				panic(err)
			}

			//有向图，如果是无向图再添加一条反向边
			graph.G[s][t] = EdgeType(weight)
		}
		i++
	}

}

//Floyd 求取多源最短路径
func Floyd(graph Graph, dist [][]EdgeType, path [][]int) error {
	for i := 0; i < graph.VNum; i++ {
		for j := 0; j < graph.VNum; j++ {
			path[i][j] = -1
			dist[i][j] = graph.G[i][j]
		}
	}

	//注意，k必须放在最外层，如果放在最里层会过早的确认两点的最短路径
	for k := 0; k < graph.VNum; k++ {
		for i := 0; i < graph.VNum; i++ {
			for j := 0; j < graph.VNum; j++ {

				//找到更短的路径
				if dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]

					//发现负值圈
					if i == j && dist[i][j] < 0 {
						return errors.New("存在负值圈")
					}
					path[i][j] = k
				}
			}
		}
	}
	return nil
}

//GetPathForFloyd 获取路径
func GetPathForFloyd(path [][]int, s, t int) (tPath []int) {
	tPath = make([]int, 1)
	tPath[0] = s
	for {
		s = path[s][t]
		if s == -1 || s == t {
			tPath = append(tPath, t)
			return tPath
		}
		tPath = append(tPath, s)
	}
}
