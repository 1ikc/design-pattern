package template

import "fmt"

type IBaseSMS interface {
	Valid(content string) error
	DoSend(phone int, content string) error
	Send(phone int, content string) error
}

type BaseSMS struct {}

func (bs *BaseSMS) Valid(content string) error {
	return nil
}

func (bs *BaseSMS) DoSend(phone int, content string) error {
	return nil
}

func (bs *BaseSMS) Send(phone int, content string) error {
	if err := bs.Valid(content); err != nil {
		return err
	}

	return bs.DoSend(phone, content)
}

type MobileSMS struct {
	IBaseSMS
}

func (m *MobileSMS) Valid(content string) error {
	fmt.Println("valid from MobileSMS")
	return nil
}

func (m *MobileSMS) DoSend(phone int, content string) error {
	fmt.Println("send from MobileSMS")
	return nil
}

func (m *MobileSMS) Send(phone int, content string) error {
	if err := m.Valid(content); err != nil {
		return err
	}

	return m.DoSend(phone, content)
}

type TelecomSMS struct {
	IBaseSMS
}

func (t *TelecomSMS) Valid(content string) error {
	fmt.Println("valid from TelecomSMS")
	return nil
}

func (t *TelecomSMS) DoSend(phone int, content string) error {
	fmt.Println("send from TelecomSMS")
	return nil
}

func (t *TelecomSMS) Send(phone int, content string) error {
	if err := t.Valid(content); err != nil {
		return err
	}

	return t.DoSend(phone, content)
}