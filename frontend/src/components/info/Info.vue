<script setup lang="ts">
const version = 'v0.0.1'
const github = 'https://github.com/RungePei/pulsar-desk'
</script>

<template>
  <el-card>
    <el-form>
      <el-form-item>
        <h3>版本：{{ version }}</h3>
      </el-form-item>
      <el-form-item>
        <el-link type="primary" href="">检查更新</el-link>
      </el-form-item>
      <el-form-item>
        该项目为练手使用，功能简陋
      </el-form-item>
      <el-form-item>
        <el-link type="primary" :href="github" target="_blank">GitHub</el-link>
      </el-form-item>
    </el-form>
  </el-card>
</template>

<style scoped>
.el-card{
  width: 100%;
}

</style>