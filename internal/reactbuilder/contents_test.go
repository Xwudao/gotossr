package reactbuilder

import (
	"strings"
	"testing"
)

func TestGenerateTanStackSPABuildContents(t *testing.T) {
	server, err := GenerateServerSPABuildContents(nil, "/app/App.tsx", "tanstack", "/app")
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		`createMemoryHistory`, `RouterProvider`, `createSSRRouter`,
		`initialEntries: [props.__requestPath || "/"]`,
	} {
		if !strings.Contains(server, want) {
			t.Errorf("server build content missing %q:\n%s", want, server)
		}
	}

	client, err := GenerateClientSPABuildContents(nil, "/app/App.tsx", "tanstack")
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		`createBrowserHistory`, `hydrateRoot`, `RouterProvider`,
		`createSSRRouter({ history: createBrowserHistory(), props: ssrProps })`,
	} {
		if !strings.Contains(client, want) {
			t.Errorf("client build content missing %q:\n%s", want, client)
		}
	}
}
