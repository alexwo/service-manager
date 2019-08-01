/*
 * Copyright 2018 The Service Manager Authors
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

package test

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gavv/httpexpect"
	"github.com/gofrs/uuid"
	"strings"
	"time"

	"github.com/Peripli/service-manager/pkg/query"

	. "github.com/onsi/ginkgo/extensions/table"

	"net/http"

	. "github.com/onsi/ginkgo"

	"github.com/Peripli/service-manager/test/common"
)

type listOpEntry struct {
	resourcesToExpectBeforeOp []common.Object

	queryTemplate               string
	queryArgs                   common.Object
	resourcesToExpectAfterOp    []common.Object
	resourcesNotToExpectAfterOp []common.Object
	expectedStatusCode          int
}

func DescribeListTestsFor(ctx *common.TestContext, t TestCase) bool {
	var r []common.Object
	var rWithMandatoryFields common.Object

	attachLabel := func(obj common.Object) common.Object {
		patchLabelsBody := make(map[string]interface{})
		patchLabels := []query.LabelChange{
			{
				Operation: query.AddLabelOperation,
				Key:       "labelKey1",
				Values:    []string{"1"},
			},
			{
				Operation: query.AddLabelOperation,
				Key:       "labelKey2",
				Values:    []string{"str"},
			},
			{
				Operation: query.AddLabelOperation,
				Key:       "labelKey3",
				Values:    []string{`{"key1": "val1", "key2": "val2"}`},
			},
		}
		patchLabelsBody["labels"] = patchLabels

		By(fmt.Sprintf("Attempting to patch resource of %s with labels as labels are declared supported", t.API))
		ctx.SMWithOAuth.PATCH(t.API + "/" + obj["id"].(string)).WithJSON(patchLabelsBody).
			Expect().
			Status(http.StatusOK)

		result := ctx.SMWithOAuth.GET(t.API + "/" + obj["id"].(string)).
			Expect().
			Status(http.StatusOK).JSON().Object()
		result.ContainsKey("labels")
		r := result.Raw()
		return r
	}

	By(fmt.Sprintf("Attempting to create a random resource of %s with mandatory fields only", t.API))
	rWithMandatoryFields = t.ResourceWithoutNullableFieldsBlueprint(ctx, ctx.SMWithOAuth)
	for i := 0; i < 10; i++ {
		By(fmt.Sprintf("Attempting to create a random resource of %s", t.API))

		gen := t.ResourceBlueprint(ctx, ctx.SMWithOAuth)
		gen = attachLabel(gen)
		delete(gen, "created_at")
		delete(gen, "updated_at")
		r = append(r, gen)
	}

	entries := []TableEntry{
		Entry("returns 200",
			listOpEntry{
				resourcesToExpectBeforeOp: []common.Object{r[0]},
				queryTemplate:             "%s = %v",
				queryArgs:                 r[0],
				resourcesToExpectAfterOp:  []common.Object{r[0]},
				expectedStatusCode:        http.StatusOK,
			},
		),
		Entry("returns 200",
			listOpEntry{
				resourcesToExpectBeforeOp:   []common.Object{r[0], r[1], r[2], r[3]},
				queryTemplate:               "%s != %v",
				queryArgs:                   r[0],
				resourcesNotToExpectAfterOp: []common.Object{r[0]},
				expectedStatusCode:          http.StatusOK,
			},
		),

		Entry("returns 200",
			listOpEntry{
				resourcesToExpectBeforeOp: []common.Object{r[0], r[1], r[2], r[3]},
				queryTemplate:             "%[1]s in [%[2]v||%[2]v||%[2]v]",
				queryArgs:                 r[0],
				resourcesToExpectAfterOp:  []common.Object{r[0]},
				expectedStatusCode:        http.StatusOK,
			},
		),

		Entry("returns 200",
			listOpEntry{
				resourcesToExpectBeforeOp: []common.Object{r[0], r[1], r[2], r[3]},
				queryTemplate:             "%s in [%v]",
				queryArgs:                 r[0],
				resourcesToExpectAfterOp:  []common.Object{r[0]},
				expectedStatusCode:        http.StatusOK,
			},
		),
		Entry("returns 200",
			listOpEntry{
				resourcesToExpectBeforeOp:   []common.Object{r[0], r[1], r[2], r[3]},
				queryTemplate:               "%[1]s notin [%[2]v||%[2]v||%[2]v]",
				queryArgs:                   r[0],
				resourcesNotToExpectAfterOp: []common.Object{r[0]},
				expectedStatusCode:          http.StatusOK,
			},
		),
		Entry("returns 200",
			listOpEntry{
				resourcesToExpectBeforeOp:   []common.Object{r[0], r[1], r[2], r[3]},
				queryTemplate:               "%s notin [%v]",
				queryArgs:                   r[0],
				resourcesNotToExpectAfterOp: []common.Object{r[0]},
				expectedStatusCode:          http.StatusOK,
			},
		),
		Entry("returns 200",
			listOpEntry{
				resourcesToExpectBeforeOp:   []common.Object{r[0], r[1], r[2], r[3]},
				queryTemplate:               "%s gt %v",
				queryArgs:                   common.RemoveNonNumericArgs(r[0]),
				resourcesNotToExpectAfterOp: []common.Object{r[0]},
				expectedStatusCode:          http.StatusOK,
			},
		),
		Entry("returns 200 for greater than or equal queries",
			listOpEntry{
				resourcesToExpectBeforeOp: []common.Object{r[0], r[1], r[2], r[3]},
				queryTemplate:             "%s gte %v",
				queryArgs:                 common.RemoveNonNumericArgs(r[0]),
				resourcesToExpectAfterOp:  []common.Object{r[0], r[1], r[2], r[3]},
				expectedStatusCode:        http.StatusOK,
			},
		),
		Entry("returns 400 for greater than or equal queries when query args are non numeric",
			listOpEntry{
				resourcesToExpectBeforeOp: []common.Object{r[0], r[1], r[2], r[3]},
				queryTemplate:             "%s gte %v",
				queryArgs:                 common.RemoveNumericArgs(r[0]),
				resourcesToExpectAfterOp:  []common.Object{r[0], r[1], r[2], r[3]},
				expectedStatusCode:        http.StatusBadRequest,
			},
		),
		Entry("returns 200 for less than or equal queries",
			listOpEntry{
				resourcesToExpectBeforeOp: []common.Object{r[0], r[1], r[2], r[3]},
				queryTemplate:             "%s lte %v",
				queryArgs:                 common.RemoveNonNumericArgs(r[0]),
				resourcesToExpectAfterOp:  []common.Object{r[0], r[1], r[2], r[3]},
				expectedStatusCode:        http.StatusOK,
			},
		),
		Entry("returns 400 for less than or equal queries when query args are non numeric",
			listOpEntry{
				resourcesToExpectBeforeOp: []common.Object{r[0], r[1], r[2], r[3]},
				queryTemplate:             "%s lte %v",
				queryArgs:                 common.RemoveNumericArgs(r[0]),
				resourcesToExpectAfterOp:  []common.Object{r[0], r[1], r[2], r[3]},
				expectedStatusCode:        http.StatusBadRequest,
			},
		),
		Entry("returns 200",
			listOpEntry{
				resourcesToExpectBeforeOp:   []common.Object{r[0], r[1], r[2], r[3]},
				queryTemplate:               "%s lt %v",
				queryArgs:                   common.RemoveNonNumericArgs(r[0]),
				resourcesNotToExpectAfterOp: []common.Object{r[0]},
				expectedStatusCode:          http.StatusOK,
			},
		),
		Entry("returns 200 for field queries",
			listOpEntry{
				resourcesToExpectBeforeOp: []common.Object{r[0], rWithMandatoryFields},
				queryTemplate:             "%s eqornil %v",
				queryArgs:                 common.RemoveNotNullableFieldAndLabels(r[0], rWithMandatoryFields),
				resourcesToExpectAfterOp:  []common.Object{r[0], rWithMandatoryFields},
				expectedStatusCode:        http.StatusOK,
			},
		),
		Entry("returns 400 for label queries with operator eqornil",
			listOpEntry{
				queryTemplate: "%s eqornil %v",
				queryArgs: common.Object{
					"labels": map[string]interface{}{
						"labelKey1": []interface{}{
							"str",
						},
					}},
				expectedStatusCode: http.StatusBadRequest,
			},
		),
		Entry("returns 200 for JSON fields with stripped new lines",
			listOpEntry{
				resourcesToExpectBeforeOp: []common.Object{r[0]},
				queryTemplate:             "%s = %v",
				queryArgs:                 common.RemoveNonJSONArgs(r[0]),
				resourcesToExpectAfterOp:  []common.Object{r[0]},
				expectedStatusCode:        http.StatusOK,
			},
		),

		Entry("returns 400 when query operator is invalid",
			listOpEntry{
				queryTemplate:      "%s @@ %v",
				queryArgs:          r[0],
				expectedStatusCode: http.StatusBadRequest,
			},
		),
		Entry("returns 400 when query is duplicated",
			listOpEntry{
				queryTemplate:      "%[1]s = %[2]v|%[1]s = %[2]v",
				queryArgs:          r[0],
				expectedStatusCode: http.StatusBadRequest,
			},
		),
		Entry("returns 400 when operator is not properly separated with right space from operands",
			listOpEntry{
				queryTemplate:      "%s =%v",
				queryArgs:          r[0],
				expectedStatusCode: http.StatusBadRequest,
			},
		),
		Entry("returns 400 when operator is not properly separated with left space from operands",
			listOpEntry{
				queryTemplate:      "%s= %v",
				queryArgs:          r[0],
				expectedStatusCode: http.StatusBadRequest,
			},
		),

		Entry("returns 400 when field query left operands are unknown",
			listOpEntry{
				queryTemplate:      "%[1]s in [%[2]v||%[2]v]",
				queryArgs:          common.Object{"unknownkey": "unknownvalue"},
				expectedStatusCode: http.StatusBadRequest,
			},
		),
		Entry("returns 200 when label query left operands are unknown",
			listOpEntry{
				resourcesToExpectBeforeOp: []common.Object{r[0], r[1], r[2], r[3]},
				queryTemplate:             "%[1]s in [%[2]v||%[2]v]",
				queryArgs: common.Object{
					"labels": map[string]interface{}{
						"unknown": []interface{}{
							"unknown",
						},
					}},
				resourcesNotToExpectAfterOp: []common.Object{r[0], r[1], r[2], r[3]},
				expectedStatusCode:          http.StatusOK,
			},
		),
		Entry("returns 400 when single value operator is used with multiple right value arguments",
			listOpEntry{
				queryTemplate:      "%[1]s != [%[2]v||%[2]v||%[2]v]",
				queryArgs:          r[0],
				expectedStatusCode: http.StatusBadRequest,
			},
		),

		Entry("returns 400 when numeric operator is used with non-numeric operands",
			listOpEntry{
				queryTemplate:      "%s < %v",
				queryArgs:          common.RemoveNumericArgs(r[0]),
				expectedStatusCode: http.StatusBadRequest,
			},
		),
	}

	verifyListOpWithAuth := func(listOpEntry listOpEntry, query string, auth *httpexpect.Expect) {
		var expectedAfterOpIDs []string
		var unexpectedAfterOpIDs []string
		expectedAfterOpIDs = common.ExtractResourceIDs(listOpEntry.resourcesToExpectAfterOp)
		unexpectedAfterOpIDs = common.ExtractResourceIDs(listOpEntry.resourcesNotToExpectAfterOp)

		By(fmt.Sprintf("[TEST]: Verifying expected %s before operation after present", t.API))
		beforeOpArray := common.List(ctx.SMWithOAuth, t.API)

		for _, v := range beforeOpArray.Iter() {
			obj := v.Object().Raw()
			delete(obj, "created_at")
			delete(obj, "updated_at")
		}

		for _, entity := range listOpEntry.resourcesToExpectBeforeOp {
			delete(entity, "created_at")
			delete(entity, "updated_at")
			beforeOpArray.Contains(entity)
		}

		By("[TEST]: ======= Expectations Summary =======")

		By(fmt.Sprintf("[TEST]: Listing %s with %s", t.API, query))
		By(fmt.Sprintf("[TEST]: Currently present resources: %v", r))
		By(fmt.Sprintf("[TEST]: Expected %s ids after operations: %s", t.API, expectedAfterOpIDs))
		By(fmt.Sprintf("[TEST]: Unexpected %s ids after operations: %s", t.API, unexpectedAfterOpIDs))
		By(fmt.Sprintf("[TEST]: Expected status code %d", listOpEntry.expectedStatusCode))

		By("[TEST]: ====================================")

		By(fmt.Sprintf("[TEST]: Verifying expected status code %d is returned from list operation", listOpEntry.expectedStatusCode))

		if listOpEntry.expectedStatusCode != http.StatusOK {
			By(fmt.Sprintf("[TEST]: Verifying error and description fields are returned after list operation"))
			req := ctx.SMWithOAuth.GET(t.API)
			if query != "" {
				req = req.WithQueryString(query)
			}
			req.Expect().Status(listOpEntry.expectedStatusCode).JSON().Object().Keys().Contains("error", "description")
		} else {
			array := common.ListWithQuery(ctx.SMWithOAuth, t.API, query)
			for _, v := range array.Iter() {
				obj := v.Object().Raw()
				delete(obj, "created_at")
				delete(obj, "updated_at")
			}

			if listOpEntry.resourcesToExpectAfterOp != nil {
				By(fmt.Sprintf("[TEST]: Verifying expected %s are returned after list operation", t.API))
				for _, entity := range listOpEntry.resourcesToExpectAfterOp {
					delete(entity, "created_at")
					delete(entity, "updated_at")
					array.Contains(entity)
				}
			}

			if listOpEntry.resourcesNotToExpectAfterOp != nil {
				By(fmt.Sprintf("[TEST]: Verifying unexpected %s are NOT returned after list operation", t.API))

				for _, entity := range listOpEntry.resourcesNotToExpectAfterOp {
					delete(entity, "created_at")
					delete(entity, "updated_at")
					array.NotContains(entity)
				}
			}
		}
	}

	verifyListOp := func(listOpEntry listOpEntry, query string) {
		verifyListOpWithAuth(listOpEntry, query, ctx.SMWithOAuth)
	}

	return Describe("List", func() {
		Context("with basic auth", func() {
			It("returns 200", func() {
				ctx.SMWithBasic.GET(t.API).
					Expect().
					Status(http.StatusOK)
			})
		})

		Context("with bearer auth", func() {
			if !t.DisableTenantResources {
				Context("when authenticating with tenant scoped token", func() {
					var rForTenant common.Object

					BeforeEach(func() {
						rForTenant = t.ResourceBlueprint(ctx, ctx.SMWithOAuthForTenant)
					})

					It("returns only tenant specific resources", func() {
						verifyListOpWithAuth(listOpEntry{
							resourcesToExpectBeforeOp: []common.Object{r[0], r[1], rForTenant},
							resourcesToExpectAfterOp:  []common.Object{rForTenant},
							expectedStatusCode:        http.StatusOK,
						}, "", ctx.SMWithOAuthForTenant)
					})

					Context("when authenticating with global token", func() {
						It("it returns all resources", func() {
							verifyListOpWithAuth(listOpEntry{
								resourcesToExpectBeforeOp: []common.Object{r[0], r[1], rForTenant},
								resourcesToExpectAfterOp:  []common.Object{r[0], r[1], rForTenant},
								expectedStatusCode:        http.StatusOK,
							}, "", ctx.SMWithOAuth)
						})
					})
				})
			}

			Context("with no field query", func() {
				It("it returns all resources", func() {
					verifyListOpWithAuth(listOpEntry{
						resourcesToExpectBeforeOp: []common.Object{r[0], r[1]},
						resourcesToExpectAfterOp:  []common.Object{r[0], r[1]},
						expectedStatusCode:        http.StatusOK,
					}, "", ctx.SMWithOAuth)
				})
			})

			Context("Paging", func() {
				Context("with max items query", func() {
					It("returns smaller pages and Link header", func() {
						pageSize := 5
						resp := ctx.SMWithOAuth.GET(t.API).WithQuery("max_items", pageSize).Expect().Status(http.StatusOK)

						resp.Header("Link").Contains(fmt.Sprintf("<%s?max_items=%d&token=", t.API, pageSize)).Contains(`>; rel="next"`)
						resp.JSON().Path("$.has_more_items").Boolean().True()
						resp.JSON().Path("$.items[*]").Array().Length().Le(pageSize)
						resp.JSON().Path("$.token").NotNull()
					})
				})

				Context("with negative max items query", func() {
					It("returns 400", func() {
						ctx.SMWithOAuth.GET(t.API).WithQuery("max_items", -1).Expect().Status(http.StatusBadRequest)
					})
				})

				PContext("with zero max items query", func() {
					It("returns count of the items only", func() {
						resp := ctx.SMWithOAuth.GET(t.API).WithQuery("max_items", 0).Expect().Status(http.StatusOK).JSON()
						resp.Path("$.items[*]").Array().Empty()
						resp.Path("$.num_items").Number().Ge(0)
						resp.Object().NotContainsKey("token")
					})
				})

				Context("with invalid token", func() {
					executeWithInvalidToken := func(token string) {
						ctx.SMWithOAuth.GET(t.API).WithQuery("token", token).Expect().Status(http.StatusNotFound)
					}

					Context("no base64 encoded", func() {
						It("returns 404", func() {
							executeWithInvalidToken("invalid")
						})
					})

					Context("no timestamp", func() {
						It("returns 404", func() {
							id, _ := uuid.NewV4()
							token := base64.StdEncoding.EncodeToString([]byte("_" + id.String()))
							executeWithInvalidToken(token)
						})
					})

					Context("no id", func() {
						It("returns 404", func() {
							token := base64.StdEncoding.EncodeToString([]byte(time.Now().Format(time.RFC3339) + "_"))
							executeWithInvalidToken(token)
						})
					})

					Context("valid token without coresponding entity in database", func() {
						It("returns 404", func() {
							id, _ := uuid.NewV4()
							token := base64.StdEncoding.EncodeToString([]byte(time.Now().Format(time.RFC3339) + "_" + id.String()))
							executeWithInvalidToken(token)
						})
					})
				})
			})

			Context("with empty field query", func() {
				It("returns 200", func() {
					verifyListOp(listOpEntry{
						resourcesToExpectBeforeOp: []common.Object{r[0], r[1]},
						resourcesToExpectAfterOp:  []common.Object{r[0], r[1]},
						expectedStatusCode:        http.StatusOK,
					}, "fieldQuery=")
				})
			})

			Context("with empty label query", func() {
				It("returns 200", func() {
					verifyListOp(listOpEntry{
						resourcesToExpectBeforeOp: []common.Object{r[0], r[1]},
						resourcesToExpectAfterOp:  []common.Object{r[0], r[1]},
						expectedStatusCode:        http.StatusOK,
					}, "labelQuery=")
				})
			})

			Context("with empty label query and field query", func() {
				It("returns 200", func() {
					verifyListOp(listOpEntry{
						resourcesToExpectBeforeOp: []common.Object{r[0], r[1]},
						resourcesToExpectAfterOp:  []common.Object{r[0], r[1]},
						expectedStatusCode:        http.StatusOK,
					}, "labelQuery=&fieldQuery=")
				})
			})

			// expand all field and label query test enties into Its wrapped by descriptive Contexts
			for i := 0; i < len(entries); i++ {
				params := entries[i].Parameters[0].(listOpEntry)
				if len(params.queryTemplate) == 0 {
					panic("query templates missing")
				}
				var multiQueryValue string
				var queryValues []string

				fields := common.CopyObject(params.queryArgs)
				delete(fields, "labels")
				multiQueryValue, queryValues = expandFieldQuery(fields, params.queryTemplate)
				fquery := "fieldQuery" + "=" + multiQueryValue

				Context("with field query=", func() {
					for _, queryValue := range queryValues {
						query := "fieldQuery" + "=" + queryValue
						DescribeTable(fmt.Sprintf("%s", queryValue), func(test listOpEntry) {
							verifyListOp(test, query)
						}, entries[i])
					}

					if len(queryValues) > 1 {
						DescribeTable(fmt.Sprintf("%s", multiQueryValue), func(test listOpEntry) {
							verifyListOp(test, fquery)
						}, entries[i])
					}
				})

				labels := params.queryArgs["labels"]
				if labels != nil {
					multiQueryValue, queryValues = expandLabelQuery(labels.(map[string]interface{}), params.queryTemplate)
					lquery := "labelQuery" + "=" + multiQueryValue

					Context("with label query=", func() {
						for _, queryValue := range queryValues {
							query := "labelQuery" + "=" + queryValue
							DescribeTable(fmt.Sprintf("%s", queryValue), func(test listOpEntry) {
								verifyListOp(test, query)
							}, entries[i])
						}

						if len(queryValues) > 1 {
							DescribeTable(fmt.Sprintf("%s", multiQueryValue), func(test listOpEntry) {
								verifyListOp(test, lquery)
							}, entries[i])
						}
					})

					Context("with multiple field and label queries", func() {
						DescribeTable(fmt.Sprintf("%s", fquery+"&"+lquery), func(test listOpEntry) {
							verifyListOp(test, fquery+"&"+lquery)
						}, entries[i])
					})
				}
			}
		})
	})
}

func expandFieldQuery(fieldQueryArgs common.Object, queryTemplate string) (string, []string) {
	var expandedMultiQuery string
	var expandedQueries []string
	for k, v := range fieldQueryArgs {
		if v == nil {
			continue
		}

		if m, ok := v.(map[string]interface{}); ok {
			bytes, err := json.Marshal(m)
			if err != nil {
				panic(err)
			}
			v = string(bytes)
		}
		if a, ok := v.([]interface{}); ok {
			bytes, err := json.Marshal(a)
			if err != nil {
				panic(err)
			}
			v = string(bytes)

		}
		expandedQueries = append(expandedQueries, fmt.Sprintf(queryTemplate, k, v))
	}

	expandedMultiQuery = strings.Join(expandedQueries, "|")
	return expandedMultiQuery, expandedQueries
}

func expandLabelQuery(labelQueryArgs map[string]interface{}, queryTemplate string) (string, []string) {
	var expandedMultiQuery string
	var expandedQueries []string

	for key, values := range labelQueryArgs {
		for _, value := range values.([]interface{}) {
			expandedQueries = append(expandedQueries, fmt.Sprintf(queryTemplate, key, value))
		}
	}

	expandedMultiQuery = strings.Join(expandedQueries, "|")
	return expandedMultiQuery, expandedQueries
}
