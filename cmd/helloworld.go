package helloworld

import (
	"net/http"

	"github.com/kobsio/kobs/pkg/kube/clusters"
	"github.com/kobsio/kobs/pkg/log"
	"github.com/kobsio/kobs/pkg/middleware/errresponse"
	"github.com/kobsio/kobs/pkg/satellite/plugins/plugin"
	"github.com/kobsio/plugin-template/pkg/instance"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"go.uber.org/zap"
)

// PluginType is the type of the plugin, how it must be specified in the configuration. The PluginType is also used as
// prefix for the returned chi.Router from the Mount function.
const PluginType = "helloworld"

// Router implements a router for the plugin. It contains all the fields and functions from the chi.Mux struct and all
// the configured instances.
type Router struct {
	*chi.Mux
	instances []instance.Instance
}

// getInstance is a helper function, which returns a instance by it's name. If we couldn't found an instance with the
// provided name and when the provided name is "default" we return the first configured instance.
func (router *Router) getInstance(name string) instance.Instance {
	for _, i := range router.instances {
		if i.GetName() == name {
			return i
		}
	}

	if name == "default" && len(router.instances) > 0 {
		return router.instances[0]
	}

	return nil
}

// getVariable is a spacial endpoint which is mounted under the "/variable" path. This endpoint can be used to use the
// plugin within the variables section of a dashboard. The endpoint must always return a slice of strings (e.g. via
// "render.JSON(w, r, values)", where values is a of type []string).
func (router *Router) getVariable(w http.ResponseWriter, r *http.Request) {
	errresponse.Render(w, r, nil, http.StatusNotImplemented, "Variable is not implemented for the helloworld plugin")
}

// getInsight is a special endpoint which is mounted under the "/insights" path. This endpoint can be used to use the
// plugin within the insights section of an application. The endpoint must always return a slice of datums, where a
// datum is defined as follows:
//   type Datum struct {
//       X int64    `json:"x"`
//       Y *float64 `json:"y"`
//   }
func (router *Router) getInsight(w http.ResponseWriter, r *http.Request) {
	errresponse.Render(w, r, nil, http.StatusNotImplemented, "Insights are not implemented for the helloworld plugin")
}

// getGreeting is a custom endpoint for the plugin, which returns the greeting message from the configuration for a
// plugin instance.
func (router *Router) getGreeting(w http.ResponseWriter, r *http.Request) {
	// We always pass the name of the plugin instance via the "x-kobs-plugin" header to all plugins. This allows us to
	// use multiple instances of an plugin within one satellite. E.g. when we are using the SQL we can use the plugin to
	// connect to multiple SQL databases.
	name := r.Header.Get("x-kobs-plugin")

	log.Debug(r.Context(), "Get greeting parameter", zap.String("name", name))

	// If the instance with the provided name was not found, we should always return an error. Normally this shouldn't
	// happen, when the frontend is implemented correctly, since a user shouldn't be able to set a custom plugin name.
	i := router.getInstance(name)
	if i == nil {
		log.Error(r.Context(), "Could not find instance name", zap.String("name", name))
		errresponse.Render(w, r, nil, http.StatusBadRequest, "Could not find instance name")
		return
	}

	data := struct {
		Greeting string `json:"string"`
	}{
		i.GetGreeting(),
	}

	render.JSON(w, r, data)
}

// Mount must be implemented by all plugins. It must return a chi.Router or an error. If no error is returned, the
// returned chi.Router is mounted under the specified PluginType route.
//
// For example the complete path for the "/greeting" endpoint can be called via the following URL in the frontend:
// "/api/plugins/helloworld/greeting".
func Mount(instances []plugin.Instance, clustersClient clusters.Client) (chi.Router, error) {
	// The following logic allows us to use multiple instances of the same plugin within one satellite. For thsi we
	// recommend to have a "instance" package for all plugins, where the real logic is implemented. The http routes
	// should then just be used for parsing request parameters or the request body and pass it to a instance function.
	//
	// For plugins which are just using the "clustersClient" to access the Kubernetes API, this doesn't make sense and
	// we recommend to not use a instance package.
	var helloworldInstances []instance.Instance

	for _, i := range instances {
		helloworldInstance, err := instance.New(i.Name, i.Options)
		if err != nil {
			return nil, err
		}

		helloworldInstances = append(helloworldInstances, helloworldInstance)
	}

	router := Router{
		chi.NewRouter(),
		helloworldInstances,
	}

	router.Post("/variable", router.getVariable)
	router.Post("/insight", router.getInsight)
	router.Get("/greeting", router.getGreeting)

	return router, nil
}
