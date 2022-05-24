package logic

type User struct {
	UID       uint64   `json:"uid"`
	PartnerID uint64   `json:"partner_id"`
	CompanyID uint64   `json:"company_id"`
	Name      string   `json:"name"`
	Follow    []uint64 `json:"follow"`
	Type      string   `json:"type"`
}
