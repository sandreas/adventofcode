package main

import (
	"strings"
	"strconv"
	"fmt"
)

func main() {
	input := `409	194	207	470	178	454	235	333	511	103	474	293	525	372	408	428
4321	2786	6683	3921	265	262	6206	2207	5712	214	6750	2742	777	5297	3764	167
3536	2675	1298	1069	175	145	706	2614	4067	4377	146	134	1930	3850	213	4151
2169	1050	3705	2424	614	3253	222	3287	3340	2637	61	216	2894	247	3905	214
99	797	80	683	789	92	736	318	103	153	749	631	626	367	110	805
2922	1764	178	3420	3246	3456	73	2668	3518	1524	273	2237	228	1826	182	2312
2304	2058	286	2258	1607	2492	2479	164	171	663	62	144	1195	116	2172	1839
114	170	82	50	158	111	165	164	106	70	178	87	182	101	86	168
121	110	51	122	92	146	13	53	34	112	44	160	56	93	82	98
4682	642	397	5208	136	4766	180	1673	1263	4757	4680	141	4430	1098	188	1451
158	712	1382	170	550	913	191	163	459	1197	1488	1337	900	1182	1018	337
4232	236	3835	3847	3881	4180	4204	4030	220	1268	251	4739	246	3798	1885	3244
169	1928	3305	167	194	3080	2164	192	3073	1848	426	2270	3572	3456	217	3269
140	1005	2063	3048	3742	3361	117	93	2695	1529	120	3480	3061	150	3383	190
489	732	57	75	61	797	266	593	324	475	733	737	113	68	267	141
3858	202	1141	3458	2507	239	199	4400	3713	3980	4170	227	3968	1688	4352	4168`

	stringLines := strings.Split(input, "\n")

	lineCount := len(stringLines)
	lines := make([][]int, lineCount)
	for i:=0;i<len(stringLines);i++ {
		currentStringLine := strings.Fields(stringLines[i])
		colCount := len(currentStringLine)
		lines[i] = make([]int, colCount)
		for j:=0;j<colCount;j++ {
			trimmed := strings.TrimSpace(currentStringLine[j])
			if  trimmed != "" {
				value, _ := strconv.Atoi(trimmed)
				// fmt.Printf("line %d: value: %d\n", i+1, value)
				lines[i][j] = value
			}
		}
	}

	checksum := 0
	divisionChecksum := 0
	for i:=0;i<len(lines);i++ {
		min, max := MinMax(lines[i])

		diff := max - min
		// fmt.Printf("line %d: min: %d, max: %d, diff:%d\n", i+1, min, max, diff)


		checksum += diff
		divisionChecksum += EvenDivide(lines[i])
	}

	fmt.Printf("checksum: %d\n", checksum)
	fmt.Printf("divisionChecksum: %d", divisionChecksum)


}


func MinMax(array []int) (int, int) {
	// fmt.Printf("%#v", array)
	var max = array[0]
	var min = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func EvenDivide(array []int) (int){
	for i:=0;i<len(array);i++ {
		for j:=0;j<len(array);j++ {
			if i==j {
				continue
			}
			if array[i] % array[j] == 0 {
				return array[i] / array[j]
			}
		}
	}
	return 0
}
