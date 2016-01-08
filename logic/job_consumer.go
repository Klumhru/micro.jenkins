package logic

import (
	"github.com/go-swagger/go-swagger/httpkit/middleware"

	"github.com/bndr/gojenkins"
	"github.com/spf13/viper"

	"github.com/IcelandairLabs/micro.jenkins/models"
	"github.com/IcelandairLabs/micro.jenkins/restapi/jenkins/job"
)

var (
	api  *gojenkins.Jenkins
	conf *viper.Viper
)

func init() {
	conf = viper.New()
}

// GetJobListHandler returns a list of jobs
func GetJobListHandler() middleware.Responder {
	color, name, url := "Hi there", "Go", "Away"
	payload := []*models.Job{
		&models.Job{
			Color: &color,
			Name:  &name,
			URL:   &url,
		},
	}
	return job.NewGetJobListOK().WithPayload(payload)
}
