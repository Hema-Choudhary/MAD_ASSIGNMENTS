package usercrudhandler

import "domain"

func transformobjListToResponse(resp []*domain.Restaurant) RestGetListRespDTO {
	responseObj := RestGetListRespDTO{}
	for _, obj := range resp {
		RestObj := RestGetRespDTO{
			DBID:        	obj.DBID,
			Name: 		obj.Name,
			Address:  	obj.Address,
			AddressLine2:   obj.AddressLine2,
			URL:		obj.URL,
			Outcode:	obj.Outcode,
			Postcode:	obj.Postcode,
			Rating:		obj.Rating,
			TypeOfFood:	obj.TypeOfFood,
		}
		responseObj.Rests = append(responseObj.Rests, RestObj)
	}
	responseObj.Count = len(responseObj.Rests)

	return responseObj
}
