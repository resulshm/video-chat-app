<script lang="ts">
export default {
  name: "HomeView",
};
</script>

<script setup lang="ts">
import { ref } from "vue";
import { createRoom } from "@/api/room";
import BaseInput from "@/components/BaseComponents/BaseInput.vue";
import { useRouter } from "vue-router";
const room = ref<string>();
const router = useRouter();
async function onSubmit() {
  const formData = new FormData();
  formData.append("name", "Web");
  const newRoom = await createRoom(formData);
  console.log(newRoom);
  router.push(`/${newRoom.data.data.room_id}`);
}
</script>

<template>
  <div>
    <form @submit.prevent="onSubmit">
      <BaseInput v-model="room" />
      <button>Create room</button>
    </form>
  </div>
</template>
