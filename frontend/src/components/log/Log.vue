<script setup lang="ts">
import {computed, nextTick, onMounted, ref, watch} from 'vue'
import {filterLevels, initLogService, logs} from "./LogService.js";

const scrollRef = ref<any>(null);

const filteredLogs = computed(() => logs.filter((log) => filterLevels.value.includes(log.Level)))

// 滚动到底部
async function scrollToBottom() {
  const wrap = scrollRef.value?.wrapRef;
  if (wrap) {
    wrap.scrollTo({
      top: wrap.scrollHeight,   // 滚动到内容的最大高度
      behavior: "smooth"        // 平滑滚动，可改成 "auto"
    });
  }
}

watch(filteredLogs, () => {
  nextTick(() => {
    scrollToBottom();
  })
})

function clearLogs() {
  logs.splice(0, logs.length)
}

onMounted(() => {
  initLogService()
  scrollToBottom()
})

function formatTime(time: number) {
  return new Date(time).toLocaleString();
}
</script>

<template>
  <div class="log-page">
    <!-- 操作栏 -->
    <div class="operate">
      <el-checkbox-group v-model="filterLevels">
        <el-checkbox label="INFO">INFO</el-checkbox>
        <el-checkbox label="DEBUG">DEBUG</el-checkbox>
        <el-checkbox label="WARN">WARN</el-checkbox>
        <el-checkbox label="ERROR">ERROR</el-checkbox>
      </el-checkbox-group>
      <el-button type="danger" @click="clearLogs">清空日志</el-button>
    </div>

    <!-- 日志展示 -->
    <el-card class="log-card">
      <el-scrollbar class="log-container" ref="scrollRef">
        <div v-for="(log, index) in filteredLogs"
             :key="index"
             :class="['log-line', log.Level.toLowerCase()]">
          {{ index + 1 }} : {{ formatTime(log.Time) }} [{{ log.Level }}] {{ log.Msg }}
        </div>
      </el-scrollbar>
    </el-card>
  </div>
</template>

<style scoped>
.log-page {
  display: flex;
  flex-direction: column;
  height: 100%;
  width: 100%;
  padding: 10px;
}

.operate {
  display: flex;
  justify-content: space-between;
  border-color: #409eff;
}
.log-card{
  height: 100%;
}

.log-container {
  flex: 1;
  font-family: monospace;
  padding: 10px;
  border-radius: 4px;
  height: 100%;
}

.log-line {
  white-space: pre-wrap;
  margin-bottom: 2px;
}

/* 日志级别颜色 */
.log-line.info {
  color: blue;
}

.log-line.debug {
  color: #6bc26b;
}

.log-line.warn {
  color: yellow;
}

.log-line.error {
  color: red;
}
</style>
