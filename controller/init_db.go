package controller

func (c *Controller) InitDB() error {
	// init empty database
	if err := c.storage.InitEmpty(); err != nil {
		return err
	}

	// fill in the database with dummy data
	if err := c.storage.InitDummyData(); err != nil {
		return err
	}
	return nil
}
