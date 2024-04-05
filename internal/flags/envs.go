package flags

import (
	"fmt"
	"strings"

	"github.com/soulteary/webhook/internal/fn"
	"github.com/soulteary/webhook/internal/hook"
)

func ParseEnvs() AppFlags {
	var flags AppFlags
	flags.Host = fn.GetEnvStr(ENV_KEY_HOST, DEFAULT_HOST)
	flags.Port = fn.GetEnvInt(ENV_KEY_PORT, DEFAULT_PORT)
	flags.Verbose = fn.GetEnvBool(ENV_KEY_VERBOSE, DEFAULT_ENABLE_VERBOSE)
	flags.LogPath = fn.GetEnvStr(ENV_KEY_LOG_PATH, DEFAULT_LOG_PATH)
	flags.Debug = fn.GetEnvBool(ENV_KEY_DEBUG, DEFAULT_ENABLE_DEBUG)
	flags.NoPanic = fn.GetEnvBool(ENV_KEY_NO_PANIC, DEFAULT_ENABLE_NO_PANIC)
	flags.HotReload = fn.GetEnvBool(ENV_KEY_HOT_RELOAD, DEFAULT_ENABLE_HOT_RELOAD)
	flags.HooksURLPrefix = fn.GetEnvStr(ENV_KEY_HOOKS_URLPREFIX, DEFAULT_URL_PREFIX)
	flags.AsTemplate = fn.GetEnvBool(ENV_KEY_TEMPLATE, DEFAULT_ENABLE_PARSE_TEMPLATE)
	flags.UseXRequestID = fn.GetEnvBool(ENV_KEY_X_REQUEST_ID, DEFAULT_ENABLE_X_REQUEST_ID)
	flags.XRequestIDLimit = fn.GetEnvInt(ENV_KEY_X_REQUEST_ID, DEFAULT_X_REQUEST_ID_LIMIT)
	flags.MaxMultipartMem = int64(fn.GetEnvInt(ENV_KEY_MAX_MPART_MEM, DEFAULT_MAX_MPART_MEM))
	flags.SetGID = fn.GetEnvInt(ENV_KEY_GID, DEFAULT_GID)
	flags.SetUID = fn.GetEnvInt(ENV_KEY_UID, DEFAULT_UID)
	flags.HttpMethods = fn.GetEnvStr(ENV_KEY_HTTP_METHODS, DEFAULT_HTTP_METHODS)
	flags.PidPath = fn.GetEnvStr(ENV_KEY_PID_FILE, DEFAULT_PID_FILE)

	// init i18n, set lang and i18n dir
	flags.Lang = fn.GetEnvStr(ENV_KEY_LANG, DEFAULT_LANG)
	flags.I18nDir = fn.GetEnvStr(ENV_KEY_I18N, DEFAULT_I18N_DIR)

	hooks := strings.Split(fn.GetEnvStr(ENV_KEY_HOOKS, ""), ",")
	var hooksFiles hook.HooksFiles
	for _, hook := range hooks {
		err := hooksFiles.Set(hook)
		if err != nil {
			fmt.Println("Error parsing hooks from environment variable: ", err)
		}
	}
	if len(hooksFiles) > 0 {
		flags.HooksFiles = hooksFiles
	}
	return flags
}
