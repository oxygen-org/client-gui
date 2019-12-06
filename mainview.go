package main

import (
	"fmt"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"github.com/oxygen-org/client-gui/utils"
	"github.com/oxygen-org/client-gui/utils/open"
)

var Tr = utils.Tr
var mainwin *ui.Window
var loginForm *ui.Window
var registerform *ui.Window

func makeBasicControlsPage() ui.Control {
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)
	vbox.Append(hbox, false)

	hbox.Append(ui.NewButton("Button"), false)
	hbox.Append(ui.NewCheckbox("Checkbox"), false)

	vbox.Append(ui.NewLabel("This is a label. Right now, labels can only span one line."), false)

	vbox.Append(ui.NewHorizontalSeparator(), false)

	group := ui.NewGroup("Entries")
	group.SetMargined(true)
	vbox.Append(group, true)

	group.SetChild(ui.NewNonWrappingMultilineEntry())

	entryForm := ui.NewForm()
	entryForm.SetPadded(true)
	group.SetChild(entryForm)

	entryForm.Append("Entry", ui.NewEntry(), false)
	entryForm.Append("Password Entry", ui.NewPasswordEntry(), false)
	entryForm.Append("Search Entry", ui.NewSearchEntry(), false)
	entryForm.Append("Multiline Entry", ui.NewMultilineEntry(), true)
	entryForm.Append("Multiline Entry No Wrap", ui.NewNonWrappingMultilineEntry(), true)

	return vbox
}

func makeNumbersPage() ui.Control {
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	group := ui.NewGroup("Numbers")
	group.SetMargined(true)
	hbox.Append(group, true)

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	group.SetChild(vbox)

	spinbox := ui.NewSpinbox(0, 100)
	slider := ui.NewSlider(0, 100)
	pbar := ui.NewProgressBar()
	spinbox.OnChanged(func(*ui.Spinbox) {
		slider.SetValue(spinbox.Value())
		pbar.SetValue(spinbox.Value())
	})
	slider.OnChanged(func(*ui.Slider) {
		spinbox.SetValue(slider.Value())
		pbar.SetValue(slider.Value())
	})
	vbox.Append(spinbox, false)
	vbox.Append(slider, false)
	vbox.Append(pbar, false)

	ip := ui.NewProgressBar()
	ip.SetValue(-1)
	vbox.Append(ip, false)

	group = ui.NewGroup("Lists")
	group.SetMargined(true)
	hbox.Append(group, true)

	vbox = ui.NewVerticalBox()
	vbox.SetPadded(true)
	group.SetChild(vbox)

	cbox := ui.NewCombobox()
	cbox.Append("Combobox Item 1")
	cbox.Append("Combobox Item 2")
	cbox.Append("Combobox Item 3")
	vbox.Append(cbox, false)

	ecbox := ui.NewEditableCombobox()
	ecbox.Append("Editable Item 1")
	ecbox.Append("Editable Item 2")
	ecbox.Append("Editable Item 3")
	vbox.Append(ecbox, false)

	rb := ui.NewRadioButtons()
	rb.Append("Radio Button 1")
	rb.Append("Radio Button 2")
	rb.Append("Radio Button 3")
	vbox.Append(rb, false)

	return hbox
}

func makeDataChoosersPage() ui.Control {
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	hbox.Append(vbox, false)

	vbox.Append(ui.NewDatePicker(), false)
	vbox.Append(ui.NewTimePicker(), false)
	vbox.Append(ui.NewDateTimePicker(), false)
	vbox.Append(ui.NewFontButton(), false)
	vbox.Append(ui.NewColorButton(), false)

	hbox.Append(ui.NewVerticalSeparator(), false)

	vbox = ui.NewVerticalBox()
	vbox.SetPadded(true)
	hbox.Append(vbox, true)

	grid := ui.NewGrid()
	grid.SetPadded(true)
	vbox.Append(grid, false)

	button := ui.NewButton("Open File")
	entry := ui.NewEntry()
	entry.SetReadOnly(true)
	button.OnClicked(func(*ui.Button) {
		filename := ui.OpenFile(mainwin)
		if filename == "" {
			filename = "(cancelled)"
		}
		entry.SetText(filename)
	})
	grid.Append(button,
		0, 0, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)
	grid.Append(entry,
		1, 0, 1, 1,
		true, ui.AlignFill, false, ui.AlignFill)

	button = ui.NewButton("Save File")
	entry2 := ui.NewEntry()
	entry2.SetReadOnly(true)
	button.OnClicked(func(*ui.Button) {
		filename := ui.SaveFile(mainwin)
		if filename == "" {
			filename = "(cancelled)"
		}
		entry2.SetText(filename)
	})
	grid.Append(button,
		0, 1, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)
	grid.Append(entry2,
		1, 1, 1, 1,
		true, ui.AlignFill, false, ui.AlignFill)

	msggrid := ui.NewGrid()
	msggrid.SetPadded(true)
	grid.Append(msggrid,
		0, 2, 2, 1,
		false, ui.AlignCenter, false, ui.AlignStart)

	button = ui.NewButton("Message Box")
	button.OnClicked(func(*ui.Button) {
		ui.MsgBox(mainwin,
			"This is a normal message box.",
			"More detailed information can be shown here.")
	})
	msggrid.Append(button,
		0, 0, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)
	button = ui.NewButton("Error Box")
	button.OnClicked(func(*ui.Button) {
		ui.MsgBoxError(mainwin,
			"This message box describes an error.",
			"More detailed information can be shown here.")
	})
	msggrid.Append(button,
		1, 0, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)

	return hbox
}

func mainview() {
	mainwin = ui.NewWindow(Tr("Oxygen计算平台客户端"), 840, 680, true)

	mainwin.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		mainwin.Destroy()
		return true
	})

	tab := ui.NewTab()
	
	// mainwin.SetChild(tab)
	mainwin.SetMargined(true)

	tab.Append(Tr("搜索添加"), searchPage())
	tab.SetMargined(0, true)

	tab.Append("Job管理", makeBasicControlsPage())
	tab.SetMargined(0, true)

	tab.Append("Job模板管理", makeNumbersPage())
	tab.SetMargined(1, true)

	tab.Append("数据管理", makeDataChoosersPage())
	tab.SetMargined(2, true)
	vbox := ui.NewVerticalBox()
	mainwin.SetChild(vbox)
	vbox.Append(tab,false)
	vbox.Append(ui.NewHorizontalSeparator(), false)
	vbox.Append(ui.NewHorizontalSeparator(), false)

	vbox.Append(ui.NewLabel("©️Oxygen-org / Tacey Wong"), false)
	mainwin.Show()
}

func login() {
	loginForm = ui.NewWindow("Oxygen: "+Tr("用户登录"), 420, 150, true)
	loginForm.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	loginForm.SetChild(vbox)

	group := ui.NewGroup(Tr("🐼：欢迎使用Oxygen计算平台"))
	group.SetMargined(true)
	vbox.Append(group, true)

	group.SetChild(ui.NewNonWrappingMultilineEntry())

	entryForm := ui.NewForm()
	entryForm.SetPadded(true)
	group.SetChild(entryForm)
	user := ui.NewEntry()
	entryForm.Append(Tr("用户名/邮箱"), user, false)
	password := ui.NewPasswordEntry()
	entryForm.Append(Tr("密码"), password, false)
	vbox.Append(ui.NewHorizontalSeparator(), false)

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)
	vbox.Append(hbox, false)
	registerBtn := ui.NewButton(Tr("注册"))
	hbox.Append(registerBtn, true)
	hbox.Append(ui.NewVerticalSeparator(), false)
	loginBtn := ui.NewButton(Tr("登录"))
	hbox.Append(loginBtn, true)

	registerBtn.OnClicked(func(*ui.Button) {
		loginForm.Destroy()
	})
	user.OnChanged(func(entry *ui.Entry) {
		fmt.Println(entry.Text())
	})
	loginBtn.OnClicked(func(*ui.Button) {
		if user.Text() != "" && password.Text() != "" {
			ui.MsgBoxError(loginForm,
				Tr("登录错误"),
				Tr("密码错误或用户不存在"))
		} else {
			ui.MsgBox(loginForm,
				Tr("恭喜！登录成功"),
				Tr("你已经通过用户密码验证"))
			loginForm.Destroy()
			mainview()
		}
	})
	registerBtn.OnClicked(func(*ui.Button) {
		open.Start("https://www.github.com/oxygen-org/client-gui")
	})
	vbox.Append(ui.NewHorizontalSeparator(), false)
	vbox.Append(ui.NewLabel("©️Oxygen-org / Tacey Wong"), false)
	vbox.Append(ui.NewHorizontalSeparator(), false)

	loginForm.Show()

}

func start() {
	ui.OnShouldQuit(func() bool {
		mainwin.Destroy()
		loginForm.Destroy()
		return true
	})
	ui.Main(login)
	fmt.Println("Start")
}
