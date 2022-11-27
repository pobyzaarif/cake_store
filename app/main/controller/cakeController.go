package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/pobyzaarif/cake_store/app/main/common"
	"github.com/pobyzaarif/cake_store/business"
	"github.com/pobyzaarif/cake_store/business/cake"

	goLogger "github.com/pobyzaarif/go-logger/logger"
)

var cclogger = goLogger.NewLog("CAKE_CONTROLLER")

func (ctrl *Controller) CakeCreate(c echo.Context) error {
	trackerID, _ := c.Get("tracker_id").(string)
	cclogger.SetTrackerID(trackerID)
	ic := business.NewInternalContext(trackerID)

	request := new(cake.CakeCreateSpec)
	if err := c.Bind(request); err != nil {
		cclogger.Error(common.ErrorBindingRequest.String(), err)
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(common.EmptyObject))
	}

	res, err := ctrl.cakeService.Create(ic, *request)
	if err != nil {
		if err == business.ErrInvalidSpec {
			cclogger.Error(common.ErrorValidationRequest.String(), err)
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(common.EmptyObject))
		}
		cclogger.Error(common.ErrorGeneral.String(), err)
		return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse(common.EmptyObject))
	}

	response := common.NewResponseCreated(res)

	return c.JSON(http.StatusOK, response)
}

func (ctrl *Controller) CakeFindAll(c echo.Context) error {
	trackerID, _ := c.Get("tracker_id").(string)
	cclogger.SetTrackerID(trackerID)
	ic := business.NewInternalContext(trackerID)

	res, err := ctrl.cakeService.FindAll(ic)
	if err != nil {
		cclogger.Error(common.ErrorGeneral.String(), err)
		return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse(common.EmptyList))
	} else if len(res) == 0 {
		return c.JSON(http.StatusOK, common.NewResponseOK([]string{}))
	}

	response := common.NewResponseOK(res)
	return c.JSON(http.StatusOK, response)
}

func (ctrl *Controller) CakeFindByID(c echo.Context) error {
	trackerID, _ := c.Get("tracker_id").(string)
	cclogger.SetTrackerID(trackerID)
	ic := business.NewInternalContext(trackerID)

	id, _ := strconv.Atoi(c.Param("id"))
	res, err := ctrl.cakeService.FindByID(ic, id)
	if err != nil {
		if err == business.ErrInvalidSpec {
			cclogger.Error(common.ErrorValidationRequest.String(), err)
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(common.EmptyObject))
		} else if err == business.ErrNotFound {
			cclogger.Error(common.ErrorGeneral.String(), err)
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse(common.EmptyObject))
		}
		cclogger.Error(common.ErrorGeneral.String(), err)
		return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse(common.EmptyObject))
	}

	response := common.NewResponseOK(res)
	return c.JSON(http.StatusOK, response)
}

func (ctrl *Controller) CakeUpdate(c echo.Context) error {
	trackerID, _ := c.Get("tracker_id").(string)
	cclogger.SetTrackerID(trackerID)
	ic := business.NewInternalContext(trackerID)

	request := new(cake.CakeUpdateSpec)
	if err := c.Bind(request); err != nil {
		cclogger.Error(common.ErrorBindingRequest.String(), err)
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(common.EmptyObject))
	}

	if c.Request().Method == "PATCH" {
		id, _ := strconv.Atoi(c.Param("id"))
		request.ID = id
	}

	err := ctrl.cakeService.Update(ic, *request)
	if err != nil {
		if err == business.ErrInvalidSpec {
			cclogger.Error(common.ErrorValidationRequest.String(), err)
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(common.EmptyObject))
		} else if err == business.ErrNotFound {
			cclogger.Error(common.ErrorGeneral.String(), err)
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse(common.EmptyObject))
		}
		cclogger.Error(common.ErrorGeneral.String(), err)
		return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse(common.EmptyObject))
	}

	response := common.NewResponseOK(request)

	return c.JSON(http.StatusOK, response)
}

func (ctrl *Controller) CakeDelete(c echo.Context) error {
	trackerID, _ := c.Get("tracker_id").(string)
	cclogger.SetTrackerID(trackerID)
	ic := business.NewInternalContext(trackerID)

	id, _ := strconv.Atoi(c.Param("id"))

	err := ctrl.cakeService.Delete(ic, id)
	if err != nil {
		if err == business.ErrInvalidSpec {
			cclogger.Error(common.ErrorValidationRequest.String(), err)
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(common.EmptyObject))
		} else if err == business.ErrNotFound {
			cclogger.Error(common.ErrorGeneral.String(), err)
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse(common.EmptyObject))
		}
		cclogger.Error(common.ErrorGeneral.String(), err)
		return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse(common.EmptyObject))
	}

	return c.JSON(http.StatusOK, common.NewResponseOK(map[string]interface{}{}))
}
