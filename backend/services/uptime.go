package services

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/Testausserveri/uptimes/types"
	"github.com/romeq/jobscheduler"
)

var (
	errReqFailed            = errors.New("connection to server failed")
	errCannotPrepareRequest = errors.New("preparing the request for sending failed")
)

func getStatusOf(domain types.Domain) *types.StatusMetadata {
	var (
		httpclient http.Client
		md         = &types.StatusMetadata{
			Reachable: false,
			Date:      time.Now(),
		}
	)

	req, err := http.NewRequest("GET", domain.Configuration.Domain, nil)
	if err != nil {
		md.Error = errCannotPrepareRequest.Error()
		return md
	}
	req.Header.Set("User-Agent", "testausuptime")

	res, err := httpclient.Do(req)
	if err != nil {
		md.Error = errReqFailed.Error()
		return md
	}

	md.Reachable = true
	if err := verifyRequest(res, domain.Configuration.Requirements); err != nil {
		md.Error = err.Error()
		return md
	}

	return md
}

func verifyRequest(res *http.Response, r types.DomainUptimeRequirements) error {
	if r.StatusCode > 0 && res.StatusCode != r.StatusCode {
		return errors.New("status codes don't match")
	}
	if r.ContentType != "" && res.Header.Get("Content-Type") != r.ContentType {
		return errors.New("content types don't match")
	}
	return nil
}

func domainUpdater(dm types.Domain, hl int) jobscheduler.Job {
	return jobscheduler.NewJob(0, dm.Configuration.UpdateInterval, func() {
		*dm.LastStatus = *getStatusOf(dm)

		dmh := *dm.History
		if len(*dm.History) >= hl {
			dmh = dmh[len(dmh)-hl+1:]
		}
		*dm.History = append(dmh, *dm.LastStatus)

		log.Printf("update: url=%s up=%v err=%s\n", dm.Configuration.Domain,
			dm.LastStatus.Reachable, dm.LastStatus.Error)

	}, true)
}

func InitStatusGroupUpdater(sg *types.StatusGroup) {
	var jobs []jobscheduler.Job
	for _, dm := range *sg.Domains {
		jobs = append(jobs, domainUpdater(dm, 50))
	}

	go jobscheduler.Run(jobs)
}
