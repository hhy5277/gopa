// Generated by ego.
// DO NOT EDIT

package pprof

import (
	"fmt"
	"github.com/infinitbyte/gopa/modules/ui/common"
	"io"
	"net/http"
)

var _ = fmt.Sprint("") // just so that we can keep the fmt import for now
func Setting(w http.ResponseWriter, r *http.Request, str string) error {
	_, _ = io.WriteString(w, "\n\n")
	_, _ = io.WriteString(w, "\n")
	_, _ = io.WriteString(w, "\n\n")
	common.Head(w, "pprof", "")
	_, _ = io.WriteString(w, "\n")
	common.Body(w)
	_, _ = io.WriteString(w, "\n")
	common.Nav(w, r, "pprof")
	_, _ = io.WriteString(w, "\n\n\n\n<div class=\"tm-middle\">\n    <div class=\"uk-container uk-container-center\">\n\n        <div class=\"uk-grid\" data-uk-grid-margin=\"\">\n            <div class=\"tm-sidebar uk-width-medium-1-4 uk-hidden-small uk-row-first\">\n\n\n\n            </div>\n            <div class=\"tm-main uk-width-medium-3-4\">\n\n                <div class=\"uk-alert\" ><span id=\"domain-alert\">pprof</span></div>\n\n            <article class=\"uk-article\">\n\n\n            </article>\n            <button class=\"uk-button uk-button-primary\" type=\"button\">Snapshot</button>\n            <button class=\"uk-button uk-button-primary\" type=\"button\">Compare</button>\n\n            </div>\n        </div>\n\n    </div>\n</div>\n\n")
	common.Footer(w)
	_, _ = io.WriteString(w, "\n")
	return nil
}
