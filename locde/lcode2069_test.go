package locde

var Dirs = []string{"North", "East", "South", "West"}
var ders = [2][4]int{{0, 1, 0, -1} /**/, {1, 0, -1, 0}}

type Robot struct {
	dir    int
	pox    int
	maxPox int
	poy    int
	maxPoy int
	limits [4]int
}

func Constructor(width int, height int) Robot {
	return Robot{
		dir:    1,
		pox:    0,
		maxPox: width,
		poy:    0,
		maxPoy: height,
		limits: [4]int{
			height, -1, -1, width,
		},
	}
}

func (this *Robot) Step(num int) {
	for i := 0; i < num; i++ {
		if this.poy+ders[1][this.dir] == this.limits[this.dir] {

		}
		this.poy += ders[1][this.dir]
		this.pox += ders[0][this.dir]
	}
}

func (this *Robot) GetPos() []int {
	return []int{this.pox, this.poy}
}

func (this *Robot) GetDir() string {
	return Dirs[this.dir]
}
