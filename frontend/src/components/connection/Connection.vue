<script setup lang="ts">

import ConnList from "./ConnList.vue";
import TopicList from "./TopicList.vue";
import MsgDetail from "./MsgDetail.vue";
</script>

<template>
  <el-container class="conn-container">
    <ConnList/>
    <TopicList/>
    <MsgDetail/>
  </el-container>
</template>

<style scoped>
.conn-container {
  width: 100%;
}
</style>