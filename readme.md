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

type TestingConfig struct {
    F1           float64      `json:"f1" default:"1.1" env:"DEFAULT_F1"`
    F2           float32      `json:"f2" default:"1.2"`
    I1           int          `json:"i1" default:"1"`
    I2           int8         `json:"i2" default:"2"`
    I3           int16        `json:"i3" default:"3"`
    I4           int32        `json:"i4" default:"4"`
    I5           int64        `json:"i5" default:"5"`
    U1           uint         `json:"u1" default:"1"`
    U2           uint8        `json:"u2" default:"2"`
    U3           uint16       `json:"u3" default:"3"`
    U4           uint32       `json:"u4" default:"4"`
    U5           uint64       `json:"u5" default:"5"`
    S1           string       `json:"s1" default:"str1"`
    B1           bool         `json:"b1" default:"true"`
    B3           bool         `json:"b3" default:"1"`
    A1           []string     `json:"a1" default:"[\"1\",\"2\"]"`
    A2           []int        `json:"a2" default:"[1,2,3]"`
    A3           []bool       `json:"a3" default:"[true,false]"`
    InnerConfig1 InnerConfig  `json:"inner_config_1"`
    InnerConfig2 *InnerConfig `json:"inner_config_2"`
    *InnerConfig
}

type InnerConfig struct {
    F1 float64  `json:"f1" default:"1.1"`
    F2 float32  `json:"f2" default:"1.2"`
    I1 int      `json:"i1" default:"1"`
    I2 int8     `json:"i2" default:"2"`
    I3 int16    `json:"i3" default:"3"`
    I4 int32    `json:"i4" default:"4"`
    I5 int64    `json:"i5" default:"5"`
    U1 uint     `json:"u1" default:"1"`
    U2 uint8    `json:"u2" default:"2"`
    U3 uint16   `json:"u3" default:"3"`
    U4 uint32   `json:"u4" default:"4"`
    U5 uint64   `json:"u5" default:"5"`
    S1 string   `json:"s1" default:"str1"`
    B1 bool     `json:"b1" default:"true"`
    B3 bool     `json:"b3" default:"1"`
    A1 []string `json:"a1" default:"[\"1\",\"2\"]"`
    A2 []int    `json:"a2" default:"[1,2,3]"`
    A3 []bool   `json:"a3" default:"[true,false]"`
}
```



