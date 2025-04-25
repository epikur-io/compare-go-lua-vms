package benchmarks

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/epikur-io/compare-go-lua-vms/internal/utils"
	lua "github.com/epikur-io/gopher-lua"
)

func BenchmarkGopherLuaTables(b *testing.B) {
	files, err := utils.ListFiles("./../lua-code/tables")
	utils.FailOnError(b, err)

	for _, file := range files {
		testname := fmt.Sprintf("tables/%s", filepath.Base(file))
		b.Run(testname, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StartTimer()
				lvm := lua.NewState()
				utils.FailOnError(b, lvm.DoFile(file))
				b.StopTimer()
			}
		})
	}

}

func BenchmarkGopherLuaMist(b *testing.B) {
	files, err := utils.ListFiles("./../lua-code/misc")
	utils.FailOnError(b, err)

	for _, file := range files {
		testname := fmt.Sprintf("misc/%s", filepath.Base(file))
		utils.FailOnError(b, err)
		b.Run(testname, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StartTimer()
				lvm := lua.NewState()
				utils.FailOnError(b, lvm.DoFile(file))
				b.StopTimer()
			}
		})
	}

}

func BenchmarkGopherLuaIssues(b *testing.B) {
	files, err := utils.ListFiles("./../lua-code/issues")
	utils.FailOnError(b, err)

	for _, file := range files {
		testname := fmt.Sprintf("issues/%s", filepath.Base(file))
		utils.FailOnError(b, err)
		b.Run(testname, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StartTimer()
				lvm := lua.NewState()
				utils.FailOnError(b, lvm.DoFile(file))
				b.StopTimer()
			}
		})
	}

}
