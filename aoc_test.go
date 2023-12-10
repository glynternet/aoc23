package aoc23

import (
	"bufio"
	"fmt"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

const prodInput = `F-|-J7F7F7F-7F7F--7-L|-L7777.FFL-|.|.FF-FF7FFFFF7--F-J7.|-FF77|-F|777-F7.F7FF|-J7|-7.FFJ77FL.7-|-|7.|.77|-F-|-F|-FL.F77.FF-77|7.7--7--7FJ--7
|7LJ7.F-JLJ7L7--7F77.|LLJ|LF7|.LLJ-|7L|FJLF77L-FJ7FJ.LFF|7L-LLF-LJF777||7L-F7LLLLJ7|L7||F-J|-|.J.F7FF7F7|J|F|-FJLF.FLL|7J|FF-7-FJ7J|J.F7L|LJ
L|J.7.L-J7FJ--..F-J|-7-|.J|LL-J.F|J|-77L|.LL7J||FJ7|.F--J7-J|7.LJJJJFFJL7JF||7.||-LJ-F7-|FFJ-|77.|L7||||7F-LJF|L7J7FF|JL7L7JJL-|---|7-JJ-FJ7
LJ.L-|--JF7J|LF7|L-7.|FFJJ-.|JL7JF7F7F7777F-JFFJJLLJFL-J-|.FLJ-7|FF7FL7FJF7||F7FF-J7F||F7-7L7L|7.L7||LJ|.LJ|FFJFF7F7FLLJLJ||LL-J7.L|7LJJFJF-
.L7J7L7JJF|-||||7LJ|LFL7J-LL|7-|-.F|LJ|F7F7JF|7J7.FFJ|-JF-7F7.F77FJ|F7||FJ||LJL-7.FFFJ|||F|7|-JFF7|||F7|7--J||.JJFFF7J7.JJL-.L7L---LF7|J|-|J
.||||-7FF---J--LJ-F|F|JL|.|.|.F||.-L-7LJLJL7-JJJ--F7FF7-|LFJL-JL7L7||||||FJL-7F-JFJJL7LJL7F--7JFJLJLJ|LJJ-FF|--J.J-L-7F7J.L7|.J.F||FJ--.|.L-
.-7-J.L7J7.F|JLLJFL--L7FLFL---FJ-7L||L7F---JJ-J.||||FJ|F-7|F---7L-J||||LJL--7|L--7LFFL7F-J|F-J.L--7F-JF77.F.|FF-LJL|LL7|.FL|LJ--7-7J|..----J
|JJJLF.L.--|-FJ7..LLF-7FL7|LJF7J-F-7F7|L--7J.F.FFFJ||FJL7|LJF--JF-7LJ||F----J|F--JF7F-J|JFJ|||-F7FJ|7L||7-J-F-|7|J|LF-JLF-7|-7J.F-J-|F|J.L|.
|J||7L-J.|FLF-FJJ-7LJF|J.J|-.7|F-L7LJLJF--JF-7-||L7|||F7||F7L7F-JJL-7LJL-7FF7||F--J|L-7|FJFJFF-J||FJ7L|L7.L7L7.F|.L-F-|7L-JJ.F7--7L-|-|-7.|F
|--7LJ.|FF7F--JJJ.F.FLJL|-|JFF-|J7L---7L-7||FJJLLFJLJLJLJLJ|-LJF7.F-JF---JFJLJLJF7FJF7||L7|F7L-7||L7F7L7|7FJ77-F7-|J|FF-JJ|7FF--J|.L7FL7J.|J
FJFJ|.|L--L-LFJJ.L7-JF|LL.L.7JLJ.FFF7FJF-JFJL7JFLL-7F7F7F-7L7F-JL-JF7L7F7-L7F---JLJFJ|||FJLJL7FJLJFJ|L-JL7J7.|-F|F7|LFJ||J|FF7.F7J77JFF|F-J|
|.JJ|FFJ.FL7L7J|FL-.FJJ|F--7-.J..LFJ|L7L7FJF-J.7JF-J|LJ|L7L-JL---7FJL7LJ|F7|L-7F7F7L7||LJF---JL7F-J|L7F7FJJFF|-JL7L|J|F-JF--L|F7|LF|7J||7.F7
L|-F-FJ7-7.L7L7F7-LJ7..-JJ-F7JFF|.|FJL|FJL7||F77.L--JF-JFJ7F7.F--J|F7L-7LJLJF-J|LJ|FJ||F-J.F7F7|L-7F-J|LJF77.||--L-L.LJJL|FL|L777.LJJ-JJ7-||
.LFJLJJ7FL7J|FJ-L77|JF-L7F7.F7FF7FJL7FJL7FJL7|L-7F7F7L7FJF7|L7L---J||F-JF---JF-JF7||FJ|L--7|||||F-JL-7L--JL7-L7.LF.L|.|.FLL7J-LL-7LJLFJLF.L|
FJJ7F---7J||JLLLJ|FJ7.LJ77FFF-7||L7FJL-7||F7||F-J||||FJL-J||FJF----J|L-7|F7F7L--J||||FJF--J|LJLJL7F7FJF----JJ||7.F77F.F-L.|7J7LJ.|.|LL|LLJ||
77.FLJ|.F.FL7JJ|FLL7F|FL|JL-L7LJL-J|F7FJLJ||LJL7FJ||LJF---J||.L7F7F7L7FJ||LJL--7F||||L7L7F7L-7F-7|||L7L----7F7-F7||FJ.7.J-7|L7.|F7F-..|LLLLL
|7--JLL---|J.L---7|F-JF-7F7..L----7|||L-7FJL---JL7|L-7|F7JL||F7LJLJL7|L7||F----JFJLJ|FJFJ|L-7||F||||FJF----J||F7FJL7JF77|F-J-.--J.|-J7JF-FJJ
.JLL|LL-JF|F77...|J-|7L-JFL7.F---7|LJ|F7|L-7F7F7FJL7FJ||L7FJLJ|F7F7FJL-JLJL7F7F7L--7|L7|FJF7LJL7LJ||L7|.F7F7|||LJF-J.|L7JJ.J-F-|F|J77L-JF|.|
F|-JL7.FF77|F7-LL.LL|F-JLF--FL7F7LJF-J|||F-J||||L7FJL7||FJ|F--J||||L--7F7F7LJ|||F--JL-J||FJL-7FJF-J|7|L-J||LJ|L-7L---JFJ.|-L7--LJLFF||..|LJJ
FJ7FJ|F--77|.L7J.-7LJLF.LF77LFJ|L-7L7FJ||L-7||||J||F-J||L7|L7F7|||L7F-J|LJ|F-J||L-----7LJL-7FJ|FJF7L7|F--J|F-JF-JF7F7FJ7-F7J|.F|JL-L7FL-J-||
LJFJLL7L||L77L7..F|7-F--7||F7L-JF7|FJL7LJF-J||||FJ|L-7||FJL7||||LJFJL--JF7||F7|L7JF7F-JF---JL-JL-JL7|||F7FJL7|L7FJLJLJF-7||LF-777..F|F|7|J|J
.FJ7-FJ.7-|F-JLF-FF7.L-7|||||-F-JLJL--JF7|F7||||L7L7FJLJL7FJLJ|L-7|F7F-7|LJ||LJFJFJLJF7||F7F7F7FF7FJLJ||LJF-JF7||F7F-7|FJ|L7L7|LJ7F7FLLJJ-|.
F.||-J|.|LFL7.FF-J||F7F|LJLJL7L--7F7F-7|LJ|||LJL7|FJL7F--JL7F7|F-JLJ|L7|L-7LJF7L7L---J|L-JLJ||L7||L-7FJ|F-JF7|||LJLJFJ||FJFJFJL7.F--7FFJ|LL7
L7|-7FL-77.FL-F-7FJLJL-JF7F-7L7F7LJLJFJ|F-J||F--J|L7FJL7-F7||LJL--7FJFJL--JF-JL-J|F---JF7F--J|FJ|L7FJL-JL-7|||LJF--7L7||L7L7|F7|JL|F|7|77||F
FLJFLLLFFJ7|7L||FJF-----JLJ.|FJ||F7F7L7LJF7LJ|F--JFJL7FJFJ||L7-F7FJL7L-7F-7L7F---7L----J|L7F7|L7L7|L7F7F--J|||F-JF7L-J||7L7|||LJJF|-JF|-JFJ-
F|7|7J.||77LJ--LL-J-F7F--7F-J|FJLJ|||7L7FJ|F-JL--7|F7||||FJL7L7||L-7L--JL7L7||F--JF7F---JFJ||L7|FJL-J|LJF7J|LJ|F-JL--7||F7|LJL--7J.|LJ.FF-JJ
FJFJ|7-.F77J-7FLF7LFJ|L-7|L-7|L--7LJ|F7||L|L-7F-7|||LJL7||F7L7||L77L---7FJJ|LJL-7F||L---7L-J|FJ|L---7L77|L7L-7LJF----J|||||F----JJ-J.|-L--F7
||-7|7-LLJ7.-JFFJL7L7|F7||F-JL7F7L-7||LJL7L--J|FJ||L-7FJ||||7|||FJF7F7FLJF7|F---JFJ|F7F7|F--J|FJF--7L7L-JFJF7|F-JJF--7|LJ|||F7F77|JFF.|J-LL7
7J7F|LFLJJJ.FF-JF7L7||||||L-7FJ|L7FJ||F--JF7F7||FJL7FJL-JLJL7|LJL7|LJ|F7FJLJL---7|FJ|LJ|||.F7|L7|F-JFJF--JFJLJ|FF7|F7|L-7LJLJLJ|F77|.7--7-|J
|.-LJF|FJ||L-|F-JL7|||||||F7|L-JFJL7|||F-7|||LJ||F-JL-7F----JL-7FJL-7||LJF-7F7F7||L7L-7LJL7|||FJ|L-7L7L7F7|F--JFJ|||LJF7|F7F--7LJL7JF-7-|J||
-7JF7-JJ.F77LLJF--JLJ||||||LJF-7L--J||||FJ||L-7|||F7|FJ|F7F7LF-JL7JFJ||F-JJ||LJLJ|FJF7L7F-J|LJL7|F7|FJFJ||||JF-JFJ|L-7||LJ||F7L---JL-||-77.-
||.L.-7||F-JJ|.L----7|||||L-7|JL---7||||L7|L-7|||LJL7L7||||L7L-7FJFJFJLJF7FJL---7||FJ|FJL-7L-7FJLJ||L7|FJLJL7|F-JL|F-J|L7F|LJL-7F--7-||.|J77
F77.L-|-LL.|-FF-----JLJLJ|F7LJF7F7J||||L7||F-J||L--7L-JLJ||FJF-JL-JFJF7FJLJF----J||L7|L7F-JF7|L-7L||FJLJF---J||LF-JL-7L7|FJF7F-J|F-J7L7FJ.LF
L-F.L7|JL|7FF7L--7F7F-7F7LJL--JLJ|FJLJL7|LJL7FJL-7LL-7F--J|L7L----7L7||L--7L-7F7FJL7||FJL-7|||F-JFJ|L-7FJF7F7|L-JF---JFJLJFJ|L--JL7-7.-|.F|J
FLL7L-|7F--FJL--7LJLJ7LJL7F7F---7|L---7|L-7FJ|F--JF77||-F7|FJF7-F7L7|||F7||F-J||L7FJ||L7F-J||||F7L7|F7||FJLJLJF--JF7F7|F7FJ.L7F-7FJF|7|F.F7J
F7FL7L|--7FL---7L7F-7F---J|LJF--JL7F-7|L7FJ|FJ|F7FJL7|L7|||L7|L7|L7|LJ|||FJ|F7||FJL7||FJ|F7|||||L7|LJ||||F---7L7F-J|||||LJF7|||FLJF7F77J.LJ7
L-|-L7JL7FJ7LJFL7LJFJL----JF7L7F-7LJ-|L7||FJ|FJ|||F-JL7||LJ|||-||FJ|F-J|||-LJ||||F-J|||-||||||||.|L7FJ||||F7|L7LJF-J|LJL7FJL7|L7F7|LJ||.F-LF
JL|.L7FL7JF||LF7L-7|F7JF---JL-J|FJF-7|FJLJL-J|FJLJL7F7||L-7FJL7LJL7||F7||L-7FJ|LJL7FJLJFJ|||LJLJFJFJL7|LJLJL7FL--JF7|F-7LJF7||FJ|||F-JJFFJL|
|L7F7J--F7LLF-JL--JLJ|FJF------JL7|FJLJF7LF--JL--7|||LJ|F7||F-JF--J|||LJ|F7|L7|F--JL--7L-J|L-7F-JFJJFLJF----JF7F--JLJL7L7FJLJ|L-JLJ|JF|7|-FJ
L-.|..JFJJ|LL-------7|L-JF7F--7F-J|L---J|FJF-7F7FJFJL-7|||||L7FJF-7||L7FJ|LJ7LJL7F7F--JF--JF7||F7L-7F--JF----JLJF7F---J.LJF-7L--7F7L-7JFL-JJ
FL-JF7FJJ.|7F7F----7|||F7||L7FJL--JF-7F7||FJ.||LJFJF7FJ|||||FJL7||LJL7|L7L---7F-J||L--7L--7|||LJL7FJ|F-7|F-----7||L-7F7F77|FJF77|||F7L--7F-.
FLJ77|J7JF--J|L---7LJL-JLJL7||F-7F-JLLJLJLJ-FJL-7|FJ|L7|||||L7FJL--7FJ|FJF7F-J|F7||F-7L7F-J||L-7FJ|7LJ|LJL7F7F7LJ|F7LJLJL-JL-JL7||LJ|F--J7.L
JJ.L7JFFFL--7L----JF7F7F7F7LJ|L7|L---------7L7F-JLJJL7||||||FJ|F7F-JL7||FJ|L7FJ|||||FJFJ|F7||F7|L7|F----7FJ|LJ|F7LJ|F-7F-------JLJJFLJ|J|L-.
||F-|.JLF---JF----7|LJLJLJ|F7L-JL----7F7F-7|FJL-----7|||||||L7||LJF7FJ||L7L-JL7||||||FJFLJ||||||FJ|L---7|L-JF7LJL-7LJL|L--7LF7.F7.L|F|JL|7JF
LL|7|7..L-7F7|F---JL7F----J||F7F-7F-7|||L7LJL7F7F---JLJLJ||L7|||F-JLJFJ|FJF---J|||||||F---J|LJLJL-J|F--JL---J|F-7FJF-7L7F7L-JL7||7JLJJ.FJL7J
|.-L7-77LLLJLJL7F---JL7F-7FJLJLJ-LJFJLJL-J.F-J||L----7LF7LJJLJLJL--7FJ7||L|F7F7|LJ||||L-7F7|F---7F-7L-------7LJ-LJ-|FJJ|||F7F7LJ||.FJ.-JJ|.F
|-J-FJFJ77LF---J|F--7.LJ.LJF7F7F7F-JF7F---7L--J|F---7L-JL-7-F7F----J|F-JL7LJLJLJF-J|LJ|FJ||||F-7|L7L7F7F----JF7F---JL-7LJLJLJL--J7F-F-7LL7-J
|-L-JF7F77JL-7F7|L-7|F7|F7FJLJLJ|L--JLJF--JF--7|L-77|F7F-7|FJLJF-7F7|L7F7L-----7L--J-F7L-JLJLJFJL7L7|||L---7|||L7F7F--J7F7F7JF7F---7-JJ.FJLJ
L-7|LJ-7L-F--J|||F-J|||FJLJF---7|F--7F7|F-7L7FJL-7L7|||L7||L7F7L7|||L7LJL7F-7F7L-7F--JL-7F7F7FJF7L-JLJ|F--7|FJ|FJ||L-7F7|LJL7|LJF7FJJF|.F7FJ
FLF7-L.-..L---JLJL-7|||L-7FJFF-JLJF7LJLJ|FJFJL--7L7||||FJLJ-LJL7|||L-JF7FJ||LJ|F-JL7F--7||LJLJFJL----7|L-7LJ|FJL7|L--J|LJF-7LJF-JLJJ-777.FF7
LJJL7.-JF7F-7F7FF7FJLJL-7LJF7L-7F-JL--7FJL-JF---J7|||||L-7F7F--J|||-F-J|L7L--7|L7-LLJF7||L---7|F7F---JL-7L--JL--JL7FF7|F-J-|F7L---7.||JF-JJL
.LL-JLJ.|FL7|||FJLJF-7F7L--JL--J|F---7LJF7F7L----7||LJL--J|LJF--JLJFJF7L-JF7FJL7L7F--JLJL----JLJ|L-----7L-7F7F7F-7L-JLJL--7LJL---7|J-F-JJ7.L
77.||L-|7J||LJ|L7F7L7LJL7F-7F7F7|L-77L--JLJL7F---J|L7F7F7JL-7L7F---JFJ|F-7|||F-JFJ|F--7F-----7F7|F-----JF7||LJLJ.|F7F7F---JF7F7||LJJ-|.LF-7.
.7-LL7F|||FJF7L-J|L7|F--J|FJ|LJLJF-JF-----7|LJF--7L-J|||L-7||FJ|F7F7|FJ|.||||L-7L7||F-J|F----J|LJL-----7||LJF7F-7LJLJLJ|F--JLJL--7J|FJLFJ-J7
F.FLJ||LL7L-JL---J|LJL---JL7|F---JF7L7F---JF77|F-JF7F|||F-JFJL7|||||||FJFJ||L7FJFJLJL7FJL7F7F7|F------7LJL-7|||FJLF7F7F7|F-------J-F77-JJF-|
--7-F-77L|L|FLF----------7-LJL7F--JL-JL----JL7|L--J|FJLJ|F7L--J||LJ||||JL7|L7|L7L-7F-JL7LLJ||||L7F-7F7L-7F7LJLJL--JLJLJLJL7F7F-7F7F||F7FJ.-J
F7J7F-7-FL.|L.L7F7F7F--7FJF7F7LJF-7F7F7F---7FJ|F---JL7F-J|||F-7LJF-J|LJFL||.||FJF7||F7FJF--J|||FJL7LJ|F7LJL--7F7F-7F------J|||FJ||FJLJ|7JF-J
F|-7|-L.|JF7|F-J|||LJF-JL-JLJ|F7L7LJ|||L--7LJFJL7F7FFJL--JL-JFJF-JF7L7-|LLJ-||L7|LJLJ|L7|F-7|LJL7FJF-J|L-7-F7LJLJJ|L-------JLJL7|||F--J|F|7|
L77JF7.7777LFL-7|LJ-FJF--7F-7LJL7L-7LJ|F-7L--JF7LJL-JF7F7F---J|L--JL7|JL.|LJ||J|L--7FJFJLJ7||F-7||-L--JF7L-JL-7F-7|F7F7F7F-----J||||F7.L-LJJ
LL|7-|-|LF-7-.FJL7F-JFJF7LJLL--7|F-JF7LJ7L--7FJL7F---JLJLJJF7F7F77|L||-|FL--LJFJF7FJL-J7LF7LJL7LJ|F----JL-----J|FJLJLJLJ|L------JLJLJ|-F|||J
FJJ|-LL|-L7|FFL--J|F7|FJL-----7||L--JL7LF-7L|L-7|L---------JLJ|||F7.LJJ.F|-||JL7|||F7FF--JL7F7L-7|L--7F-----7F-JL-7F7F-7|F7F-7F-7F7F-J7--7|7
L-777...FFS|7F7JF7LJLJL------7LJL-----JFJFJFJF-J|F7F7F-7F7F7F7LJLJ|.JJFL-JFJ77||||LJL7L---7LJ|F7LJJF7LJF----J|F---J|LJFJLJLJJLJLLJ|L7J.J.--J
.LF7J..FFL7L7||FJL7-F----77F7L---------JFJFJFJFFJ|LJLJ-||LJ||L7F-7|-J.77FFFJJFFJ|L-7FJFF-7L-7LJL---JL-7|F7F--JL----JF7L7-F7JF----7L-J77JF7.F
J.|L7-|7|||FJ||L-7|FJF--7|FJL-----7F7F7FJ7L-JF7L7L7F7F7|L-7||FJL7LJ-|J|FFJ|L7LL-JJJLJ|FL7L-7|F-7F----7|LJLJF7F------JL-JFJL-JF---JJF7F7J||-|
|F7.J.|LLFJL-J|F7|LJFJF-JLJF-7F7F7LJ||LJFF7F-JL7L7LJLJ||F-JLJL-7L-7J|LJ-L-J7|.LLJ..|L-F-JF7LJ|FJ|F7.FJL----J|L-7F-7F----JF---JJF-7FJLJL777.J
FLL7F|7-FL--7FJ||L-7|FJF7F-J-LJLJL-7||F--JLJF-7L7L7F--JLJF7.||FJF-J-7.L7L--J.|J.F-FL.L|F-JL7FJL-J|L-JF-7F--7|F7LJFJ|F----JF77F7|FJL7F--J7L7J
7J|.F-|.FF--JL-JL--J||FJLJF7F7F7F7FJLJ|F7F-7L7|FJFJL----7|L7.FL7L7.|L..JJ777L7.LF7-L7-LJF7-LJF7FFJF--J7LJF-J||L7FJFJL----7|L-J||L7FJL----7|.
|F7.L7L-FJF7F-7F7F--JLJF--JLJLJLJ|L---J||L7L-JLJ-L7F----J|FJ7.FJFJF|-77|F7L-7|-|||LF|FLFJL---JL7L-JF7F7F7L7FLJLLJ.|F-----J|F--J|FJL7F7F--JJ7
|LJF.|L-L-JLJ7LJLJF7F-7L7F------7||F---JL-JF--7|F7LJF7F7L|L7-.L-JLFL7|J-LJ7.F7-FJ|7F7F-JF---7F7L---JLJ||L-JF7F7F7FJL---7F7||F--JL-7||LJJ|7LJ
|LL7.J||.FF-----77||L7|FJ|F-----JL-JF----77L-7L7||-FJLJL7|FJFF7|-L7-|JL-7LF7J.L|FJFJ|L-7|F--J|L7F7F--7|L---JLJLJLJF----J|||LJF-7F7LJL7LF|7.|
-77LF-F77LL----7L-JL-J|L7|L7F--7F7F-JF---JF-7L7||L7|F---J||-F7-|.7L-.77L|-J|F-J|L-JFJF7LJL---J7LJLJF7|L---------7FJF--7||||F-JJLJL-7FJ--FJ-L
|JL-7FL-|LF----JF----7|-||FJ|F-J|||F7|F77FJFJFJ||FJ||F7F-JL-J|JL-|7-|-L.|LFLJ.F|F-7L-J|-F7F7F7F7F7FJLJF7F7F7F--7LJFJF7L7|LJL----7F7||JJ-L--J
-7|LL7J.LFL-7F7FJF---JL7LJL-JL--JLJ|LJ|L7L7L7L7||L-JLJLJF7F-7|L7JLJ.LLJ-LJ.-7F-LJLL7F7L7|LJLJLJLJLJF--JLJLJ||F-JF-JFJL7LJF7F7F--J||LJJ....L-
L77.||FJ-J.LLJLJ.L----7L----7F7.F7FJF-JFJ7|FJFJLJF------JLJL||FJJFFJ.||J|FL-F-7-L--LJL7LJF-------7FJF7F7FF-J||F7|F-J.FJF7|LJLJLF7|L7J7F-.F7|
LL--JF-7|JFLF------7F7L----7LJL-J|L-JF7L--JL-JF-7L-7F7F7F--7LJJ--L-7F7-7F|JF|.FF7--JF-JF7|F----7FJ|FJLJL7L--JLJLJL--7L-J|L--7F7|LJFJJL7|.|L7
L|.|L|L|J-J7L-7F7F7LJL----7L----7L---JL---7F7FJFJF7LJLJLJF7L7JJ-LJ7L7.L|7||F-L7.L7JFL-7|LJL7F-7LJFJL7F-7L-7F7F7F7F--JF-7L7F7LJ||F7L-7.F7.LJ.
.|F77L-J77|LF-J|LJL7F7F-7FJF7F-7L----7F--7|||L7L7|L7F-7F-JL-J.|-JJFJLFJF7J-|FJ77L|FF7FJ|F7J|L7L7FJF-J|FJF7LJ|||||L--7|FJFLJ|F7|||L-7|-7|7.|.
LFLLLJJFF-LJL7FJF77LJLJF||FJLJFJF---7LJF-JLJL-J-LJJ|L7LJF---7F7|.-JJF|-||7-LL|LL-F-JLJFJ||FJFJFLJFL7FJL-JL7|LJ||L7F7LJL77F-J|LJ|L-7LJ.LL|-|7
.|L||-F|J|J-LLJF||F-----J|L--7L7L--7L-7L7JF7LF7F7F7L-JLFJF-7||L7-|J7FF7|L777.|7||L-7F-JFJ|L-J-F--7|LJ7F---JF-7||JLJ|F-7L7L-7L--JF-JF|F7.F..F
|LFLJF-777JLLF-7||L---7F7L---JFJF-7L-7L7L7||FJ||LJL7F--JFJ-LJL7L-7|FF||L7|F--7J-FF-JL7|L7|F7F7L-7L7F-7L----JFJLJF7LLJ-L-JF7L7F7FJ-LL-7..|-FF
|-J.F|FJF7|.LL7|||F7F7LJL--7F7L7L7|F7L7L-J|||FJ|F-7LJF-7L-7F7FJF-J7-FJ|FJ||F-JJ.FL-7FJF7|LJLJL-7L7|L7L7F---7L---JL-------JL7|||L-77L-..-J-FL
-.L7FJL-JL7-|FJLJLJLJ|FF7F7LJL7L-JLJL-J.F7|LJL7|L7L--JL|F7LJLJFJF7F7L7||FJ||7FL7.LLLJFJLJF--7F7L-JL7L7LJF7-L7F---7F7F---7F-JLJ|F7|7F|F7JJ.J|
L7-FL7F7F7L7FJF-7F--7|FJLJL7F7L---------JLJF-7LJFJF7-F7LJL----J||||L7||||FJ|F7J|||FF7L---JF7LJL-7F7L-JF-JL7FJL-7FLJ|L--7|L7F77LJLJF-7-J-F-.L
FF-LFJ|LJL7LJFJFJ|F-JLJF--7LJL---------7F7FJL|F7L-JL-JL-----7F7FJ|L7||||LJFJ|L-7J-FJL---7FJL7F-7|||F--JF7FJL7F7|F7|L-7FJL-J|L7F7|FJFJJL77J.J
||LFL-J.F-JF7L7L7|L7F7FJF-JF7F7F-----7J|||L7FJ||F-7F7F7F----J||L7|L||||L7FJ7|F7|F7L----7|L7FJL7LJ|LJF--JLJF7LJ|LJL--7|L7|F-JFJ|L7|FJJ--|F-|.
.LFJ-|LFJF7|L7|-LJFJ|||FJF7|||||F---7L7LJ|FJL7|||J||LJLJLF7F7||FJL-J|||FJL7FJ|LJ||7LF7J|L-JL-7L--JF7L----7|L-7L-7F7FJL7L7|F-JFJFJ||F77F|L7L7
7.-JLL|L-J|L7||F--JFJLJL-JLJLJLJL--7|FJF7LJFFJ||L7|L-7F--JLJLJ||F-7FJ||L-7|L7L7FJL7FJL7L---7FJFF7.|L-7F-7LJF-JF7LJ||F7L7LJL--JFJ.|LJL---77-J
LJJF7-|FL.|FJ||L---JJF7FF7F-7F-7F--JLJFJL--7|FJ|FJ|F7|L7F7F7F7|LJFJL7||F7|L-JFJ|F7|L-7L-7F7||F-JL7L-7LJFJF7L7F|L-7|LJL7|F7F---J-FJF-7F7FJJ.|
.LLJJ.L7F-LJ7LJF7F---JL7||L7|L7|L-----JF7F-JLJ-LJ7LJLJFJ|||LJLJF7L-7||LJ|L7F-JF||LJF-JF7LJLJ|L7F7L-7L-7L-JL7L-JF-JL-7FJLJ|L---7FJFJLLJ||J.F7
F.FJ.J---J||7F-J|L----7||L-JL-JL--7F7F-J|L------7F7F7-L-J|L-7F-JL--JLJF7|FJL-7FJL7-L--JL---7L7LJL-7|F7L--7FJF--JJF-7LJF-7L----JL7L-7L|LJLLLF
|-7.LJ.|.LL|FL-7|F7F--JLJF7F7F-7F7LJLJF7L----7F7||LJ|F7F-JF-JL------7FJLJL-7FJL7FJF7F7F7JF7L7L7F--JLJL---JL-JF7F7L7L7J|FJ.F7F---JF7L77LJ|.F|
||F-77-777FJJF-JLJLJF7F-7|||LJFJ||F7F7||F7F77||LJ|F-J||L-7L-7F--7-F7|L7F7F7||7|||F|LJLJ|FJ|||FJL----------7F7|LJL-JFJFJL7FJ||F-7FJL-J7J7L7LJ
LFL7||.|-F7JJL------J||FJ|LJLFJFJLJLJLJLJLJL7LJF7||7FJ|F7|F-J|F7L-J|L7LJ||||L7FJL7|F7F-JL7|FJ|7F7F--------J|||F----J|L7FJL7||L7|L--7.J.|FJ.|
F-7J-J-FJFJFF7F7F77F7LJL-JF-7L-JF----------7L7FJLJL7L7||LJL7FLJL--7L7|F-J|||FJ|F-JLJ|L-77||L7|FJ|L--------7|LJL---7F7FJL77|||FJL---J7JFLJF7-
LLJJJ|..FF--JLJLJL-JL-----JFJF7-L---------7|FJL---7L-J|L--7L-7LF7LL7||L-7|||L7||F7F-JF-JFJL-J||FJLF7F7F---J|F-----J|LJF-JFJLJL-7F7JF7F7.L7J|
F|F|F--7-L--7F-7F---7F7F--7L-JL7F----7F7F-J||F-7F7|F-7L-7.L-7|FJL--J||F-J||L7LJLJ|L-7L-7L7F-7|||F7||||L7F77|L-7F7F7|F-JF-JF-7F7LJ|FJ|||77|.|
-7|FL|7L|L7FJL7||F--J|LJF7L---7|L---7||LJF7LJL7LJ||L7L--JF7FJ|L-7F-7LJL--JL7L-7F-JF7L7FJ-LJFJ||LJ||||L7LJL-JF7LJLJLJL--JF-J.||L--J|FJ||LF|-|
LFLJJF--FJFL7FJLJL---JF-JL----J|F7F-JLJF7|L--7|F-J|FJF7F7||L7|F7LJJL--7F-7FJF-J|F-J|FJ|7FF7L7LJF-J||L7L---7FJ|F--7F--7F7L--7LJF---JL-JL7-|F|
|L|77J7.|LF7LJF7F----7L--7F7F-7|||L7F-7|||F--J|L7FJL7|||LJ|.||||F7F--7|L7LJ7L-7LJF-JL7L7FJL7L7FJ-FJ|FJF--7LJFJL-7|L-7||L7F7L--JF7F7F7F-JJJL7
|FJL-.|7-FJL7L||L---7L7F7||LJFJLJL-J|FJ||||F-7L7|L7FJ||L7FJFJLJLJ|L7FJ|FJ.F---JF-JF7|L7|L-7|.||F7|FJL-JF7L-7L---JL--JLJFJ|L7F--JLJ||LJ.L7JF-
F|FJJ-L7FL-7L-JL---7L7LJ|LJF7L7F----JL-JLJ||FJFJL7|L7||FJ|LL-7F--JFJL-J|F7L---7|F7||F7||F-JL7|||||L-7F7|L--JF--7F7F7F7-L7L7|L----7||JJ-|--|J
LL--J.|J-|-L7F--7F7L-JF7L7FJL-J|FF7F7F7F--J|L7L-7LJFJ|||FJF7FJL7F7L-7F-J|L7F77||||||||||L7F7||LJ||F-J||L--7.|F-J|||||L-7|FJL-----JLJ..FJ.|.F
F7|LLL|7FL7J||F-J|L-7FJL7LJF7F7|FJLJLJLJJF7|FJF-JF-JF||||F||L7FJ||F7|L-7|FJ||FJLJLJ|||||FLJ||L7FJ|L7FJL---JFJL--JLJLJF7||L-7F77FF7J.7F||.FFJ
L--F-7LL-.|JLJ|F-JF-J|F7L-7|LJ||L-----7F7|||L7L-7|F-7||||FJL-J|FJ||LJF-J|L7||L-7F--J|||L7F7|L7|L-JFJL------JF7F------JLJ|F7LJL--JL777-J-F-F.
||.|L-.LL-|-F-J|F-JF7||L--J|F-J|F-----J||||L7|F7||L7||||||F---JL7|L7FJF7L7|||F7||F77|||FJ||L7|L7F-JF7F7F7F-7||L--------7LJL7F----7L7--|-L.|J
FJ7JL|F-|7|LL-7|L--JLJL7F-7|L--JL--7F7FJLJ|FJ||||L7|||||||L--7F7|L7||FJL7|||||||LJ|FJ|||FJL7||FJ|F7|LJ|||L7LJ|F--7F7F7FJ.F7|L---7|FJJFLJ|FLJ
7L7.FF7-FF7LLL||F--7F--J|FJ|F-7F---J||L--7LJ-LJLJFJ||||||L7F-J|LJF||LJF-J|||LJ|L-7|L7||||F-J|||FJ|||F-J|L7L7||L7LLJLJLJF-J|L--7FJ|L7|FJFFLJ7
||.F--|-7.-J|LLJL-7|L---JL7|L7|L7F-7||F-7L7F-----JFJLJLJ|FJL-7L--7|L-7|F7|||F-JF7||FJ||LJL7FJ|||FJ||L-7L7L7L7L7L---7-F-JF-JF-7|L7L-J-JL||||7
F|-7.LL7LFLFFJJLF-JL7|F7F7||FJL7LJFJ|||7L-JL---7F7L-7F--J||F7|F--JL7FJ||||||L7FJ||||FJL-7|||FJ||L7|L7FJFJ.|FJJL-7F7L7L7FJF-JFJ|FJJF|-F.LF|J7
.J-7FL|7.7-LJF-JL--7L7||||LJ|F7|.FJFJ|L----7.F-J|L7FJ|F-7|FJLJL--7FJL7||||||FJL7||LJ|F7FJFJ||FJL-J|FJL-JF-JL---7||L7L-JL-JF7||LJJF7F7-77||.|
JF777.FJ.J-|77J-||FL7LJLJ|F7LJ||FJFJFJF7F-7L7L7FJFJ|JLJFJ||F7F7F-J|F-J||LJ||L7FJ|L7FJ||L7|FJ|L7F7L|L7F--JF-7F7FJ|L7|F-7F--JLJF-7FJLJL-77LF77
F7|.F7LF-J..7-F--7F7L---7LJL7FJ|L7|-L7|LJ7L7|FJ|FJFJF7FJFJ|||||L-7|L7LLJ|FJ|FJ|FJFJL7||FJLJJL7||L-JFJ|F-7|.|||L7L7|LJFJL---7JL7||F--7FJ-7J||
L|L77L-F7L7.J|L-7LJL----JF-7|L7L-JL-7|L-7F7LJL7|L7|FJLJFJFJ||||F7|L7L---7|FJL7|L7|F7|||L---7-||L7F7|.LJFJL7LJ|FJ7|L-7L7F7F7L--JLJ|J-LJ.L|7||
L|.LF7..-F|7JL7.L-------7|FJL-JF7F7FJL-7LJL--7|L7LJL--7|FJFJ|||||L7|F---J|||FJL7|||LJ||F--7|FJ|FJ|||F--JF7L7FJL-7|F7|FJ||||F-7F-7L-7FJ77.L7J
-F-LLJ-FJLL7-7F.F------7LJL7F7FJ|||L--7|F-7F-JL-JF----J|L7|FJ|LJ|FJ|L--7FJL7|F7|||L7FJ|L-7||L7|L7|LJL7F7|L-JL7F-JLJLJL7||||L7||FJF-J-J-FJ-JL
|.L.|7.L7-7|F|7-L-----7L---J|||FJ||F7FJ|L7|L--7JFJF-7F7L7||L7L7||L7|F7FJL7FJ||LJLJFJL7|F7||L7||FJ|F--J||L---7|L---7F7FJ||||FJ||L7L7.JJL|7L7.
JJ.|F|7.F--7FJF7F7F7F7L----7||||FJ||LJFL7||F--JFJFJFJ|L7|||FJFJFJFJ|||L7FJ|L||7F--JF-J||LJ|FJ||L7|L-7FJ|F---JL7F-7LJ||FJ||LJ-LJL|FJ7.-7LLF|J
|F7J7||-J-LLJF|LJLJLJL-----J|||||F|L--7FJ||L--7L-JL|FJFJ|||L-JFJFJFJ|L-JL7L7||FJF7FJF7||F-J|FJL7LJF-J|FJL7F--7|L7|F-JLJ-|||JJ.F-JL7J77J.|LL7
-FLJLL|||7.LF-L-7F---------7||||L7|F--JL-JL-7FJF---JL7L-J||F--JFJ-L7L---7L7|||L7|||J|LJ|L-7||F7|F7|F7|L-7LJF7||FJ|L----7|L7J.FJF--JF77|-L7||
L|J.7-L7J--.L-LLLJ7F7F-----J|||L-J|L------7FJ|FJF7F--JF--J||F7FJF7F|F7F7L7||LJL||LJFJF7L7J||LJ|||LJ|||F-JF-JLJ|L7|F7F7FJ|FJ-7L-J|LJJLF-.|LFJ
LJ.L|-LJL7|.|||JJF-JLJF7F--7||L-7FJF7F7F--JL-JL7||L-7JL-7FJ|||L-J|FJ|||L7|LJF--JL7FJFJ|FJFJ|F7LJL-7|LJL-7L-7F-JFJLJLJ|L7||-L|7|.L.|.|L--L-7.
.F-7L-JJ.||.F|..FJF7F7|||F-J||F7||FJ|||L---7-F7LJ|F7L-7FJL7|||F--JL7|||FJ|F7|F-7FJL7L7||FJFJ||F7F-JL-7LFJF-JL-7L--7LFJFJ||77|F|7F7-7F7L7L--.
-7LJ||||7LJ7LL7FL7|LJ|||||F7||||||L7|||F-7FJFJL--J||F-JL7FJLJ||F--7LJ||L7LJ|||FJ|F-JFJ||L7|FJLJLJF-7FJFJFJJF-7L7F7L7|FJLLJ--L-|-7J7F|7|J7|FJ
FJ.LFF7F|7JFJ.F--LJF-J|||||||LJ|||FJ||||-|L7|F-7F7|||JF-JL-7FJLJF-JF7LJLL-7|LJ|FJ|F-J.||FJ|L7F7F-JFJ|FL7|F7|FJ-|||FJ||-|J-|||FJ-LF-7.FJFL77.
F-7-LJ-F----..F-7JFL-7||||||L7FLJ||FJ||L7L-J||FJ|||||FJF--7||F-7L--JL-7F--JL-7|L7||L|-|||FJ-LJ||F-JFJF-JLJLJL-7||LJ7|L-7J7.FJL-F-JFLL|7FLL-7
F||.J.F|||7|.-77F-L.|||LJ||L7L7F7LJL7|L7L-7FJ|L-J|||||FJF-J|||FJF-7F7FJL--7F7|L-JLJ|.FLJ||F---J||F7L7L7F7F7F7FJ|L7-LL7FJ-F|L-L-J--J7LLJ77L-F
FL-77-F|7-L|7F77.L|-FLJ-FJL7L7LJL--7LJ.L7FJ|FJF--J||||L7L-7|||L7||||LJJ-L-LJLJJJ|LJ7J-JJLJL---7||||FJL||||||LJF|FJ..LLJ7-|.F7-L||7.J7J.LL|-J
|JF--7FJ.|LL7F||7.|J|LL-L7FJFJF-7F7|.FJL||7|L7L---JLJL7L-7||||FJL7|L--7FFL7L7|J.F7J.LFJ|7.LF--J||||L-7LJ|||L7J-||77F|JL7.FL777.--7J-JJFF|.|7
|F7JL|J.FJ-.F7LJF7JLLJ|F-JL7L7L7LJ|L-7J7||FJFJF-------JF-J|||||F7||F-7L77-LJLJL7|--7-F7|J-7L-7FJ|||F-JLL|||FJJFLJ--JL7JJF-7LJ-7JFLJL|-LJJ--J
-|JFFJ7-LLJ-J-.LL-..LF7L7F7|7L7L7JL--JLF|||FJFJF7F7F7F7L-7|||||||||L7L-J-7|..-7.|-FFJL7..FJ-L||FJ||L-77JLJ||LJ7-JJ||7|7|J.LJLJFJ..|F77-FJL||
F|F.|-F.FJJ.|.F7.|-F.F7J||LJJ.L7L-7LJJFLLJLJFJFJ|||||||F7|LJLJLJLJL-JJJF|7J|J-|-|.LJ.FL-J77FFJ||FJ|F7||7L-LJ7L7-J.JJ-L|-.F.7LF-7F-FL||.LJ.|7
L-|-77L7F-7-J-J|.|J.-|L-LJL|7LFJF-J.L-|-JJ-LL-JL||||||||||FJ|7|-J-LJ.|FJJF-L7LF7J7LJF-7FL|7F|FJ||JLJLJ-L.||FFF7L7-J7JL7F|.FJ-L-J7JF7LL77|-F7
J.F..F-F|LJ77|FJF7..LL7|.JFLF7L7L77F7.|FJ|-LL|JJLJ||||||||JLL-.F||LL-JJF-7|LL.LJFF7F7.|.-F--JL7LJ-JF|||L-7J|J|J-JLF|-|L-L-.JL|J7JL--7J|FL-F7
.FF7LL7.|-|-7FJ.|JF|.J|J-|F-LJ.L7L7L-.7JF77.||-F|FLJLJ||LJ.F|--J77-LJ-FJ---.|-L7-77FJ-77|L----J-JJ.|7.F|-|.L7L7FJ.F.F7-LLJ7.F--J-F|F|.|7|LJF
.FL|J.|F77J.F---F7-|-7L7|FF.7.-.L-J-L.|LF|L7LJ-F|FJ|.L||LLJF777L77F7J7FJ-LLF77.J-LJJ.L|-F|||JFJFL7.777.L7F.F|J---7LLF7|FJ.FJ7L7JFL7-L-LLJJLL
-77JLFF7|J7.7||F|J-J.FLL7.77.7.|JLJ|.F|-FF---JJJLJ.F7FLJ-.|FJ---77-L7-J.7||LJ7F-F-.F-.|L-7-JJF.7-FF||-|FJJ-LJ|JF77JJ|F7.|77F--J-L-J|.F|7LJF7
FLF--7L-J7|-LL|-|J.J7--JF|77F|.FJLFJ.L|-FJ7.|F7-|7.7.LJJFF.JJJFL.|||FJ|-L--7JF--|FFF|---7J|LF-JJF--7--JL.L7.L|.|L|JLF7JF7-FJ|.F.F|.J.-7J|F-7
F7|-FJ|L|-J7.F--F7|FJF|7L|||F77J-FFFL7.F|J|F-FJ|.F.--------|.|F.---7L-LF|-L7L7JL|L7.|.|7F-J.JJ.--|7|||.JF-F7|L|.|.|.FL-7L|L|7..|7-|.|LJ||7FL
LJJJ||L-J7LJ7F|-LLJ|7L--7L-J-.|F--FJ.|7LL.|L--7-7L7|L||-7JJL-JJLLJFF7J.LL77..J.F|.|7LF---J.FL.JLF|LJJ-F-F..F--|.77F7-JJL77.||.|F7.LFL77|J7FJ
LJJ..LL-JJ-L---7J-F77LL.LFJJLL-J-7JJ.L-LL7-L7.LLJL|JJF7-J-JLLJLLJ-7LJJ..LLJ-LJJ.LFL-FJJL|.FLLF-FL|LL.J.FL|-JJ.JJL--JLL--JJ.|J-L|J--J.L7..JJJ
`

const testInput0 = `......
..S-7.
..|.|.
..L-J.
......`
const testInput = `.....
.S-7.
.|.|.
.L-J.
.....`

const testInput2 = `...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........`

const testInput3 = `.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...`

var inputs = [][2]string{
	{"testInput0", testInput0},
	{"testInput", testInput},
	{"testInput2", testInput2},
	{"testInput3", testInput3},
	//{"prodInput", prodInput},
}

type dir int

const (
	North dir = 1 << 0
	East  dir = 1 << 1
	South dir = 1 << 2
	West  dir = 1 << 3
	AllD      = North | South | East | West
)

func part(t *testing.T, input string) {
	var lines []string
	var startX, startY int
	var startFound bool
	var y int
	scanLines(t, input, skipBlankLines, func(line string) bool {
		lines = append(lines, line)
		s := strings.Index(line, "S")
		if s != -1 {
			if startFound {
				panic("already found")
			}
			startFound = true
			startY = y
			startX = s
		}
		y++
		return true
	})
	fmt.Println(startX, startY)

	types := biMapFromMap(map[rune]dir{
		'|': North | South,
		'-': East | West,
		'L': North | East,
		'J': North | West,
		'7': South | West,
		'F': South | East,
	})

	dirsString := func(d dir) string {
		if d&AllD == AllD {
			return "ALL"
		}
		if d&AllD == 0 {
			return "NONE"
		}
		var ss []string
		for i := 0; i <= 1<<3; i++ {
			switch d & (1 << i) {
			case North:
				ss = append(ss, "N")
			case South:
				ss = append(ss, "S")
			case East:
				ss = append(ss, "E")
			case West:
				ss = append(ss, "W")
			default:
				ss = append(ss, " ")
			}
		}
		return strings.Join(ss, "")
	}
	_ = dirsString

	oppositeDirs := biMapFromMap(map[dir]dir{
		North: South,
		South: North,
		West:  East,
		East:  West,
	})

	var dirDeltas biMap[dir, [2]int]
	dirDeltas.add(North, [2]int{0, -1})
	dirDeltas.add(East, [2]int{1, 0})
	dirDeltas.add(South, [2]int{0, 1})
	dirDeltas.add(West, [2]int{-1, 0})

	var next [2]int
	for _, delta := range dirDeltas.to {
		x := delta[0]
		y := delta[1]
		searchY := startY + y
		searchX := startX + x
		pipeSegment, ok := types.to[rune(lines[searchY][searchX])]
		if !ok {
			continue
		}
		opposite := dirDeltas.from[[2]int{-x, -y}]
		if pipeSegment&opposite > 0 {
			next = [2]int{searchX, searchY}
			break
		}
		//fmt.Println("found", searchX, searchY, dir, opposite, doesJoin)
	}
	if next == [2]int{0, 0} {
		panic("did not find next")
	}

	delta := func(from [2]int, to [2]int) [2]int {
		return [2]int{
			to[0] - from[0],
			to[1] - from[1],
		}
	}

	add := func(a, b [2]int) [2]int {
		return [2]int{
			a[0] + b[0],
			a[1] + b[1],
		}
	}

	findNextPipeSegment := func(previous, current [2]int) [2]int {
		delta := delta(previous, current)
		cameTravellingIn, ok := dirDeltas.from[delta]
		if !ok {
			panic("nope")
		}
		cameTravellingFrom := oppositeDirs.to[cameTravellingIn]

		char := lines[current[1]][current[0]]
		currentPipeSegment := types.to[rune(char)]

		nextDir := currentPipeSegment ^ cameTravellingFrom
		nextDelta := dirDeltas.to[nextDir]
		return [2]int{current[0] + nextDelta[0], current[1] + nextDelta[1]}
	}

	pipeLocations := make(map[[2]int]dir)
	start := [2]int{startX, startY}
	previous := [2]int{startX, startY}
	first := next
	current := next
	pipeLocations[previous] = types.to[rune(lines[previous[1]][previous[0]])]
	pipeLocations[current] = types.to[rune(lines[current[1]][current[0]])]
	//t.Log(previous, current)
	for iterations := 2; next != start; iterations++ {
		next = findNextPipeSegment(previous, current)
		previous = current
		current = next
		pipeLocations[previous] = types.to[rune(lines[previous[1]][previous[0]])]
		pipeLocations[current] = types.to[rune(lines[current[1]][current[0]])]
		//t.Log(iterations, iterations/2, previous, current)
	}

	// first and current should join to start
	firstDelta := delta(start, first)
	lastDelta := delta(start, previous)
	startDirs := dirDeltas.from[firstDelta] | dirDeltas.from[lastDelta]
	fmt.Println("S is", firstDelta, lastDelta, startDirs, string(types.from[startDirs]))

	lines[startY] = strings.Replace(lines[startY], "S", string(types.from[startDirs]), 1)

	printLines := func() {
		for _, line := range lines {
			fmt.Println(line)
		}
	}

	printLines()

	getChar := func(p [2]int) rune {
		return rune(lines[p[1]][p[0]])
	}
	_ = getChar

	previousOutsideFaces := make(map[[2]int]dir)
	lookupOutsideFaces := func(t *testing.T, p [2]int) dir {
		t.Helper()
		d, ok := previousOutsideFaces[p]
		if !ok {
			require.FailNow(t, fmt.Sprint(p, "not known"))
		}
		return d
	}
	_ = lookupOutsideFaces

	updateChar := func(p [2]int, c rune) {
		currentLine := lines[p[1]]
		before := currentLine[:p[0]]
		after := currentLine[p[0]+1:]
		lines[p[1]] = before + string(c) + after
	}

	for x := 0; x < len(lines[0]); x++ {
		previousOutsideFaces[[2]int{x, -1}] = AllD
	}
	for y := 0; y < len(lines); y++ {
		previousOutsideFaces[[2]int{-1, y}] = AllD
	}

	partOfPipeline := func(p [2]int) bool {
		_, ok := pipeLocations[p]
		return ok
	}
	for y := 0; y < len(lines); y++ {
		//previously := "outside"

		t.Log("LINE", y)

		for x := 0; x < len(lines[y]); x++ {
			p := [2]int{x, y}

			var resolved bool
			for _, dir := range []dir{North, West} {
				adjacentPos := add(p, dirDeltas.to[dir])
				adjacentOutsideFaces := lookupOutsideFaces(t, adjacentPos)
				oppositeDir := oppositeDirs.to[dir]
				adjacentSideOutside := (adjacentOutsideFaces & oppositeDir) == oppositeDir
				//t.Log(p, "adjacent outside", dirsString(dir), dirsString(adjacentOutsideFaces), adjacentSideOutside)

				//t.Log(p, dirsString(dir), "adjacent is", adjacentPos, lookupOutsideFaces(t, adjacentPos), dirsString(lookupOutsideFaces(t, adjacentPos)))
				if !adjacentSideOutside {
					continue
				}
				if !partOfPipeline(p) {
					updateChar(p, 'O')
					previousOutsideFaces[p] = AllD
					resolved = true
				} else {
					// handle when part of pipeline
					// should be able to determine what faces of the pipe are outside given that we know
					// the direction of an outside and we know what shape the pipe is
					// might need to track 2 things, outside sides, inside sides, because there are pipe sides which are neither inside nor outside
				}
			}

			if !resolved {
				previousOutsideFaces[p] = 0
			}
		}
	}

	fmt.Println()
	printLines()
	fmt.Println()
}

func biMapFromMap[k comparable, v comparable](m map[k]v) biMap[k, v] {
	var out biMap[k, v]
	for key, value := range m {
		out.add(key, value)
	}
	return out
}

type biMap[k comparable, v comparable] struct {
	to   map[k]v
	from map[v]k
}

func (m *biMap[k, v]) add(key k, value v) {
	if m.to == nil {
		m.to = make(map[k]v)
	}
	if m.from == nil {
		m.from = make(map[v]k)
	}

	if _, ok := m.to[key]; ok {
		panic("key already exists")
	}
	if _, ok := m.from[value]; ok {
		panic("value already exists")
	}

	m.to[key] = value
	m.from[value] = key
}

func TestPart1(t *testing.T) {
	for _, input := range inputs {
		fmt.Println(input[0])
		part(t, input[1])
		fmt.Println()
	}
}

func scanWords(t *testing.T, input string, processors ...func(string) bool) {
	reader := bufio.NewScanner(strings.NewReader(input))
	reader.Split(bufio.ScanWords)
	for reader.Scan() {
		word := reader.Text()
		var processed bool
		for _, p := range processors {
			if p(word) {
				processed = true
				break
			}
		}
		require.True(t, processed, "not all input was processed for word:", word)
	}
}

func scanLines(t *testing.T, input string, processors ...func(string) bool) {
	reader := bufio.NewScanner(strings.NewReader(input))
	for reader.Scan() {
		line := reader.Text()
		var processed bool
		for _, p := range processors {
			if p(line) {
				processed = true
				break
			}
		}
		require.True(t, processed, "not all input was processed for line:", line)
	}
}

func skipBlankLines(line string) bool {
	return line == ""
}

func all[t any](values []t, predicate func(t) bool) bool {
	for _, value := range values {
		if !predicate(value) {
			return false
		}
	}
	return true
}

func mapSlice[in any, out any](mapIt func(in) out, inSlice []in) []out {
	outSlice := make([]out, len(inSlice))
	for i, inV := range inSlice {
		outSlice[i] = mapIt(inV)
	}
	return outSlice
}

func must[I any, O any](t *testing.T, fn func(I) (O, error), in I) O {
	t.Helper()
	o, err := fn(in)
	require.NoError(t, err)
	return o
}
