<script setup lang="ts">
import { computed } from "vue";
const props = defineProps({ modelValue: String });
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
    <input class="border border-c-black" v-model="value" />
  </div>
</template>
