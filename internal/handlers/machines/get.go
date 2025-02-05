package machines

import (
    "dockMon/internal/domain/dto"
    "dockMon/internal/domain/interfaces/services"
    "dockMon/pkg/http/response"
    "go.uber.org/zap"
    "log"
    "net/http"
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
        machines, err := m.Machines(r.Context())
        if err != nil {
            logger.Error(err.Error(), zap.String("handler", "GetMachines"))
            response.InternalServerError(w, placeholder)
            return
        }
        response.WriteJSON(w, http.StatusOK, dto.Machines{List: machines})
    }
}
