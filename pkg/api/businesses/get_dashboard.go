package businesses

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/mewben/arrstate/internal/enums"
	"github.com/mewben/arrstate/pkg/errors"
)

// DashboardParams -
type DashboardParams struct {
	SalesFrom string `query:"salesFrom"`
	SalesTo   string `query:"salesTo"`
}

// TotalStruct -
type TotalStruct struct {
	Total int64 `bson:"total"`
}

// GetDashboard -
func (h *Handler) GetDashboard(params *DashboardParams) (fiber.Map, error) {
	var err error
	var tz = "Asia/Manila"
	var now = time.Now()
	var salesFrom time.Time
	var salesTo time.Time

	loc, err := time.LoadLocation(tz)
	if err != nil {
		return nil, err
	}

	if params.SalesFrom == "" {
		salesFrom = now.Add(-30 * 24 * time.Hour).In(loc)
	} else {
		salesFrom, err = time.ParseInLocation(enums.DateTimeForm, params.SalesFrom, loc)
		if err != nil {
			return nil, err
		}
	}

	if params.SalesTo == "" {
		salesTo = now.In(loc)
	} else {
		salesTo, err = time.ParseInLocation(enums.DateTimeForm, params.SalesTo, loc)
		if err != nil {
			return nil, err
		}
	}

	salesData, err := h.getSalesData(salesFrom, salesTo, tz)
	if err != nil {
		return nil, err
	}
	countsData, err := h.getEntityCounts()
	if err != nil {
		return nil, err
	}

	return fiber.Map{
		"sales":  salesData,
		"counts": countsData,
	}, err

}

// getSalesData
/*
var today = moment().startOf('day');
  // "2018-12-05T00:00:00.00
  var tomorrow = moment(today).endOf('day');
  // ("2018-12-05T23:59:59.999
*/
func (h *Handler) getSalesData(from, to time.Time, tz string) (interface{}, error) {
	matchStage := bson.D{
		{
			Key: "$match",
			Value: bson.D{
				{
					Key:   "businessID",
					Value: h.Business.ID,
				},
				{
					Key:   "status",
					Value: enums.StatusPaid,
				},
				{
					Key: "paidAt",
					Value: bson.M{
						"$gte": from,
						"$lte": to,
					},
				},
			},
		},
	}

	groupStage := bson.D{
		{
			Key: "$group",
			Value: bson.D{
				{
					Key: "_id",
					Value: bson.D{
						{
							Key: "month",
							Value: bson.M{
								"$month": bson.M{
									"date":     "$paidAt",
									"timezone": tz,
								},
							},
						},
						{
							Key: "day",
							Value: bson.M{
								"$dayOfMonth": bson.M{
									"date":     "$paidAt",
									"timezone": tz,
								},
							},
						},
						{
							Key: "year",
							Value: bson.M{
								"$year": bson.M{
									"date":     "$paidAt",
									"timezone": tz,
								},
							},
						},
					},
				},
				{
					Key: "amount",
					Value: bson.M{
						"$sum": "$total",
					},
				},
				{
					Key: "count",
					Value: bson.M{
						"$sum": 1,
					},
				},
			},
		},
	}

	sortStage := bson.D{
		{
			Key: "$sort",
			Value: bson.M{
				"_id": 1,
			},
		},
	}

	cursor, err := h.DB.Aggregate(h.Ctx, enums.CollInvoices, mongo.Pipeline{matchStage, groupStage, sortStage})
	if err != nil {
		return nil, errors.NewHTTPError(errors.ErrDefault, err)
	}

	var response []bson.M
	if err = cursor.All(h.Ctx, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (h *Handler) getEntityCounts() (interface{}, error) {
	filter := bson.M{
		"businessID": h.Business.ID,
	}

	totalProjects, err := h.DB.Count(h.Ctx, enums.CollProjects, filter)
	if err != nil {
		return nil, err
	}

	totalProperties, err := h.DB.Count(h.Ctx, enums.CollProperties, filter)
	if err != nil {
		return nil, err
	}

	totalPeople, err := h.DB.Count(h.Ctx, enums.CollPeople, filter)
	if err != nil {
		return nil, err
	}

	matchSalesStage := bson.D{
		{
			Key: "$match",
			Value: bson.D{
				{
					Key:   "businessID",
					Value: h.Business.ID,
				},
				{
					Key:   "status",
					Value: enums.StatusPaid,
				},
			},
		},
	}

	groupSalesStage := bson.D{
		{
			Key: "$group",
			Value: bson.M{
				"_id": nil,
				"total": bson.M{
					"$sum": "$total",
				},
			},
		},
	}

	cursor, err := h.DB.Aggregate(h.Ctx, enums.CollInvoices, mongo.Pipeline{matchSalesStage, groupSalesStage})
	if err != nil {
		return nil, err
	}

	var responseSales []TotalStruct
	if err = cursor.All(h.Ctx, &responseSales); err != nil {
		return nil, err
	}

	var totalSales int64
	if len(responseSales) > 0 {
		totalSales = responseSales[0].Total
	}

	return fiber.Map{
		"projects":   totalProjects,
		"properties": totalProperties,
		"people":     totalPeople,
		"sales":      totalSales,
	}, err

}
