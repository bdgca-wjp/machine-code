/*
author: superl[N.S.T]
github: https://github.com/super-l/
*/
package main

import (
	"fmt"

	"github.com/bdgca-wjp/machine-code/machine"
)

// https://www.icode9.com/content-3-710187.html  go 获取linux cpuId 的方法
func main() {
	machineData := machine.GetMachineData()

	fmt.Printf("%+v\n", machineData)
}
