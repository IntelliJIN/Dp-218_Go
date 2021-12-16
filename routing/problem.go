package routing

import (
	"Dp218Go/models"
	"Dp218Go/services"
	"Dp218Go/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

var problemService *services.ProblemService
var problemIDKey = "problemID"

var keyProblemRoutes = []Route{
	{
		Uri:     `/problems`,
		Method:  http.MethodGet,
		Handler: getAllProblems,
	},
	{
		Uri:     `/problem/{` + problemIDKey + `}`,
		Method:  http.MethodGet,
		Handler: getProblemInfo,
	},
	{
		Uri:     `/problem`,
		Method:  http.MethodPost,
		Handler: addProblem,
	},
}

func AddProblemHandler(router *mux.Router, service *services.ProblemService) {
	problemService = service
	for _, rt := range keyProblemRoutes {
		router.Path(rt.Uri).HandlerFunc(rt.Handler).Methods(rt.Method)
		router.Path(APIprefix + rt.Uri).HandlerFunc(rt.Handler).Methods(rt.Method)
	}
}

func getAllProblems(w http.ResponseWriter, r *http.Request) {

	var problems *models.ProblemList
	var err error
	var userID, typeID, dateFrom, dateTo interface{}
	format := GetFormatFromRequest(r)

	userID, err = GetParameterFromRequest(r, "UserID", utils.ConvertStringToInt())
	if err == nil {
		problems, err = problemService.GetProblemsByUserID(userID.(int))
		if err != nil {
			ServerErrorRender(format, w)
			return
		}
	}

	if err != nil {
		typeID, err = GetParameterFromRequest(r, "TypeID", utils.ConvertStringToInt())
		if err == nil {
			problems, err = problemService.GetProblemsByTypeID(typeID.(int))
			if err != nil {
				ServerErrorRender(format, w)
				return
			}
		}
	}

	if err != nil {
		dateFrom, err = GetParameterFromRequest(r, "DateFrom", utils.ConvertStringToTime())
		if err == nil {
			dateTo, err = GetParameterFromRequest(r, "DateTo", utils.ConvertStringToTime())
			if err == nil {
				problems, err = problemService.GetProblemsByTimePeriod(dateFrom.(time.Time), dateTo.(time.Time))
				if err != nil {
					ServerErrorRender(format, w)
					return
				}
			}
		}
	}

	if err != nil {
		problems, err = problemService.GetProblemsByBeingSolved(false)
		if err != nil {
			ServerErrorRender(format, w)
			return
		}
	}

	EncodeAnswer(format, w, problems, HTMLPath+"problems.html")
}

func getProblemInfo(w http.ResponseWriter, r *http.Request) {
	format := GetFormatFromRequest(r)

	problemID, err := strconv.Atoi(mux.Vars(r)[problemIDKey])
	if err != nil {
		EncodeError(format, w, ErrorRendererDefault(err))
		return
	}

	problem, err := problemService.GetProblemByID(problemID)
	if err != nil {
		EncodeError(FormatHTML, w, ErrorRendererDefault(err))
		return
	}

	EncodeAnswer(format, w, problem, HTMLPath+"problem.html")
}

func addProblem(w http.ResponseWriter, r *http.Request) {
	format := GetFormatFromRequest(r)

	problemData := models.Problem{}
	DecodeRequest(format, w, r, &problemData, DecodeProblemAddRequest)
	err := problemService.AddNewProblem(&problemData)
	if err != nil {
		ServerErrorRender(format, w)
		return
	}

	EncodeAnswer(format, w, problemData, HTMLPath+"problem-add.html")
}

func DecodeProblemAddRequest(r *http.Request, data interface{}) error {

	var err error

	problemData := data.(*models.Problem)

	description, _ := GetParameterFromRequest(r, "Description", utils.ConvertStringToString())
	scooterID, _ := GetParameterFromRequest(r, "ScooterID", utils.ConvertStringToInt())
	userID, err := GetParameterFromRequest(r, "UserID", utils.ConvertStringToInt())
	if err != nil {
		return err
	}
	typeID, err := GetParameterFromRequest(r, "TypeID", utils.ConvertStringToInt())
	if err != nil {
		return err
	}

	problemData.Description = description.(string)
	problemData.IsSolved = false
	err = problemService.AddProblemComplexFields(problemData, typeID.(int), scooterID.(int), userID.(int))
	if err != nil {
		return err
	}

	return problemService.AddNewProblem(problemData)
}