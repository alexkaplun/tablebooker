package controller

func (c *Controller) InitDB() error {
	if err := c.storage.InitEmpty(); err != nil {
		return err
	}

	if err := c.storage.InitDummyData(); err != nil {
		return err
	}
	return nil
}
