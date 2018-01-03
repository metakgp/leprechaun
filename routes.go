package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
    Route{
        "BeginAuth",
        "POST",
        "/auth",
        BeginAuth,
    },
    Route{
        "VerifyStep1",
        "GET",
        "/verify1/{token}",
        VerifyStep1,
    },
}
