<script setup lang="ts">
import parentLayout from "./layouts/parentLayout.vue";

</script>

<template>
  <parent-layout :LayoutMode="'column'" />
</template>

<style lang="less">
html,
body {
  background-color: var(--background-primary);
}
</style>
