package collector

import (
	"mastodon_prometheus_exporter/mastodon"

	"log"

	"github.com/VictoriaMetrics/metrics"
)

type MastodonCollector struct {
	Instance                   string
	AccountNumber              *metrics.Counter
	StatusesCount              *metrics.Counter
	DomainCount                *metrics.Counter
	WeeklyStatusesCount        *metrics.Counter
	WeeklyLoginsCount          *metrics.Counter
	WeeklyRegistrationsCount   *metrics.Counter
	MonthlyActiveUsersCount    *metrics.Counter
	HalfYearlyActiveUsersCount *metrics.Counter
}

func InitMastodonStats(instance string) MastodonCollector {
	var collector MastodonCollector

	collector.Instance = instance

	log.Println("initialize mastodon counters... ")
	collector.AccountNumber = metrics.NewCounter("mastodon_accounts_number")
	collector.StatusesCount = metrics.NewCounter("mastodon_statuses_count")
	collector.DomainCount = metrics.NewCounter("mastodon_domain_count")
	collector.WeeklyStatusesCount = metrics.NewCounter("mastodon_week_statuses_count")
	collector.WeeklyLoginsCount = metrics.NewCounter("mastodon_week_logins_count")
	collector.WeeklyRegistrationsCount = metrics.NewCounter("mastodon_week_registrations_count")
	collector.MonthlyActiveUsersCount = metrics.NewCounter("mastodon_month_active_count")
	collector.HalfYearlyActiveUsersCount = metrics.NewCounter("mastodon_half_year_active_count")

	return collector
}

func GetMastodonStats(collector MastodonCollector) {
	log.Println("retrive generic stats")
	genericStats, err := mastodon.GetGenericStats(collector.Instance)
	if err != nil {
		log.Println(err.Error() + " values for mastodon_accounts_number, mastodon_statuses_count, mastodon_domain_count will not be updated")
	} else {
		collector.AccountNumber.Set(mastodon.GetAccountNumber(genericStats))
		collector.StatusesCount.Set(mastodon.GetStatusesCount(genericStats))
		collector.DomainCount.Set(mastodon.GetDomainCount(genericStats))
	}

	log.Println("retrive activities")
	activities, err := mastodon.GetActivities(collector.Instance)
	if err != nil {
		log.Println(err.Error() + " values for mastodon_week_statuses_count, mastodon_week_logins_count, mastodon_week_registrations_count will not be updated")
	} else {
		weeklyStatuses, err := mastodon.GetWeeklyStatusesCount(activities)
		if err != nil {
			log.Println(err.Error())
		} else {
			collector.WeeklyStatusesCount.Set(weeklyStatuses)
		}

		weeklyLogins, err := mastodon.GetWeeklyLoginsCount(activities)
		if err != nil {
			log.Println(err.Error())
		} else {
			collector.WeeklyLoginsCount.Set(weeklyLogins)
		}

		weeklyRegistrations, err := mastodon.GetWeeklyRegistrationsCount(activities)
		if err != nil {
			log.Println(err.Error())
		} else {
			collector.WeeklyRegistrationsCount.Set(weeklyRegistrations)
		}
	}

	log.Println("retrive nodeInfo")
	nodeInfo, err := mastodon.GetNodeInfo(collector.Instance)
	if err != nil {
		log.Println(err.Error() + " values for mastodon_month_active_count, mastodon_half_year_active_count will not be updated")
	} else {
		collector.MonthlyActiveUsersCount.Set(mastodon.GetMonthlyActiveUsersCount(nodeInfo))
		collector.HalfYearlyActiveUsersCount.Set(mastodon.GetHalfYearlyActiveUsersCount(nodeInfo))
	}
}
