package benchmarks

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/epikur-io/compare-go-lua-vms/internal/utils"
	lua "github.com/epikur-io/go-lua"
)

func BenchmarkGoLuaTables(b *testing.B) {
	files, err := utils.ListFiles("./../lua-code/tables")
	utils.FailOnError(b, err)

	for _, file := range files {
		testname := fmt.Sprintf("tables/%s", filepath.Base(file))
		utils.FailOnError(b, err)
		b.Run(testname, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StartTimer()
				lvm := lua.NewState()
				lua.OpenLibraries(lvm)
				utils.FailOnError(b, lua.DoFile(lvm, file))
				b.StopTimer()

			}
		})
	}

}

func BenchmarkGoLuaMisc(b *testing.B) {
	files, err := utils.ListFiles("./../lua-code/misc")
	utils.FailOnError(b, err)

	for _, file := range files {
		fb := filepath.Base(file)
		if fb == "url-encode.lua" {
			// ignore for now until gsub is fixed for go-lua
			continue
		}
		testname := fmt.Sprintf("misc/%s", fb)
		utils.FailOnError(b, err)
		b.Run(testname, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StartTimer()
				lvm := lua.NewState()
				lua.OpenLibraries(lvm)
				utils.FailOnError(b, lua.DoFile(lvm, file))
				b.StopTimer()
			}
		})
	}

}

func BenchmarkGoLuaIssues(b *testing.B) {
	files, err := utils.ListFiles("./../lua-code/issues")
	utils.FailOnError(b, err)

	for _, file := range files {
		fb := filepath.Base(file)
		testname := fmt.Sprintf("issues/%s", fb)
		utils.FailOnError(b, err)
		b.Run(testname, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StartTimer()
				lvm := lua.NewState()
				lua.OpenLibraries(lvm)
				utils.FailOnError(b, lua.DoFile(lvm, file))
				b.StopTimer()
			}
		})
	}

}
