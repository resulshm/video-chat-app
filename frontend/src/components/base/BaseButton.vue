<script setup lang="ts">
import { onMounted, ref } from "vue";
const { color } = defineProps({
  color: {
    type: String,
    default: "blue",
  },
});
const colorCode = ref<string>("#007bff");
onMounted(() => {
  if (color === "red") {
    colorCode.value = "#ff4444";
  }
});
</script>

<template>
  <button :style="{ 'background-color': colorCode }">
    <slot></slot>
  </button>
</template>

<style scoped>
button {
  color: #ffff;
  border: #ffff;
  padding: 7.5px;
  border-radius: 3px;
  cursor: pointer;
}

button:hover {
  opacity: 0.8;
}
</style>
