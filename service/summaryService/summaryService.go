package summaryService

import (
	"WakaTImeGo/model/entity"
	v1 "WakaTImeGo/model/entity/wakatime/v1"
	"WakaTImeGo/service/duration"
	"WakaTImeGo/utils"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

/**
 * @Author: Kim
 * @Description: Summary service that generate summary after heartbeat saved.
 * @File:  summaryService
 * @Date: 12/13/2022 4:44 PM
 */

// GetSummaryByTimeRange summary service
func GetSummaryByTimeRange(userId string, start, end time.Time) v1.Summary {
	durations := duration.GetDurationByTime(userId, start, end)
	summary := v1.Summary{}
	summary.Data = v1.SummaryData{}
	// prepare summary data
	var categories = make(map[string]v1.SummaryItem)
	var editors = make(map[string]v1.SummaryItem)
	var languages = make(map[string]v1.SummaryItem)
	var machines = make(map[string]v1.SummaryItem)
	var operatingSystems = make(map[string]v1.SummaryItem)
	var projects = make(map[string]v1.SummaryItem)
	//var dependencies = make(map[string]v1.SummaryItem)
	summary.Data.GrandTotal = v1.GrandTotal{}
	summary.Data.Range = v1.Range{}
	for _, duration := range durations {
		log.Info(duration)
		durationTotal := GetDurationTotal(duration)
		summary.Data.GrandTotal = mergeTwoGrandTotal(summary.Data.GrandTotal, durationTotal)
		summaryItems := getSummaryItem(duration, durationTotal)
		//get if category exist
		if _, ok := categories[duration.Category]; ok {
			//category exist
			categories[duration.Category] = mergeTwoSummaryItem(categories[duration.Category], summaryItems[0])
		} else {
			//category not exist
			categories[duration.Category] = summaryItems[0]
		}

		//get if editor exist
		if _, ok := editors[duration.Editor]; ok {
			//editor exist
			editors[duration.Editor] = mergeTwoSummaryItem(editors[duration.Editor], summaryItems[1])
		} else {
			//editor not exist
			editors[duration.Editor] = summaryItems[1]
		}

		//get if language exist
		if _, ok := languages[duration.Language]; ok {
			//language exist
			languages[duration.Language] = mergeTwoSummaryItem(languages[duration.Language], summaryItems[2])
		} else {
			//language not exist
			languages[duration.Language] = summaryItems[2]
		}

		//get if machine exist
		if _, ok := machines[duration.Machine]; ok {
			//machine exist
			machines[duration.Machine] = mergeTwoSummaryItem(machines[duration.Machine], summaryItems[3])
		} else {
			//machine not exist
			machines[duration.Machine] = summaryItems[3]
		}

		//get if operatingSystem exist
		if _, ok := operatingSystems[duration.OperatingSystem]; ok {
			//operatingSystem exist
			operatingSystems[duration.OperatingSystem] = mergeTwoSummaryItem(operatingSystems[duration.OperatingSystem], summaryItems[4])
		} else {
			//operatingSystem not exist
			operatingSystems[duration.OperatingSystem] = summaryItems[4]
		}

		//get if project exist
		if _, ok := projects[duration.Project]; ok {
			//project exist
			projects[duration.Project] = mergeTwoSummaryItem(projects[duration.Project], summaryItems[5])
		} else {
			//project not exist
			projects[duration.Project] = summaryItems[5]
		}
	}
	// convert map to array
	categoriesArray := make([]v1.SummaryItem, 0)
	for _, value := range categories {
		categoriesArray = append(categoriesArray, value)
	}
	editorsArray := make([]v1.SummaryItem, 0)
	for _, value := range editors {
		editorsArray = append(editorsArray, value)
	}
	languagesArray := make([]v1.SummaryItem, 0)
	for _, value := range languages {
		languagesArray = append(languagesArray, value)
	}
	machinesArray := make([]v1.SummaryItem, 0)
	for _, value := range machines {
		machinesArray = append(machinesArray, value)
	}
	operatingSystemsArray := make([]v1.SummaryItem, 0)
	for _, value := range operatingSystems {
		operatingSystemsArray = append(operatingSystemsArray, value)
	}
	projectsArray := make([]v1.SummaryItem, 0)
	for _, value := range projects {
		projectsArray = append(projectsArray, value)
	}
	summary.Data.Categories = categoriesArray
	summary.Data.Editors = editorsArray
	summary.Data.Languages = languagesArray
	summary.Data.Machines = machinesArray
	summary.Data.OperatingSystems = operatingSystemsArray
	summary.Data.Projects = projectsArray
	summary.Data.Range.Date = time.Now().Format("2006-01-02")
	summary.Data.Range.Text = summary.Data.GrandTotal.Text
	summary.Data.Range.End = end
	summary.Data.Range.Start = start
	summary.Data.Range.Timezone = ""
	return summary
}

// getSummaryItem get summary item by duration, will return 6 items that summary need
// the 6 items are: categories, editors, languages, machines, operatingSystems, projects
// param duration: duration entity
// param durationTotal: duration's total
func getSummaryItem(duration entity.Duration, total v1.GrandTotal) []v1.SummaryItem {
	category := v1.SummaryItem{
		Name:         duration.Category,
		Hours:        total.Hours,
		Minutes:      total.Minutes,
		Seconds:      total.Seconds,
		Digital:      total.Digital,
		Decimal:      total.Decimal,
		Text:         total.Text,
		TotalSeconds: total.TotalSeconds,
	}
	editor := v1.SummaryItem{
		Name:         duration.Editor,
		Hours:        total.Hours,
		Minutes:      total.Minutes,
		Seconds:      total.Seconds,
		Digital:      total.Digital,
		Decimal:      total.Decimal,
		Text:         total.Text,
		TotalSeconds: total.TotalSeconds,
	}
	language := v1.SummaryItem{
		Name:         duration.Language,
		Hours:        total.Hours,
		Minutes:      total.Minutes,
		Seconds:      total.Seconds,
		Digital:      total.Digital,
		Decimal:      total.Decimal,
		Text:         total.Text,
		TotalSeconds: total.TotalSeconds,
	}
	machine := v1.SummaryItem{
		Name:         duration.Machine,
		Hours:        total.Hours,
		Minutes:      total.Minutes,
		Seconds:      total.Seconds,
		Digital:      total.Digital,
		Decimal:      total.Decimal,
		Text:         total.Text,
		TotalSeconds: total.TotalSeconds,
	}
	operatingSystem := v1.SummaryItem{
		Name:         duration.OperatingSystem,
		Hours:        total.Hours,
		Minutes:      total.Minutes,
		Seconds:      total.Seconds,
		Digital:      total.Digital,
		Decimal:      total.Decimal,
		Text:         total.Text,
		TotalSeconds: total.TotalSeconds,
	}
	project := v1.SummaryItem{
		Name:         duration.Project,
		Hours:        total.Hours,
		Minutes:      total.Minutes,
		Seconds:      total.Seconds,
		Digital:      total.Digital,
		Decimal:      total.Decimal,
		Text:         total.Text,
		TotalSeconds: total.TotalSeconds,
	}
	summaryItems := []v1.SummaryItem{category, editor, language, machine, operatingSystem, project}
	return summaryItems
}

func mergeTwoSummaryItem(item1 v1.SummaryItem, item2 v1.SummaryItem) v1.SummaryItem {
	// merge two summary item
	// if item1 and item2 are the same, then merge them
	// if not, return item1
	if item1.Name == item2.Name {
		item1.Hours += item2.Hours
		item1.Minutes += item2.Minutes
		if item1.Minutes >= 60 {
			item1.Hours += item1.Minutes / 60
			item1.Minutes = item1.Minutes % 60
		}
		item1.Seconds += item2.Seconds
		if item1.Seconds >= 60 {
			item1.Minutes += item1.Seconds / 60
			item1.Seconds = item1.Seconds % 60
		}
		item1.TotalSeconds += item2.TotalSeconds
		item1.Text = strconv.Itoa(item1.Hours) + " hrs " + strconv.Itoa(item1.Minutes) + " mins " + strconv.Itoa(item1.Seconds) + " secs"
		item1.Decimal = strconv.FormatFloat(float64(item1.Hours)+float64(item1.Minutes)/60+float64(item1.Seconds)/3600, 'f', 2, 64)
		item1.Digital = strconv.Itoa(item1.Hours) + ":" + strconv.Itoa(item1.Minutes) + ":" + strconv.Itoa(item1.Seconds)
		return item1
	} else {
		return item1
	}
}

func GetDurationTotal(duration entity.Duration) v1.GrandTotal {
	hours, minutes, seconds := utils.GetInfoByDuration(duration.Duration)
	hoursStr := strconv.Itoa(hours)
	minutesStr := strconv.Itoa(minutes)
	secondsStr := strconv.Itoa(seconds)
	hoursInDecimal := float64(hours) + float64(minutes)/60 + float64(seconds)/3600
	digital := strconv.Itoa(hours) + ":" + strconv.Itoa(minutes) + ":" + strconv.Itoa(seconds)
	text := hoursStr + " hrs " + minutesStr + " mins " + secondsStr + " secs"
	decimal := strconv.FormatFloat(hoursInDecimal, 'f', 2, 64)
	totalSeconds := float64(duration.Time.Nanosecond()) / 1000.0
	return v1.GrandTotal{
		Hours:        hours,
		Minutes:      minutes,
		Seconds:      seconds,
		Digital:      digital,
		Text:         text,
		Decimal:      decimal,
		TotalSeconds: totalSeconds,
	}
}

func mergeTwoGrandTotal(total1 v1.GrandTotal, total2 v1.GrandTotal) v1.GrandTotal {
	total1.Hours += total2.Hours
	total1.Minutes += total2.Minutes
	if total1.Minutes >= 60 {
		total1.Hours += 1
		total1.Minutes -= 60
	}
	total1.Seconds += total2.Seconds
	if total1.Seconds >= 60 {
		total1.Minutes += 1
		total1.Seconds -= 60
	}
	total1.TotalSeconds += total2.TotalSeconds
	total1.Text = strconv.Itoa(total1.Hours) + " hrs " + strconv.Itoa(total1.Minutes) + " mins " + strconv.Itoa(total1.Seconds) + " secs"
	total1.Decimal = strconv.FormatFloat(float64(total1.Hours)+float64(total1.Minutes)/60+float64(total1.Seconds)/3600, 'f', 2, 64)
	total1.Digital = strconv.Itoa(total1.Hours) + ":" + strconv.Itoa(total1.Minutes) + ":" + strconv.Itoa(total1.Seconds)
	return total1
}
