package api

type AdminIdInput struct {
	UnixID string `uri:"admin_id" binding:"required"`
}
