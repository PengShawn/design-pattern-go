package main

import "fmt"

func main() {
	smsOtp := &Sms{}
	o := Otp{
		iOtp: smsOtp,
	}
	o.getAndSendOTP(4)

	fmt.Println("")
	emailOtp := &Email{}
	o = Otp{
		iOtp: emailOtp,
	}
	o.getAndSendOTP(4)
}
