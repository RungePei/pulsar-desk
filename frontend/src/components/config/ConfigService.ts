import {ref, watch} from "vue";
import {QueryConfigs} from "../../../wailsjs/go/backend/DbService.js";
import {backend} from "../../../wailsjs/go/models.js";
import {useDark, usePreferredDark, useToggle} from "@vueuse/core";
import Config = backend.Config;

enum ConfigId {
    TIMEOUT = 1,
    SUBSCRIPTION_TYPE,
    THEME
}

enum SubscriptionType {
    Exclusive = "0",
    Shared = "1",
    Failover = "2",
    KeyShared = "3"
}

enum Theme {
    Auto = "跟随系统",
    Light = "明亮",
    Dark = "黑暗",
}

export const subscriptionType = Object.entries(SubscriptionType);

export const themeEnum = Object.entries(Theme);

export const configs = ref<Config[]>([])
export const timeout = ref<Config>(new Config())
export const subType = ref<Config>(new Config())
export const theme = ref<Config>(new Config())

// 获取用户系统是否偏好暗色模式
const prefersDark = usePreferredDark()
// 用 useDark 控制 html.dark 类
const isDark = useDark()
const darkHandle = watch(prefersDark, (newVal) => {
    isDark.value = newVal
    useToggle(isDark)
    console.log('系统主题变化为：', newVal ? '暗黑' : '明亮')
})

export function switchDarkMode() {
    switch (theme.value.Value) {
        case "Auto": {
            darkHandle.resume()
            isDark.value = usePreferredDark().value
            break
        }
        case "Light": {
            darkHandle.pause()
            isDark.value = false
            useToggle(isDark)
            break
        }
        case "Dark": {
            darkHandle.pause()
            isDark.value = true
            useToggle(isDark)
            break
        }
    }
}

let init = false;

export async function initConfigService() {
    if (init)
        return
    configs.value = await QueryConfigs()
    timeout.value = configs.value.find((c) => c.Id === ConfigId.TIMEOUT) || new Config()
    subType.value = configs.value.find((c) => c.Id === ConfigId.SUBSCRIPTION_TYPE) || new Config()
    theme.value = <Config>configs.value.find((c) => c.Id === ConfigId.THEME)
    init = true
}