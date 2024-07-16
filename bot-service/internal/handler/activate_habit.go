package handler

import (
	"encoding/json"
	"main/internal/usecase/activatehabit"
	"net/http"
)

// CreateUserHandler godoc
//
//	@Summary	Activate user habit
//	@Tags		handler
//	@Accept		json
//	@Produce	json
//	@Param		habitActivation	body	HabitActivationModel	true	"HabitActivationRequest"
//	@Router		/user/activateHabit [put]
//	@Success	200
//	@Failure	400	{string}	Bad		Request
//	@Failure	500	{string}	Server	Error
func (h *HttpHandler) ActivateHabit() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var am HabitActivationModel

		err := json.NewDecoder(r.Body).Decode(&am)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		cmd, err := activatehabit.NewCommand(am.UserId, am.HabitId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = h.ActivateHabitHandler.Handle(r.Context(), cmd)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})
}
