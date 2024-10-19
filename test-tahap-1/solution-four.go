package test_tahap_1

import (
	"fmt"
	"time"
)

func SolutionFour() {
	var jointHoliday int
	var joinDate string
	var leaveDatePlan string
	var totalLeavePlan int
	isEligibleLeave := false
	eligibleDaysAfterJoin := 180
	fmt.Print("Jumlah cuti bersama: ")
	fmt.Scan(&jointHoliday)

	fmt.Print("Tanggal join karyawan: ")
	fmt.Scan(&joinDate)

	fmt.Print("Tanggal rencana cuti: ")
	fmt.Scan(&leaveDatePlan)

	fmt.Print("Durasi cuti (hari): ")
	fmt.Scan(&totalLeavePlan)

	totalLeaveEligible, reason := getTotalLeave(jointHoliday, joinDate, leaveDatePlan, eligibleDaysAfterJoin, totalLeavePlan)
	if totalLeaveEligible >= totalLeavePlan {
		isEligibleLeave = true
	}
	fmt.Println(isEligibleLeave)
	if len(reason) > 0 {
		fmt.Println("Alasan: " + reason)
	}
}

func getTotalLeave(jointHoliday int, joinDateString string, leaveDatePlan string, eligibleDaysAfterJoin int, totalLeavePlan int) (totalLeave int, reason string) {
	joinDate := getDate(joinDateString)
	leaveDate := getDate(leaveDatePlan)
	leaveStartDate := getLeaveStartDate(joinDate, eligibleDaysAfterJoin)
	totalDayAfterJoin := getTotalJoinDate(joinDate, leaveDate)

	if totalDayAfterJoin <= eligibleDaysAfterJoin {
		reason = fmt.Sprintf("Karena belum %d hari sejak tanggal join karyawan", eligibleDaysAfterJoin)
		return
	}

	totalEligibleInYear := getTotalEligibleLeaveDayInYear(leaveStartDate)

	result := float64(float64(totalEligibleInYear)) / 365 * float64(jointHoliday)
	totalLeave = int(result)

	if totalLeavePlan > totalLeave {
		reason = fmt.Sprintf("Karena hanya boleh mengambil %d hari cuti", totalLeave)
		return
	}
	return
}

func getTotalEligibleLeaveDayInYear(leaveStartDate time.Time) int {
	endYear := getEndOfYear(leaveStartDate)
	duration := endYear.Sub(leaveStartDate)

	return int(duration.Hours() / 24)
}

func getTotalJoinDate(joinDate time.Time, leaveStartDate time.Time) int {
	duration := leaveStartDate.Sub(joinDate)
	return int(duration.Hours() / 24)
}

func getLeaveStartDate(joinDate time.Time, totalFutureDay int) time.Time {
	return joinDate.AddDate(0, 0, totalFutureDay)
}

func getDate(dateString string) time.Time {
	layout := "2006-01-02"

	parsedTime, _ := time.Parse(layout, dateString)

	return parsedTime
}

func getEndOfYear(date time.Time) time.Time {
	endOfYear := time.Date(date.Year(), time.December, 31, 23, 59, 59, 999999999, date.Location())
	return endOfYear
}
