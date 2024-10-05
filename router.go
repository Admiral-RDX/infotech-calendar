package main

import (
	handlers "infotech-calendar/handlers"
	auth_helpers "infotech-calendar/helpers"
	"infotech-calendar/web/base"
	"infotech-calendar/web/routes/calendar"
	"infotech-calendar/web/routes/dashboard"
	"infotech-calendar/web/routes/login"
	"net/http"
)

func registerLoginRoutes(mux *http.ServeMux) {
	loginFrag := login.LoginFrag()
	loginPage := base.Base(loginFrag)

	// Login page
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if auth_helpers.HasValidSession(r) {
			// Redirect authenticated users to the dashboard
			http.Redirect(w, r, "/dashboard", http.StatusFound)
			return
		}

		ctx := r.Context()
		if err := loginPage.Render(ctx, w); err != nil {
			http.Error(w, "Error rendering page", http.StatusInternalServerError)
		}
	})

	// Login submit handler
	mux.HandleFunc("/login/submit", handlers.HandleLogin)
}

func registerDashboardRoutes(mux *http.ServeMux) {
	dashboardFrag := dashboard.DashboardFrag()
	dashboardPage := base.Base(dashboardFrag)

	// Handle the dashboard page
	mux.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		if !auth_helpers.HasValidSession(r) {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		ctx := r.Context()
		dashboardPage.Render(ctx, w)
	})

	mux.HandleFunc("/dashboard/events", handlers.GetEventData)
}

func registerCalendarRoutes(mux *http.ServeMux) {
	calendarFrag := calendar.CalendarFrag()
	calendarPage := base.Base(calendarFrag)

	// Calendar page
	mux.HandleFunc("/calendar", func(w http.ResponseWriter, r *http.Request) {
		if !auth_helpers.HasValidSession(r) {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		ctx := r.Context()
		calendarPage.Render(ctx, w)
	})
}
