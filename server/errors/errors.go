package errorsApp

import "log"

func ErrorWhatStopTheApp(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func ErrorWhatShowAWarning(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func ErrorWhatReturnAMessage(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
