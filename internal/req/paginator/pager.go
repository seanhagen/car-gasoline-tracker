package paginator

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/seanhagen/gas-web/internal/db"
	"github.com/shopspring/decimal"
)

const defaultPerPage = 50
const maxPerPage = 100
const magicTimeNumber = 10000.0

type (
	// Pager TODO
	Pager struct {
		after   *time.Time
		before  *time.Time
		perPage int
		args    url.Values
		route   string
		ctx     *gin.Context
		fetcher Fetcher
	}

	// Paginatable TODO
	Paginatable interface {
		UUID() string
		Created() time.Time
	}

	// PagerColl TODO
	PagerColl interface {
		Count() int
		First() Paginatable
		Last() Paginatable
		Itterator() *Itterator
		Item(int) Paginatable
		FirstQuery(isAdmin bool) string
		LastQuery(isAdmin bool) string
		LastTimeQuery(isAdmin bool) string
	}

	// Fetcher TODO
	Fetcher interface {
		Getter(i interface{}, queryName string, args ...interface{}) error
	}
)

// CreatePager uses the context of the current request to determine the page
// and per-page values
func CreatePager(ctx *gin.Context, fetch Fetcher) Pager {
	p := &Pager{
		ctx:     ctx,
		fetcher: fetch,
	}
	p.setBeforeAfter()
	p.setPerPage()
	p.setArgs()
	p.setRoute()
	return *p
}

// PerPage returns how many items per page there are
func (p Pager) PerPage() int {
	return p.perPage
}

// GetArgs TODO
func (p Pager) GetArgs() url.Values {
	a := p.args
	if p.after != nil {
		x := timeToProperFloat(*p.after)
		a.Add("after", x)
	}

	if p.before != nil {
		x := timeToProperFloat(*p.before)
		a.Add("before", x)
	}

	a.Add("per_page", strconv.Itoa(p.perPage))
	return a
}

func (p *Pager) setBeforeAfter() {
	ctx := p.ctx
	b, _ := ctx.GetQuery("before")
	p.before = ParseTime(b)

	a, _ := ctx.GetQuery("after")
	p.after = ParseTime(a)
}

func ParseTime(a string) *time.Time {
	if a != "" {
		i, err := strconv.ParseFloat(a, 64)
		if err == nil {
			x := decimal.NewFromFloat(i / (magicTimeNumber * 10))
			seconds := x.IntPart()
			nanos := int64(x.Exponent())
			tm := time.Unix(seconds, nanos)
			return &tm
		}
	}
	return nil
}

func (p *Pager) setPerPage() {
	var iPer int
	var err error
	sPer, ok := p.ctx.GetQuery("per_page")
	if ok {
		iPer, err = strconv.Atoi(sPer)
		if err != nil {
			iPer = defaultPerPage
		} else if iPer > maxPerPage {
			iPer = maxPerPage
		}
	} else {
		iPer = defaultPerPage
	}

	if iPer <= 0 {
		iPer = defaultPerPage
	}

	p.perPage = iPer
}

func (p *Pager) setRoute() {
	p.route = p.ctx.Request.URL.Path
}

func (p *Pager) setArgs() {
	args := p.ctx.Request.URL.Query()

	p.args = args
}

// SetCountHeader TODO
func (p Pager) setCountHeader(filter db.FilterArgs, db db.Storage) (int, error) {
	c, err := db.PageCount(filter)
	if err != nil {
		return 0, err
	}

	count := strconv.Itoa(c)
	p.ctx.Header("Count", count)
	return c, nil
}

// BaseArgs TODO
func (p Pager) BaseArgs() url.Values {
	t := p.args.Encode()
	args, _ := url.ParseQuery(t)
	args.Del("before")
	args.Del("after")
	args.Del("per_page")
	args.Add("per_page", strconv.Itoa(p.perPage))
	return args
}

func (p Pager) createLink(tm time.Time, rel string) string {
	args := p.BaseArgs()

	stamp := timeToProperFloat(tm)
	if rel == "first" || rel == "next" {
		args.Add("after", stamp)
	} else {
		args.Add("before", stamp)
	}

	host := p.ctx.Request.Header.Get("X-Forwarded-Host")
	proto := p.ctx.Request.Header.Get("X-Forwarded-Proto")

	return fmt.Sprintf("<%v://%v%v?%v>, rel='%v'", proto, host, p.route, args.Encode(), rel)
}

func timeToProperFloat(tm time.Time) string {
	return strconv.FormatFloat(float64(tm.UnixNano())/(magicTimeNumber*magicTimeNumber*10), 'f', 5, 64)
}
