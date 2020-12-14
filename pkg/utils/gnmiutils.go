// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package utils

import (
	"encoding/json"
	"fmt"
	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/openconfig/gnmi/proto/gnmi_ext"
	"github.com/openconfig/ygot/ygot"
	"reflect"
	"strings"
)

// NewGnmiGetRequest creates a GetRequest from a REST call
func NewGnmiGetRequest(openapiPath string, target string, pathParams ...string) (*gnmi.GetRequest, error) {
	gnmiGet := new(gnmi.GetRequest)
	gnmiGet.Path = make([]*gnmi.Path, 1)
	elems, err := BuildElems(openapiPath, 4, pathParams...)
	if err != nil {
		return nil, fmt.Errorf("error creating new update set request %v", err)
	}

	gnmiGet.Path[0] = &gnmi.Path{
		Elem:   elems,
		Target: target,
	}
	return gnmiGet, nil
}

// GetResponseUpdate -- extract the single Update from the GetResponse
func GetResponseUpdate(gr *gnmi.GetResponse, err error) (*gnmi.TypedValue_JsonVal, error) {
	if err != nil {
		return nil, err
	}
	if len(gr.Notification) != 1 {
		return nil, fmt.Errorf("unexpected number of GetResponse notifications %d", len(gr.Notification))
	}
	n0 := gr.Notification[0]
	if len(n0.Update) != 1 {
		return nil, fmt.Errorf("unexpected number of GetResponse notification updates %d", len(n0.Update))
	}
	u0 := n0.Update[0]
	if u0.Val == nil {
		return nil, nil
	}
	switch valueTyped := u0.Val.Value.(type) {
	case *gnmi.TypedValue_JsonVal:
		return valueTyped, nil
	default:
		return nil, fmt.Errorf("unhandled - non JsonVal response from onos-config %v", valueTyped)
	}
}

// NewGnmiSetDeleteRequest a single delete in a Set request
func NewGnmiSetDeleteRequest(openapiPath string, target string, pathParams ...string) (*gnmi.SetRequest, error) {
	gnmiSet := new(gnmi.SetRequest)
	gnmiSet.Delete = make([]*gnmi.Path, 1)
	elems, err := BuildElems(openapiPath, 4, pathParams...)
	if err != nil {
		return nil, fmt.Errorf("error creating new update set request %v", err)
	}

	gnmiSet.Delete[0] = &gnmi.Path{
		Elem:   elems,
		Target: target,
	}
	return gnmiSet, nil
}

// NewGnmiSetUpdateRequest a single delete in a Set request
// Deprecated
func NewGnmiSetUpdateRequest(openapiPath string, target string, gnmiObj interface{},
	pathParams ...string) (*gnmi.SetRequest, error) {

	gnmiSet := new(gnmi.SetRequest)
	gnmiSet.Extension = buildExtensions(openapiPath)
	gnmiSet.Update = make([]*gnmi.Update, 1)
	elems, err := BuildElems(openapiPath, 4, pathParams...)
	if err != nil {
		return nil, fmt.Errorf("error creating new update set request %v", err)
	}
	gnmiJSONVal, err := json.Marshal(gnmiObj)
	if err != nil {
		return nil, fmt.Errorf("error marshalling gNMI obj to JSON %v", err)
	}
	gnmiSet.Update[0] = &gnmi.Update{
		Path: &gnmi.Path{
			Elem:   elems,
			Target: target,
		},
		Val: &gnmi.TypedValue{
			Value: &gnmi.TypedValue_JsonVal{JsonVal: []byte(strings.ToLower(string(gnmiJSONVal)))},
		},
	}

	return gnmiSet, nil
}

// NewGnmiSetUpdateRequestUpdates a single delete in a Set request
func NewGnmiSetUpdateRequestUpdates(openapiPath string, target string,
	update []*gnmi.Update, pathParams ...string) (*gnmi.SetRequest, error) {

	gnmiSet := new(gnmi.SetRequest)
	gnmiSet.Extension = buildExtensions(openapiPath)
	gnmiSet.Update = make([]*gnmi.Update, 1)
	elems, err := BuildElems(openapiPath, 4, pathParams...)
	if err != nil {
		return nil, fmt.Errorf("error creating new update set request %v", err)
	}
	gnmiSet.Prefix = &gnmi.Path{
		Elem:   elems,
		Target: target,
	}
	gnmiSet.Update = update

	return gnmiSet, nil
}

// ExtractExtension100 - the name of the change will be returned as extension 100
func ExtractExtension100(gnmiResponse *gnmi.SetResponse) *string {
	for _, ext := range gnmiResponse.Extension {
		switch extTyped := ext.Ext.(type) {
		case *gnmi_ext.Extension_RegisteredExt:
			if extTyped.RegisteredExt.Id == 100 {
				changeName := string(extTyped.RegisteredExt.Msg)
				return &changeName
			}
		}
	}

	return nil
}

// BuildElems - create a set of gnmi PathElems
// For start at this is the element in the path at the ith position (remembering that 0 is empty)
func BuildElems(openapiPath string, startAt int, pathParams ...string) ([]*gnmi.PathElem, error) {
	if !strings.HasPrefix(openapiPath, "/") {
		return nil, fmt.Errorf("openapipath must begin with '/'. Got %s", openapiPath)
	}
	oapiParts := strings.Split(openapiPath, "/")
	if len(oapiParts) < startAt+1 {
		return nil, fmt.Errorf("expected path to have >= %d parts e.g. api,ver,device,path Got %v", startAt, oapiParts)
	}
	elemCount := 0
	paramCount := 0
	elems := make([]*gnmi.PathElem, 0)

	for i := startAt; i < len(oapiParts); i++ {
		if strings.Contains(oapiParts[i], "{") { // Is a key
			keyName := oapiParts[i]
			keyName = keyName[1 : len(keyName)-1]
			if elems[elemCount-1].Key == nil {
				elems[elemCount-1].Key = make(map[string]string)
			}
			elems[elemCount-1].Key[keyName] = pathParams[paramCount]
			paramCount++
		} else {
			pathElem := gnmi.PathElem{
				Name: oapiParts[i],
			}
			elems = append(elems, &pathElem)
			elemCount++
		}
	}

	return elems, nil
}

func buildExtensions(openapiPath string) []*gnmi_ext.Extension {
	oapiParts := strings.Split(openapiPath, "/")
	if len(oapiParts) < 3 {
		return nil
	}
	// First 2 fields should give us the modelType and modelVersion
	modelType := strings.Title(oapiParts[1]) // Change to title case
	modelVersion := oapiParts[2][1:]         // Remove the "v" at the start of "v1.0.0"

	extensions := make([]*gnmi_ext.Extension, 0)
	if modelVersion != "" {
		ext101 := gnmi_ext.Extension{
			Ext: &gnmi_ext.Extension_RegisteredExt{
				RegisteredExt: &gnmi_ext.RegisteredExtension{
					Id:  101,
					Msg: []byte(modelVersion),
				},
			},
		}
		extensions = append(extensions, &ext101)
	}
	if modelType != "" {
		ext102 := gnmi_ext.Extension{
			Ext: &gnmi_ext.Extension_RegisteredExt{
				RegisteredExt: &gnmi_ext.RegisteredExtension{
					Id:  102,
					Msg: []byte(modelType),
				},
			},
		}
		extensions = append(extensions, &ext102)
	}
	return extensions
}

// UpdateForElement -- create a gnmi.Update for a Json element
func UpdateForElement(value interface{}, path string, pathParams ...string) (*gnmi.Update, error) {
	reflectValue := reflect.ValueOf(value)
	update := new(gnmi.Update)
	update.Path = new(gnmi.Path)
	var err error
	if update.Path.Elem, err = BuildElems(path, 1, pathParams...); err != nil {
		return nil, err
	}
	update.Val = new(gnmi.TypedValue)

	switch reflectValue.Type().String() {
	case "*string":
		update.Val.Value = &gnmi.TypedValue_StringVal{StringVal: reflect.Indirect(reflectValue).String()}
	case "[]string":
		valueStrArr := value.([]string)
		llVals := make([]*gnmi.TypedValue, 0)
		for _, str := range valueStrArr {
			llVal := gnmi.TypedValue{
				Value: &gnmi.TypedValue_StringVal{StringVal: str},
			}
			llVals = append(llVals, &llVal)
		}
		update.Val.Value = &gnmi.TypedValue_LeaflistVal{
			LeaflistVal: &gnmi.ScalarArray{
				Element: llVals,
			},
		}
	case "*uint32":
		update.Val.Value = &gnmi.TypedValue_UintVal{UintVal: reflect.Indirect(reflectValue).Uint()}
	case "*bool":
		update.Val.Value = &gnmi.TypedValue_BoolVal{BoolVal: reflect.Indirect(reflectValue).Bool()}
	default:
		switch reflectValue.Kind().String() {
		case "int64":
			update.Val.Value = &gnmi.TypedValue_IntVal{IntVal: reflect.Indirect(reflectValue).Int()}
		default:
			n := reflectValue.Type().String()
			k := reflectValue.Type().Kind()
			return nil, fmt.Errorf("unhandled type %s %v", n, k)
		}
	}

	return update, nil
}

// ExtractGnmiAttribute given a YGOT gNMI object decode the given attribute 'attr'
// at path 'parent' using keys 'params'
func ExtractGnmiAttribute(modelPluginDevice interface{}, parent string, attr string, params []string) (string, error) {
	fmt.Printf("testing %T %s %s\n", modelPluginDevice, parent, attr)
	parentParts := strings.Split(parent, "_")
	parentParts = append(parentParts, attr)
	return recurseGnmiPath(modelPluginDevice, parentParts[1:], params)
}

func recurseGnmiPath(element interface{}, pathParts []string, params []string) (string, error) {
	skipPathParts := 0
	skipParams := 0
	value := reflect.ValueOf(element)
	var field reflect.Value
	switch value.Kind() {
	case reflect.String:
		return value.String(), nil
	case reflect.Int64:
		return fmt.Sprintf("%d", value.Int()), nil
	case reflect.Struct:
		field = value.FieldByName(pathParts[0])
		if !field.IsValid() {
			return "", fmt.Errorf("error getting fieldname %s on %v", pathParts[0], element)
		}
		skipPathParts++
	case reflect.Ptr:
		field = value.Elem()
	case reflect.Map:
		if len(params) == 0 {
			return "", fmt.Errorf("at least 1 param needed to decode map of %T", element)
		}
		p := reflect.ValueOf(params[0])
		field = value.MapIndex(p)
		if !field.IsValid() {
			return "", fmt.Errorf("error getting map index %s on %v", params[0], element)
		}
		skipParams++
	case reflect.Slice:
		values := make([]string, value.Len())
		for i := 0; i < value.Len(); i++ {
			res, err := recurseGnmiPath(value.Index(i).Interface(), pathParts[skipPathParts:], params[skipParams:])
			if err != nil {
				return "", err
			}
			values[i] = res
		}
		return strings.Join(values, "\n"), nil
	default:
		return "", fmt.Errorf("unhandled %v", value.Kind())
	}

	return recurseGnmiPath(field.Interface(), pathParts[skipPathParts:], params[skipParams:])
}

// ExtractGnmiListKeyMap - get the keys of a map
func ExtractGnmiListKeyMap(gnmiElement interface{}) (map[string]interface{}, error) {
	valuesMap := make(map[string]interface{})
	value := reflect.ValueOf(gnmiElement)
	keysMethod := value.MethodByName("ΛListKeyMap")
	if !keysMethod.IsZero() {
		methodReturn := keysMethod.Call(make([]reflect.Value, 0))
		if len(methodReturn) != 2 {
			return nil, fmt.Errorf("expecting 2 values back from method ΛListKeyMap")
		}
		if !methodReturn[1].IsNil() {
			return nil, fmt.Errorf("error calling ΛListKeyMap")
		}
		yangListKeysIf := methodReturn[0].Interface()
		yangListKeys, ok := yangListKeysIf.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("unable to cast to a map")
		}
		return yangListKeys, nil
	}

	return valuesMap, nil
}

// ExtractGnmiEnumMap - extract an enum value from YGOT
func ExtractGnmiEnumMap(gnmiElement interface{}, path string, attr string, oaiValue int) (string, *ygot.EnumDefinition, error) {

	pathParts := strings.Split(path, "_")
	pathParts = append(pathParts, attr)
	enumPath := fmt.Sprintf("/%s", strings.ToLower(strings.Join(pathParts[1:], "/")))
	value := reflect.ValueOf(gnmiElement)
	keysMethod := value.MethodByName("ΛEnumTypeMap")
	if !keysMethod.IsZero() {
		methodReturn := keysMethod.Call(make([]reflect.Value, 0))
		if len(methodReturn) != 1 {
			return "", nil, fmt.Errorf("expecting 2 values back from method ΛListKeyMap")
		}
		yangEnumTypeMapIf := methodReturn[0].Interface()
		yangEnumTypeMap, ok := yangEnumTypeMapIf.(map[string][]reflect.Type)
		if !ok {
			return "", nil, fmt.Errorf("unable to cast to a map")
		}
		enums, ok := yangEnumTypeMap[enumPath]
		if !ok {
			return "", nil, fmt.Errorf("could not find Enum %s in device enum map", enumPath)
		}
		for _, e := range enums {
			eVal := reflect.Zero(e)
			lambdaMap := eVal.MethodByName("ΛMap")
			if !lambdaMap.IsZero() {
				lambdaMapReturn := lambdaMap.Call(make([]reflect.Value, 0))
				if len(lambdaMapReturn) != 1 {
					return "", nil, fmt.Errorf("expecting 2 values back from method ΛListKeyMap")
				}
				lambdaMapIf := lambdaMapReturn[0].Interface()
				yangEnumTypeMap, ok := lambdaMapIf.(map[string]map[int64]ygot.EnumDefinition)
				if !ok {
					return "", nil, fmt.Errorf("unable to cast to a map")
				}
				mapDefs, ok := yangEnumTypeMap[e.Name()]
				if !ok {
					return "", nil, fmt.Errorf("enum %s not present", e.Name())
				}
				def, ok := mapDefs[int64(oaiValue)]
				if !ok {
					return "", nil, fmt.Errorf("value %d in enum %s not present", oaiValue, e.Name())
				}
				return e.Name(), &def, nil
			}
		}
		return "", nil, fmt.Errorf("expected to find enum values")
	}

	return "", nil, fmt.Errorf("expected to find enum values")
}
