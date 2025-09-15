<script lang="ts" setup>
import AsideMenu from "./components/AsideMenu.vue";
</script>

<template>
  <el-container class="main-container" >
    <AsideMenu/>
    <RouterView></RouterView>
  </el-container>
</template>

<style>
html, body, #app {
  height: 100%;
  width: 100%;
  margin: 0;
  padding: 0;
  overflow: hidden;
}


.main-container {
  height: 100%;
  overflow: hidden;
}
</style>
