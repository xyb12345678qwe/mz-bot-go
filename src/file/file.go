package file
import (
	"fmt"
    "gopkg.in/yaml.v3"
	"os"
)


var FilePaths = map[string]string{
   "botConfig": "config/bot_config.yaml",
}

/**
 * 读取 YAML 文件，数据结构不固定
 * @param key 文件路径
 * @return 解析后的配置数据和错误信息
 */
func ReadFileByYAML(key string) (map[string]interface{}, error) {
    dataBytes, err := os.ReadFile(FilePaths[key])
	if err != nil {
		return nil, fmt.Errorf("读取文件失败: %w", err)
	}

	var config map[string]interface{}
	if err = yaml.Unmarshal(dataBytes, &config); err != nil {
		return nil, fmt.Errorf("解析 yaml 文件失败: %w", err)
	}

	return config, nil
}

// 写入 YAML 文件
// @param config 需要写入的数据
// @param key 文件路径的键名
// @return 错误信息
func WriteFileByYAML(config map[string]interface{}, key string) error {
	dataBytes, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("序列化配置到 YAML 失败: %w", err)
	}

	filePath := FilePaths[key]
	if err := os.WriteFile(filePath, dataBytes, 0644); err != nil {
		return fmt.Errorf("写入文件失败: %w", err)
	}

	fmt.Println("成功写入 YAML 文件:", filePath)
	return nil
}