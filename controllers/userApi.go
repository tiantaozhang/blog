package controllers

func (c *UserController) API_Profile() {

	type user struct {
		UserId string   `json:"userid"`
		Tag    string   `json:"tags"`
		Email  string   `json:"email"`
		Hobby  []string `json:"hobbys"`
	}

	c.Data["json"] = user{
		"jike",
		"shen",
		"2123@q.com",
		[]string{"chess", "code"},
	}

	c.ServeJson()
}
