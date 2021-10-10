package celestrak_fetcher

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"text/template"
	"tle_manager/shared/util"
	"tle_manager/tle_fetcher/lambda/model"

	"git.darknebu.la/Satellite/tle"
)

const (
	templateCelestrakName = "celestrak"
)

var (
	templateCelestrak = template.Must(template.New(templateCelestrakName).Parse("https://celestrak.com/NORAD/elements/gp.php?NAME={{ .Name }}&FORMAT=TLE"))
)

func FetchTLEForSatellites(satName []string) (map[string]*model.Tle, map[string]error) {
	tleBySat := make(map[string]*model.Tle, len(satName))
	errors := make(map[string]error)
	for _, sat := range satName {
		tle, err := fetchTLEForSatellite(sat)
		if err != nil {
			errors[sat] = err
			continue
		}
		tleBySat[sat] = tle
	}
	return tleBySat, errors
}

func fetchTLEForSatellite(satName string) (*model.Tle, error) {
	url, err := util.ExecTempl(templateCelestrak, struct {
		Name string
	}{
		Name: satName,
	})
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	lines, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	splitLines := strings.Split(strings.Trim(string(lines), "\n"), "\n")
	if len(splitLines) != 3 {
		return nil, errors.New("TLE should be 3 lines")
	}
	newTLE, err := tle.NewTLE(string(lines))
	if err != nil {
		return nil, fmt.Errorf("can't read tle: %w", err)
	}
	return &model.Tle{
		Lines: splitLines,
		TLE:   &newTLE,
	}, nil
}
