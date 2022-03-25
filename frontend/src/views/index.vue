<script lang="ts">
export default {
  name: "HomeView",
};
</script>

<script setup lang="ts">
import { ref } from "vue";
import BaseInput from "@/components/BaseComponents/BaseInput.vue";
const name = ref<string>();
const room = ref<string>();
</script>

<template>
  <div>
    <form>
      <BaseInput v-model="name" />
      <BaseInput v-model="room" />
    </form>
  </div>
</template>
