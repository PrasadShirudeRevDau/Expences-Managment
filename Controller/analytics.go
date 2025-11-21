package controller

import (
	models "ExpencesManagment/Models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var CurrentTime = time.Now()

// @Summary get total expense per category
// @Description get expenses total as per category
// @Tags analytics
// @Accept json
// @Produce json
// @Security BearerAuth
// @Router /analytics/category [get]
func TotalByCategory(c *gin.Context) {
	var expense []struct {
		Category    string
		TotalAmount float64
	}
	userid := c.GetUint("user_id")

	db.Model(&models.Expense{}).
		Select("user_id, category, SUM(amount) as total_amount").
		Where("user_id = ? ", userid).
		Group("category,user_id").
		Scan(&expense)

	c.JSON(http.StatusOK, expense)
}

// @Summary get total expense by date
// @Description get expenses total as per date range
// @Tags analytics
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param from query string true "from date"
// @Param to query string true "to date"
// @Router /analytics/range [get]
func TotalByDateRange(c *gin.Context) {
	user_ID := c.GetUint("user_id")
	from := c.Query("from")
	to := c.Query("to")

	var total float64

	FromDate, err := time.Parse("2006-01-02", from)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid from date formate. use YYYY-MM-DD"})
		return
	}

	ToDate, err := time.Parse("2006-01-02", to)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid to date formate. use YYYY-MM-DD"})
		return
	}

	db.Model(&models.Expense{}).
		Select("SUM(amount)").
		Where("user_id = ? and date between ? and ?", user_ID, FromDate, ToDate).
		Scan(&total)

	c.JSON(http.StatusOK, gin.H{"total": total})
}

// @Summary get summary of expense by monthly/daly/weekly
// @Description get expenses summary as per daily/weekly/monthly
// @Tags analytics
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param period query string true "Period" Enums(daily,weekly,monthly)
// @Router /analytics/summary [get]
func SummaryOfExpenses(c *gin.Context) {
	userID := c.GetUint("user_id")
	period := c.Query("period")

	year, week := CurrentTime.ISOWeek()

	jan4 := time.Date(year, 1, 4, 0, 0, 0, 0, time.UTC)
	weekday := int(jan4.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	mondayWeek1 := jan4.AddDate(0, 0, -weekday+1)
	startOfWeek := mondayWeek1.AddDate(0, 0, (week-2)*7)
	endOfWeek := startOfWeek.AddDate(0, 0, 6)

	startOfMonth := time.Date(year, CurrentTime.Month(), 1, 0, 0, 0, 0, time.UTC)
	endOfMonth := startOfMonth.AddDate(0, 1, -1)

	type SummaryResult struct {
		Period string  `json:"Period"`
		Amount float64 `json:"Amount"`
	}
	result := SummaryResult{}

	switch period {
	case "daily":
		db.Model(&models.Expense{}).
			Select("SUM(amount) as amount").
			Where("user_id = ? and date = ?", userID, CurrentTime.Format("2006-01-02")).
			Scan(&result)
		result.Period = "Today"
	case "weekly":
		db.Model(&models.Expense{}).
			Select("SUM(amount) as amount").
			Where("user_id = ? and date between ? and ?", userID, startOfWeek.Format("2006-01-02"), endOfWeek.Format("2006-01-02")).
			Scan(&result)
		result.Period = "weekly"
	case "monthly":
		db.Model(&models.Expense{}).
			Select("sum(amount) as amount").
			Where("user_id = ? and date between ? and ?", userID, startOfMonth.Format("2006-01-02"), endOfMonth.Format("2006-01-02")).
			Scan(&result)
		result.Period = "monthly"
	default:
		c.JSON(http.StatusNotFound, "enter valid value")
	}

	c.JSON(http.StatusOK, result)
}

// @Summary get expenses summary per month
// @Description get expenses total how many you have spend in a month
// @Tags analytics
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param months query string true "month" Enums(January,February,March,April,May,June,July,August,September,October,November,December)
// @Router /analytics/permonth [get]
func SummaryPerMonth(c *gin.Context) {
	userId := c.GetUint("user_id")
	Months := c.Query("months")

	var StartDate time.Time
	var EndDate time.Time

	type SummaryResult struct {
		Month  string  `json:"month"`
		Amount float64 `json:"Amount"`
	}
	result := SummaryResult{}

	switch Months {
	case "January":
		StartDate = time.Date(CurrentTime.Year(), 1, 1, 0, 0, 0, 0, time.UTC)
		EndDate = StartDate.AddDate(0, 1, -1)
	case "February":
		StartDate = time.Date(CurrentTime.Year(), 2, 1, 0, 0, 0, 0, time.UTC)
		EndDate = StartDate.AddDate(0, 1, -1)
	case "March":
		StartDate = time.Date(CurrentTime.Year(), 3, 1, 0, 0, 0, 0, time.UTC)
		EndDate = StartDate.AddDate(0, 1, -1)
	case "April":
		StartDate = time.Date(CurrentTime.Year(), 4, 1, 0, 0, 0, 0, time.UTC)
		EndDate = StartDate.AddDate(0, 1, -1)
	case "May":
		StartDate = time.Date(CurrentTime.Year(), 5, 1, 0, 0, 0, 0, time.UTC)
		EndDate = StartDate.AddDate(0, 1, -1)
	case "June":
		StartDate = time.Date(CurrentTime.Year(), 6, 1, 0, 0, 0, 0, time.UTC)
		EndDate = StartDate.AddDate(0, 1, -1)
	case "July":
		StartDate = time.Date(CurrentTime.Year(), 7, 1, 0, 0, 0, 0, time.UTC)
		EndDate = StartDate.AddDate(0, 1, -1)
	case "August":
		StartDate = time.Date(CurrentTime.Year(), 8, 1, 0, 0, 0, 0, time.UTC)
		EndDate = StartDate.AddDate(0, 1, -1)
	case "September":
		StartDate = time.Date(CurrentTime.Year(), 9, 1, 0, 0, 0, 0, time.UTC)
		EndDate = StartDate.AddDate(0, 1, -1)
	case "October":
		StartDate = time.Date(CurrentTime.Year(), 10, 1, 0, 0, 0, 0, time.UTC)
		EndDate = StartDate.AddDate(0, 1, -1)
	case "November":
		StartDate = time.Date(CurrentTime.Year(), 11, 1, 0, 0, 0, 0, time.UTC)
		EndDate = StartDate.AddDate(0, 1, -1)
	case "December":
		StartDate = time.Date(CurrentTime.Year(), 12, 1, 0, 0, 0, 0, time.UTC)
		EndDate = StartDate.AddDate(0, 1, -1)
	}

	db.Model(&models.Expense{}).
		Select("sum(amount) as amount").
		Where("user_id = ? and date between ? and ?", userId, StartDate.Format("2006-01-02"), EndDate.Format("2006-01-02")).
		Scan(&result)
	result.Month = Months

	c.JSON(http.StatusOK, result)
}