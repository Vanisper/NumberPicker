<script setup lang="ts">
import { onMounted } from "vue";
import parentLayout from "./layouts/parentLayout.vue";
import { WindowMaximise } from "@wailsjs/runtime";

onMounted(() => {
  WindowMaximise();
});
</script>

<template>
  <parent-layout :LayoutMode="'column'" />
</template>

<style lang="less">
@tailwind base;
@tailwind components;
@tailwind utilities;

html,
body {
  background-color: var(--background-primary);
}

/*禁止复制*/
div,
view,
span,
button,
select {
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
}
</style>
