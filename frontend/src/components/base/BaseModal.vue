<script setup>
import { GDialog } from "gitart-vue-dialog";
import "gitart-vue-dialog/dist/style.css";
import BaseButton from "./BaseButton.vue";
defineProps({
  dialogState: {
    type: Boolean,
    default: false,
  },
});
</script>

<template>
  <GDialog v-model="dialogState" max-width="400" persistent>
    <div class="wrapper">
      <div class="title">
        <slot name="title"></slot>
      </div>
      <div class="body">
        <slot name="body"> </slot>
      </div>
      <div class="actions">
        <BaseButton color="red" @click="$emit('close')">Close</BaseButton>
        <slot name="actions"></slot>
      </div>
    </div>
  </GDialog>
</template>

<style scoped>
.wrapper {
  color: #000;
}

.title {
  font-size: 22.5px;
  font-weight: 700;
  margin: 20px;
}
.body {
  margin: 10px 20px;
}
.actions {
  display: flex;
  gap: 10px;
  justify-content: flex-end;
  padding: 10px 15px;
}
</style>
