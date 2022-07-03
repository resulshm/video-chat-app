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
  <div>
    <input
      :placeholder="placeholder"
      class="w-full h-12 border border-gray-400 outline-none px-3 py-3 rounded-[5px] focus:border-blue-500 focus:border-[2px] placeholder:text-gray-600 box-border"
      v-model="value"
    />
  </div>
</template>
