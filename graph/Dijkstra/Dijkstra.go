package Dijkstra

import (
	"bufio"
	"container/list"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

//INF 无穷大
const INF = 0xfffff

// EdgeType 边的权值类型
type EdgeType int

// VextexType 顶点类型定义
type VextexType int

// VextexDataType 顶点值类型定义
type VextexDataType int

//EdgeNode 边的节点
type EdgeNode struct {
	Weight EdgeType   //权值
	V      VextexType //指向储存该顶点的下标
	Next   *EdgeNode  //指向下一条边
}

//VextexNode 顶点节点定义
type VextexNode struct {
	data      VextexDataType //顶点的值
	FisrtEdge *EdgeNode      //该顶点指向的第一条边
}

//Graph 图
type Graph struct {
	VNum, ENum int          //顶点数目，边数目
	G          []VextexNode //邻接表
}

//CreateGraph 创建邻接表
func CreateGraph(VNum int) (graph Graph) {
	graph.VNum = VNum
	graph.G = make([]VextexNode, VNum)
	for i := 0; i < VNum; i++ {
		graph.G[i] = VextexNode{}
	}
	return graph
}

//AddEdge 添加边
func (graph Graph) AddEdge(s, t VextexType, weight EdgeType) {
	edge := &EdgeNode{V: t, Weight: weight}

	//添加边到头部
	edge.Next = graph.G[s].FisrtEdge
	graph.G[s].FisrtEdge = edge
}

//BuildGraph 通过读取文件建图
//文件格式要求:
//顶点个数 边数
//顶点v1 顶点V2 边的权重
//...
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
			graph.AddEdge(VextexType(s), VextexType(t), EdgeType(weight))
		}
		i++
	}

}

//Dijkstra 算法
//一种求单源最短路径的算法
func Dijkstra(graph Graph, s VextexType, dist []EdgeType, path []VextexType) {
	visited := make([]bool, graph.VNum)
	//初始化
	for i := 0; i < graph.VNum; i++ {
		dist[i] = INF //距离为无穷大
		path[i] = -1  //没有上一个节点
		visited[i] = false
	}
	path[s] = s
	dist[s] = 0

	//使用list实现一个队列操作
	q := list.New()

	//将点s入队
	q.PushBack(s)
	for q.Len() != 0 {
		u := q.Front().Value.(VextexType)
		q.Remove(q.Front())
		//如果该点周围的点已经走过，则无需再走
		if visited[u] {
			continue
		}

		//将该点加入已观察
		visited[u] = true

		e := graph.G[u].FisrtEdge

		for e != nil {
			//这条边下的顶点
			v := e.V

			//如果该点尚未走过，并且当前点的距离加上边的距离小于之前该点的距离，那么就更新该点的距离
			if visited[v] == false && dist[v] > dist[u]+e.Weight {
				dist[v] = dist[u] + e.Weight //更新该点距离
				path[v] = u                  //更新父节点
				q.PushBack(v)                //将该点入队
			}
			e = e.Next
		}

	}

}

//GetPath 通过路径获得到指定目的节点的路径
func GetPath(path []VextexType, t VextexType) ([]VextexType, error) {
	tPath := make([]VextexType, 0)
	for {
		tPath = append(tPath, t)
		if path[t] == -1 {
			return nil, errors.New("不存在到该节点的路径")
		}
		if t == path[t] {
			return tPath, nil
		}
		t = path[t]
	}
}
