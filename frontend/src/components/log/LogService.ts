import {reactive, ref} from 'vue'
import {backend} from "../../../wailsjs/go/models.js";
import {EventsOn} from "../../../wailsjs/runtime/runtime.js";
import LogMsg = backend.LogMsg;

// 全局日志数组
export const logs = reactive<LogMsg[]>([])

export const filterLevels = ref<string[]>(['INFO', 'DEBUG', 'WARN', 'ERROR'])

// 初始化日志服务，保证只执行一次
let initialized = false
export function initLogService() {
    if (initialized) return
    initialized = true

    EventsOn('logMsg', (log: LogMsg) => {
        logs.push(log)
    })
}
