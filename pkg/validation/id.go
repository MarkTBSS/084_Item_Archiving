package validation

import (
	_adminException "github.com/MarkTBSS/084_Item_Archiving/pkg/admin/exception"
	"github.com/labstack/echo/v4"
)

func AdminIDGetting(pctx echo.Context) (string, error) {
	if adminID, ok := pctx.Get("adminID").(string); !ok || adminID == "" {
		return "", &_adminException.AdminCreating{AdminID: "Unknown"}
	} else {
		return adminID, nil
	}
}
