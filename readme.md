### Configure

使用: 

```shell
  go get github.com/make-money-fast/xconfig
```

### 声明使用

* 支持 `yaml` 解析 
* 支持 `json` 解析
* 支持设置环境变量: `env:"ENV_NAME"`
* 支持 `设置默认值`, 设置如下:

- [ ] ~~不支持 Map 配置~~

```go
// 增加 default:"" 设置默认值
type Configure struct {
    FloatValue float32 `default:"1.1" json:"float_value"`
    ... 
}
```



