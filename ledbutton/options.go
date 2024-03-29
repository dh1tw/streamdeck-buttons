package ledbutton

import "image/color"

// TextColor is a functional option which sets the text color.
func TextColor(c color.Color) func(*LedButton) {
	return func(btn *LedButton) {
		btn.textColor.C = c
	}
}

// LedColor is a functional option to set the color of the LED.
func LedColor(color LEDColor) func(*LedButton) {
	return func(btn *LedButton) {
		btn.ledColor = color
	}
}

// Text is a functional option for providing the initial text on the LED Button.
// Max 5 characters.
func Text(text string) func(*LedButton) {
	return func(btn *LedButton) {
		btn.text = text
	}
}

// State is a functional option for providing the initial state of the LED Button.
func State(on bool) func(*LedButton) {
	return func(btn *LedButton) {
		btn.state = on
	}
}
