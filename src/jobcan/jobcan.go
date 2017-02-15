package jobcan

import (
	"errors"
	"time"

	"github.com/sclevine/agouti"
)

func login(input userInput, page *agouti.Page) error {
	if err := page.Navigate(loginURL()); err != nil {
		return err
	}
	if err := page.FindByID("email").Fill(input.Email); err != nil {
		return err
	}
	if err := page.FindByID("password").Fill(input.Password); err != nil {
		return err
	}
	if err := page.FindByButton("ログイン").Click(); err != nil {
		return err
	}

	time.Sleep(1500 * time.Millisecond)

	return nil
}

func editAttendance(input attendanceInput, page *agouti.Page) error {
	const url = "https://ssl.jobcan.jp/employee/attendance/edit"
	if err := page.Navigate(url); err != nil {
		return err
	}
	if err := page.All("select").At(0).Select(input.Year); err != nil {
		return err
	}
	if err := page.All("select").At(1).Select(input.Month); err != nil {
		return err
	}

	time.Sleep(1500 * time.Millisecond)

	return nil
}

func apply(input applyInput, page *agouti.Page) error {
	// TODO 出勤,退勤,休憩 -> 保存
	return errors.New("Not implement apply logic")
}

func Do(page *agouti.Page) (int, error) {
	if err := login(newLoginUser(), page); err != nil {
		return 1, err
	}
	if err := editAttendance(newEditAttendance(), page); err != nil {
		return 1, err
	}
	if err := apply(newApply(), page); err != nil {
		return 1, err
	}
	return 0, nil
}
