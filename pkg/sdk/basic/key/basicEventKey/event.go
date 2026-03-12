package basicEventKey

// EventTenantNo 事件租户编号
func EventTenantNo(tenantNo, eventNo string) string {
	return "basicEvent:event:" + tenantNo + ":" + eventNo
}

// EventTenantNoByEventNo 事件租户编号,event = no
func EventTenantNoByEventNo(tenantNo, eventNo string) string {
	return "basicEvent:eventNo:" + tenantNo + ":" + eventNo
}

// EventFieldTenantNo 事件字段租户编号
func EventFieldTenantNo(tenantNo, eventNo, fieldNo string) string {
	return "basicEvent:field:" + tenantNo + ":" + eventNo + ":" + fieldNo
}

// EventFieldTenantNoByFieldNo 事件字段: filed = no,字段 和 编号相互转换
func EventFieldTenantNoByFieldNo(tenantNo, event, field string) string {
	return "basicEvent:fieldNo:" + tenantNo + ":" + event + ":" + field
}

// RulesByEventFieldTenantNo 规则字段租户编号
func RulesByEventFieldTenantNo(tenantNo, eventNo, fieldNo string) string {
	return "basicRules:eventField:" + tenantNo + ":" + eventNo + ":" + fieldNo
}

// EventTenantNoKeys 事件租户编号所有键
func EventTenantNoKeys(tenantNo string) string {
	return "basicEvent:eventKeys:" + tenantNo
}

// EventTenantNoByCode 事件租户编号通过编码
func EventTenantNoByCode(tenantNo, code string) string {
	return "basicEvent:eventNoCode:" + tenantNo + ":" + code
}

// RulesByFieldTenantNo 规则租户:字段编号
func RulesByFieldTenantNo(tenantNo, fieldNo string) string {
	return "basicRules:rules:" + tenantNo + ":" + fieldNo
}

// RulesTenantNoKeys 规则租户编号所有键
func RulesTenantNoKeys(tenantNo string) string {
	return "basicRules:rulesKeys:" + tenantNo
}
