package dkron

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSchedule(t *testing.T) {
	sched := NewScheduler()

	assert.False(t, sched.Started)

	testJob1 := &Job{
		Name:       "cron_job",
		Schedule:   "@every 2s",
		Command:    "echo 'test1'",
		Owner:      "John Dough",
		OwnerEmail: "foo@bar.com",
		Shell:      true,
	}
	sched.Start([]*Job{testJob1})

	assert.True(t, sched.Started)
	now := time.Now().Truncate(time.Second)

	assert.Equal(t, now.Add(time.Second*2), sched.GetEntry(testJob1).Next)

	testJob2 := &Job{
		Name:       "cron_job",
		Schedule:   "@every 5s",
		Command:    "echo 'test2'",
		Owner:      "John Dough",
		OwnerEmail: "foo@bar.com",
		Shell:      true,
	}
	sched.Restart([]*Job{testJob2})

	assert.True(t, sched.Started)
	assert.Len(t, sched.Cron.Entries(), 1)
}
