package mediator

func ExampleDialog_HandleEvent() {
	usernameInput := Input("username input")
	passwordInput := Input("password input")
	repeatPwdInput := Input("repeat password input")

	selection := Selection("登录")
	d := &Dialog{
		Selection:           &selection,
		UsernameInput:       &usernameInput,
		PasswordInput:       &passwordInput,
		RepeatPasswordInput: &repeatPwdInput,
	}

	d.HandleEvent(&selection)
	regSelection := Selection("注册")
	d.Selection = &regSelection
	d.HandleEvent(&regSelection)

	// Output:
	// select login
	// not show: repeat password input
	// select register
	// show: repeat password input
}