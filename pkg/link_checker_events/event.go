package link_checker_events

import om "github.com/shshimamo/delinkcious/pkg/object_model"

type Event struct {
	Username string
	Url      string
	Status   om.LinkStatus
}
