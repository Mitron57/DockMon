package machines

import (
    "context"
    "dockMon/internal/domain/interfaces/services"
    "dockMon/internal/domain/models"
    "dockMon/pkg/http/response"
    "dockMon/pkg/marshalizers"
    "go.uber.org/zap"
    "log"
    "net/http"
    "time"
)

const placeholder = "Something went wrong"

func PutMachine(m services.Manager) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        logger, ok := r.Context().Value("logger").(*zap.Logger)
        if !ok {
            log.Println("No logger in context")
            response.InternalServerError(w, placeholder)
            return
        }
        defer logger.Sync()
        machine, err := marshalizers.UnmarshalJson[models.Machine](r.Body)
        if err != nil {
            logger.Info("BadRequest", zap.Error(err))
            response.ErrorResponse(w, http.StatusBadRequest, "Incorrect json format")
            return
        }
        ctx, cancel := context.WithTimeout(r.Context(), time.Second*5)
        defer cancel()
        err = m.Save(ctx, machine)
        if err != nil {
            logger.Error(err.Error(), zap.String("handler", "PutMachine"))
            response.InternalServerError(w, placeholder)
            return
        }
        response.WriteJSON(w, http.StatusOK, nil)
    }
}
