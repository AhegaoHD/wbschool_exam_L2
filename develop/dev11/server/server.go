package server

import (
	"develop/develop/dev11/event"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func loggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s\n", r.Method, r.URL)
		next(w, r)
	}
}

func RunServer() error {
	//GET
	http.HandleFunc("/events_for_day", loggerMiddleware(handleEvent))
	http.HandleFunc("/events_for_week", loggerMiddleware(handleEvent))
	http.HandleFunc("/events_for_month", loggerMiddleware(handleEvent))

	//POST
	http.HandleFunc("/create_event", loggerMiddleware(handleEvent))
	http.HandleFunc("/update_event", loggerMiddleware(handleEvent))
	http.HandleFunc("/delete_event", loggerMiddleware(handleEvent))

	return http.ListenAndServe("localhost:8080", nil)
}

func handleEvent(w http.ResponseWriter, r *http.Request) {
	url_resuest := r.URL.Path
	switch r.Method {
	case http.MethodGet:
		switch url_resuest {
		case "/events_for_day":
			events_for_day(w, r)
		case "/events_for_week":
			events_for_week(w, r)
		case "/events_for_month":
			events_for_month(w, r)
		}
	case http.MethodPost:
		switch url_resuest {
		case "/create_event":
			createEvent(w, r)
		case "/update_event":
			update_event(w, r)
		case "/delete_event":
			delete_event(w, r)
		}
	}
}

// GET /events_for_month
func events_for_month(w http.ResponseWriter, r *http.Request) {
	// curl -v -X GET -H "date:2023-04-21, userid:1" localhost:8080/events_for_day
	date, err := time.Parse("2006-01-02", r.URL.Query().Get("date"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userid, err := strconv.Atoi(r.URL.Query().Get("userid"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for key := range event.Events {
		if date.Month() == event.Events[key].Date.Month() && userid == event.Events[key].UserID {
			resp, err := json.Marshal(event.Events)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Write(resp)
		}
	}
}

// GET /events_for_week
func events_for_week(w http.ResponseWriter, r *http.Request) {
	// curl -v -X GET -H "date:"2023-04-21", userid:1" localhost:8080/events_for_week
	date, err := time.Parse("2006-01-02", r.URL.Query().Get("date"))
	if err != nil {
		fmt.Printf("Ошибка в парсинге даты %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userid, err := strconv.Atoi(r.URL.Query().Get("userid"))
	if err != nil {
		fmt.Printf("Ошибка в поиске event по id пользователя %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for key := range event.Events {
		if int(date.Weekday()) == int(event.Events[key].Date.Weekday()) && userid == event.Events[key].UserID {
			resp, err := json.Marshal(event.Events)
			if err != nil {
				fmt.Printf("Ошибка конвертации map в json %s\n", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Write(resp)
		}
	}
}

// GET /events_for_day
func events_for_day(w http.ResponseWriter, r *http.Request) {
	// curl -v -X GET -H "date:2023-04-21, userid:1" localhost:8080/events_for_day
	date, err := time.Parse("2006-01-02", r.URL.Query().Get("date"))
	if err != nil {
		fmt.Printf("Ошибка в парсинге даты %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userid, err := strconv.Atoi(r.URL.Query().Get("userid"))
	if err != nil {
		fmt.Printf("Ошибка в поиске event по id пользователя %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for key := range event.Events {
		if date.Day() == event.Events[key].Date.Day() && userid == event.Events[key].UserID {
			resp, err := json.Marshal(event.Events)
			if err != nil {
				fmt.Printf("Ошибка конвертации map в json %s\n", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Write(resp)
		}
	}
}

// POST /create_event
func createEvent(w http.ResponseWriter, r *http.Request) {
	// curl -v -X POST localhost:8080/create_event -d '{"event_id":1, "user_id":1, "name": "Событие №1","description": "Очень важное событие 1"}'

	// curl -v -X POST localhost:8080/create_event -d '{"event_id":1, "user_id":1, "name": "Событие №1","description": "Очень важное событие 1", "date": "2023-04-21"}'

	var e event.Event
	if err := e.Decode(r.Body); err != nil {
		log.Println("Ошибка декодирования json в структуру")
		w.WriteHeader(http.StatusBadGateway)
		return
	}

	event.Events[e.EventID] = e
}

// POST /update_event
func update_event(w http.ResponseWriter, r *http.Request) {
	// curl -v -X POST localhost:8080/update_event -d '{"event_id":1, "user_id":2, "name": "Событие №2","description": "Очень важное событие 2", "date": "2010-01-01"}'

	var e event.Event
	if err := e.Decode(r.Body); err != nil {
		log.Println("Ошибка декодирования json в структуру")
		w.WriteHeader(http.StatusBadGateway)
		return
	}

	for key := range event.Events {
		if e.EventID == event.Events[key].EventID {
			event.Events[key] = e
		}
	}
}

// POST /delete_event
func delete_event(w http.ResponseWriter, r *http.Request) {
	// curl -v -X POST localhost:8080/delete_event -d '{"event_id":1}'
	resBytest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var newEvent event.Event
	if err = json.Unmarshal(resBytest, &newEvent); err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}

	for key := range event.Events {
		if newEvent.EventID == event.Events[key].EventID {
			delete(event.Events, newEvent.EventID)
		}
	}
}
