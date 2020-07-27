/**
* TestCabDriversCleanUp
* Deletes the CabDrivers
*
* Must pass
*/
func TestCabDriversCleanUp(t *testing.T) {
	var err error
	_, err = db.GetDB().DB().Exec("TRUNCATE TABLE cabdrivers")
	if err != nil {
		t.Error(err)
	}
}
