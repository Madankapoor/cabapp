


/**
* TestCreatePassenger
* Test passenger create
*
* Must return response code 200
 */
func TestCreatePassenger(t *testing.T) {
	testRouter := SetupRouter()
	type InputPassenger struct {
		Email    string `json:"email"`
		Name     string `json:"name"`
		MobileNo string `json:"mobileno"`
	}
	var passenger InputPassenger
	passenger.Name = "MadanKapoor"
	passenger.Email = fmt.Sprintf("madankapoor%d@gmail.com", rand.Intn(100))
	passenger.MobileNo = fmt.Sprintf("877%d2946", rand.Intn(100))

	data, _ := json.Marshal(passenger)
	req, err := http.NewRequest("POST", "/api/v1.0/passenger/", bytes.NewBufferString(string(data)))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()

	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

/**
* TestDeletePassenger
* Test deleting an passenger
*
* Must return response code 204
 */
func TestDeletePassenger(t *testing.T) {
	testRouter := SetupRouter()
	url := fmt.Sprintf("/api/v1.0/passenger/%d", 1)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusNoContent, resp.Code)
}

/**
* TestCabsCleanUp
* Deletes the cabs
*
* Must pass
 */
func TestCabsCleanUp(t *testing.T) {
	var err error
	_, err = db.GetDB().DB().Exec("TRUNCATE TABLE cabs")
	if err != nil {
		t.Error(err)
	}
}
