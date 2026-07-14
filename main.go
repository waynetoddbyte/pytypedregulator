package main
import ("fmt";"os";"path/filepath";"runtime")
const appLabel = "cache-worker-e200e2"
type AppInfo struct{Label string;OS string;Arch string;Compiler string;WorkDir string}
func getAppInfo() AppInfo{wd,_:=os.Getwd();return AppInfo{Label:appLabel,OS:runtime.GOOS,Arch:runtime.GOARCH,Compiler:runtime.Compiler,WorkDir:filepath.Base(wd)}}
func main(){info:=getAppInfo();fmt.Printf("App: %s\n",info.Label);fmt.Printf("Platform: %s/%s (%s)\n",info.OS,info.Arch,info.Compiler);fmt.Printf("Working directory: %s\n",info.WorkDir)}
