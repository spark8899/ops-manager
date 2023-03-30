package core

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/spark8899/ops-manager/server/core/internal"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"

	"github.com/spark8899/ops-manager/server/global"
	_ "github.com/spark8899/ops-manager/server/packfile"
)

// Viper //
// priority: command Line > environment variable > defaults
// Author spark8899
func Viper(path ...string) *viper.Viper {
	var config string

	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" { // check command Line
			if configEnv := os.Getenv(internal.ConfigEnv); configEnv == "" { // decision internal.ConfigEnv environment variable is empty
				switch gin.Mode() {
				case gin.DebugMode:
					config = internal.ConfigDefaultFile
					fmt.Printf("gin model %s environment name, config's path is %s\n", gin.EnvGinMode, internal.ConfigDefaultFile)
				case gin.ReleaseMode:
					config = internal.ConfigReleaseFile
					fmt.Printf("gin model %s environment name, config's path is %s\n", gin.EnvGinMode, internal.ConfigReleaseFile)
				case gin.TestMode:
					config = internal.ConfigTestFile
					fmt.Printf("gin model %s environment name, config's path is %s\n", gin.EnvGinMode, internal.ConfigTestFile)
				}
			} else { // internal.ConfigEnv environment variable is empty, assign to config
				config = configEnv
				fmt.Printf("you are using the %s environment variable, config's path %s\n", internal.ConfigEnv, config)
			}
		} else { // The command line parameter is not empty, assign the value to config
			fmt.Printf("the value you are passing with the -c parameter of the command line, config' path is '%s\n", config)
		}
	} else { // The first value of the variable parameter passed by the function is assigned to config
		config = path[0]
		fmt.Printf("you are using the value passed by func Viper(), config's path is %s\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.OPM_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.OPM_CONFIG); err != nil {
		fmt.Println(err)
	}

	// root adaptability Find the corresponding migration location according to the root location to ensure that the root path is valid
	global.OPM_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	return v
}
