<script lang="ts">
export default {
  name: "HomeView",
};
</script>

<script setup lang="ts">
import { ref, watch } from "vue";
import { createRoom } from "@/api/room";
import BaseInput from "@/components/base/BaseInput.vue";
import BaseButton from "@/components/base/BaseButton.vue";
import { useRouter } from "vue-router";
const meetingCode = ref<string>("");
const visibility = ref<"hidden" | "visible">("visible");
watch(
  meetingCode,
  (v) => {
    if (v.length > 0) {
      visibility.value = "visible";
      return;
    }
    visibility.value = "hidden";
  },
  {
    immediate: true,
  }
);
const router = useRouter();
async function onSubmit() {
  const formData = new FormData();
  formData.append("name", "New meet");
  const response = await createRoom(formData);
  if (!response.success) {
    console.error(response.err);
    return;
  }
  router.push(`/${response.data.room_id}`);
}
function join() {
  router.push(`/${meetingCode.value}`);
}
</script>

<template>
  <form @submit.prevent="onSubmit">
    <div>
      <BaseButton> New meeting</BaseButton>
    </div>
    <div>
      <BaseInput v-model="meetingCode" placeholder="Enter code" />
    </div>
    <div :style="{ visibility }" @click="join" style="cursor: pointer">
      Join
    </div>
  </form>
</template>

<style>
form {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 10px;
}
</style>
