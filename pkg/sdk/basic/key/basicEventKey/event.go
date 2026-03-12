package basicEventKey

// EventTenantNo 事件租户编号
func EventTenantNo(tenantNo, eventNo string) string {
	return "basicEvent:event:" + tenantNo + ":" + eventNo
}

// EventFieldTenantNo 事件字段租户编号
func EventFieldTenantNo(tenantNo, eventNo, fieldNo string) string {
	return "basicEvent:field:" + tenantNo + ":" + eventNo + ":" + fieldNo
}

// RulesByEventFieldTenantNo 规则字段租户编号
func RulesByEventFieldTenantNo(tenantNo, eventNo, fieldNo string) string {
	return "basicRules:eventField:" + tenantNo + ":" + eventNo + ":" + fieldNo
}
