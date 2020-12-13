package Core

func InitRequire()  error{
	for _,require :=range requireMap{
		err :=require.callback(featureMap[require.needType])
		if err != nil{
			return err
		}
	}
	return nil

}
