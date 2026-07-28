package blogKeyPg

// ArticleCategoryTenantNo 文章分类租户编号
func ArticleCategoryTenantNo(tenantNo, code string) string {
	return "blog:artCate:" + tenantNo + ":" + code
}

// ArticleCategoryTenantNoKeys 文章分类租户编号,所有键
func ArticleCategoryTenantNoKeys(tenantNo string) string {
	return "blog:artCateKeys:" + tenantNo
}

// ArticleCategoryTenantNoAndNoByCode 文章分类租户编号和编号
func ArticleCategoryTenantNoAndNoByCode(tenantNo, code string) string {
	return "blog:artCateNoCode:" + tenantNo + ":" + code
}
