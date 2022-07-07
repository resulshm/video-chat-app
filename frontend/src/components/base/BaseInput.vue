<script setup lang="ts">
import { computed } from "vue";
const props = defineProps({ modelValue: String, placeholder: String });
const emits = defineEmits(["update:modelValue"]);
const value = computed({
  get: () => {
    return props.modelValue;
  },
  set: (value) => {
    emits("update:modelValue", value);
  },
});
</script>
<template>
  <input :placeholder="placeholder" v-model="value" />
</template>

<style>
input {
  border: 1px solid #ccc;
  padding: 7.5px;
  border-radius: 3px;
}
input:focus {
  outline: none;
}
</style>
