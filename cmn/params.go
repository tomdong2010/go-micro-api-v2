/**
 * Created by Goland.
 * User: yan.wang5
 * Date: 2019/11/03
 */
package cmn

// 关于环境的定义
const (
	ENV_PROD = "prod" // 生产环境
	ENV_PRE  = "pre"  // 预发环境
	ENV_TEST = "test" // 测试环境
	ENV_DEV  = "dev"  // 开发环境
)

// 服务名称
const (
	APP_CONF_PREFIX = "/micro/config/demo" // 配置文件前缀
	APP_NAME_PREFIX = "demo."              // 命名空间前缀

	APP_API_USER = "api.user" // 用户网关app名称
	APP_SRV_USER = "srv.user" // 用户服务app名称
)

// 账号类型
const (
	ACCOUNT_TYPE_PHONE = iota + 1
	ACCOUNT_TYPE_EMAIL
	ACCOUNT_TYPE_USERNAME
	ACCOUNT_TYPE_GOOGLE
	ACCOUNT_TYPE_FACEBOOK
	ACCOUNT_TYPE_TWITTER
)

// redis相关的key
const (
	REDIS_EMAIL_BACKTIMES = "demo:email:backtimes_"
)

// 参数校验使用的正则表达式
const (
	REGEXP_USERNAME = `^[a-zA-Z][a-zA-Z0-9]{2,23}$`
	REGEXP_CC       = `^[1-9][0-9]{0,3}$`
	REGEXP_PHONE    = `^[0-9]{6,20}$`
)
