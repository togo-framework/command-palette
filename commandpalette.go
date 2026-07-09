// Package commandpalette is the Command Palette togo plugin: A cmd-K Raycast-style overlay: client-side fuzzy ranking over cached resource lists + quick actions (theme, locale, CV, contact), recents, full keyboard nav, en/ar + RTL.
//
// It self-registers a provider on blank-import and mounts its routes onto the
// kernel. The concrete implementation is ported from the fadymondy.com app under
// internal/server — this scaffold wires the provider + a health route.
package commandpalette

import (
	"net/http"

	"github.com/togo-framework/togo"
)

func init() {
	togo.RegisterProviderFunc("command-palette", togo.PriorityService, func(k *togo.Kernel) error {
		k.Router.Get("/api/command-palette/health", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"plugin":"command-palette","status":"ok"}`))
		})
		return nil
	})
}
