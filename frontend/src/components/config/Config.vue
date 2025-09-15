<script setup lang="ts">
import {onMounted} from "vue";
import {
  initConfigService,
  subscriptionType,
  subType,
  switchDarkMode,
  theme,
  themeEnum,
  timeout
} from "./ConfigService.js";
import {UpdateConfig} from "../../../wailsjs/go/backend/DbService.js";
import {ElMessage} from "element-plus";

async function updateConfig() {
  await UpdateConfig(timeout.value)
  await UpdateConfig(subType.value)
  await UpdateConfig(theme.value)
  ElMessage.success("更新成功")
  switchDarkMode()
}

onMounted(() => initConfigService())
</script>

<template>
  <el-card class="config-card">
    <el-form
        label-position="left"
        label-width="140px"
        class="config-form">
      <!-- 超时时间 -->
      <el-form-item label="连接超时时间">
        <el-input
            v-model="timeout.Value"
            class="input-number"/>
        <span class="unit">秒</span>
      </el-form-item>

      <!-- 订阅类型 -->
      <el-form-item label="订阅类型">
        <el-select
            v-model="subType.Value"
            class="select">
          <el-option
              v-for="[key, value] in subscriptionType"
              :key="key"
              :label="key"
              :value="value"/>
        </el-select>
      </el-form-item>

      <!--夜间模式-->
      <el-form-item label="外观">
        <el-select
            v-model="theme.Value"
            class="select">
          <el-option
              v-for="[key, value] in themeEnum"
              :key="value"
              :label="value"
              :value="key"/>
        </el-select>
      </el-form-item>

      <!-- 保存按钮 -->
      <el-form-item class="form-actions">
        <el-button type="primary" @click="updateConfig">保存配置</el-button>
      </el-form-item>

    </el-form>
  </el-card>
</template>

<style scoped>
.config-card {
  width: 100%;
  margin: 10px auto;
  padding: 20px 25px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
}

.config-form .el-form-item {
  margin-bottom: 20px;
}

.input-number {
  width: 120px;
}

.unit {
  margin-left: 8px;
  color: #606266;
  font-size: 14px;
}

.select {
  width: 180px;
}

.form-actions {
  text-align: right;
  margin-top: 10px;
}
</style>
