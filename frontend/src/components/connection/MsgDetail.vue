<script setup lang="ts">
import {computed, nextTick, ref, watch, watchEffect} from "vue"
import {backend} from "../../../wailsjs/go/models.js";
import {idStore} from "../../config/Store.js";
import {QueryByTopic, SendMsg} from "../../../wailsjs/go/backend/DbService.js";
import {ElMessage} from "element-plus";
import {EventsOn} from "../../../wailsjs/runtime/runtime.js";
import Msg = backend.Msg;

const messages = ref<Msg[]>([])
const topicId = computed(() => idStore().topicId);
const input = ref("")

const scrollRef = ref<any>(null);

watch(messages, () => {
  nextTick(scrollToBottom);
}, {"deep": true});

watchEffect(async () => messages.value = (await QueryByTopic(topicId.value) || []))

// 发送消息
async function sendMessage() {
  if (!input.value.trim()) return
  const id = topicId.value
  const msg = new Msg({
    Content: input.value,
    TopicId: id,
    Type: 'to',
    Time: Date.now()
  })
  try {
    messages.value.push(msg)
    await SendMsg(msg)
  } catch (err: any) {
    console.log(err)
    ElMessage.error(err)
  }
  input.value = ""
}

// 格式化时间
const formatTime = (timestamp: number) => {
  const date = new Date(timestamp)
  return date.toLocaleString()
}

// 监听后端发送的事件
EventsOn("mqMsg", (msg: Msg) => {
  messages.value.push(msg)
})

// 滚动到底部
function scrollToBottom() {
  const wrap = scrollRef.value?.wrapRef;
  if (wrap) {
    wrap.scrollTo({
      top: wrap.scrollHeight,   // 滚动到内容的最大高度
      behavior: "smooth"        // 平滑滚动，可改成 "auto"
    });
  }
}
</script>

<template>
  <div class="chat-container">
    <!-- 聊天消息区 -->
    <div class="chat-card">
      <el-scrollbar ref="scrollRef">
        <div v-for="msg in messages"
             class="chat-message"
             :class="msg.Type === 'to' ? 'from-me' : 'from-other'">
          <div class="bubble">{{ msg.Content }}</div>
          <div class="time">{{ formatTime(msg.Time) }}</div>
        </div>
      </el-scrollbar>
    </div>

    <!-- 输入框区 -->
    <div class="chat-input">
      <el-input
          v-model="input"
          placeholder="输入消息..."
          @keyup.enter="sendMessage"
          clearable/>
      <el-button type="primary" @click="sendMessage">发送</el-button>
    </div>
  </div>
</template>

<style scoped>
/* 整体容器 */
.chat-container {
  display: flex;
  flex-direction: column;
  height: 100vh; /* 全屏高度，可根据实际需求调整 */
  width: 100%;
  gap: 8px;
  padding: 10px;
  box-sizing: border-box;
  border-color: #999999;
}

/* 聊天消息区 */
.chat-card {
  flex: 1; /* 占满剩余空间 */
  overflow: hidden;
  border-radius: 12px;
  padding: 10px;
  display: flex;
  flex-direction: column;
}

/* 消息列表 */
.chat-message {
  display: flex;
  flex-direction: column; /* 垂直排列气泡和时间 */
  align-items: flex-start; /* 默认靠左 */
  margin: 6px 0;
}

.chat-message.from-me {
  align-items: flex-end; /* 我发送的消息靠右 */
}

/* 消息气泡 */
.bubble {
  max-width: 60%;
  padding: 10px 14px;
  border-radius: 12px;
  color: black;
  background-color: #fff;
  word-break: break-word;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
}

/* 我发送的消息气泡 */
.from-me .bubble {
  background-color: #409eff;
  color: white;
}

/* 时间显示 */
.time {
  font-size: 12px;
  color: #999;
  margin-top: 4px; /* 消息与时间间距 */
}

/* 输入框区域 */
.chat-input {
  display: flex;
  gap: 8px;
  padding-top: 5px;
}

/* 输入框和按钮高度一致 */
.el-input, .el-button {
  height: 48px;
  box-sizing: border-box;
}

/* 输入框自动占满剩余空间 */
.el-input {
  flex: 1;
}

</style>
