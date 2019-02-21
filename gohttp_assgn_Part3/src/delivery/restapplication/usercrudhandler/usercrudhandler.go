package usercrudhandler

import (	
	"fmt"
	"encoding/json"
	"io/ioutil"
	logger "log"
	"net/http"
	"domain"
	"github.com/gorilla/mux"
	"dbrepo/userrepo"
	customerrors "delivery/restapplication/packages/errors"
	"delivery/restapplication/packages/httphandlers"
	mthdroutr "delivery/restapplication/packages/mthdrouter"
	"delivery/restapplication/packages/resputl"
)

type RestCrudHandler struct {
	httphandlers.BaseHandler
	usersvc userrepo.Repository
}

func NewRestCrudHandler(usersvc userrepo.Repository) *RestCrudHandler {
	return &RestCrudHandler{usersvc: usersvc}
}

func (p *RestCrudHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := mthdroutr.RouteAPICall(p, r)
	response.RenderResponse(w)
}

//Get http method to get data
func (p *RestCrudHandler) Get(r *http.Request) resputl.SrvcRes {
	usVals := r.URL.Query()
	fmt.Println(len(usVals))
	i,ok := usVals["name"]
	if ok {
		resp, err := p.usersvc.FindByName(i[0])
		if err != nil {
			return resputl.ReponseCustomError(err)
		}
		responseObj := transformobjListToResponse(resp)
		return resputl.Response200OK(responseObj)
	

		
	}
	
	i,ok = usVals["typeOfFood"]
	if ok {
		resp, err := p.usersvc.FindByTypeOfFood(i[0])
		if err != nil {
			return resputl.ReponseCustomError(err)
		}
		responseObj := transformobjListToResponse(resp)
		return resputl.Response200OK(responseObj)
		
	}

	i,ok = usVals["search_term"]
	if ok {
		resp, err := p.usersvc.Search(i[0])
		if err != nil {
			return resputl.ReponseCustomError(err)
		}
		responseObj := transformobjListToResponse(resp)
		return resputl.Response200OK(responseObj)
		
	}else{
	pathParam := mux.Vars(r)
	usID := pathParam["id"]
	if usID == "" {
		//return resputl.Response200OK(generateSampleResponseObj())
		resp, err := p.usersvc.GetAll()
		if err != nil {
			return resputl.ReponseCustomError(err)
		}
		responseObj := transformobjListToResponse(resp)
		return resputl.Response200OK(responseObj)

	} else {
		obj, err := p.usersvc.Get(usID)	
		if err != nil {
			return resputl.ProcessError(customerrors.NotFoundError("User Object Not found"), "")
		}
		RestObj := RestGetRespDTO{
			DBID:        	obj.DBID,
			Name:		obj.Name,
			Address:  	obj.Address,
			AddressLine2: 	obj.AddressLine2,
			URL:		obj.URL,
			Postcode: 	obj.Postcode,
			Outcode:	obj.Outcode,
			Rating:		obj.Rating,
			TypeOfFood:	obj.TypeOfFood,
		}

		return resputl.Response200OK(RestObj)

		}
	}
	return resputl.Response200OK("")
}


//Post method creates new temporary schedule
func (p *RestCrudHandler) Post(r *http.Request) resputl.SrvcRes {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resputl.ReponseCustomError(err)
	}
	//e, err := ValidateUserCreateUpdateRequest(string(body))
	/*if e == false {
		return resputl.ProcessError(err, body)
		//return resputl.SimpleBadRequest("Invalid Input Data")

	}*/
	logger.Printf("Received POST request to Create schedule %s ", string(body))
	var obj *RestCreateReqDTO
	err = json.Unmarshal(body, &obj)
	if err != nil {
		resputl.SimpleBadRequest("Error unmarshalling Data")
	}
	restobj :=&domain.Restaurant{
			Name:		obj.Name,
			Address:  	obj.Address,
			AddressLine2: 	obj.AddressLine2,
			URL:		obj.URL,
			Postcode: 	obj.Postcode,
			Outcode:	obj.Outcode,
			Rating:		obj.Rating,
			TypeOfFood:	obj.TypeOfFood,
	} 
	
	id, err := p.usersvc.Store(restobj)
	if err != nil {
		logger.Fatalf("Error while creating in DB: %v", err)
		return resputl.ProcessError(customerrors.UnprocessableEntityError("Error in writing to DB"), "")
	}
	return resputl.Response200OK(&RestCreateRespDTO{ID: id})
}

//Put method modifies temporary schedule contents
func (p *RestCrudHandler) Put(r *http.Request) resputl.SrvcRes {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resputl.ReponseCustomError(err)
	}
	/*e, err := ValidateUserCreateUpdateRequest(string(body))
	if e == false {
		return resputl.ProcessError(err, body)
		//return resputl.SimpleBadRequest("Invalid Input Data")

	}*/
	logger.Printf("Received PUT request to Update schedule %s ", string(body))
		
	var obj *RestUpdateDTO
	err = json.Unmarshal(body, &obj)
	if err != nil {
		resputl.SimpleBadRequest("Error unmarshalling Data")
	}

	restObj :=&domain.Restaurant{
			DBID:		obj.DBID,
			Name:		obj.Name,
			Address:  	obj.Address,
			AddressLine2: 	obj.AddressLine2,
			URL:		obj.URL,
			Postcode: 	obj.Postcode,
			Outcode:	obj.Outcode,
			Rating:		obj.Rating,
			TypeOfFood:	obj.TypeOfFood,
	} 
	
	_,err = p.usersvc.Store(restObj)
	if err != nil {
		logger.Fatalf("Error while updating in DB: %v", err)
		return resputl.ProcessError(customerrors.UnprocessableEntityError("Error in updating to DB"), "")
	}
	return resputl.Response200OK(&RestUpdateRespDTO{ID: restObj.DBID})
	
}

//Delete method removes temporary schedule from db
func (p *RestCrudHandler) Delete(r *http.Request) resputl.SrvcRes {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resputl.ReponseCustomError(err)
	}
	logger.Printf("Received DELETE request to Delete schedule %s ", string(body))
		
	var requestdata *RestDeleteDTO
	err = json.Unmarshal(body, &requestdata)
	if err != nil {
		resputl.SimpleBadRequest("Error unmarshalling Data")
	}
	err =p.usersvc.Delete(requestdata.DBID)
	if err != nil {
		logger.Fatalf("Error while Deleting in DB: %v", err)
		return resputl.ProcessError(customerrors.UnprocessableEntityError("Error in Deleting to DB"), "")
	}
	return resputl.Response200OK(&RestDeleteRespDTO{ID: requestdata.DBID})

}
