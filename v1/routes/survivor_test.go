package routes

import (
	"reflect"
	"serve/v1/helpers"
	"strings"
	"testing"
)

func TestSurvivorRouters(t *testing.T) {
	methods := []string{}

	for _, route := range GetSurvivorRouters() {
		if !strings.Contains(route.URI, "survivor") {
			t.Errorf("There isn't word users in uri of users routes of method %s", route.Method)
		}

		if !route.RequireAuthentication {
			t.Errorf("Should there are user authentication in method %s", route.Method)
		}

		if reflect.TypeOf(route.DomainFunc).Kind() != reflect.Func {
			t.Errorf("DomainFunc of method %s isn't a function", route.Method)
		}
		methods = append(methods, route.Method)
	}
	if !helpers.Contains(methods, "GET") || !helpers.Contains(methods, "POST") || !helpers.Contains(methods, "PUT") || !helpers.Contains(methods, "DELETE") {
		t.Errorf("There aren't all the methods of user routers")
	}

}
