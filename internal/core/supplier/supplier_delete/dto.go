package supplier_delete

import (
	"github.com/google/uuid"
)

type InputSupplierDelete struct {
	SupplierID uuid.UUID `path:"id" json:"id" doc:"supplierID to delete" format:"uuid"`
}