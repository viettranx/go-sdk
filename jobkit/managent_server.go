package jobkit

import (
	"fmt"

	"github.com/blend/go-sdk/cron"
	"github.com/blend/go-sdk/web"
)

// NewManagementServer returns a new management server that lets you
// trigger jobs or look at job statuses via. a json api.
func NewManagementServer(jm *cron.JobManager, cfg *Config) *web.App {
	app := web.NewFromConfig(&cfg.Web)
	app.Views().AddLiterals(headerTemplate, footerTemplate, indexTemplate)
	app.GET("/", func(r *web.Ctx) web.Result {
		return r.View().View("index", jm.Status())
	})
	app.GET("/healthz", func(_ *web.Ctx) web.Result {
		if jm.IsRunning() {
			return web.JSON.OK()
		}
		return web.JSON.InternalError(fmt.Errorf("job manager is stopped or in an inconsistent state"))
	})
	app.GET("/api/jobs", func(_ *web.Ctx) web.Result {
		return web.JSON.Result(jm.Status())
	})
	app.GET("/api/job.status/:jobName", func(r *web.Ctx) web.Result {
		jobName, err := r.RouteParam("jobName")
		if err != nil {
			return web.JSON.BadRequest(err)
		}
		status, err := jm.Job(jobName)
		if err := jm.RunJob(jobName); err != nil {
			return web.JSON.BadRequest(err)
		}
		return web.JSON.Result(status)
	})
	app.POST("/api/job.run/:jobName", func(r *web.Ctx) web.Result {
		jobName, err := r.RouteParam("jobName")
		if err != nil {
			return web.JSON.BadRequest(err)
		}
		if err := jm.RunJob(jobName); err != nil {
			return web.JSON.BadRequest(err)
		}
		return web.JSON.OK()
	})
	app.POST("/api/job.cancel/:jobName", func(r *web.Ctx) web.Result {
		jobName, err := r.RouteParam("jobName")
		if err != nil {
			return web.JSON.BadRequest(err)
		}
		if err := jm.CancelJob(jobName); err != nil {
			return web.JSON.BadRequest(err)
		}
		return web.JSON.OK()
	})
	app.POST("/api/job.disable/:jobName", func(r *web.Ctx) web.Result {
		jobName, err := r.RouteParam("jobName")
		if err != nil {
			return web.JSON.BadRequest(err)
		}
		if err := jm.DisableJob(jobName); err != nil {
			return web.JSON.BadRequest(err)
		}
		return web.JSON.Result(fmt.Sprintf("%s disabled", jobName))
	})
	app.POST("/api/job.enable/:jobName", func(r *web.Ctx) web.Result {
		jobName, err := r.RouteParam("jobName")
		if err != nil {
			return web.JSON.BadRequest(err)
		}
		if err := jm.EnableJob(jobName); err != nil {
			return web.JSON.BadRequest(err)
		}
		return web.JSON.Result(fmt.Sprintf("%s enabled", jobName))
	})
	return app
}
