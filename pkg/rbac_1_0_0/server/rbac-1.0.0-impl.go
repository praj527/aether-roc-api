// Code generated by oapi-codegen. DO NOT EDIT.
// Package server provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package server

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// server-interface template override

import (
	"github.com/onosproject/aether-roc-api/pkg/rbac_1_0_0/types"
	"github.com/onosproject/aether-roc-api/pkg/southbound"
	"github.com/onosproject/aether-roc-api/pkg/utils"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"reflect"
)

// Implement the Server Interface for access to gNMI
var log = logging.GetLogger("rbac")

// ServerImpl -
type ServerImpl struct {
	GnmiClient southbound.GnmiClient
}

// DeleteRbacV100targetRbac impl of gNMI access at /rbac/v1.0.0/{target}/rbac
func (i *ServerImpl) DeleteRbacV100targetRbac(ctx echo.Context, target types.Target) error {

	var response interface{}
	var err error

	// Response
	err = i.gnmiDeleteRbacV100targetRbac(context.Background(), "/rbac/v1.0.0/{target}/rbac", target)

	if err != nil {
		if strings.HasPrefix(err.Error(), "rpc error: code = Internal desc = rpc error: code = InvalidArgument") {
			return echo.NewHTTPError(http.StatusNoContent, err.Error())
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	// It's not enough to check if response==nil - see https://medium.com/@glucn/golang-an-interface-holding-a-nil-value-is-not-nil-bb151f472cc7
	if reflect.ValueOf(response).Kind() == reflect.Ptr && reflect.ValueOf(response).IsNil() {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	log.Infof("DeleteRbacV100targetRbac")
	return ctx.JSON(http.StatusOK, response)
}

// GetRbacV100targetRbac impl of gNMI access at /rbac/v1.0.0/{target}/rbac
func (i *ServerImpl) GetRbacV100targetRbac(ctx echo.Context, target types.Target) error {

	var response interface{}
	var err error

	// Response GET OK 200
	response, err = i.gnmiGetRbacV100targetRbac(context.Background(), "/rbac/v1.0.0/{target}/rbac", target)

	if err != nil {
		if strings.HasPrefix(err.Error(), "rpc error: code = Internal desc = rpc error: code = InvalidArgument") {
			return echo.NewHTTPError(http.StatusNoContent, err.Error())
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	// It's not enough to check if response==nil - see https://medium.com/@glucn/golang-an-interface-holding-a-nil-value-is-not-nil-bb151f472cc7
	if reflect.ValueOf(response).Kind() == reflect.Ptr && reflect.ValueOf(response).IsNil() {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	log.Infof("GetRbacV100targetRbac")
	return ctx.JSON(http.StatusOK, response)
}

// PostRbacV100targetRbac impl of gNMI access at /rbac/v1.0.0/{target}/rbac
func (i *ServerImpl) PostRbacV100targetRbac(ctx echo.Context, target types.Target) error {

	var response interface{}
	var err error

	// Response created

	body, err := utils.ReadRequestBody(ctx.Request().Body)
	if err != nil {
		return err
	}
	extension100, err := i.gnmiPostRbacV100targetRbac(context.Background(), body, "/rbac/v1.0.0/{target}/rbac", target)
	if err == nil {
		log.Infof("Post succeded %s", *extension100)
		return ctx.JSON(http.StatusOK, extension100)
	}

	if err != nil {
		if strings.HasPrefix(err.Error(), "rpc error: code = Internal desc = rpc error: code = InvalidArgument") {
			return echo.NewHTTPError(http.StatusNoContent, err.Error())
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	// It's not enough to check if response==nil - see https://medium.com/@glucn/golang-an-interface-holding-a-nil-value-is-not-nil-bb151f472cc7
	if reflect.ValueOf(response).Kind() == reflect.Ptr && reflect.ValueOf(response).IsNil() {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	log.Infof("PostRbacV100targetRbac")
	return ctx.JSON(http.StatusOK, response)
}

// DeleteRbacV100targetRbacGroup impl of gNMI access at /rbac/v1.0.0/{target}/rbac/group/{groupid}
func (i *ServerImpl) DeleteRbacV100targetRbacGroup(ctx echo.Context, target types.Target, groupid string) error {

	var response interface{}
	var err error

	// Response
	err = i.gnmiDeleteRbacV100targetRbacGroup(context.Background(), "/rbac/v1.0.0/{target}/rbac/group/{groupid}", target, groupid)

	if err != nil {
		if strings.HasPrefix(err.Error(), "rpc error: code = Internal desc = rpc error: code = InvalidArgument") {
			return echo.NewHTTPError(http.StatusNoContent, err.Error())
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	// It's not enough to check if response==nil - see https://medium.com/@glucn/golang-an-interface-holding-a-nil-value-is-not-nil-bb151f472cc7
	if reflect.ValueOf(response).Kind() == reflect.Ptr && reflect.ValueOf(response).IsNil() {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	log.Infof("DeleteRbacV100targetRbacGroup")
	return ctx.JSON(http.StatusOK, response)
}

// GetRbacV100targetRbacGroup impl of gNMI access at /rbac/v1.0.0/{target}/rbac/group/{groupid}
func (i *ServerImpl) GetRbacV100targetRbacGroup(ctx echo.Context, target types.Target, groupid string) error {

	var response interface{}
	var err error

	// Response GET OK 200
	response, err = i.gnmiGetRbacV100targetRbacGroup(context.Background(), "/rbac/v1.0.0/{target}/rbac/group/{groupid}", target, groupid)

	if err != nil {
		if strings.HasPrefix(err.Error(), "rpc error: code = Internal desc = rpc error: code = InvalidArgument") {
			return echo.NewHTTPError(http.StatusNoContent, err.Error())
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	// It's not enough to check if response==nil - see https://medium.com/@glucn/golang-an-interface-holding-a-nil-value-is-not-nil-bb151f472cc7
	if reflect.ValueOf(response).Kind() == reflect.Ptr && reflect.ValueOf(response).IsNil() {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	log.Infof("GetRbacV100targetRbacGroup")
	return ctx.JSON(http.StatusOK, response)
}

// PostRbacV100targetRbacGroup impl of gNMI access at /rbac/v1.0.0/{target}/rbac/group/{groupid}
func (i *ServerImpl) PostRbacV100targetRbacGroup(ctx echo.Context, target types.Target, groupid string) error {

	var response interface{}
	var err error

	// Response created

	body, err := utils.ReadRequestBody(ctx.Request().Body)
	if err != nil {
		return err
	}
	extension100, err := i.gnmiPostRbacV100targetRbacGroup(context.Background(), body, "/rbac/v1.0.0/{target}/rbac/group/{groupid}", target, groupid)
	if err == nil {
		log.Infof("Post succeded %s", *extension100)
		return ctx.JSON(http.StatusOK, extension100)
	}

	if err != nil {
		if strings.HasPrefix(err.Error(), "rpc error: code = Internal desc = rpc error: code = InvalidArgument") {
			return echo.NewHTTPError(http.StatusNoContent, err.Error())
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	// It's not enough to check if response==nil - see https://medium.com/@glucn/golang-an-interface-holding-a-nil-value-is-not-nil-bb151f472cc7
	if reflect.ValueOf(response).Kind() == reflect.Ptr && reflect.ValueOf(response).IsNil() {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	log.Infof("PostRbacV100targetRbacGroup")
	return ctx.JSON(http.StatusOK, response)
}

// DeleteRbacV100targetRbacGroupRole impl of gNMI access at /rbac/v1.0.0/{target}/rbac/group/{groupid}/role/{roleid}
func (i *ServerImpl) DeleteRbacV100targetRbacGroupRole(ctx echo.Context, target types.Target, groupid string, roleid string) error {

	var response interface{}
	var err error

	// Response
	err = i.gnmiDeleteRbacV100targetRbacGroupRole(context.Background(), "/rbac/v1.0.0/{target}/rbac/group/{groupid}/role/{roleid}", target, groupid, roleid)

	if err != nil {
		if strings.HasPrefix(err.Error(), "rpc error: code = Internal desc = rpc error: code = InvalidArgument") {
			return echo.NewHTTPError(http.StatusNoContent, err.Error())
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	// It's not enough to check if response==nil - see https://medium.com/@glucn/golang-an-interface-holding-a-nil-value-is-not-nil-bb151f472cc7
	if reflect.ValueOf(response).Kind() == reflect.Ptr && reflect.ValueOf(response).IsNil() {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	log.Infof("DeleteRbacV100targetRbacGroupRole")
	return ctx.JSON(http.StatusOK, response)
}

// GetRbacV100targetRbacGroupRole impl of gNMI access at /rbac/v1.0.0/{target}/rbac/group/{groupid}/role/{roleid}
func (i *ServerImpl) GetRbacV100targetRbacGroupRole(ctx echo.Context, target types.Target, groupid string, roleid string) error {

	var response interface{}
	var err error

	// Response GET OK 200
	response, err = i.gnmiGetRbacV100targetRbacGroupRole(context.Background(), "/rbac/v1.0.0/{target}/rbac/group/{groupid}/role/{roleid}", target, groupid, roleid)

	if err != nil {
		if strings.HasPrefix(err.Error(), "rpc error: code = Internal desc = rpc error: code = InvalidArgument") {
			return echo.NewHTTPError(http.StatusNoContent, err.Error())
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	// It's not enough to check if response==nil - see https://medium.com/@glucn/golang-an-interface-holding-a-nil-value-is-not-nil-bb151f472cc7
	if reflect.ValueOf(response).Kind() == reflect.Ptr && reflect.ValueOf(response).IsNil() {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	log.Infof("GetRbacV100targetRbacGroupRole")
	return ctx.JSON(http.StatusOK, response)
}

// PostRbacV100targetRbacGroupRole impl of gNMI access at /rbac/v1.0.0/{target}/rbac/group/{groupid}/role/{roleid}
func (i *ServerImpl) PostRbacV100targetRbacGroupRole(ctx echo.Context, target types.Target, groupid string, roleid string) error {

	var response interface{}
	var err error

	// Response created

	body, err := utils.ReadRequestBody(ctx.Request().Body)
	if err != nil {
		return err
	}
	extension100, err := i.gnmiPostRbacV100targetRbacGroupRole(context.Background(), body, "/rbac/v1.0.0/{target}/rbac/group/{groupid}/role/{roleid}", target, groupid, roleid)
	if err == nil {
		log.Infof("Post succeded %s", *extension100)
		return ctx.JSON(http.StatusOK, extension100)
	}

	if err != nil {
		if strings.HasPrefix(err.Error(), "rpc error: code = Internal desc = rpc error: code = InvalidArgument") {
			return echo.NewHTTPError(http.StatusNoContent, err.Error())
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	// It's not enough to check if response==nil - see https://medium.com/@glucn/golang-an-interface-holding-a-nil-value-is-not-nil-bb151f472cc7
	if reflect.ValueOf(response).Kind() == reflect.Ptr && reflect.ValueOf(response).IsNil() {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	log.Infof("PostRbacV100targetRbacGroupRole")
	return ctx.JSON(http.StatusOK, response)
}

// DeleteRbacV100targetRbacRole impl of gNMI access at /rbac/v1.0.0/{target}/rbac/role/{roleid}
func (i *ServerImpl) DeleteRbacV100targetRbacRole(ctx echo.Context, target types.Target, roleid string) error {

	var response interface{}
	var err error

	// Response
	err = i.gnmiDeleteRbacV100targetRbacRole(context.Background(), "/rbac/v1.0.0/{target}/rbac/role/{roleid}", target, roleid)

	if err != nil {
		if strings.HasPrefix(err.Error(), "rpc error: code = Internal desc = rpc error: code = InvalidArgument") {
			return echo.NewHTTPError(http.StatusNoContent, err.Error())
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	// It's not enough to check if response==nil - see https://medium.com/@glucn/golang-an-interface-holding-a-nil-value-is-not-nil-bb151f472cc7
	if reflect.ValueOf(response).Kind() == reflect.Ptr && reflect.ValueOf(response).IsNil() {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	log.Infof("DeleteRbacV100targetRbacRole")
	return ctx.JSON(http.StatusOK, response)
}

// GetRbacV100targetRbacRole impl of gNMI access at /rbac/v1.0.0/{target}/rbac/role/{roleid}
func (i *ServerImpl) GetRbacV100targetRbacRole(ctx echo.Context, target types.Target, roleid string) error {

	var response interface{}
	var err error

	// Response GET OK 200
	response, err = i.gnmiGetRbacV100targetRbacRole(context.Background(), "/rbac/v1.0.0/{target}/rbac/role/{roleid}", target, roleid)

	if err != nil {
		if strings.HasPrefix(err.Error(), "rpc error: code = Internal desc = rpc error: code = InvalidArgument") {
			return echo.NewHTTPError(http.StatusNoContent, err.Error())
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	// It's not enough to check if response==nil - see https://medium.com/@glucn/golang-an-interface-holding-a-nil-value-is-not-nil-bb151f472cc7
	if reflect.ValueOf(response).Kind() == reflect.Ptr && reflect.ValueOf(response).IsNil() {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	log.Infof("GetRbacV100targetRbacRole")
	return ctx.JSON(http.StatusOK, response)
}

// PostRbacV100targetRbacRole impl of gNMI access at /rbac/v1.0.0/{target}/rbac/role/{roleid}
func (i *ServerImpl) PostRbacV100targetRbacRole(ctx echo.Context, target types.Target, roleid string) error {

	var response interface{}
	var err error

	// Response created

	body, err := utils.ReadRequestBody(ctx.Request().Body)
	if err != nil {
		return err
	}
	extension100, err := i.gnmiPostRbacV100targetRbacRole(context.Background(), body, "/rbac/v1.0.0/{target}/rbac/role/{roleid}", target, roleid)
	if err == nil {
		log.Infof("Post succeded %s", *extension100)
		return ctx.JSON(http.StatusOK, extension100)
	}

	if err != nil {
		if strings.HasPrefix(err.Error(), "rpc error: code = Internal desc = rpc error: code = InvalidArgument") {
			return echo.NewHTTPError(http.StatusNoContent, err.Error())
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	// It's not enough to check if response==nil - see https://medium.com/@glucn/golang-an-interface-holding-a-nil-value-is-not-nil-bb151f472cc7
	if reflect.ValueOf(response).Kind() == reflect.Ptr && reflect.ValueOf(response).IsNil() {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	log.Infof("PostRbacV100targetRbacRole")
	return ctx.JSON(http.StatusOK, response)
}

// DeleteRbacV100targetRbacRolePermission impl of gNMI access at /rbac/v1.0.0/{target}/rbac/role/{roleid}/permission
func (i *ServerImpl) DeleteRbacV100targetRbacRolePermission(ctx echo.Context, target types.Target, roleid string) error {

	var response interface{}
	var err error

	// Response
	err = i.gnmiDeleteRbacV100targetRbacRolePermission(context.Background(), "/rbac/v1.0.0/{target}/rbac/role/{roleid}/permission", target, roleid)

	if err != nil {
		if strings.HasPrefix(err.Error(), "rpc error: code = Internal desc = rpc error: code = InvalidArgument") {
			return echo.NewHTTPError(http.StatusNoContent, err.Error())
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	// It's not enough to check if response==nil - see https://medium.com/@glucn/golang-an-interface-holding-a-nil-value-is-not-nil-bb151f472cc7
	if reflect.ValueOf(response).Kind() == reflect.Ptr && reflect.ValueOf(response).IsNil() {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	log.Infof("DeleteRbacV100targetRbacRolePermission")
	return ctx.JSON(http.StatusOK, response)
}

// GetRbacV100targetRbacRolePermission impl of gNMI access at /rbac/v1.0.0/{target}/rbac/role/{roleid}/permission
func (i *ServerImpl) GetRbacV100targetRbacRolePermission(ctx echo.Context, target types.Target, roleid string) error {

	var response interface{}
	var err error

	// Response GET OK 200
	response, err = i.gnmiGetRbacV100targetRbacRolePermission(context.Background(), "/rbac/v1.0.0/{target}/rbac/role/{roleid}/permission", target, roleid)

	if err != nil {
		if strings.HasPrefix(err.Error(), "rpc error: code = Internal desc = rpc error: code = InvalidArgument") {
			return echo.NewHTTPError(http.StatusNoContent, err.Error())
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	// It's not enough to check if response==nil - see https://medium.com/@glucn/golang-an-interface-holding-a-nil-value-is-not-nil-bb151f472cc7
	if reflect.ValueOf(response).Kind() == reflect.Ptr && reflect.ValueOf(response).IsNil() {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	log.Infof("GetRbacV100targetRbacRolePermission")
	return ctx.JSON(http.StatusOK, response)
}

// PostRbacV100targetRbacRolePermission impl of gNMI access at /rbac/v1.0.0/{target}/rbac/role/{roleid}/permission
func (i *ServerImpl) PostRbacV100targetRbacRolePermission(ctx echo.Context, target types.Target, roleid string) error {

	var response interface{}
	var err error

	// Response created

	body, err := utils.ReadRequestBody(ctx.Request().Body)
	if err != nil {
		return err
	}
	extension100, err := i.gnmiPostRbacV100targetRbacRolePermission(context.Background(), body, "/rbac/v1.0.0/{target}/rbac/role/{roleid}/permission", target, roleid)
	if err == nil {
		log.Infof("Post succeded %s", *extension100)
		return ctx.JSON(http.StatusOK, extension100)
	}

	if err != nil {
		if strings.HasPrefix(err.Error(), "rpc error: code = Internal desc = rpc error: code = InvalidArgument") {
			return echo.NewHTTPError(http.StatusNoContent, err.Error())
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	// It's not enough to check if response==nil - see https://medium.com/@glucn/golang-an-interface-holding-a-nil-value-is-not-nil-bb151f472cc7
	if reflect.ValueOf(response).Kind() == reflect.Ptr && reflect.ValueOf(response).IsNil() {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	log.Infof("PostRbacV100targetRbacRolePermission")
	return ctx.JSON(http.StatusOK, response)
}

// register template override
