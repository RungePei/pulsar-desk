<script setup lang="ts">

import {computed, onMounted, ref} from "vue";
import CollapseBtn from "./connection/CollapseBtn.vue";

const isCollapse = ref(false)

const asideWidth = computed(() => isCollapse.value ? "60px" : "120px");

onMounted(() => {
  console.log("aside collapsed", isCollapse.value);
})

</script>

<template>
  <el-aside :width="asideWidth">
    <h4>{{ isCollapse ? "Pulsar" : "Pulsar客户端" }}</h4>
    <el-menu :collapse="isCollapse"
             default-active="/"
             router>
      <el-menu-item index="/">
        <el-icon>
          <Connection/>
        </el-icon>
        <span>连接</span>
      </el-menu-item>
      <el-menu-item index="/log">
        <el-icon>
          <Memo/>
        </el-icon>
        <span>日志</span>
      </el-menu-item>
      <el-menu-item index="/config">
        <el-icon>
          <Setting/>
        </el-icon>
        <span>设置</span>
      </el-menu-item>
      <el-menu-item index="/info">
        <el-icon>
          <InfoFilled/>
        </el-icon>
        <span>信息</span>
      </el-menu-item>
    </el-menu>

    <CollapseBtn :isCollapse="isCollapse" @collapse="isCollapse =!isCollapse"></CollapseBtn>
  </el-aside>
</template>

<style scoped>
.el-aside, .el-menu {
  background-color: #3d525e;
}

.el-aside {
  height: 100%;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.el-menu {
  flex: 1;
}

.el-menu-item {
  color: white;
}

h4 {
  color: white;
  text-align: center;
}

.btn-div {
  text-align: right;
}
</style>