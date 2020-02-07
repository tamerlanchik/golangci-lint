package golinters

import (
	provider "github.com/Rikkuru/myerrorlint"
	"golang.org/x/tools/go/analysis"

	"github.com/golangci/golangci-lint/pkg/config"
	"github.com/golangci/golangci-lint/pkg/golinters/goanalysis"
	"github.com/golangci/golangci-lint/pkg/lint/linter"
)

func getConfigFromSettings(settings config.MyErrorLintSettings) provider.Config {
	return provider.Config{
		AllowedTypes:  settings.AllowedTypes,
		OurPackages:   settings.OurPackages,
		ReportUnknown: settings.ReportUnknown,
	}
}

func NewMyErrorLint() *goanalysis.Linter {
	analyzer := provider.NewAnalyzerWithoutRun()
	return goanalysis.NewLinter(
		provider.Name,
		provider.Doc,
		[]*analysis.Analyzer{analyzer},
		nil,
	).WithLoadMode(goanalysis.LoadModeTypesInfo).
		WithContextSetter(func(lintCtx *linter.Context) {
			analyzer.Run = provider.NewRun(getConfigFromSettings(lintCtx.Settings().MyErrorLint))
		})
}
