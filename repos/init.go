package repos

func Init(path string) error {
	pIns := NewPostDaoInstance()
	if err := pIns.initPostIndexMap(path + "/post"); err != nil {
		return err
	}

	if err := initTopicIndexMap(path + "/topic"); err != nil {
		return err
	}
	return nil
}
