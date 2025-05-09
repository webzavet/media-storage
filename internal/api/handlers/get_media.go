package handlers

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"github.com/hs-zavet/comtools/httpkit"
	"github.com/hs-zavet/comtools/httpkit/problems"
	"github.com/hs-zavet/media-storage/internal/api/responses"
	"github.com/hs-zavet/media-storage/internal/app/ape"
)

func (h *Handler) GetMedia(w http.ResponseWriter, r *http.Request) {
	mediaId, err := uuid.Parse(chi.URLParam(r, "media_id"))
	if err != nil {
		h.log.WithError(err).Warn("error parsing request")
		httpkit.RenderErr(w, problems.BadRequest(validation.Errors{
			"media_id": validation.NewError("media_id", "invalid UUID format"),
		})...)
		return
	}

	media, err := h.app.GetMedia(r.Context(), mediaId)
	if err != nil {
		switch {
		case errors.Is(err, ape.ErrMediaNotFound):
			httpkit.RenderErr(w, problems.NotFound("media not found"))
		default:
			httpkit.RenderErr(w, problems.InternalError())
		}

		h.log.WithError(err).Error("error getting media")
		return
	}

	httpkit.Render(w, responses.Media(media))
}
