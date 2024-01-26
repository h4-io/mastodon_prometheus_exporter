package mastodon

import (
	"errors"
	"strconv"
)

func GetAccountNumber(stats GenericStats) uint64 {
	return uint64(stats.Stats.UserCount)
}

func GetStatusesCount(stats GenericStats) uint64 {
	return uint64(stats.Stats.StatusCount)
}

func GetDomainCount(stats GenericStats) uint64 {
	return uint64(stats.Stats.DomainCount)
}

func GetWeeklyStatusesCount(stats []Activities) (uint64, error) {
	if len(stats) == 0 {
		return 0, errors.New("cannot retrive weekly statuses count, set it to 0")
	}
	weeklyStatusesCountConverted, err := strconv.ParseUint(stats[0].Statuses, 0, 64)
	if err != nil {
		return 0, err
	}
	return weeklyStatusesCountConverted, nil
}

func GetWeeklyLoginsCount(stats []Activities) (uint64, error) {
	if len(stats) == 0 {
		return 0, errors.New("cannot retrive weekly logins count, set it to 0")
	}
	weeklyLoginsCountConverted, err := strconv.ParseUint(stats[0].Logins, 0, 64)
	if err != nil {
		return 0, err
	}
	return weeklyLoginsCountConverted, nil
}

func GetWeeklyRegistrationsCount(stats []Activities) (uint64, error) {
	if len(stats) == 0 {
		return 0, errors.New("cannot retrive weekly registration count, set it to 0")
	}
	weeklyRegistrationsCountConverted, err := strconv.ParseUint(stats[0].Registrations, 0, 64)
	if err != nil {
		return 0, err
	}
	return weeklyRegistrationsCountConverted, nil
}

func GetMonthlyActiveUsersCount(stats NodeInfo) uint64 {
	return uint64(stats.Usage.Users.ActiveMonth)
}

func GetHalfYearlyActiveUsersCount(stats NodeInfo) uint64 {
	return uint64(stats.Usage.Users.ActiveHalfYear)
}
