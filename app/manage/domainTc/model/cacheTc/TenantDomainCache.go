package cacheTc

import (
	"github.com/foxiswho/blog-go/app/manage/domainTc/model/dtoTenantDomain"
	"github.com/foxiswho/blog-go/pkg/configPg"
	"github.com/foxiswho/blog-go/pkg/tools/mapPg"
	"strings"
)

// TenantDomainCache
// @Description: 域名缓存
type TenantDomainCache struct {
	DomainData *mapPg.SafeMap[string, dtoTenantDomain.Dto]
	Domain     *mapPg.SafeMap[string, string]
}

// IsLocalHostExist
//
//	@Description: 本地域名
//	@receiver c
//	@param host
//	@return bool
func (c *TenantDomainCache) IsLocalHostExist(host string) bool {
	if strings.Contains(host, "localhost") {
		return true
	}
	if strings.Contains(host, "192.168.") {
		return true
	}
	return false
}

// IsServerHostExist
//
//	@Description: 指定域名
//	@receiver c
//	@param host
//	@param server
//	@return bool
func (c *TenantDomainCache) IsServerHostExist(host string, server configPg.Server) bool {
	str := strings.Replace(server.Domain, "http://", "", -1)
	str = strings.Replace(str, "https://", "", -1)
	// 查找最后一个冒号的位置
	lastColonIndex := strings.LastIndex(str, ":")
	// 如果找到了冒号，并且冒号后面还有字符，则截取冒号前面的部分
	if lastColonIndex != -1 && lastColonIndex < len(str)-1 {
		str = str[:lastColonIndex]
	}
	if strings.Contains(host, str) {
		return true
	}
	return false
}
