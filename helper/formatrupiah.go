package helper

import (
	"github.com/yudapc/go-rupiah"
)

func FormatRupiah(total uint64) string {
	amountFloat := float64(total)
	formatRupiah := rupiah.FormatRupiah(amountFloat)
	return formatRupiah
}
