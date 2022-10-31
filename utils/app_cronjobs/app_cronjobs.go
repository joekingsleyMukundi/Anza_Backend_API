package appcronjobs

import (
	"github.com/jasonlvhit/gocron"
)

// jobs include -- database backup, start new month ,delete notifcation

func StartServerCronJobs() {
	//checks if day date is first;

	// gocron.Every(1).Day().At("01:00").Do(updateNewMonthData)
	// gocron.Every(1).Day().At("01:00").Do(deleteReadUserNotification)
	gocron.Start()
}
