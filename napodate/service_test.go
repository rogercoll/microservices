package napodate

import (
	"time"
	"context"
	"testing"
	"github.com/magiconair/properties/assert"
)

func setup() (srv Service, ctx context.Context) {
	return NewService(), context.Background()
}

func TestStatus(t *testing.T) {
	expectedOutput := "ok"
	srv, ctx := setup()
	s, err := srv.Status(ctx)
	if err != nil {
		t.Errorf("Error: %s\n", err)
	}
	assert.Equal(t,expectedOutput,s)
}

func TestGet(t *testing.T) {
	time := time.Now()
    expectedOutput := time.Format("02/01/2006")
	srv, ctx := setup()
    d, err := srv.Get(ctx)
    if err != nil {
        t.Errorf("Error: %s", err)
	}
	assert.Equal(t,expectedOutput,d)
}

func TestValidate(t *testing.T) {
	srv, ctx := setup()
    b, err := srv.Validate(ctx, "31/12/2019")
    if err != nil {
        t.Errorf("Error: %s", err)
    }
    if !b {
        t.Errorf("date should be valid")
    }
    b, err = srv.Validate(ctx, "31/31/2019")
    if b {
        t.Errorf("date should be invalid")
    }
    b, err = srv.Validate(ctx, "12/31/2019")
    if b {
        t.Errorf("USA date should be invalid")
    }
}