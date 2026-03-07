package blogKeyPg

// ArticleCategoryTenantNo 文章分类租户编号
func ArticleCategoryTenantNo(tenantNo, code string) string {
	return "blog:artCate:" + tenantNo + ":" + code
}

// ArticleCategoryTenantNoAndNoByCode 文章分类租户编号和编号
func ArticleCategoryTenantNoAndNoByCode(tenantNo, code string) string {
	return "blog:artCateNoCode:" + tenantNo + ":" + code
}
