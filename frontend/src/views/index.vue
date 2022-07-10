<script lang="ts">
export default {
  name: "HomeView",
};
</script>

<script setup lang="ts">
import { ref, watch } from "vue";
import { useFetch } from "@vueuse/core";
import { useRouter } from "vue-router";

const router = useRouter();
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

async function onSubmit() {
  const { data, error } = await useFetch("api/v1/room", {
    method: "POST",
  }).json();
  if (error.value) {
    console.error(error);
    return;
  }
  router.push(`/${data.value.room_id}`);
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
