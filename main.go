// MIT License
//
// Copyright (c) 2018 SpiralScout
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"github.com/sirupsen/logrus"
	rr "github.com/spiral/roadrunner/cmd/rr/cmd"

	// services (plugins)
	"github.com/spiral/grpc"

	"github.com/spiral/roadrunner/service/env"
	"github.com/spiral/roadrunner/service/http"
	"github.com/spiral/roadrunner/service/rpc"
	"github.com/spiral/roadrunner/service/static"

	"github.com/spiral/jobs"
	"github.com/spiral/jobs/broker/beanstalk"
	"github.com/spiral/jobs/broker/local"
	"github.com/spiral/jobs/broker/sqs"

	// additional commands and debug handlers
	_ "github.com/spiral/grpc/cmd/rr-grpc/grpc"
	_ "github.com/spiral/jobs/cmd/rr-jobs/jobs"
	_ "github.com/spiral/roadrunner/cmd/rr/http"
)

func main() {
	rr.Container.Register(env.ID, &env.Service{})
	rr.Container.Register(rpc.ID, &rpc.Service{})
	rr.Container.Register(http.ID, &http.Service{})
	rr.Container.Register(grpc.ID, &grpc.Service{})
	rr.Container.Register(static.ID, &static.Service{})

	rr.Container.Register(jobs.ID, &jobs.Service{
		Brokers: map[string]jobs.Broker{
			"local":     &local.Broker{},
			"beanstalk": &beanstalk.Broker{},
			"sqs":       &sqs.Broker{},
		},
	})

	rr.Logger.Formatter = &logrus.TextFormatter{ForceColors: true}

	rr.Execute()
}
