package machines

import (
    "context"
    "dockMon/internal/domain/dto"
    "dockMon/internal/domain/interfaces/services"
    "dockMon/pkg/http/response"
    "go.uber.org/zap"
    "log"
    "net/http"
    "time"
)

func GetMachines(m services.Manager) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        logger, ok := r.Context().Value("logger").(*zap.Logger)
        if !ok {
            log.Println("No logger in context")
            response.InternalServerError(w, "Something went wrong")
            return
        }
        defer logger.Sync()
        ctx, cancel := context.WithTimeout(r.Context(), time.Second*5)
        defer cancel()
        machines, err := m.Machines(ctx)
        if err != nil {
            logger.Error(err.Error(), zap.String("handler", "GetMachines"))
            response.InternalServerError(w, placeholder)
            return
        }
        response.WriteJSON(w, http.StatusOK, dto.Machines{List: machines})
    }
}
