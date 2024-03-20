package wiringPi

//#cgo LDFLAGS: -L/usr/local/lib -lwiringPi
/*
#include "wiringPi.h"

void init() {
	wiringPiSetup();
    pinMode(1,PWM_OUTPUT);  //only wiringPi pin 1 (BCM_GPIO 18) supports PWM
    pwmSetMode(PWM_MODE_MS); // Set to mode: mark:space
    pwmSetClock(192); //PWM clock: 19.2MHz, divisor 192:100KHz
    pwmSetRange(2000); //period 20ms
}

void lock() {
    pwmWrite(1,250);
}

void unlock() {
	pwmWrite(1,50);
}
*/
import "C"
import "time"

func init() {
	C.init()
	Lock()
}

func Unlock() {
	C.unlock()
}

func Lock() {
	C.lock()
}

func Pass() {
	Unlock()
	time.Sleep(time.Duration(3) * time.Second)
	Lock()
}
