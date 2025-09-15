<script setup lang="ts">

defineProps({
  'isCollapse' : Boolean,
})

defineEmits<{
  (e: 'collapse'): void
}>()
</script>

<template>
  <div class="collapse-button">
    <el-button @click="$emit('collapse')" type="text" size="large">
      <el-icon color="gray">
        <template v-if="isCollapse">
          <Expand/>
        </template>
        <template v-else>
          <Fold/>
        </template>
      </el-icon>
    </el-button>
  </div>
</template>

<style scoped>
.collapse-button {
  text-align: right;
  padding: 5px 15px;
}
</style>