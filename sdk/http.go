package sdk

import "net/http"

type appHandler func(http.ResponseWriter, *http.Request) error

/*
func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		http.Error(w, err.Error(), 500)
	}
}*/

type appError struct {
	Error   error
	Message string
	Code    int
}

/*
func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil { // e is *appError, not os.Error.
		c := appengine.NewContext(r)
		c.Errorf("%v", e.Error)
		http.Error(w, e.Message, e.Code)
	}
}



func init() {
	http.Handle("/view", appHandler(viewRecord))
}

func viewRecord(w http.ResponseWriter, r *http.Request) error {
	c := appengine.NewContext(r)
	key := datastore.NewKey(c, "Record", r.FormValue("id"), 0, nil)
	record := new(Record)
	if err := datastore.Get(c, key, record); err != nil {
		return err
	}
	return viewTemplate.Execute(w, record)
}
*/
