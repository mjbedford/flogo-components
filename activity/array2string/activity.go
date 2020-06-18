package array2string

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/metadata"
)

func init() {
	_ = activity.Register(&Activity{}, New)
}

const (
	methodPOST  = "POST"
	methodPUT   = "PUT"
	methodPATCH = "PATCH"
)

// Q query struct
type Q struct {
	Query string `json:"query"`
}

var activityMd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

// New activity
func New(ctx activity.InitContext) (activity.Activity, error) {
	s := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), s, true)
	if err != nil {
		return nil, err
	}

	act := &Activity{settings: s}

	return act, nil
}

// Activity is an activity that is used to invoke a REST Operation
// settings : {method, uri, headers, proxy, skipSSL}
// input    : {pathParams, queryParams, headers, content}
// outputs  : {status, result}
type Activity struct {
	settings *Settings
	// containsParam bool
	// client        *http.Client
}

// Metadata Activity
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Invokes a REST Operation
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}
	err = ctx.GetInputObject(input)
	if err != nil {
		return false, err
	}
	delimeter := a.settings.Delimeter
	prefix := a.settings.Prefix
	suffix := a.settings.Suffix

	// 	uri = uri + "?" + qp.Encode()
	// }

	logger := ctx.Logger()

	if logger.DebugEnabled() {
		logger.Debugf("Eval called: [%s] %s", a.settings.Delimeter, delimeter)
		logger.Debugf("Eval called: [%s] %s", a.settings.Prefix, prefix)
		logger.Debugf("Eval called: [%s] %s", a.settings.Suffix, suffix)
	}

	var result string

	if prefix != "" {
		result = prefix
	}

	for key, value := range input.InputArray {
		// qp.Set(key, value)
		logger.Debugf("Eval called: [%s] %s", a.settings.Delimeter, key)
		varType := reflect.TypeOf(value)
		t := reflect.TypeOf(varType)
		fmt.Println(strings.Repeat("\t", 1), "Type is", t.Name(), "and kind is", t.Kind())
		// var n, k string
		// n = t.Name()
		// k = t.Kind().String()

		// result = result + n + k
		nval, err := coerce.ToString(value)
		if err != nil {
			return false, err
		}
		bval := []byte(nval)
		var query Q
		if err := json.Unmarshal(bval, &query); err != nil {
			//log.Println("----------------------------------------------------")
			//
			fmt.Println(err)
			//log.Println(string(mesg))
			// shot.Query
			//log.Println("----------------------------------------------------")

		}
		fmt.Println(key)
		fmt.Println(query.Query)
		if key == 0 {
			result = result + query.Query
		} else {
			result = result + delimeter + query.Query
		}

	}
	if suffix != "" {
		result = result + suffix
	}


	output := &Output{ResultString: result} //coerce.ToString(result)}

	err = ctx.SetOutputObject(output)
	if err != nil {
		return false, err
	}

	return true, nil
}
