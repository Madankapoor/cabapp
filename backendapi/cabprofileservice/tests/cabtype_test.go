/**
* TestCabTypesCleanUp
* Deletes the CabTypes
*
* Must pass
 */
 func TestCabTypesCleanUp(t *testing.T) {
	var err error
	_, err = db.GetDB().DB().Exec("TRUNCATE TABLE cabtypes")
	if err != nil {
		t.Error(err)
	}
}
