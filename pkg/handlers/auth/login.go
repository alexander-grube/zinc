/* Copyright 2022 Zinc Labs Inc. and Contributors
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*     http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/zinclabs/zinc/pkg/auth"
	"github.com/zinclabs/zinc/pkg/meta"
)

func Login(c *gin.Context) {
	var user meta.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loggedInUser, validationResult := auth.VerifyCredentials(user.ID, user.Password)
	resUser := gin.H{}
	if validationResult {
		resUser = gin.H{
			"_id":  loggedInUser.ID,
			"name": loggedInUser.Name,
			"role": loggedInUser.Role,
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"validated": validationResult,
		"user":      resUser,
	})
}
