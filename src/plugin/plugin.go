package plugin

import (
	"fmt"
	"os"
	_"plugin"
	_"io/ioutil"
	 "os/exec"
)

func PluginLoad(dir string) ([]string, error) {
	// 检查目录是否存在
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return nil, fmt.Errorf("目录不存在: %s", dir)
	}

	// 读取目录条目
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("无法读取目录: %s, 错误: %v", dir, err)
	}

	var dirs []string
	// 筛选出所有一级文件夹
	for _, entry := range entries {
		if entry.IsDir() { // 确认是目录
			dirs = append(dirs, entry.Name()) // 添加文件夹名称
			runCmd := exec.Command("plugins/" + entry.Name() + "/plugin.plugin")
			go func() {
				output, err := runCmd.CombinedOutput()
				fmt.Printf(string(output))
				if err != nil {
        			fmt.Println("[Error][load]" + "[" + entry.Name() + "]", err)
        			return
    			}
			}()
			
		}
	}

	return dirs, nil
}//response
