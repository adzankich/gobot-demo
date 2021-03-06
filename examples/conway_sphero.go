package main

import (
	"fmt"
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot-sphero"
)

type conway struct {
	alive    bool
	age      int
	contacts int
	sphero   *gobotSphero.SpheroDriver
}

func main() {

	master := gobot.GobotMaster()

	var robots []*gobot.Robot
	spheros := []string{
		"/dev/rfcomm0",
		"/dev/rfcomm1",
		"/dev/rfcomm2",
		"/dev/rfcomm3",
		"/dev/rfcomm4",
	}

	for s := range spheros {
		spheroAdaptor := new(gobotSphero.SpheroAdaptor)
		spheroAdaptor.Name = "Sphero"
		spheroAdaptor.Port = spheros[s]

		sphero := gobotSphero.NewSphero(spheroAdaptor)
		sphero.Name = "Sphero" + spheros[s]

		work := func() {

			conway := new(conway)
			conway.sphero = sphero

			conway.birth()

			gobot.On(sphero.Events["Collision"], func(data interface{}) {
				conway.contact()
			})

			gobot.Every("3s", func() {
				if conway.alive == true {
					conway.movement()
				}
			})

			gobot.Every("10s", func() {
				if conway.alive == true {
					conway.birthday()
				}
			})
		}

		robots = append(robots, &gobot.Robot{Connections: []gobot.Connection{spheroAdaptor}, Devices: []gobot.Device{sphero}, Work: work})
	}

	master.Robots = robots

	master.Start()
}

func (c *conway) resetContacts() {
	c.contacts = 0
}

func (c *conway) contact() {
	c.contacts += 1
}

func (c *conway) rebirth() {
	fmt.Println("Welcome back", c.sphero.Name, "!")
	c.life()
}

func (c *conway) birth() {
	c.resetContacts()
	c.age = 0
	c.life()
	c.movement()
}

func (c *conway) life() {
	c.sphero.SetRGB(0, 255, 0)
	c.alive = true
}

func (c *conway) death() {
	fmt.Println(c.sphero.Name, "died :(")
	c.alive = false
	c.sphero.SetRGB(255, 0, 0)
	c.sphero.Stop()
}

func (c *conway) enoughContacts() bool {
	if c.contacts >= 2 && c.contacts < 7 {
		return true
	} else {
		return false
	}
}

func (c *conway) birthday() {
	c.age += 1

	fmt.Println("Happy birthday", c.sphero.Name, "you are", c.age, "and had", c.contacts, "contacts.")

	if c.enoughContacts() == true {
		if c.alive == false {
			c.rebirth()
		}
	} else {
		c.death()
	}

	c.resetContacts()
}

func (c *conway) movement() {
	if c.alive == true {
		c.sphero.Roll(70, uint16(gobot.Rand(360)))
	}
}
