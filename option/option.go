// Copyright (C) 2014-2018 Goodrain Co., Ltd.
// RAINBOND, Application Management Platform

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version. For any non-GPL usage of Rainbond,
// one or multiple Commercial Licenses authorized by Goodrain Co., Ltd.
// must be obtained first.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package option

import (
	"github.com/spf13/pflag"
)

var Conf *Config

func init() {
	Conf = NewConfig()
}

//Config config
type Config struct {
	ResultPath    string
	ImageSavePath string
}

//NewAPIServer new server
func NewConfig() *Config {
	return &Config{}
}

//AddFlags config
func (c *Config) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&c.ResultPath, "result-path", "/run/yolo-tensorflow/result", "the path where detection result json saved")
	fs.StringVar(&c.ImageSavePath, "image-savepath", "/run/images", "the path where image saved from websocket")
}
