// Copyright 2020 Coinbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Generated by: OpenAPI Generator (https://openapi-generator.tech)

package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/coinbase/rosetta-sdk-go/models"
)

// A ConstructionAPIController binds http requests to an api service and writes the service results
// to the http response
type ConstructionAPIController struct {
	service ConstructionAPIServicer
}

// NewConstructionAPIController creates a default api controller
func NewConstructionAPIController(s ConstructionAPIServicer) Router {
	return &ConstructionAPIController{service: s}
}

// Routes returns all of the api route for the ConstructionAPIController
func (c *ConstructionAPIController) Routes() Routes {
	return Routes{
		{
			"TransactionConstruction",
			strings.ToUpper("Post"),
			"/construction/metadata",
			c.TransactionConstruction,
		},
		{
			"TransactionSubmit",
			strings.ToUpper("Post"),
			"/construction/submit",
			c.TransactionSubmit,
		},
	}
}

// TransactionConstruction - Get Transaction Construction Metadata
func (c *ConstructionAPIController) TransactionConstruction(
	w http.ResponseWriter,
	r *http.Request,
) {
	transactionConstructionRequest := &models.TransactionConstructionRequest{}
	// TODO: Assert required params are present
	if err := json.NewDecoder(r.Body).Decode(&transactionConstructionRequest); err != nil {
		err = EncodeJSONResponse(&models.Error{
			Message: err.Error(),
		}, http.StatusBadRequest, w)
		if err != nil {
			log.Fatal(err)
		}

		return
	}

	result, serviceErr := c.service.TransactionConstruction(*transactionConstructionRequest)
	if serviceErr != nil {
		err := EncodeJSONResponse(serviceErr, http.StatusInternalServerError, w)
		if err != nil {
			log.Fatal(err)
		}

		return
	}

	if err := EncodeJSONResponse(result, http.StatusOK, w); err != nil {
		log.Fatal(err)
	}
}

// TransactionSubmit - Submit a Signed Transaction
func (c *ConstructionAPIController) TransactionSubmit(w http.ResponseWriter, r *http.Request) {
	transactionSubmitRequest := &models.TransactionSubmitRequest{}
	// TODO: Assert required params are present
	if err := json.NewDecoder(r.Body).Decode(&transactionSubmitRequest); err != nil {
		err = EncodeJSONResponse(&models.Error{
			Message: err.Error(),
		}, http.StatusBadRequest, w)
		if err != nil {
			log.Fatal(err)
		}

		return
	}

	result, serviceErr := c.service.TransactionSubmit(*transactionSubmitRequest)
	if serviceErr != nil {
		err := EncodeJSONResponse(serviceErr, http.StatusInternalServerError, w)
		if err != nil {
			log.Fatal(err)
		}

		return
	}

	if err := EncodeJSONResponse(result, http.StatusOK, w); err != nil {
		log.Fatal(err)
	}
}
