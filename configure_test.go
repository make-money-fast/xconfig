package xconfig

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

type TestingConfig struct {
	F1           float64      `json:"f1" default:"1.1"`
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

func TestParseFromData(t *testing.T) {
	var config TestingConfig
	err := ParseFromData(nil, &config)
	require.Nil(t, err)

	require.EqualValues(t, config.F1, float64(1.1))
	require.EqualValues(t, config.F2, float32(1.2))
	require.EqualValues(t, config.I1, 1)
	require.EqualValues(t, config.I2, 2)
	require.EqualValues(t, config.I3, 3)
	require.EqualValues(t, config.I4, 4)
	require.EqualValues(t, config.I5, 5)

	require.EqualValues(t, config.U1, 1)
	require.EqualValues(t, config.U2, 2)
	require.EqualValues(t, config.U3, 3)
	require.EqualValues(t, config.U4, 4)
	require.EqualValues(t, config.U5, 5)

	require.EqualValues(t, config.S1, "str1")
	require.EqualValues(t, config.B1, true)
	require.EqualValues(t, config.B3, true)
	require.EqualValues(t, config.A1, []string{"1", "2"})
	require.EqualValues(t, config.A2, []int{1, 2, 3})
	require.EqualValues(t, config.A3, []bool{true, false})

	data, _ := json.MarshalIndent(config, "", "\t")
	fmt.Println(string(data))
}
