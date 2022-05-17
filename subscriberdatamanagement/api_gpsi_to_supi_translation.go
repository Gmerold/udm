// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service
 *
 * API version: 2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package subscriberdatamanagement

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/omec-project/http_wrapper"
	"github.com/omec-project/openapi"
	"github.com/omec-project/openapi/models"
	"github.com/omec-project/udm/logger"
	"github.com/omec-project/udm/producer"
)

// GetIdTranslationResult - retrieve a UE's SUPI
func HTTPGetIdTranslationResult(c *gin.Context) {
	req := http_wrapper.NewRequest(c.Request, nil)
	req.Params["gpsi"] = c.Params.ByName("gpsi")
	req.Query.Set("SupportedFeatures", c.Query("supported-features"))

	rsp := producer.HandleGetIdTranslationResultRequest(req)

	responseBody, err := openapi.Serialize(rsp.Body, "application/json")
	if err != nil {
		logger.SdmLog.Errorln(err)
		problemDetails := models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Cause:  "SYSTEM_FAILURE",
			Detail: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, problemDetails)
	} else {
		c.Data(rsp.Status, "application/json", responseBody)
	}
}
